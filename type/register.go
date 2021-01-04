/*---------------------------------------------------------------------------------------------
 *  Copyright (c) Peter Bjorklund. All rights reserved.
 *  Licensed under the MIT License. See LICENSE in the project root for license information.
 *--------------------------------------------------------------------------------------------*/

package swampopcodetype

import "fmt"

type Register struct {
	id uint8
}

func NewRegister(a uint8) Register {
	return Register{id: a}
}

func (r Register) Value() uint8 {
	return r.id
}

func (r Register) String() string {
	return fmt.Sprintf("%v", r.id)
}
