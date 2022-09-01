package message

import (
	"os"
	"path"

	"github.com/NicoNex/echotron/v3"
)

// MediaInfo countain all the infos about medias, Polls and so on contained into a message
type MediaInfo struct {
	MediaGroupID    string                    `json:"media_group_id,omitempty"`
	Animation       *echotron.Animation       `json:"animation,omitempty"`
	Audio           *echotron.Audio           `json:"audio,omitempty"`
	Document        *echotron.Document        `json:"document,omitempty"`
	Photo           []*echotron.PhotoSize     `json:"photo,omitempty"`
	Sticker         *echotron.Sticker         `json:"sticker,omitempty"`
	Video           *echotron.Video           `json:"video,omitempty"`
	VideoNote       *echotron.VideoNote       `json:"video_note,omitempty"`
	Voice           *echotron.Voice           `json:"voice,omitempty"`
	Caption         string                    `json:"caption,omitempty"`
	CaptionEntities []*echotron.MessageEntity `json:"caption_entities,omitempty"`
	Contact         *echotron.Contact         `json:"contact,omitempty"`
	Dice            *echotron.Dice            `json:"dice,omitempty"`
	Game            *echotron.Game            `json:"game,omitempty"`
	Poll            *echotron.Poll            `json:"poll,omitempty"`
	Venue           *echotron.Venue           `json:"venue,omitempty"`
	Location        *echotron.Location        `json:"location,omitempty"`
}

// ExtractFileID try to extracts the FileID from the Info of a media
func (m MediaInfo) ExtractFileID() (id *FileID) {
    switch {
    case m.Animation != nil:
        *id = GrabAnimationFileID(m.Animation)
    case m.Audio != nil:
        *id = GrabAudioFileID(m.Audio)
    case m.Document != nil:
        *id = GrabDocumentFileID(m.Document)
    case m.Photo != nil && len(m.Photo) > 0:
        *id = GrabPhotoFileID(m.Photo[len(m.Photo) - 1])
    case m.Sticker != nil:
        *id = GrabStickerFileID(m.Sticker)
    case m.Video != nil:
        *id = GrabVideoFileID(m.Video)
    case m.VideoNote != nil:
        *id = GrabVideoNoteFileID(m.VideoNote)
    case m.Voice != nil:
        *id = GrabVoiceFileID(m.Voice)
    }

    return
}

// FileID is the identifier of a sent or recived File of any kind
type FileID string

// RetrieveInfo retrieve the info of a particular file from Telegram servers
func (id FileID) RetrieveInfo() (file *echotron.File, err error) {
    var res echotron.APIResponseFile

	res, err = api.GetFile(string(id))
    err = parseResponseError(res, err)
    if err == nil {
        file = res.Result
    }
    return
}

// FetchFile fetch the file content from Telegram servers
func (id FileID) FetchFile() (content []byte, err error) {
    var file *echotron.File

    file, err = id.RetrieveInfo()
    if err != nil {
        return
    }

	return api.DownloadFile(file.FilePath)
}

// SaveFile downloads a file in the given directory at the same relative path
// specified by Telegram and returns the complete path where has been saved locally,
// it's content and error
func (id FileID) SaveFile(directory string) (filePath string, content []byte, err error) {
    if res, e := id.RetrieveInfo(); e != nil {
        err = e
    } else {
        filePath = res.FilePath
    }

    content, err = api.DownloadFile(filePath)
    if err != nil {
        return
    }

    filePath = path.Join(directory, filePath)
	err = os.WriteFile(filePath, content, os.ModePerm)
	return
}


// GrabAnimationFileID grabs the FileID from the given media
func GrabAnimationFileID(media *echotron.Animation) FileID {
    return FileID(media.FileID)
}

// GrabAudioFileID grabs the FileID from the given media
func GrabAudioFileID(media *echotron.Audio) FileID {
    return FileID(media.FileID)
}

// GrabDocumentFileID grabs the FileID from the given media
func GrabDocumentFileID(media *echotron.Document) FileID {
    return FileID(media.FileID)
}

// GrabPhotoFileID grabs the FileID from the given media
func GrabPhotoFileID(media *echotron.PhotoSize) FileID {
    return FileID(media.FileID)
}

// GrabStickerFileID grabs the FileID from the given media
func GrabStickerFileID(media *echotron.Sticker) FileID {
    return FileID(media.FileID)
}

// GrabVideoFileID grabs the FileID from the given media
func GrabVideoFileID(media *echotron.Video) FileID {
    return FileID(media.FileID)
}

// GrabVideoNoteFileID grabs the FileID from the given media
func GrabVideoNoteFileID(media *echotron.VideoNote) FileID {
    return FileID(media.FileID)
}

// GrabVoiceFileID grabs the FileID from the given media
func GrabVoiceFileID(media *echotron.Voice) FileID {
    return FileID(media.FileID)
}
