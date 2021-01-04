/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package swampopcodetype

import "fmt"

type Field struct {
	index uint8
}

func (f Field) Value() uint8 {
	return f.index
}

func NewField(a uint8) Field {
	return Field{index: a}
}

func (f Field) String() string {
	return fmt.Sprintf("#%v", f.index)
}
