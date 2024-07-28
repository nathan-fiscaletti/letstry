package editors

import "time"

const (
	EditorNameNotepadPlusPlus EditorName = "notepad++"
)

func NotepadPlusPlusEditor() Editor {
	return Editor{
		Name:                EditorNameNotepadPlusPlus,
		RunType:             EditorRunTypeStart,
		ExecPath:            "C:\\Program Files\\Notepad++\\notepad++.exe",
		Args:                "-multiInst -openFoldersAsWorkspace",
		ProcessCaptureDelay: time.Second * 2,
		TrackingType:        TrackingTypeProcess,
	}
}
