package States

import "Mercury/pkg/Type"

func PassState(task Type.Task, events []byte) (Type.Task, []byte) {
	// Pass State Implementation
	return task, events
}
