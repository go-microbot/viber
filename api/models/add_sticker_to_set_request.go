package models

import "github.com/go-microbot/telegram/form"

// AddStickerToSetRequest represents `addStickerToSet` request body.
type AddStickerToSetRequest struct {
	// User identifier of sticker set owner.
	UserID form.PartAny `form:"user_id"`
	// Sticker set name.
	Name form.PartText `form:"name"`
	// Optional. PNG image with the sticker, must be up to 512 kilobytes in size,
	// dimensions must not exceed 512px, and either width or height must be exactly 512px.
	// Pass a file_id as a String to send a file that already exists on the Telegram servers,
	// pass an HTTP URL as a String for Telegram to get a file from the Internet,
	// or upload a new one using multipart/form-data.
	PngSticker form.Part `form:"png_sticker,omitempty"`
	// Optional. TGS animation with the sticker, uploaded using multipart/form-data.
	// See https://core.telegram.org/animated_stickers#technical-requirements
	// for technical requirements.
	TgsSticker form.PartFile `form:"tgs_sticker,omitempty"`
	// One or more emoji corresponding to the sticker.
	Emojis form.PartText `form:"emojis"`
	// Optional. A JSON-serialized object for position where the mask should be placed on faces.
	MaskPosition form.PartJSON `form:"mask_position,omitempty"`
}
