package main


func main(){
	_ , err := ReadJsonFile("reza.json")

	if err != nil {
		if errors.As(err, FileNotFound{})
		fmt.Println(err.Error())
	}
}
func ReadJsonFile(fileName string) (map[string]any, error) {
	m := map[string]any{}
	// step 1 :
	// read file
	// errors : file does not exist
	return m, FileNotFound{FileName: fileName}
	// errors : can not read file

	// step 2
	// convert to hashmap
}

type FileNotFound struct {
	FileName string
}

func (e FileNotFound) Error() string {
	return e.FileName + " not found"
}
