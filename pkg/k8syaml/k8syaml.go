package k8syaml

import (
	"fmt"
	"io"
	"strings"

	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	yamlUtil "k8s.io/apimachinery/pkg/util/yaml"
	"sigs.k8s.io/yaml"
)

// Decode returns a slice of parse unstructured Kubernetes API objects from a
// specified JSON or YAML file.
func Decode(in io.Reader) ([]unstructured.Unstructured, error) {
	decoder := yamlUtil.NewYAMLToJSONDecoder(in)
	var result []unstructured.Unstructured
	var err error
	for err == nil || isEmptyYamlError(err) {
		var out unstructured.Unstructured
		err = decoder.Decode(&out)
		if err == nil {
			result = append(result, out)
		}
	}

	if err != io.EOF {
		return nil, err
	}

	return result, nil
}

// Encode prints the specified documents encode as YAML into the writer.
func Encode(out io.Writer, documents []unstructured.Unstructured) error {
	for _, doc := range documents {
		bytes, err := yaml.Marshal(doc.Object)
		if err != nil {
			return err
		}

		fmt.Fprintln(out, "---")
		out.Write(bytes)
	}

	return nil
}

func isEmptyYamlError(err error) bool {
	return strings.Contains(err.Error(), "is missing in 'null'")
}