/*
  This go program parses Spark AR patch json data
*/

package arputil

func getPort(patch Patch, portType string) PatchModelChildren {
  var result PatchModelChildren

  grandchildren := patch.Children[0].Children

  for _, grandchild := range grandchildren {
    if grandchild.ModelName == portType {
      result = grandchild
      break
    }
  }

  return result
}

func getPortNames(portType PatchModelChildren) []string {
  var result []string

  for _, port := range portType.Children {
    result = append(result, port.Name)
  }

  return result
}

// Returns an array of the input port names
func InputPorts(patch Patch) []string {
  return getPortNames(getPort(patch, "patch_input_ports_model"))
}

// Returns an array of the output port names
func OutputPorts(patch Patch) []string {
  return getPortNames(getPort(patch, "patch_output_ports_model"))
}
