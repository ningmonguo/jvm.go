package lang

import (
	"github.com/zxh0/jvm.go/rtda"
	"github.com/zxh0/jvm.go/rtda/heap"
	"github.com/zxh0/jvm.go/vmutils"
)

func init() {
	_class(getRawAnnotations, "getRawAnnotations", "()[B")
}

// native byte[] getRawAnnotations();
// ()[B
func getRawAnnotations(frame *rtda.Frame) {
	this := frame.GetThis()

	class := this.Extra.(*heap.Class)
	goBytes := class.AnnotationData
	if goBytes != nil {
		jBytes := vmutils.CastBytesToInt8s(goBytes)
		byteArr := heap.NewByteArray(jBytes)
		frame.PushRef(byteArr)
		return
	}

	frame.PushRef(nil)
}

// native byte[] getRawTypeAnnotations();
// ()[B
func getRawTypeAnnotations(frame *rtda.Frame) {
	// todo
}
