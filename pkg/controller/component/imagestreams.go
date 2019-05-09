package component

import (
	"github.com/snowdrop/component-operator/pkg/apis/component/v1alpha2"
	imagev1 "github.com/openshift/api/image/v1"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"strings"
)

const (
	SUPERVISOR_IMAGE_NAME = "copy-supervisord"
)

func init() {
	image["java"] = "quay.io/snowdrop/spring-boot-s2i"
	image["nodejs"] = "nodeshift/centos7-s2i-nodejs:10.x"
	image["supervisord"] = "quay.io/snowdrop/supervisord"
}

//buildImageStream returns the service resource
func (r *ReconcileComponent) buildImageStream(c *v1alpha2.Component, imageName string) *imagev1.ImageStream {
	ls := r.getAppLabels(c.Name)
	imageSpecInfo := r.getImageInfoSpec(c, imageName)
	ser := &imagev1.ImageStream{
		TypeMeta: metav1.TypeMeta{
			APIVersion: "v1",
			Kind:       "Service",
		},
		ObjectMeta: metav1.ObjectMeta{
			Name:      imageName,
			Namespace: c.Namespace,
			Labels:    ls,
		},
		Spec: imagev1.ImageStreamSpec{
			LookupPolicy: imagev1.ImageLookupPolicy{false},
			Tags: []imagev1.TagReference{
				{
					Annotations: r.generateAnnotations(imageSpecInfo.Name),
					From:      &corev1.ObjectReference{Kind: "DockerImage", Name: imageSpecInfo.Repo},
					Name:      imageSpecInfo.Tag,
					Reference: true,
				},
			},
		},
	}
	return ser
}

func (r *ReconcileComponent) generateAnnotations(name string) map[string]string{
	if ( name == SUPERVISOR_IMAGE_NAME ) {
		return map[string]string{"cmds": "run-java:/usr/local/s2i/run;run-node:/usr/libexec/s2i;compile-java:/usr/local/s2i/assemble;build:/deployments/buildapp"}
	} else {
		return map[string]string{}
	}
}

func (r *ReconcileComponent) getDevImageNames(c *v1alpha2.Component) []string {
	return []string{
		SUPERVISOR_IMAGE_NAME,
		r.getRuntimeImageName(c),
	}
}

func (r *ReconcileComponent) getRuntimeImageName(c *v1alpha2.Component) string {
	return strings.Join([]string{"dev-runtime", strings.ToLower(c.Spec.Runtime)}, "-")
}

// Get the key of the image stream to of the runtime
func (r *ReconcileComponent) getRuntimeKey(c *v1alpha2.Component) string {
	switch r := c.Spec.Runtime; r {
	case "spring-boot", "vert.x", "thorntail":
		return "java"
	case "nodejs":
		return "nodejs"
	default:
		return "java"
	}
}

func (r *ReconcileComponent) getImageInfoSpec(c *v1alpha2.Component, name string) *v1alpha2.Image {
	if ( name == SUPERVISOR_IMAGE_NAME ) {
		return r.buildSupervisordImage()
	} else {
		return r.buildImage(true, c.Spec.RuntimeName, "latest", image[r.getRuntimeKey(c)], false)
	}
}

func (r *ReconcileComponent) buildSupervisordImage() *v1alpha2.Image {
	return r.buildImage(true, "copy-supervisord", "latest", image["supervisord"], true)
}

func (r *ReconcileComponent) buildImage(dockerImage bool, name string, tag string, repo string, annotationCmd bool) *v1alpha2.Image {
	return &v1alpha2.Image{
		DockerImage:    dockerImage,
		Name:           name,
		Repo:           repo,
		AnnotationCmds: annotationCmd,
		Tag:            tag,
	}
}