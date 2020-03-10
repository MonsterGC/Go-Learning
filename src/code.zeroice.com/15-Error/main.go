package main

import "fmt"


type DivideError struct{
	divided int
	divider int
}

func (de *DivideError)Error() string {
	strError := `
	Cannot proceed, the divider is zero.
    dividee: %d
    divider: 0
	`
	// return fmt.Printf(strError,de.divider)
	return fmt.Sprintf(strError,de.divided)
}

func Divide(varDivided int , varDivider int)(result int , errorMsg string)  {
	if varDivider  == 0 {
		errorData := DivideError{
			divided : varDivided,
			divider : varDivider,
		}
		errorMsg = errorData.Error()
		return
	}else{
		return varDivided / varDivider , ""
	}
}

func main()  {
	// 正常情况
    if result, errorMsg := Divide(100, 10); errorMsg == "" {
		fmt.Println("100/10 = ", result)
	}
	// 当除数为零的时候会返回错误信息
	if _, errorMsg := Divide(100, 0); errorMsg != "" {
			fmt.Println("errorMsg is: ", errorMsg)
	}
}