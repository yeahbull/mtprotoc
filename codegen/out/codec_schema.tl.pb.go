/*
 * WARNING! All changes made in this file will be lost!
 * Created from 'scheme.tl' by 'codegen_encode_decode.py'
 *
 *  Copyright (c) 2017, https://github.com/nebulaim
 *  All rights reserved.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *   http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

// ConstructorList
// RequestList

package mtproto

import (
// "encoding/binary"
// "fmt"
// "github.com/golang/protobuf/proto"
)

///////////////////////////////////////////////////////////////////////////////
// Help_Support <--
//  + TL_HelpSupport
//

func (m *Help_Support) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_help_support:
		t := m.To_HelpSupport()
		return t.Encode()

	default:
		return []byte{}
	}
}

// help.support#17c6b5f6 phone_number:string user:User = help.Support;
func (m *Help_Support) To_HelpSupport() *TLHelpSupport {
	return &TLHelpSupport{
		Data2: m.Data2,
	}
}

// help.support#17c6b5f6 phone_number:string user:User = help.Support;
func (m *TLHelpSupport) To_Help_Support() *Help_Support {
	return &Help_Support{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLHelpSupport) SetPhoneNumber(v string) { m.Data2.PhoneNumber = v }
func (m *TLHelpSupport) GetPhoneNumber() string  { return m.Data2.PhoneNumber }

func (m *TLHelpSupport) SetUser(v *User) { m.Data2.User = v }
func (m *TLHelpSupport) GetUser() *User  { return m.Data2.User }

func (m *TLHelpSupport) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_help_support))

	x.StringBytes(m.Data2.PhoneNumber)
	x.Bytes(m.Data2User.Encode())

	return x.buf
}

func (m *TLHelpSupport) Decode(dbuf *DecodeBuf) error {
	m.Data2.PhoneNumber = x.StringBytes()
	m2 := &User{}
	m2.Decode(dbuf)
	m.SetUser(m2)

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// WebPage <--
//  + TL_WebPageEmpty
//  + TL_WebPagePending
//  + TL_WebPage
//  + TL_WebPageNotModified
//

func (m *WebPage) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_webPageEmpty:
		t := m.To_WebPageEmpty()
		return t.Encode()
	case TLConstructor_CRC32_webPagePending:
		t := m.To_WebPagePending()
		return t.Encode()
	case TLConstructor_CRC32_webPage:
		t := m.To_WebPage()
		return t.Encode()
	case TLConstructor_CRC32_webPageNotModified:
		t := m.To_WebPageNotModified()
		return t.Encode()

	default:
		return []byte{}
	}
}

// webPageEmpty#eb1477e8 id:long = WebPage;
func (m *WebPage) To_WebPageEmpty() *TLWebPageEmpty {
	return &TLWebPageEmpty{
		Data2: m.Data2,
	}
}

// webPagePending#c586da1c id:long date:int = WebPage;
func (m *WebPage) To_WebPagePending() *TLWebPagePending {
	return &TLWebPagePending{
		Data2: m.Data2,
	}
}

// webPage#5f07b4bc flags:# id:long url:string display_url:string hash:int type:flags.0?string site_name:flags.1?string title:flags.2?string description:flags.3?string photo:flags.4?Photo embed_url:flags.5?string embed_type:flags.5?string embed_width:flags.6?int embed_height:flags.6?int duration:flags.7?int author:flags.8?string document:flags.9?Document cached_page:flags.10?Page = WebPage;
func (m *WebPage) To_WebPage() *TLWebPage {
	return &TLWebPage{
		Data2: m.Data2,
	}
}

// webPageNotModified#85849473 = WebPage;
func (m *WebPage) To_WebPageNotModified() *TLWebPageNotModified {
	return &TLWebPageNotModified{
		Data2: m.Data2,
	}
}

// webPageEmpty#eb1477e8 id:long = WebPage;
func (m *TLWebPageEmpty) To_WebPage() *WebPage {
	return &WebPage{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLWebPageEmpty) SetId(v int64) { m.Data2.Id = v }
func (m *TLWebPageEmpty) GetId() int64  { return m.Data2.Id }

func (m *TLWebPageEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_webPageEmpty))

	x.Long(m.Data2.Id)

	return x.buf
}

func (m *TLWebPageEmpty) Decode(dbuf *DecodeBuf) error {
	m.Data2.Id = x.Long()

	return dbuf.err
}

// webPagePending#c586da1c id:long date:int = WebPage;
func (m *TLWebPagePending) To_WebPage() *WebPage {
	return &WebPage{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLWebPagePending) SetId(v int64) { m.Data2.Id = v }
func (m *TLWebPagePending) GetId() int64  { return m.Data2.Id }

func (m *TLWebPagePending) SetDate(v int32) { m.Data2.Date = v }
func (m *TLWebPagePending) GetDate() int32  { return m.Data2.Date }

func (m *TLWebPagePending) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_webPagePending))

	x.Long(m.Data2.Id)
	x.Int(m.Data2.Date)

	return x.buf
}

func (m *TLWebPagePending) Decode(dbuf *DecodeBuf) error {
	m.Data2.Id = x.Long()
	m.Data2.Date = x.Int()

	return dbuf.err
}

// webPage#5f07b4bc flags:# id:long url:string display_url:string hash:int type:flags.0?string site_name:flags.1?string title:flags.2?string description:flags.3?string photo:flags.4?Photo embed_url:flags.5?string embed_type:flags.5?string embed_width:flags.6?int embed_height:flags.6?int duration:flags.7?int author:flags.8?string document:flags.9?Document cached_page:flags.10?Page = WebPage;
func (m *TLWebPage) To_WebPage() *WebPage {
	return &WebPage{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLWebPage) SetId(v int64) { m.Data2.Id = v }
func (m *TLWebPage) GetId() int64  { return m.Data2.Id }

func (m *TLWebPage) SetUrl(v string) { m.Data2.Url = v }
func (m *TLWebPage) GetUrl() string  { return m.Data2.Url }

func (m *TLWebPage) SetDisplayUrl(v string) { m.Data2.DisplayUrl = v }
func (m *TLWebPage) GetDisplayUrl() string  { return m.Data2.DisplayUrl }

func (m *TLWebPage) SetHash(v int32) { m.Data2.Hash = v }
func (m *TLWebPage) GetHash() int32  { return m.Data2.Hash }

func (m *TLWebPage) SetType(v string) { m.Data2.Type = v }
func (m *TLWebPage) GetType() string  { return m.Data2.Type }

func (m *TLWebPage) SetSiteName(v string) { m.Data2.SiteName = v }
func (m *TLWebPage) GetSiteName() string  { return m.Data2.SiteName }

func (m *TLWebPage) SetTitle(v string) { m.Data2.Title = v }
func (m *TLWebPage) GetTitle() string  { return m.Data2.Title }

func (m *TLWebPage) SetDescription(v string) { m.Data2.Description = v }
func (m *TLWebPage) GetDescription() string  { return m.Data2.Description }

func (m *TLWebPage) SetPhoto(v *Photo) { m.Data2.Photo = v }
func (m *TLWebPage) GetPhoto() *Photo  { return m.Data2.Photo }

func (m *TLWebPage) SetEmbedUrl(v string) { m.Data2.EmbedUrl = v }
func (m *TLWebPage) GetEmbedUrl() string  { return m.Data2.EmbedUrl }

func (m *TLWebPage) SetEmbedType(v string) { m.Data2.EmbedType = v }
func (m *TLWebPage) GetEmbedType() string  { return m.Data2.EmbedType }

func (m *TLWebPage) SetEmbedWidth(v int32) { m.Data2.EmbedWidth = v }
func (m *TLWebPage) GetEmbedWidth() int32  { return m.Data2.EmbedWidth }

func (m *TLWebPage) SetEmbedHeight(v int32) { m.Data2.EmbedHeight = v }
func (m *TLWebPage) GetEmbedHeight() int32  { return m.Data2.EmbedHeight }

func (m *TLWebPage) SetDuration(v int32) { m.Data2.Duration = v }
func (m *TLWebPage) GetDuration() int32  { return m.Data2.Duration }

func (m *TLWebPage) SetAuthor(v string) { m.Data2.Author = v }
func (m *TLWebPage) GetAuthor() string  { return m.Data2.Author }

func (m *TLWebPage) SetDocument(v *Document) { m.Data2.Document = v }
func (m *TLWebPage) GetDocument() *Document  { return m.Data2.Document }

func (m *TLWebPage) SetCachedPage(v *Page) { m.Data2.CachedPage = v }
func (m *TLWebPage) GetCachedPage() *Page  { return m.Data2.CachedPage }

func (m *TLWebPage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_webPage))

	// flags
	var flags uint32 = 0
	if m.GetType() != "" {
		flags |= 1 << 5
	}
	if m.GetSiteName() != "" {
		flags |= 1 << 6
	}
	if m.GetTitle() != "" {
		flags |= 1 << 7
	}
	if m.GetDescription() != "" {
		flags |= 1 << 8
	}
	if m.GetPhoto() != nil {
		flags |= 1 << 9
	}
	if m.GetEmbedUrl() != "" {
		flags |= 1 << 10
	}
	if m.GetEmbedType() != "" {
		flags |= 1 << 11
	}
	if m.GetEmbedWidth() != 0 {
		flags |= 1 << 12
	}
	if m.GetEmbedHeight() != 0 {
		flags |= 1 << 13
	}
	if m.GetDuration() != 0 {
		flags |= 1 << 14
	}
	if m.GetAuthor() != "" {
		flags |= 1 << 15
	}
	if m.GetDocument() != nil {
		flags |= 1 << 16
	}
	if m.GetCachedPage() != nil {
		flags |= 1 << 17
	}
	x.UInt(flags)

	x.Long(m.Data2.Id)
	x.StringBytes(m.Data2.Url)
	x.StringBytes(m.Data2.DisplayUrl)
	x.Int(m.Data2.Hash)
	if m.GetType() != 0 {
		x.StringBytes(m.Data2.Type)
	}
	if m.GetSiteName() != 0 {
		x.StringBytes(m.Data2.SiteName)
	}
	if m.GetTitle() != 0 {
		x.StringBytes(m.Data2.Title)
	}
	if m.GetDescription() != 0 {
		x.StringBytes(m.Data2.Description)
	}
	if m.GetPhoto() != 0 {
		x.Bytes(m.Data2Photo.Encode())
	}
	if m.GetEmbedUrl() != 0 {
		x.StringBytes(m.Data2.EmbedUrl)
	}
	if m.GetEmbedType() != 0 {
		x.StringBytes(m.Data2.EmbedType)
	}
	if m.GetEmbedWidth() != 0 {
		x.Int(m.Data2.EmbedWidth)
	}
	if m.GetEmbedHeight() != 0 {
		x.Int(m.Data2.EmbedHeight)
	}
	if m.GetDuration() != 0 {
		x.Int(m.Data2.Duration)
	}
	if m.GetAuthor() != 0 {
		x.StringBytes(m.Data2.Author)
	}
	if m.GetDocument() != 0 {
		x.Bytes(m.Data2Document.Encode())
	}
	if m.GetCachedPage() != 0 {
		x.Bytes(m.Data2CachedPage.Encode())
	}

	return x.buf
}

func (m *TLWebPage) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	m.Data2.Id = x.Long()
	m.Data2.Url = x.StringBytes()
	m.Data2.DisplayUrl = x.StringBytes()
	m.Data2.Hash = x.Int()
	if (flags & (1 << 6)) != 0 {
		m.Data2.Type = x.StringBytes()
	}
	if (flags & (1 << 7)) != 0 {
		m.Data2.SiteName = x.StringBytes()
	}
	if (flags & (1 << 8)) != 0 {
		m.Data2.Title = x.StringBytes()
	}
	if (flags & (1 << 9)) != 0 {
		m.Data2.Description = x.StringBytes()
	}
	if (flags & (1 << 10)) != 0 {
		m10 := &Photo{}
		m10.Decode(dbuf)
		m.SetPhoto(m10)
	}
	if (flags & (1 << 11)) != 0 {
		m.Data2.EmbedUrl = x.StringBytes()
	}
	if (flags & (1 << 12)) != 0 {
		m.Data2.EmbedType = x.StringBytes()
	}
	if (flags & (1 << 13)) != 0 {
		m.Data2.EmbedWidth = x.Int()
	}
	if (flags & (1 << 14)) != 0 {
		m.Data2.EmbedHeight = x.Int()
	}
	if (flags & (1 << 15)) != 0 {
		m.Data2.Duration = x.Int()
	}
	if (flags & (1 << 16)) != 0 {
		m.Data2.Author = x.StringBytes()
	}
	if (flags & (1 << 17)) != 0 {
		m17 := &Document{}
		m17.Decode(dbuf)
		m.SetDocument(m17)
	}
	if (flags & (1 << 18)) != 0 {
		m18 := &CachedPage{}
		m18.Decode(dbuf)
		m.SetCachedPage(m18)
	}

	return dbuf.err
}

// webPageNotModified#85849473 = WebPage;
func (m *TLWebPageNotModified) To_WebPage() *WebPage {
	return &WebPage{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLWebPageNotModified) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_webPageNotModified))

	return x.buf
}

func (m *TLWebPageNotModified) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// StickerSet <--
//  + TL_StickerSet
//

func (m *StickerSet) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_stickerSet:
		t := m.To_StickerSet()
		return t.Encode()

	default:
		return []byte{}
	}
}

// stickerSet#cd303b41 flags:# installed:flags.0?true archived:flags.1?true official:flags.2?true masks:flags.3?true id:long access_hash:long title:string short_name:string count:int hash:int = StickerSet;
func (m *StickerSet) To_StickerSet() *TLStickerSet {
	return &TLStickerSet{
		Data2: m.Data2,
	}
}

// stickerSet#cd303b41 flags:# installed:flags.0?true archived:flags.1?true official:flags.2?true masks:flags.3?true id:long access_hash:long title:string short_name:string count:int hash:int = StickerSet;
func (m *TLStickerSet) To_StickerSet() *StickerSet {
	return &StickerSet{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLStickerSet) SetInstalled(v bool) { m.Data2.Installed = v }
func (m *TLStickerSet) GetInstalled() bool  { return m.Data2.Installed }

func (m *TLStickerSet) SetArchived(v bool) { m.Data2.Archived = v }
func (m *TLStickerSet) GetArchived() bool  { return m.Data2.Archived }

func (m *TLStickerSet) SetOfficial(v bool) { m.Data2.Official = v }
func (m *TLStickerSet) GetOfficial() bool  { return m.Data2.Official }

func (m *TLStickerSet) SetMasks(v bool) { m.Data2.Masks = v }
func (m *TLStickerSet) GetMasks() bool  { return m.Data2.Masks }

func (m *TLStickerSet) SetId(v int64) { m.Data2.Id = v }
func (m *TLStickerSet) GetId() int64  { return m.Data2.Id }

func (m *TLStickerSet) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLStickerSet) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLStickerSet) SetTitle(v string) { m.Data2.Title = v }
func (m *TLStickerSet) GetTitle() string  { return m.Data2.Title }

func (m *TLStickerSet) SetShortName(v string) { m.Data2.ShortName = v }
func (m *TLStickerSet) GetShortName() string  { return m.Data2.ShortName }

func (m *TLStickerSet) SetCount(v int32) { m.Data2.Count = v }
func (m *TLStickerSet) GetCount() int32  { return m.Data2.Count }

func (m *TLStickerSet) SetHash(v int32) { m.Data2.Hash = v }
func (m *TLStickerSet) GetHash() int32  { return m.Data2.Hash }

func (m *TLStickerSet) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_stickerSet))

	// flags
	var flags uint32 = 0
	if m.GetInstalled() == true {
		flags |= 1 << 1
	}
	if m.GetArchived() == true {
		flags |= 1 << 2
	}
	if m.GetOfficial() == true {
		flags |= 1 << 3
	}
	if m.GetMasks() == true {
		flags |= 1 << 4
	}
	x.UInt(flags)

	x.Long(m.Data2.Id)
	x.Long(m.Data2.AccessHash)
	x.StringBytes(m.Data2.Title)
	x.StringBytes(m.Data2.ShortName)
	x.Int(m.Data2.Count)
	x.Int(m.Data2.Hash)

	return x.buf
}

func (m *TLStickerSet) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m.Data2.Installed = true
	}
	if (flags & (1 << 3)) != 0 {
		m.Data2.Archived = true
	}
	if (flags & (1 << 4)) != 0 {
		m.Data2.Official = true
	}
	if (flags & (1 << 5)) != 0 {
		m.Data2.Masks = true
	}
	m.Data2.Id = x.Long()
	m.Data2.AccessHash = x.Long()
	m.Data2.Title = x.StringBytes()
	m.Data2.ShortName = x.StringBytes()
	m.Data2.Count = x.Int()
	m.Data2.Hash = x.Int()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Messages_FeaturedStickers <--
//  + TL_MessagesFeaturedStickersNotModified
//  + TL_MessagesFeaturedStickers
//

func (m *Messages_FeaturedStickers) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_messages_featuredStickersNotModified:
		t := m.To_MessagesFeaturedStickersNotModified()
		return t.Encode()
	case TLConstructor_CRC32_messages_featuredStickers:
		t := m.To_MessagesFeaturedStickers()
		return t.Encode()

	default:
		return []byte{}
	}
}

// messages.featuredStickersNotModified#4ede3cf = messages.FeaturedStickers;
func (m *Messages_FeaturedStickers) To_MessagesFeaturedStickersNotModified() *TLMessagesFeaturedStickersNotModified {
	return &TLMessagesFeaturedStickersNotModified{
		Data2: m.Data2,
	}
}

// messages.featuredStickers#f89d88e5 hash:int sets:Vector<StickerSetCovered> unread:Vector<long> = messages.FeaturedStickers;
func (m *Messages_FeaturedStickers) To_MessagesFeaturedStickers() *TLMessagesFeaturedStickers {
	return &TLMessagesFeaturedStickers{
		Data2: m.Data2,
	}
}

// messages.featuredStickersNotModified#4ede3cf = messages.FeaturedStickers;
func (m *TLMessagesFeaturedStickersNotModified) To_Messages_FeaturedStickers() *Messages_FeaturedStickers {
	return &Messages_FeaturedStickers{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessagesFeaturedStickersNotModified) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_featuredStickersNotModified))

	return x.buf
}

func (m *TLMessagesFeaturedStickersNotModified) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// messages.featuredStickers#f89d88e5 hash:int sets:Vector<StickerSetCovered> unread:Vector<long> = messages.FeaturedStickers;
func (m *TLMessagesFeaturedStickers) To_Messages_FeaturedStickers() *Messages_FeaturedStickers {
	return &Messages_FeaturedStickers{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessagesFeaturedStickers) SetHash(v int32) { m.Data2.Hash = v }
func (m *TLMessagesFeaturedStickers) GetHash() int32  { return m.Data2.Hash }

func (m *TLMessagesFeaturedStickers) SetSets(v []*StickerSetCovered) { m.Data2.Sets = v }
func (m *TLMessagesFeaturedStickers) GetSets() []*StickerSetCovered  { return m.Data2.Sets }

func (m *TLMessagesFeaturedStickers) SetUnread(v []int64) { m.Data2.Unread = v }
func (m *TLMessagesFeaturedStickers) GetUnread() []int64  { return m.Data2.Unread }

func (m *TLMessagesFeaturedStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_featuredStickers))

	x.Int(m.Data2.Hash)

	return x.buf
}

func (m *TLMessagesFeaturedStickers) Decode(dbuf *DecodeBuf) error {
	m.Data2.Hash = x.Int()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// InputPeer <--
//  + TL_InputPeerEmpty
//  + TL_InputPeerSelf
//  + TL_InputPeerChat
//  + TL_InputPeerUser
//  + TL_InputPeerChannel
//

func (m *InputPeer) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_inputPeerEmpty:
		t := m.To_InputPeerEmpty()
		return t.Encode()
	case TLConstructor_CRC32_inputPeerSelf:
		t := m.To_InputPeerSelf()
		return t.Encode()
	case TLConstructor_CRC32_inputPeerChat:
		t := m.To_InputPeerChat()
		return t.Encode()
	case TLConstructor_CRC32_inputPeerUser:
		t := m.To_InputPeerUser()
		return t.Encode()
	case TLConstructor_CRC32_inputPeerChannel:
		t := m.To_InputPeerChannel()
		return t.Encode()

	default:
		return []byte{}
	}
}

// inputPeerEmpty#7f3b18ea = InputPeer;
func (m *InputPeer) To_InputPeerEmpty() *TLInputPeerEmpty {
	return &TLInputPeerEmpty{
		Data2: m.Data2,
	}
}

// inputPeerSelf#7da07ec9 = InputPeer;
func (m *InputPeer) To_InputPeerSelf() *TLInputPeerSelf {
	return &TLInputPeerSelf{
		Data2: m.Data2,
	}
}

// inputPeerChat#179be863 chat_id:int = InputPeer;
func (m *InputPeer) To_InputPeerChat() *TLInputPeerChat {
	return &TLInputPeerChat{
		Data2: m.Data2,
	}
}

// inputPeerUser#7b8e7de6 user_id:int access_hash:long = InputPeer;
func (m *InputPeer) To_InputPeerUser() *TLInputPeerUser {
	return &TLInputPeerUser{
		Data2: m.Data2,
	}
}

// inputPeerChannel#20adaef8 channel_id:int access_hash:long = InputPeer;
func (m *InputPeer) To_InputPeerChannel() *TLInputPeerChannel {
	return &TLInputPeerChannel{
		Data2: m.Data2,
	}
}

// inputPeerEmpty#7f3b18ea = InputPeer;
func (m *TLInputPeerEmpty) To_InputPeer() *InputPeer {
	return &InputPeer{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputPeerEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputPeerEmpty))

	return x.buf
}

func (m *TLInputPeerEmpty) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// inputPeerSelf#7da07ec9 = InputPeer;
func (m *TLInputPeerSelf) To_InputPeer() *InputPeer {
	return &InputPeer{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputPeerSelf) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputPeerSelf))

	return x.buf
}

func (m *TLInputPeerSelf) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// inputPeerChat#179be863 chat_id:int = InputPeer;
func (m *TLInputPeerChat) To_InputPeer() *InputPeer {
	return &InputPeer{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputPeerChat) SetChatId(v int32) { m.Data2.ChatId = v }
func (m *TLInputPeerChat) GetChatId() int32  { return m.Data2.ChatId }

func (m *TLInputPeerChat) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputPeerChat))

	x.Int(m.Data2.ChatId)

	return x.buf
}

func (m *TLInputPeerChat) Decode(dbuf *DecodeBuf) error {
	m.Data2.ChatId = x.Int()

	return dbuf.err
}

// inputPeerUser#7b8e7de6 user_id:int access_hash:long = InputPeer;
func (m *TLInputPeerUser) To_InputPeer() *InputPeer {
	return &InputPeer{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputPeerUser) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLInputPeerUser) GetUserId() int32  { return m.Data2.UserId }

func (m *TLInputPeerUser) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLInputPeerUser) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLInputPeerUser) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputPeerUser))

	x.Int(m.Data2.UserId)
	x.Long(m.Data2.AccessHash)

	return x.buf
}

func (m *TLInputPeerUser) Decode(dbuf *DecodeBuf) error {
	m.Data2.UserId = x.Int()
	m.Data2.AccessHash = x.Long()

	return dbuf.err
}

// inputPeerChannel#20adaef8 channel_id:int access_hash:long = InputPeer;
func (m *TLInputPeerChannel) To_InputPeer() *InputPeer {
	return &InputPeer{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputPeerChannel) SetChannelId(v int32) { m.Data2.ChannelId = v }
func (m *TLInputPeerChannel) GetChannelId() int32  { return m.Data2.ChannelId }

func (m *TLInputPeerChannel) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLInputPeerChannel) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLInputPeerChannel) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputPeerChannel))

	x.Int(m.Data2.ChannelId)
	x.Long(m.Data2.AccessHash)

	return x.buf
}

func (m *TLInputPeerChannel) Decode(dbuf *DecodeBuf) error {
	m.Data2.ChannelId = x.Int()
	m.Data2.AccessHash = x.Long()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// InputContact <--
//  + TL_InputPhoneContact
//

func (m *InputContact) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_inputPhoneContact:
		t := m.To_InputPhoneContact()
		return t.Encode()

	default:
		return []byte{}
	}
}

// inputPhoneContact#f392b7f4 client_id:long phone:string first_name:string last_name:string = InputContact;
func (m *InputContact) To_InputPhoneContact() *TLInputPhoneContact {
	return &TLInputPhoneContact{
		Data2: m.Data2,
	}
}

// inputPhoneContact#f392b7f4 client_id:long phone:string first_name:string last_name:string = InputContact;
func (m *TLInputPhoneContact) To_InputContact() *InputContact {
	return &InputContact{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputPhoneContact) SetClientId(v int64) { m.Data2.ClientId = v }
func (m *TLInputPhoneContact) GetClientId() int64  { return m.Data2.ClientId }

func (m *TLInputPhoneContact) SetPhone(v string) { m.Data2.Phone = v }
func (m *TLInputPhoneContact) GetPhone() string  { return m.Data2.Phone }

func (m *TLInputPhoneContact) SetFirstName(v string) { m.Data2.FirstName = v }
func (m *TLInputPhoneContact) GetFirstName() string  { return m.Data2.FirstName }

func (m *TLInputPhoneContact) SetLastName(v string) { m.Data2.LastName = v }
func (m *TLInputPhoneContact) GetLastName() string  { return m.Data2.LastName }

func (m *TLInputPhoneContact) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputPhoneContact))

	x.Long(m.Data2.ClientId)
	x.StringBytes(m.Data2.Phone)
	x.StringBytes(m.Data2.FirstName)
	x.StringBytes(m.Data2.LastName)

	return x.buf
}

func (m *TLInputPhoneContact) Decode(dbuf *DecodeBuf) error {
	m.Data2.ClientId = x.Long()
	m.Data2.Phone = x.StringBytes()
	m.Data2.FirstName = x.StringBytes()
	m.Data2.LastName = x.StringBytes()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// InputPhoto <--
//  + TL_InputPhotoEmpty
//  + TL_InputPhoto
//

func (m *InputPhoto) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_inputPhotoEmpty:
		t := m.To_InputPhotoEmpty()
		return t.Encode()
	case TLConstructor_CRC32_inputPhoto:
		t := m.To_InputPhoto()
		return t.Encode()

	default:
		return []byte{}
	}
}

// inputPhotoEmpty#1cd7bf0d = InputPhoto;
func (m *InputPhoto) To_InputPhotoEmpty() *TLInputPhotoEmpty {
	return &TLInputPhotoEmpty{
		Data2: m.Data2,
	}
}

// inputPhoto#fb95c6c4 id:long access_hash:long = InputPhoto;
func (m *InputPhoto) To_InputPhoto() *TLInputPhoto {
	return &TLInputPhoto{
		Data2: m.Data2,
	}
}

// inputPhotoEmpty#1cd7bf0d = InputPhoto;
func (m *TLInputPhotoEmpty) To_InputPhoto() *InputPhoto {
	return &InputPhoto{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputPhotoEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputPhotoEmpty))

	return x.buf
}

func (m *TLInputPhotoEmpty) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// inputPhoto#fb95c6c4 id:long access_hash:long = InputPhoto;
func (m *TLInputPhoto) To_InputPhoto() *InputPhoto {
	return &InputPhoto{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputPhoto) SetId(v int64) { m.Data2.Id = v }
func (m *TLInputPhoto) GetId() int64  { return m.Data2.Id }

func (m *TLInputPhoto) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLInputPhoto) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLInputPhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputPhoto))

	x.Long(m.Data2.Id)
	x.Long(m.Data2.AccessHash)

	return x.buf
}

func (m *TLInputPhoto) Decode(dbuf *DecodeBuf) error {
	m.Data2.Id = x.Long()
	m.Data2.AccessHash = x.Long()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Photos_Photo <--
//  + TL_PhotosPhoto
//

func (m *Photos_Photo) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_photos_photo:
		t := m.To_PhotosPhoto()
		return t.Encode()

	default:
		return []byte{}
	}
}

// photos.photo#20212ca8 photo:Photo users:Vector<User> = photos.Photo;
func (m *Photos_Photo) To_PhotosPhoto() *TLPhotosPhoto {
	return &TLPhotosPhoto{
		Data2: m.Data2,
	}
}

// photos.photo#20212ca8 photo:Photo users:Vector<User> = photos.Photo;
func (m *TLPhotosPhoto) To_Photos_Photo() *Photos_Photo {
	return &Photos_Photo{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPhotosPhoto) SetPhoto(v *Photo) { m.Data2.Photo = v }
func (m *TLPhotosPhoto) GetPhoto() *Photo  { return m.Data2.Photo }

func (m *TLPhotosPhoto) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLPhotosPhoto) GetUsers() []*User  { return m.Data2.Users }

func (m *TLPhotosPhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_photos_photo))

	x.Bytes(m.Data2Photo.Encode())

	return x.buf
}

func (m *TLPhotosPhoto) Decode(dbuf *DecodeBuf) error {
	m1 := &Photo{}
	m1.Decode(dbuf)
	m.SetPhoto(m1)

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// ShippingOption <--
//  + TL_ShippingOption
//

func (m *ShippingOption) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_shippingOption:
		t := m.To_ShippingOption()
		return t.Encode()

	default:
		return []byte{}
	}
}

// shippingOption#b6213cdf id:string title:string prices:Vector<LabeledPrice> = ShippingOption;
func (m *ShippingOption) To_ShippingOption() *TLShippingOption {
	return &TLShippingOption{
		Data2: m.Data2,
	}
}

// shippingOption#b6213cdf id:string title:string prices:Vector<LabeledPrice> = ShippingOption;
func (m *TLShippingOption) To_ShippingOption() *ShippingOption {
	return &ShippingOption{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLShippingOption) SetId(v string) { m.Data2.Id = v }
func (m *TLShippingOption) GetId() string  { return m.Data2.Id }

func (m *TLShippingOption) SetTitle(v string) { m.Data2.Title = v }
func (m *TLShippingOption) GetTitle() string  { return m.Data2.Title }

func (m *TLShippingOption) SetPrices(v []*LabeledPrice) { m.Data2.Prices = v }
func (m *TLShippingOption) GetPrices() []*LabeledPrice  { return m.Data2.Prices }

func (m *TLShippingOption) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_shippingOption))

	x.StringBytes(m.Data2.Id)
	x.StringBytes(m.Data2.Title)

	return x.buf
}

func (m *TLShippingOption) Decode(dbuf *DecodeBuf) error {
	m.Data2.Id = x.StringBytes()
	m.Data2.Title = x.StringBytes()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Error <--
//  + TL_Error
//

func (m *Error) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_error:
		t := m.To_Error()
		return t.Encode()

	default:
		return []byte{}
	}
}

// error#c4b9f9bb code:int text:string = Error;
func (m *Error) To_Error() *TLError {
	return &TLError{
		Data2: m.Data2,
	}
}

// error#c4b9f9bb code:int text:string = Error;
func (m *TLError) To_Error() *Error {
	return &Error{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLError) SetCode(v int32) { m.Data2.Code = v }
func (m *TLError) GetCode() int32  { return m.Data2.Code }

func (m *TLError) SetText(v string) { m.Data2.Text = v }
func (m *TLError) GetText() string  { return m.Data2.Text }

func (m *TLError) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_error))

	x.Int(m.Data2.Code)
	x.StringBytes(m.Data2.Text)

	return x.buf
}

func (m *TLError) Decode(dbuf *DecodeBuf) error {
	m.Data2.Code = x.Int()
	m.Data2.Text = x.StringBytes()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// MessagesFilter <--
//  + TL_InputMessagesFilterEmpty
//  + TL_InputMessagesFilterPhotos
//  + TL_InputMessagesFilterVideo
//  + TL_InputMessagesFilterPhotoVideo
//  + TL_InputMessagesFilterPhotoVideoDocuments
//  + TL_InputMessagesFilterDocument
//  + TL_InputMessagesFilterUrl
//  + TL_InputMessagesFilterGif
//  + TL_InputMessagesFilterVoice
//  + TL_InputMessagesFilterMusic
//  + TL_InputMessagesFilterChatPhotos
//  + TL_InputMessagesFilterPhoneCalls
//  + TL_InputMessagesFilterRoundVoice
//  + TL_InputMessagesFilterRoundVideo
//  + TL_InputMessagesFilterMyMentions
//

func (m *MessagesFilter) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_inputMessagesFilterEmpty:
		t := m.To_InputMessagesFilterEmpty()
		return t.Encode()
	case TLConstructor_CRC32_inputMessagesFilterPhotos:
		t := m.To_InputMessagesFilterPhotos()
		return t.Encode()
	case TLConstructor_CRC32_inputMessagesFilterVideo:
		t := m.To_InputMessagesFilterVideo()
		return t.Encode()
	case TLConstructor_CRC32_inputMessagesFilterPhotoVideo:
		t := m.To_InputMessagesFilterPhotoVideo()
		return t.Encode()
	case TLConstructor_CRC32_inputMessagesFilterPhotoVideoDocuments:
		t := m.To_InputMessagesFilterPhotoVideoDocuments()
		return t.Encode()
	case TLConstructor_CRC32_inputMessagesFilterDocument:
		t := m.To_InputMessagesFilterDocument()
		return t.Encode()
	case TLConstructor_CRC32_inputMessagesFilterUrl:
		t := m.To_InputMessagesFilterUrl()
		return t.Encode()
	case TLConstructor_CRC32_inputMessagesFilterGif:
		t := m.To_InputMessagesFilterGif()
		return t.Encode()
	case TLConstructor_CRC32_inputMessagesFilterVoice:
		t := m.To_InputMessagesFilterVoice()
		return t.Encode()
	case TLConstructor_CRC32_inputMessagesFilterMusic:
		t := m.To_InputMessagesFilterMusic()
		return t.Encode()
	case TLConstructor_CRC32_inputMessagesFilterChatPhotos:
		t := m.To_InputMessagesFilterChatPhotos()
		return t.Encode()
	case TLConstructor_CRC32_inputMessagesFilterPhoneCalls:
		t := m.To_InputMessagesFilterPhoneCalls()
		return t.Encode()
	case TLConstructor_CRC32_inputMessagesFilterRoundVoice:
		t := m.To_InputMessagesFilterRoundVoice()
		return t.Encode()
	case TLConstructor_CRC32_inputMessagesFilterRoundVideo:
		t := m.To_InputMessagesFilterRoundVideo()
		return t.Encode()
	case TLConstructor_CRC32_inputMessagesFilterMyMentions:
		t := m.To_InputMessagesFilterMyMentions()
		return t.Encode()

	default:
		return []byte{}
	}
}

// inputMessagesFilterEmpty#57e2f66c = MessagesFilter;
func (m *MessagesFilter) To_InputMessagesFilterEmpty() *TLInputMessagesFilterEmpty {
	return &TLInputMessagesFilterEmpty{
		Data2: m.Data2,
	}
}

// inputMessagesFilterPhotos#9609a51c = MessagesFilter;
func (m *MessagesFilter) To_InputMessagesFilterPhotos() *TLInputMessagesFilterPhotos {
	return &TLInputMessagesFilterPhotos{
		Data2: m.Data2,
	}
}

// inputMessagesFilterVideo#9fc00e65 = MessagesFilter;
func (m *MessagesFilter) To_InputMessagesFilterVideo() *TLInputMessagesFilterVideo {
	return &TLInputMessagesFilterVideo{
		Data2: m.Data2,
	}
}

// inputMessagesFilterPhotoVideo#56e9f0e4 = MessagesFilter;
func (m *MessagesFilter) To_InputMessagesFilterPhotoVideo() *TLInputMessagesFilterPhotoVideo {
	return &TLInputMessagesFilterPhotoVideo{
		Data2: m.Data2,
	}
}

// inputMessagesFilterPhotoVideoDocuments#d95e73bb = MessagesFilter;
func (m *MessagesFilter) To_InputMessagesFilterPhotoVideoDocuments() *TLInputMessagesFilterPhotoVideoDocuments {
	return &TLInputMessagesFilterPhotoVideoDocuments{
		Data2: m.Data2,
	}
}

// inputMessagesFilterDocument#9eddf188 = MessagesFilter;
func (m *MessagesFilter) To_InputMessagesFilterDocument() *TLInputMessagesFilterDocument {
	return &TLInputMessagesFilterDocument{
		Data2: m.Data2,
	}
}

// inputMessagesFilterUrl#7ef0dd87 = MessagesFilter;
func (m *MessagesFilter) To_InputMessagesFilterUrl() *TLInputMessagesFilterUrl {
	return &TLInputMessagesFilterUrl{
		Data2: m.Data2,
	}
}

// inputMessagesFilterGif#ffc86587 = MessagesFilter;
func (m *MessagesFilter) To_InputMessagesFilterGif() *TLInputMessagesFilterGif {
	return &TLInputMessagesFilterGif{
		Data2: m.Data2,
	}
}

// inputMessagesFilterVoice#50f5c392 = MessagesFilter;
func (m *MessagesFilter) To_InputMessagesFilterVoice() *TLInputMessagesFilterVoice {
	return &TLInputMessagesFilterVoice{
		Data2: m.Data2,
	}
}

// inputMessagesFilterMusic#3751b49e = MessagesFilter;
func (m *MessagesFilter) To_InputMessagesFilterMusic() *TLInputMessagesFilterMusic {
	return &TLInputMessagesFilterMusic{
		Data2: m.Data2,
	}
}

// inputMessagesFilterChatPhotos#3a20ecb8 = MessagesFilter;
func (m *MessagesFilter) To_InputMessagesFilterChatPhotos() *TLInputMessagesFilterChatPhotos {
	return &TLInputMessagesFilterChatPhotos{
		Data2: m.Data2,
	}
}

// inputMessagesFilterPhoneCalls#80c99768 flags:# missed:flags.0?true = MessagesFilter;
func (m *MessagesFilter) To_InputMessagesFilterPhoneCalls() *TLInputMessagesFilterPhoneCalls {
	return &TLInputMessagesFilterPhoneCalls{
		Data2: m.Data2,
	}
}

// inputMessagesFilterRoundVoice#7a7c17a4 = MessagesFilter;
func (m *MessagesFilter) To_InputMessagesFilterRoundVoice() *TLInputMessagesFilterRoundVoice {
	return &TLInputMessagesFilterRoundVoice{
		Data2: m.Data2,
	}
}

// inputMessagesFilterRoundVideo#b549da53 = MessagesFilter;
func (m *MessagesFilter) To_InputMessagesFilterRoundVideo() *TLInputMessagesFilterRoundVideo {
	return &TLInputMessagesFilterRoundVideo{
		Data2: m.Data2,
	}
}

// inputMessagesFilterMyMentions#c1f8e69a = MessagesFilter;
func (m *MessagesFilter) To_InputMessagesFilterMyMentions() *TLInputMessagesFilterMyMentions {
	return &TLInputMessagesFilterMyMentions{
		Data2: m.Data2,
	}
}

// inputMessagesFilterEmpty#57e2f66c = MessagesFilter;
func (m *TLInputMessagesFilterEmpty) To_MessagesFilter() *MessagesFilter {
	return &MessagesFilter{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputMessagesFilterEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputMessagesFilterEmpty))

	return x.buf
}

func (m *TLInputMessagesFilterEmpty) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// inputMessagesFilterPhotos#9609a51c = MessagesFilter;
func (m *TLInputMessagesFilterPhotos) To_MessagesFilter() *MessagesFilter {
	return &MessagesFilter{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputMessagesFilterPhotos) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputMessagesFilterPhotos))

	return x.buf
}

func (m *TLInputMessagesFilterPhotos) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// inputMessagesFilterVideo#9fc00e65 = MessagesFilter;
func (m *TLInputMessagesFilterVideo) To_MessagesFilter() *MessagesFilter {
	return &MessagesFilter{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputMessagesFilterVideo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputMessagesFilterVideo))

	return x.buf
}

func (m *TLInputMessagesFilterVideo) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// inputMessagesFilterPhotoVideo#56e9f0e4 = MessagesFilter;
func (m *TLInputMessagesFilterPhotoVideo) To_MessagesFilter() *MessagesFilter {
	return &MessagesFilter{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputMessagesFilterPhotoVideo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputMessagesFilterPhotoVideo))

	return x.buf
}

func (m *TLInputMessagesFilterPhotoVideo) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// inputMessagesFilterPhotoVideoDocuments#d95e73bb = MessagesFilter;
func (m *TLInputMessagesFilterPhotoVideoDocuments) To_MessagesFilter() *MessagesFilter {
	return &MessagesFilter{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputMessagesFilterPhotoVideoDocuments) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputMessagesFilterPhotoVideoDocuments))

	return x.buf
}

func (m *TLInputMessagesFilterPhotoVideoDocuments) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// inputMessagesFilterDocument#9eddf188 = MessagesFilter;
func (m *TLInputMessagesFilterDocument) To_MessagesFilter() *MessagesFilter {
	return &MessagesFilter{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputMessagesFilterDocument) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputMessagesFilterDocument))

	return x.buf
}

func (m *TLInputMessagesFilterDocument) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// inputMessagesFilterUrl#7ef0dd87 = MessagesFilter;
func (m *TLInputMessagesFilterUrl) To_MessagesFilter() *MessagesFilter {
	return &MessagesFilter{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputMessagesFilterUrl) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputMessagesFilterUrl))

	return x.buf
}

func (m *TLInputMessagesFilterUrl) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// inputMessagesFilterGif#ffc86587 = MessagesFilter;
func (m *TLInputMessagesFilterGif) To_MessagesFilter() *MessagesFilter {
	return &MessagesFilter{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputMessagesFilterGif) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputMessagesFilterGif))

	return x.buf
}

func (m *TLInputMessagesFilterGif) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// inputMessagesFilterVoice#50f5c392 = MessagesFilter;
func (m *TLInputMessagesFilterVoice) To_MessagesFilter() *MessagesFilter {
	return &MessagesFilter{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputMessagesFilterVoice) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputMessagesFilterVoice))

	return x.buf
}

func (m *TLInputMessagesFilterVoice) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// inputMessagesFilterMusic#3751b49e = MessagesFilter;
func (m *TLInputMessagesFilterMusic) To_MessagesFilter() *MessagesFilter {
	return &MessagesFilter{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputMessagesFilterMusic) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputMessagesFilterMusic))

	return x.buf
}

func (m *TLInputMessagesFilterMusic) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// inputMessagesFilterChatPhotos#3a20ecb8 = MessagesFilter;
func (m *TLInputMessagesFilterChatPhotos) To_MessagesFilter() *MessagesFilter {
	return &MessagesFilter{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputMessagesFilterChatPhotos) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputMessagesFilterChatPhotos))

	return x.buf
}

func (m *TLInputMessagesFilterChatPhotos) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// inputMessagesFilterPhoneCalls#80c99768 flags:# missed:flags.0?true = MessagesFilter;
func (m *TLInputMessagesFilterPhoneCalls) To_MessagesFilter() *MessagesFilter {
	return &MessagesFilter{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputMessagesFilterPhoneCalls) SetMissed(v bool) { m.Data2.Missed = v }
func (m *TLInputMessagesFilterPhoneCalls) GetMissed() bool  { return m.Data2.Missed }

func (m *TLInputMessagesFilterPhoneCalls) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputMessagesFilterPhoneCalls))

	// flags
	var flags uint32 = 0
	if m.GetMissed() == true {
		flags |= 1 << 1
	}
	x.UInt(flags)

	return x.buf
}

func (m *TLInputMessagesFilterPhoneCalls) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m.Data2.Missed = true
	}

	return dbuf.err
}

// inputMessagesFilterRoundVoice#7a7c17a4 = MessagesFilter;
func (m *TLInputMessagesFilterRoundVoice) To_MessagesFilter() *MessagesFilter {
	return &MessagesFilter{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputMessagesFilterRoundVoice) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputMessagesFilterRoundVoice))

	return x.buf
}

func (m *TLInputMessagesFilterRoundVoice) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// inputMessagesFilterRoundVideo#b549da53 = MessagesFilter;
func (m *TLInputMessagesFilterRoundVideo) To_MessagesFilter() *MessagesFilter {
	return &MessagesFilter{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputMessagesFilterRoundVideo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputMessagesFilterRoundVideo))

	return x.buf
}

func (m *TLInputMessagesFilterRoundVideo) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// inputMessagesFilterMyMentions#c1f8e69a = MessagesFilter;
func (m *TLInputMessagesFilterMyMentions) To_MessagesFilter() *MessagesFilter {
	return &MessagesFilter{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputMessagesFilterMyMentions) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputMessagesFilterMyMentions))

	return x.buf
}

func (m *TLInputMessagesFilterMyMentions) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Messages_AffectedMessages <--
//  + TL_MessagesAffectedMessages
//

func (m *Messages_AffectedMessages) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_messages_affectedMessages:
		t := m.To_MessagesAffectedMessages()
		return t.Encode()

	default:
		return []byte{}
	}
}

// messages.affectedMessages#84d19185 pts:int pts_count:int = messages.AffectedMessages;
func (m *Messages_AffectedMessages) To_MessagesAffectedMessages() *TLMessagesAffectedMessages {
	return &TLMessagesAffectedMessages{
		Data2: m.Data2,
	}
}

// messages.affectedMessages#84d19185 pts:int pts_count:int = messages.AffectedMessages;
func (m *TLMessagesAffectedMessages) To_Messages_AffectedMessages() *Messages_AffectedMessages {
	return &Messages_AffectedMessages{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessagesAffectedMessages) SetPts(v int32) { m.Data2.Pts = v }
func (m *TLMessagesAffectedMessages) GetPts() int32  { return m.Data2.Pts }

func (m *TLMessagesAffectedMessages) SetPtsCount(v int32) { m.Data2.PtsCount = v }
func (m *TLMessagesAffectedMessages) GetPtsCount() int32  { return m.Data2.PtsCount }

func (m *TLMessagesAffectedMessages) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_affectedMessages))

	x.Int(m.Data2.Pts)
	x.Int(m.Data2.PtsCount)

	return x.buf
}

func (m *TLMessagesAffectedMessages) Decode(dbuf *DecodeBuf) error {
	m.Data2.Pts = x.Int()
	m.Data2.PtsCount = x.Int()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// PageBlock <--
//  + TL_PageBlockUnsupported
//  + TL_PageBlockTitle
//  + TL_PageBlockSubtitle
//  + TL_PageBlockAuthorDate
//  + TL_PageBlockHeader
//  + TL_PageBlockSubheader
//  + TL_PageBlockParagraph
//  + TL_PageBlockPreformatted
//  + TL_PageBlockFooter
//  + TL_PageBlockDivider
//  + TL_PageBlockAnchor
//  + TL_PageBlockList
//  + TL_PageBlockBlockquote
//  + TL_PageBlockPullquote
//  + TL_PageBlockPhoto
//  + TL_PageBlockVideo
//  + TL_PageBlockCover
//  + TL_PageBlockEmbed
//  + TL_PageBlockEmbedPost
//  + TL_PageBlockCollage
//  + TL_PageBlockSlideshow
//  + TL_PageBlockChannel
//  + TL_PageBlockAudio
//

func (m *PageBlock) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_pageBlockUnsupported:
		t := m.To_PageBlockUnsupported()
		return t.Encode()
	case TLConstructor_CRC32_pageBlockTitle:
		t := m.To_PageBlockTitle()
		return t.Encode()
	case TLConstructor_CRC32_pageBlockSubtitle:
		t := m.To_PageBlockSubtitle()
		return t.Encode()
	case TLConstructor_CRC32_pageBlockAuthorDate:
		t := m.To_PageBlockAuthorDate()
		return t.Encode()
	case TLConstructor_CRC32_pageBlockHeader:
		t := m.To_PageBlockHeader()
		return t.Encode()
	case TLConstructor_CRC32_pageBlockSubheader:
		t := m.To_PageBlockSubheader()
		return t.Encode()
	case TLConstructor_CRC32_pageBlockParagraph:
		t := m.To_PageBlockParagraph()
		return t.Encode()
	case TLConstructor_CRC32_pageBlockPreformatted:
		t := m.To_PageBlockPreformatted()
		return t.Encode()
	case TLConstructor_CRC32_pageBlockFooter:
		t := m.To_PageBlockFooter()
		return t.Encode()
	case TLConstructor_CRC32_pageBlockDivider:
		t := m.To_PageBlockDivider()
		return t.Encode()
	case TLConstructor_CRC32_pageBlockAnchor:
		t := m.To_PageBlockAnchor()
		return t.Encode()
	case TLConstructor_CRC32_pageBlockList:
		t := m.To_PageBlockList()
		return t.Encode()
	case TLConstructor_CRC32_pageBlockBlockquote:
		t := m.To_PageBlockBlockquote()
		return t.Encode()
	case TLConstructor_CRC32_pageBlockPullquote:
		t := m.To_PageBlockPullquote()
		return t.Encode()
	case TLConstructor_CRC32_pageBlockPhoto:
		t := m.To_PageBlockPhoto()
		return t.Encode()
	case TLConstructor_CRC32_pageBlockVideo:
		t := m.To_PageBlockVideo()
		return t.Encode()
	case TLConstructor_CRC32_pageBlockCover:
		t := m.To_PageBlockCover()
		return t.Encode()
	case TLConstructor_CRC32_pageBlockEmbed:
		t := m.To_PageBlockEmbed()
		return t.Encode()
	case TLConstructor_CRC32_pageBlockEmbedPost:
		t := m.To_PageBlockEmbedPost()
		return t.Encode()
	case TLConstructor_CRC32_pageBlockCollage:
		t := m.To_PageBlockCollage()
		return t.Encode()
	case TLConstructor_CRC32_pageBlockSlideshow:
		t := m.To_PageBlockSlideshow()
		return t.Encode()
	case TLConstructor_CRC32_pageBlockChannel:
		t := m.To_PageBlockChannel()
		return t.Encode()
	case TLConstructor_CRC32_pageBlockAudio:
		t := m.To_PageBlockAudio()
		return t.Encode()

	default:
		return []byte{}
	}
}

// pageBlockUnsupported#13567e8a = PageBlock;
func (m *PageBlock) To_PageBlockUnsupported() *TLPageBlockUnsupported {
	return &TLPageBlockUnsupported{
		Data2: m.Data2,
	}
}

// pageBlockTitle#70abc3fd text:RichText = PageBlock;
func (m *PageBlock) To_PageBlockTitle() *TLPageBlockTitle {
	return &TLPageBlockTitle{
		Data2: m.Data2,
	}
}

// pageBlockSubtitle#8ffa9a1f text:RichText = PageBlock;
func (m *PageBlock) To_PageBlockSubtitle() *TLPageBlockSubtitle {
	return &TLPageBlockSubtitle{
		Data2: m.Data2,
	}
}

// pageBlockAuthorDate#baafe5e0 author:RichText published_date:int = PageBlock;
func (m *PageBlock) To_PageBlockAuthorDate() *TLPageBlockAuthorDate {
	return &TLPageBlockAuthorDate{
		Data2: m.Data2,
	}
}

// pageBlockHeader#bfd064ec text:RichText = PageBlock;
func (m *PageBlock) To_PageBlockHeader() *TLPageBlockHeader {
	return &TLPageBlockHeader{
		Data2: m.Data2,
	}
}

// pageBlockSubheader#f12bb6e1 text:RichText = PageBlock;
func (m *PageBlock) To_PageBlockSubheader() *TLPageBlockSubheader {
	return &TLPageBlockSubheader{
		Data2: m.Data2,
	}
}

// pageBlockParagraph#467a0766 text:RichText = PageBlock;
func (m *PageBlock) To_PageBlockParagraph() *TLPageBlockParagraph {
	return &TLPageBlockParagraph{
		Data2: m.Data2,
	}
}

// pageBlockPreformatted#c070d93e text:RichText language:string = PageBlock;
func (m *PageBlock) To_PageBlockPreformatted() *TLPageBlockPreformatted {
	return &TLPageBlockPreformatted{
		Data2: m.Data2,
	}
}

// pageBlockFooter#48870999 text:RichText = PageBlock;
func (m *PageBlock) To_PageBlockFooter() *TLPageBlockFooter {
	return &TLPageBlockFooter{
		Data2: m.Data2,
	}
}

// pageBlockDivider#db20b188 = PageBlock;
func (m *PageBlock) To_PageBlockDivider() *TLPageBlockDivider {
	return &TLPageBlockDivider{
		Data2: m.Data2,
	}
}

// pageBlockAnchor#ce0d37b0 name:string = PageBlock;
func (m *PageBlock) To_PageBlockAnchor() *TLPageBlockAnchor {
	return &TLPageBlockAnchor{
		Data2: m.Data2,
	}
}

// pageBlockList#3a58c7f4 ordered:Bool items:Vector<RichText> = PageBlock;
func (m *PageBlock) To_PageBlockList() *TLPageBlockList {
	return &TLPageBlockList{
		Data2: m.Data2,
	}
}

// pageBlockBlockquote#263d7c26 text:RichText caption:RichText = PageBlock;
func (m *PageBlock) To_PageBlockBlockquote() *TLPageBlockBlockquote {
	return &TLPageBlockBlockquote{
		Data2: m.Data2,
	}
}

// pageBlockPullquote#4f4456d3 text:RichText caption:RichText = PageBlock;
func (m *PageBlock) To_PageBlockPullquote() *TLPageBlockPullquote {
	return &TLPageBlockPullquote{
		Data2: m.Data2,
	}
}

// pageBlockPhoto#e9c69982 photo_id:long caption:RichText = PageBlock;
func (m *PageBlock) To_PageBlockPhoto() *TLPageBlockPhoto {
	return &TLPageBlockPhoto{
		Data2: m.Data2,
	}
}

// pageBlockVideo#d9d71866 flags:# autoplay:flags.0?true loop:flags.1?true video_id:long caption:RichText = PageBlock;
func (m *PageBlock) To_PageBlockVideo() *TLPageBlockVideo {
	return &TLPageBlockVideo{
		Data2: m.Data2,
	}
}

// pageBlockCover#39f23300 cover:PageBlock = PageBlock;
func (m *PageBlock) To_PageBlockCover() *TLPageBlockCover {
	return &TLPageBlockCover{
		Data2: m.Data2,
	}
}

// pageBlockEmbed#cde200d1 flags:# full_width:flags.0?true allow_scrolling:flags.3?true url:flags.1?string html:flags.2?string poster_photo_id:flags.4?long w:int h:int caption:RichText = PageBlock;
func (m *PageBlock) To_PageBlockEmbed() *TLPageBlockEmbed {
	return &TLPageBlockEmbed{
		Data2: m.Data2,
	}
}

// pageBlockEmbedPost#292c7be9 url:string webpage_id:long author_photo_id:long author:string date:int blocks:Vector<PageBlock> caption:RichText = PageBlock;
func (m *PageBlock) To_PageBlockEmbedPost() *TLPageBlockEmbedPost {
	return &TLPageBlockEmbedPost{
		Data2: m.Data2,
	}
}

// pageBlockCollage#8b31c4f items:Vector<PageBlock> caption:RichText = PageBlock;
func (m *PageBlock) To_PageBlockCollage() *TLPageBlockCollage {
	return &TLPageBlockCollage{
		Data2: m.Data2,
	}
}

// pageBlockSlideshow#130c8963 items:Vector<PageBlock> caption:RichText = PageBlock;
func (m *PageBlock) To_PageBlockSlideshow() *TLPageBlockSlideshow {
	return &TLPageBlockSlideshow{
		Data2: m.Data2,
	}
}

// pageBlockChannel#ef1751b5 channel:Chat = PageBlock;
func (m *PageBlock) To_PageBlockChannel() *TLPageBlockChannel {
	return &TLPageBlockChannel{
		Data2: m.Data2,
	}
}

// pageBlockAudio#31b81a7f audio_id:long caption:RichText = PageBlock;
func (m *PageBlock) To_PageBlockAudio() *TLPageBlockAudio {
	return &TLPageBlockAudio{
		Data2: m.Data2,
	}
}

// pageBlockUnsupported#13567e8a = PageBlock;
func (m *TLPageBlockUnsupported) To_PageBlock() *PageBlock {
	return &PageBlock{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPageBlockUnsupported) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_pageBlockUnsupported))

	return x.buf
}

func (m *TLPageBlockUnsupported) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// pageBlockTitle#70abc3fd text:RichText = PageBlock;
func (m *TLPageBlockTitle) To_PageBlock() *PageBlock {
	return &PageBlock{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPageBlockTitle) SetText(v *RichText) { m.Data2.Text = v }
func (m *TLPageBlockTitle) GetText() *RichText  { return m.Data2.Text }

func (m *TLPageBlockTitle) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_pageBlockTitle))

	x.Bytes(m.Data2Text.Encode())

	return x.buf
}

func (m *TLPageBlockTitle) Decode(dbuf *DecodeBuf) error {
	m1 := &Text{}
	m1.Decode(dbuf)
	m.SetText(m1)

	return dbuf.err
}

// pageBlockSubtitle#8ffa9a1f text:RichText = PageBlock;
func (m *TLPageBlockSubtitle) To_PageBlock() *PageBlock {
	return &PageBlock{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPageBlockSubtitle) SetText(v *RichText) { m.Data2.Text = v }
func (m *TLPageBlockSubtitle) GetText() *RichText  { return m.Data2.Text }

func (m *TLPageBlockSubtitle) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_pageBlockSubtitle))

	x.Bytes(m.Data2Text.Encode())

	return x.buf
}

func (m *TLPageBlockSubtitle) Decode(dbuf *DecodeBuf) error {
	m1 := &Text{}
	m1.Decode(dbuf)
	m.SetText(m1)

	return dbuf.err
}

// pageBlockAuthorDate#baafe5e0 author:RichText published_date:int = PageBlock;
func (m *TLPageBlockAuthorDate) To_PageBlock() *PageBlock {
	return &PageBlock{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPageBlockAuthorDate) SetAuthor(v *RichText) { m.Data2.Author_2 = v }
func (m *TLPageBlockAuthorDate) GetAuthor() *RichText  { return m.Data2.Author_2 }

func (m *TLPageBlockAuthorDate) SetPublishedDate(v int32) { m.Data2.PublishedDate = v }
func (m *TLPageBlockAuthorDate) GetPublishedDate() int32  { return m.Data2.PublishedDate }

func (m *TLPageBlockAuthorDate) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_pageBlockAuthorDate))

	x.Bytes(m.Data2Author.Encode())
	x.Int(m.Data2.PublishedDate)

	return x.buf
}

func (m *TLPageBlockAuthorDate) Decode(dbuf *DecodeBuf) error {
	m1 := &Author{}
	m1.Decode(dbuf)
	m.SetAuthor(m1)
	m.Data2.PublishedDate = x.Int()

	return dbuf.err
}

// pageBlockHeader#bfd064ec text:RichText = PageBlock;
func (m *TLPageBlockHeader) To_PageBlock() *PageBlock {
	return &PageBlock{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPageBlockHeader) SetText(v *RichText) { m.Data2.Text = v }
func (m *TLPageBlockHeader) GetText() *RichText  { return m.Data2.Text }

func (m *TLPageBlockHeader) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_pageBlockHeader))

	x.Bytes(m.Data2Text.Encode())

	return x.buf
}

func (m *TLPageBlockHeader) Decode(dbuf *DecodeBuf) error {
	m1 := &Text{}
	m1.Decode(dbuf)
	m.SetText(m1)

	return dbuf.err
}

// pageBlockSubheader#f12bb6e1 text:RichText = PageBlock;
func (m *TLPageBlockSubheader) To_PageBlock() *PageBlock {
	return &PageBlock{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPageBlockSubheader) SetText(v *RichText) { m.Data2.Text = v }
func (m *TLPageBlockSubheader) GetText() *RichText  { return m.Data2.Text }

func (m *TLPageBlockSubheader) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_pageBlockSubheader))

	x.Bytes(m.Data2Text.Encode())

	return x.buf
}

func (m *TLPageBlockSubheader) Decode(dbuf *DecodeBuf) error {
	m1 := &Text{}
	m1.Decode(dbuf)
	m.SetText(m1)

	return dbuf.err
}

// pageBlockParagraph#467a0766 text:RichText = PageBlock;
func (m *TLPageBlockParagraph) To_PageBlock() *PageBlock {
	return &PageBlock{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPageBlockParagraph) SetText(v *RichText) { m.Data2.Text = v }
func (m *TLPageBlockParagraph) GetText() *RichText  { return m.Data2.Text }

func (m *TLPageBlockParagraph) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_pageBlockParagraph))

	x.Bytes(m.Data2Text.Encode())

	return x.buf
}

func (m *TLPageBlockParagraph) Decode(dbuf *DecodeBuf) error {
	m1 := &Text{}
	m1.Decode(dbuf)
	m.SetText(m1)

	return dbuf.err
}

// pageBlockPreformatted#c070d93e text:RichText language:string = PageBlock;
func (m *TLPageBlockPreformatted) To_PageBlock() *PageBlock {
	return &PageBlock{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPageBlockPreformatted) SetText(v *RichText) { m.Data2.Text = v }
func (m *TLPageBlockPreformatted) GetText() *RichText  { return m.Data2.Text }

func (m *TLPageBlockPreformatted) SetLanguage(v string) { m.Data2.Language = v }
func (m *TLPageBlockPreformatted) GetLanguage() string  { return m.Data2.Language }

func (m *TLPageBlockPreformatted) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_pageBlockPreformatted))

	x.Bytes(m.Data2Text.Encode())
	x.StringBytes(m.Data2.Language)

	return x.buf
}

func (m *TLPageBlockPreformatted) Decode(dbuf *DecodeBuf) error {
	m1 := &Text{}
	m1.Decode(dbuf)
	m.SetText(m1)
	m.Data2.Language = x.StringBytes()

	return dbuf.err
}

// pageBlockFooter#48870999 text:RichText = PageBlock;
func (m *TLPageBlockFooter) To_PageBlock() *PageBlock {
	return &PageBlock{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPageBlockFooter) SetText(v *RichText) { m.Data2.Text = v }
func (m *TLPageBlockFooter) GetText() *RichText  { return m.Data2.Text }

func (m *TLPageBlockFooter) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_pageBlockFooter))

	x.Bytes(m.Data2Text.Encode())

	return x.buf
}

func (m *TLPageBlockFooter) Decode(dbuf *DecodeBuf) error {
	m1 := &Text{}
	m1.Decode(dbuf)
	m.SetText(m1)

	return dbuf.err
}

// pageBlockDivider#db20b188 = PageBlock;
func (m *TLPageBlockDivider) To_PageBlock() *PageBlock {
	return &PageBlock{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPageBlockDivider) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_pageBlockDivider))

	return x.buf
}

func (m *TLPageBlockDivider) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// pageBlockAnchor#ce0d37b0 name:string = PageBlock;
func (m *TLPageBlockAnchor) To_PageBlock() *PageBlock {
	return &PageBlock{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPageBlockAnchor) SetName(v string) { m.Data2.Name = v }
func (m *TLPageBlockAnchor) GetName() string  { return m.Data2.Name }

func (m *TLPageBlockAnchor) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_pageBlockAnchor))

	x.StringBytes(m.Data2.Name)

	return x.buf
}

func (m *TLPageBlockAnchor) Decode(dbuf *DecodeBuf) error {
	m.Data2.Name = x.StringBytes()

	return dbuf.err
}

// pageBlockList#3a58c7f4 ordered:Bool items:Vector<RichText> = PageBlock;
func (m *TLPageBlockList) To_PageBlock() *PageBlock {
	return &PageBlock{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPageBlockList) SetOrdered(v *Bool) { m.Data2.Ordered = v }
func (m *TLPageBlockList) GetOrdered() *Bool  { return m.Data2.Ordered }

func (m *TLPageBlockList) SetItems(v []*RichText) { m.Data2.Items_7 = v }
func (m *TLPageBlockList) GetItems() []*RichText  { return m.Data2.Items_7 }

func (m *TLPageBlockList) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_pageBlockList))

	x.Bytes(m.Data2Ordered.Encode())

	return x.buf
}

func (m *TLPageBlockList) Decode(dbuf *DecodeBuf) error {
	m1 := &Ordered{}
	m1.Decode(dbuf)
	m.SetOrdered(m1)

	return dbuf.err
}

// pageBlockBlockquote#263d7c26 text:RichText caption:RichText = PageBlock;
func (m *TLPageBlockBlockquote) To_PageBlock() *PageBlock {
	return &PageBlock{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPageBlockBlockquote) SetText(v *RichText) { m.Data2.Text = v }
func (m *TLPageBlockBlockquote) GetText() *RichText  { return m.Data2.Text }

func (m *TLPageBlockBlockquote) SetCaption(v *RichText) { m.Data2.Caption = v }
func (m *TLPageBlockBlockquote) GetCaption() *RichText  { return m.Data2.Caption }

func (m *TLPageBlockBlockquote) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_pageBlockBlockquote))

	x.Bytes(m.Data2Text.Encode())
	x.Bytes(m.Data2Caption.Encode())

	return x.buf
}

func (m *TLPageBlockBlockquote) Decode(dbuf *DecodeBuf) error {
	m1 := &Text{}
	m1.Decode(dbuf)
	m.SetText(m1)
	m2 := &Caption{}
	m2.Decode(dbuf)
	m.SetCaption(m2)

	return dbuf.err
}

// pageBlockPullquote#4f4456d3 text:RichText caption:RichText = PageBlock;
func (m *TLPageBlockPullquote) To_PageBlock() *PageBlock {
	return &PageBlock{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPageBlockPullquote) SetText(v *RichText) { m.Data2.Text = v }
func (m *TLPageBlockPullquote) GetText() *RichText  { return m.Data2.Text }

func (m *TLPageBlockPullquote) SetCaption(v *RichText) { m.Data2.Caption = v }
func (m *TLPageBlockPullquote) GetCaption() *RichText  { return m.Data2.Caption }

func (m *TLPageBlockPullquote) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_pageBlockPullquote))

	x.Bytes(m.Data2Text.Encode())
	x.Bytes(m.Data2Caption.Encode())

	return x.buf
}

func (m *TLPageBlockPullquote) Decode(dbuf *DecodeBuf) error {
	m1 := &Text{}
	m1.Decode(dbuf)
	m.SetText(m1)
	m2 := &Caption{}
	m2.Decode(dbuf)
	m.SetCaption(m2)

	return dbuf.err
}

// pageBlockPhoto#e9c69982 photo_id:long caption:RichText = PageBlock;
func (m *TLPageBlockPhoto) To_PageBlock() *PageBlock {
	return &PageBlock{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPageBlockPhoto) SetPhotoId(v int64) { m.Data2.PhotoId = v }
func (m *TLPageBlockPhoto) GetPhotoId() int64  { return m.Data2.PhotoId }

func (m *TLPageBlockPhoto) SetCaption(v *RichText) { m.Data2.Caption = v }
func (m *TLPageBlockPhoto) GetCaption() *RichText  { return m.Data2.Caption }

func (m *TLPageBlockPhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_pageBlockPhoto))

	x.Long(m.Data2.PhotoId)
	x.Bytes(m.Data2Caption.Encode())

	return x.buf
}

func (m *TLPageBlockPhoto) Decode(dbuf *DecodeBuf) error {
	m.Data2.PhotoId = x.Long()
	m2 := &Caption{}
	m2.Decode(dbuf)
	m.SetCaption(m2)

	return dbuf.err
}

// pageBlockVideo#d9d71866 flags:# autoplay:flags.0?true loop:flags.1?true video_id:long caption:RichText = PageBlock;
func (m *TLPageBlockVideo) To_PageBlock() *PageBlock {
	return &PageBlock{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPageBlockVideo) SetAutoplay(v bool) { m.Data2.Autoplay = v }
func (m *TLPageBlockVideo) GetAutoplay() bool  { return m.Data2.Autoplay }

func (m *TLPageBlockVideo) SetLoop(v bool) { m.Data2.Loop = v }
func (m *TLPageBlockVideo) GetLoop() bool  { return m.Data2.Loop }

func (m *TLPageBlockVideo) SetVideoId(v int64) { m.Data2.VideoId = v }
func (m *TLPageBlockVideo) GetVideoId() int64  { return m.Data2.VideoId }

func (m *TLPageBlockVideo) SetCaption(v *RichText) { m.Data2.Caption = v }
func (m *TLPageBlockVideo) GetCaption() *RichText  { return m.Data2.Caption }

func (m *TLPageBlockVideo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_pageBlockVideo))

	// flags
	var flags uint32 = 0
	if m.GetAutoplay() == true {
		flags |= 1 << 1
	}
	if m.GetLoop() == true {
		flags |= 1 << 2
	}
	x.UInt(flags)

	x.Long(m.Data2.VideoId)
	x.Bytes(m.Data2Caption.Encode())

	return x.buf
}

func (m *TLPageBlockVideo) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m.Data2.Autoplay = true
	}
	if (flags & (1 << 3)) != 0 {
		m.Data2.Loop = true
	}
	m.Data2.VideoId = x.Long()
	m5 := &Caption{}
	m5.Decode(dbuf)
	m.SetCaption(m5)

	return dbuf.err
}

// pageBlockCover#39f23300 cover:PageBlock = PageBlock;
func (m *TLPageBlockCover) To_PageBlock() *PageBlock {
	return &PageBlock{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPageBlockCover) SetCover(v *PageBlock) { m.Data2.Cover = v }
func (m *TLPageBlockCover) GetCover() *PageBlock  { return m.Data2.Cover }

func (m *TLPageBlockCover) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_pageBlockCover))

	x.Bytes(m.Data2Cover.Encode())

	return x.buf
}

func (m *TLPageBlockCover) Decode(dbuf *DecodeBuf) error {
	m1 := &Cover{}
	m1.Decode(dbuf)
	m.SetCover(m1)

	return dbuf.err
}

// pageBlockEmbed#cde200d1 flags:# full_width:flags.0?true allow_scrolling:flags.3?true url:flags.1?string html:flags.2?string poster_photo_id:flags.4?long w:int h:int caption:RichText = PageBlock;
func (m *TLPageBlockEmbed) To_PageBlock() *PageBlock {
	return &PageBlock{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPageBlockEmbed) SetFullWidth(v bool) { m.Data2.FullWidth = v }
func (m *TLPageBlockEmbed) GetFullWidth() bool  { return m.Data2.FullWidth }

func (m *TLPageBlockEmbed) SetAllowScrolling(v bool) { m.Data2.AllowScrolling = v }
func (m *TLPageBlockEmbed) GetAllowScrolling() bool  { return m.Data2.AllowScrolling }

func (m *TLPageBlockEmbed) SetUrl(v string) { m.Data2.Url = v }
func (m *TLPageBlockEmbed) GetUrl() string  { return m.Data2.Url }

func (m *TLPageBlockEmbed) SetHtml(v string) { m.Data2.Html = v }
func (m *TLPageBlockEmbed) GetHtml() string  { return m.Data2.Html }

func (m *TLPageBlockEmbed) SetPosterPhotoId(v int64) { m.Data2.PosterPhotoId = v }
func (m *TLPageBlockEmbed) GetPosterPhotoId() int64  { return m.Data2.PosterPhotoId }

func (m *TLPageBlockEmbed) SetW(v int32) { m.Data2.W = v }
func (m *TLPageBlockEmbed) GetW() int32  { return m.Data2.W }

func (m *TLPageBlockEmbed) SetH(v int32) { m.Data2.H = v }
func (m *TLPageBlockEmbed) GetH() int32  { return m.Data2.H }

func (m *TLPageBlockEmbed) SetCaption(v *RichText) { m.Data2.Caption = v }
func (m *TLPageBlockEmbed) GetCaption() *RichText  { return m.Data2.Caption }

func (m *TLPageBlockEmbed) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_pageBlockEmbed))

	// flags
	var flags uint32 = 0
	if m.GetFullWidth() == true {
		flags |= 1 << 1
	}
	if m.GetAllowScrolling() == true {
		flags |= 1 << 2
	}
	if m.GetUrl() != "" {
		flags |= 1 << 3
	}
	if m.GetHtml() != "" {
		flags |= 1 << 4
	}
	if m.GetPosterPhotoId() != 0 {
		flags |= 1 << 5
	}
	x.UInt(flags)

	if m.GetUrl() != 0 {
		x.StringBytes(m.Data2.Url)
	}
	if m.GetHtml() != 0 {
		x.StringBytes(m.Data2.Html)
	}
	if m.GetPosterPhotoId() != 0 {
		x.Long(m.Data2.PosterPhotoId)
	}
	x.Int(m.Data2.W)
	x.Int(m.Data2.H)
	x.Bytes(m.Data2Caption.Encode())

	return x.buf
}

func (m *TLPageBlockEmbed) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m.Data2.FullWidth = true
	}
	if (flags & (1 << 3)) != 0 {
		m.Data2.AllowScrolling = true
	}
	if (flags & (1 << 4)) != 0 {
		m.Data2.Url = x.StringBytes()
	}
	if (flags & (1 << 5)) != 0 {
		m.Data2.Html = x.StringBytes()
	}
	if (flags & (1 << 6)) != 0 {
		m.Data2.PosterPhotoId = x.Long()
	}
	m.Data2.W = x.Int()
	m.Data2.H = x.Int()
	m9 := &Caption{}
	m9.Decode(dbuf)
	m.SetCaption(m9)

	return dbuf.err
}

// pageBlockEmbedPost#292c7be9 url:string webpage_id:long author_photo_id:long author:string date:int blocks:Vector<PageBlock> caption:RichText = PageBlock;
func (m *TLPageBlockEmbedPost) To_PageBlock() *PageBlock {
	return &PageBlock{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPageBlockEmbedPost) SetUrl(v string) { m.Data2.Url = v }
func (m *TLPageBlockEmbedPost) GetUrl() string  { return m.Data2.Url }

func (m *TLPageBlockEmbedPost) SetWebpageId(v int64) { m.Data2.WebpageId = v }
func (m *TLPageBlockEmbedPost) GetWebpageId() int64  { return m.Data2.WebpageId }

func (m *TLPageBlockEmbedPost) SetAuthorPhotoId(v int64) { m.Data2.AuthorPhotoId = v }
func (m *TLPageBlockEmbedPost) GetAuthorPhotoId() int64  { return m.Data2.AuthorPhotoId }

func (m *TLPageBlockEmbedPost) SetAuthor(v string) { m.Data2.Author_23 = v }
func (m *TLPageBlockEmbedPost) GetAuthor() string  { return m.Data2.Author_23 }

func (m *TLPageBlockEmbedPost) SetDate(v int32) { m.Data2.Date = v }
func (m *TLPageBlockEmbedPost) GetDate() int32  { return m.Data2.Date }

func (m *TLPageBlockEmbedPost) SetBlocks(v []*PageBlock) { m.Data2.Blocks = v }
func (m *TLPageBlockEmbedPost) GetBlocks() []*PageBlock  { return m.Data2.Blocks }

func (m *TLPageBlockEmbedPost) SetCaption(v *RichText) { m.Data2.Caption = v }
func (m *TLPageBlockEmbedPost) GetCaption() *RichText  { return m.Data2.Caption }

func (m *TLPageBlockEmbedPost) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_pageBlockEmbedPost))

	x.StringBytes(m.Data2.Url)
	x.Long(m.Data2.WebpageId)
	x.Long(m.Data2.AuthorPhotoId)
	x.StringBytes(m.Data2.Author)
	x.Int(m.Data2.Date)

	x.Bytes(m.Data2Caption.Encode())

	return x.buf
}

func (m *TLPageBlockEmbedPost) Decode(dbuf *DecodeBuf) error {
	m.Data2.Url = x.StringBytes()
	m.Data2.WebpageId = x.Long()
	m.Data2.AuthorPhotoId = x.Long()
	m.Data2.Author = x.StringBytes()
	m.Data2.Date = x.Int()

	m7 := &Caption{}
	m7.Decode(dbuf)
	m.SetCaption(m7)

	return dbuf.err
}

// pageBlockCollage#8b31c4f items:Vector<PageBlock> caption:RichText = PageBlock;
func (m *TLPageBlockCollage) To_PageBlock() *PageBlock {
	return &PageBlock{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPageBlockCollage) SetItems(v []*PageBlock) { m.Data2.Items_26 = v }
func (m *TLPageBlockCollage) GetItems() []*PageBlock  { return m.Data2.Items_26 }

func (m *TLPageBlockCollage) SetCaption(v *RichText) { m.Data2.Caption = v }
func (m *TLPageBlockCollage) GetCaption() *RichText  { return m.Data2.Caption }

func (m *TLPageBlockCollage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_pageBlockCollage))

	x.Bytes(m.Data2Caption.Encode())

	return x.buf
}

func (m *TLPageBlockCollage) Decode(dbuf *DecodeBuf) error {

	m2 := &Caption{}
	m2.Decode(dbuf)
	m.SetCaption(m2)

	return dbuf.err
}

// pageBlockSlideshow#130c8963 items:Vector<PageBlock> caption:RichText = PageBlock;
func (m *TLPageBlockSlideshow) To_PageBlock() *PageBlock {
	return &PageBlock{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPageBlockSlideshow) SetItems(v []*PageBlock) { m.Data2.Items_26 = v }
func (m *TLPageBlockSlideshow) GetItems() []*PageBlock  { return m.Data2.Items_26 }

func (m *TLPageBlockSlideshow) SetCaption(v *RichText) { m.Data2.Caption = v }
func (m *TLPageBlockSlideshow) GetCaption() *RichText  { return m.Data2.Caption }

func (m *TLPageBlockSlideshow) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_pageBlockSlideshow))

	x.Bytes(m.Data2Caption.Encode())

	return x.buf
}

func (m *TLPageBlockSlideshow) Decode(dbuf *DecodeBuf) error {

	m2 := &Caption{}
	m2.Decode(dbuf)
	m.SetCaption(m2)

	return dbuf.err
}

// pageBlockChannel#ef1751b5 channel:Chat = PageBlock;
func (m *TLPageBlockChannel) To_PageBlock() *PageBlock {
	return &PageBlock{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPageBlockChannel) SetChannel(v *Chat) { m.Data2.Channel = v }
func (m *TLPageBlockChannel) GetChannel() *Chat  { return m.Data2.Channel }

func (m *TLPageBlockChannel) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_pageBlockChannel))

	x.Bytes(m.Data2Channel.Encode())

	return x.buf
}

func (m *TLPageBlockChannel) Decode(dbuf *DecodeBuf) error {
	m1 := &Channel{}
	m1.Decode(dbuf)
	m.SetChannel(m1)

	return dbuf.err
}

// pageBlockAudio#31b81a7f audio_id:long caption:RichText = PageBlock;
func (m *TLPageBlockAudio) To_PageBlock() *PageBlock {
	return &PageBlock{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPageBlockAudio) SetAudioId(v int64) { m.Data2.AudioId = v }
func (m *TLPageBlockAudio) GetAudioId() int64  { return m.Data2.AudioId }

func (m *TLPageBlockAudio) SetCaption(v *RichText) { m.Data2.Caption = v }
func (m *TLPageBlockAudio) GetCaption() *RichText  { return m.Data2.Caption }

func (m *TLPageBlockAudio) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_pageBlockAudio))

	x.Long(m.Data2.AudioId)
	x.Bytes(m.Data2Caption.Encode())

	return x.buf
}

func (m *TLPageBlockAudio) Decode(dbuf *DecodeBuf) error {
	m.Data2.AudioId = x.Long()
	m2 := &Caption{}
	m2.Decode(dbuf)
	m.SetCaption(m2)

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Account_PasswordInputSettings <--
//  + TL_AccountPasswordInputSettings
//

func (m *Account_PasswordInputSettings) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_account_passwordInputSettings:
		t := m.To_AccountPasswordInputSettings()
		return t.Encode()

	default:
		return []byte{}
	}
}

// account.passwordInputSettings#86916deb flags:# new_salt:flags.0?bytes new_password_hash:flags.0?bytes hint:flags.0?string email:flags.1?string = account.PasswordInputSettings;
func (m *Account_PasswordInputSettings) To_AccountPasswordInputSettings() *TLAccountPasswordInputSettings {
	return &TLAccountPasswordInputSettings{
		Data2: m.Data2,
	}
}

// account.passwordInputSettings#86916deb flags:# new_salt:flags.0?bytes new_password_hash:flags.0?bytes hint:flags.0?string email:flags.1?string = account.PasswordInputSettings;
func (m *TLAccountPasswordInputSettings) To_Account_PasswordInputSettings() *Account_PasswordInputSettings {
	return &Account_PasswordInputSettings{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLAccountPasswordInputSettings) SetNewSalt(v []byte) { m.Data2.NewSalt = v }
func (m *TLAccountPasswordInputSettings) GetNewSalt() []byte  { return m.Data2.NewSalt }

func (m *TLAccountPasswordInputSettings) SetNewPasswordHash(v []byte) { m.Data2.NewPasswordHash = v }
func (m *TLAccountPasswordInputSettings) GetNewPasswordHash() []byte  { return m.Data2.NewPasswordHash }

func (m *TLAccountPasswordInputSettings) SetHint(v string) { m.Data2.Hint = v }
func (m *TLAccountPasswordInputSettings) GetHint() string  { return m.Data2.Hint }

func (m *TLAccountPasswordInputSettings) SetEmail(v string) { m.Data2.Email = v }
func (m *TLAccountPasswordInputSettings) GetEmail() string  { return m.Data2.Email }

func (m *TLAccountPasswordInputSettings) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_account_passwordInputSettings))

	// flags
	var flags uint32 = 0
	if m.GetNewSalt() != nil {
		flags |= 1 << 1
	}
	if m.GetNewPasswordHash() != nil {
		flags |= 1 << 2
	}
	if m.GetHint() != "" {
		flags |= 1 << 3
	}
	if m.GetEmail() != "" {
		flags |= 1 << 4
	}
	x.UInt(flags)

	if m.GetNewSalt() != 0 {
		x.StringBytes(m.Data2.NewSalt)
	}
	if m.GetNewPasswordHash() != 0 {
		x.StringBytes(m.Data2.NewPasswordHash)
	}
	if m.GetHint() != 0 {
		x.StringBytes(m.Data2.Hint)
	}
	if m.GetEmail() != 0 {
		x.StringBytes(m.Data2.Email)
	}

	return x.buf
}

func (m *TLAccountPasswordInputSettings) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m.Data2.NewSalt = x.StringBytes()
	}
	if (flags & (1 << 3)) != 0 {
		m.Data2.NewPasswordHash = x.StringBytes()
	}
	if (flags & (1 << 4)) != 0 {
		m.Data2.Hint = x.StringBytes()
	}
	if (flags & (1 << 5)) != 0 {
		m.Data2.Email = x.StringBytes()
	}

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// InputBotInlineMessage <--
//  + TL_InputBotInlineMessageMediaAuto
//  + TL_InputBotInlineMessageText
//  + TL_InputBotInlineMessageMediaGeo
//  + TL_InputBotInlineMessageMediaVenue
//  + TL_InputBotInlineMessageMediaContact
//  + TL_InputBotInlineMessageGame
//

func (m *InputBotInlineMessage) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_inputBotInlineMessageMediaAuto:
		t := m.To_InputBotInlineMessageMediaAuto()
		return t.Encode()
	case TLConstructor_CRC32_inputBotInlineMessageText:
		t := m.To_InputBotInlineMessageText()
		return t.Encode()
	case TLConstructor_CRC32_inputBotInlineMessageMediaGeo:
		t := m.To_InputBotInlineMessageMediaGeo()
		return t.Encode()
	case TLConstructor_CRC32_inputBotInlineMessageMediaVenue:
		t := m.To_InputBotInlineMessageMediaVenue()
		return t.Encode()
	case TLConstructor_CRC32_inputBotInlineMessageMediaContact:
		t := m.To_InputBotInlineMessageMediaContact()
		return t.Encode()
	case TLConstructor_CRC32_inputBotInlineMessageGame:
		t := m.To_InputBotInlineMessageGame()
		return t.Encode()

	default:
		return []byte{}
	}
}

// inputBotInlineMessageMediaAuto#292fed13 flags:# caption:string reply_markup:flags.2?ReplyMarkup = InputBotInlineMessage;
func (m *InputBotInlineMessage) To_InputBotInlineMessageMediaAuto() *TLInputBotInlineMessageMediaAuto {
	return &TLInputBotInlineMessageMediaAuto{
		Data2: m.Data2,
	}
}

// inputBotInlineMessageText#3dcd7a87 flags:# no_webpage:flags.0?true message:string entities:flags.1?Vector<MessageEntity> reply_markup:flags.2?ReplyMarkup = InputBotInlineMessage;
func (m *InputBotInlineMessage) To_InputBotInlineMessageText() *TLInputBotInlineMessageText {
	return &TLInputBotInlineMessageText{
		Data2: m.Data2,
	}
}

// inputBotInlineMessageMediaGeo#f4a59de1 flags:# geo_point:InputGeoPoint reply_markup:flags.2?ReplyMarkup = InputBotInlineMessage;
func (m *InputBotInlineMessage) To_InputBotInlineMessageMediaGeo() *TLInputBotInlineMessageMediaGeo {
	return &TLInputBotInlineMessageMediaGeo{
		Data2: m.Data2,
	}
}

// inputBotInlineMessageMediaVenue#aaafadc8 flags:# geo_point:InputGeoPoint title:string address:string provider:string venue_id:string reply_markup:flags.2?ReplyMarkup = InputBotInlineMessage;
func (m *InputBotInlineMessage) To_InputBotInlineMessageMediaVenue() *TLInputBotInlineMessageMediaVenue {
	return &TLInputBotInlineMessageMediaVenue{
		Data2: m.Data2,
	}
}

// inputBotInlineMessageMediaContact#2daf01a7 flags:# phone_number:string first_name:string last_name:string reply_markup:flags.2?ReplyMarkup = InputBotInlineMessage;
func (m *InputBotInlineMessage) To_InputBotInlineMessageMediaContact() *TLInputBotInlineMessageMediaContact {
	return &TLInputBotInlineMessageMediaContact{
		Data2: m.Data2,
	}
}

// inputBotInlineMessageGame#4b425864 flags:# reply_markup:flags.2?ReplyMarkup = InputBotInlineMessage;
func (m *InputBotInlineMessage) To_InputBotInlineMessageGame() *TLInputBotInlineMessageGame {
	return &TLInputBotInlineMessageGame{
		Data2: m.Data2,
	}
}

// inputBotInlineMessageMediaAuto#292fed13 flags:# caption:string reply_markup:flags.2?ReplyMarkup = InputBotInlineMessage;
func (m *TLInputBotInlineMessageMediaAuto) To_InputBotInlineMessage() *InputBotInlineMessage {
	return &InputBotInlineMessage{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputBotInlineMessageMediaAuto) SetCaption(v string) { m.Data2.Caption = v }
func (m *TLInputBotInlineMessageMediaAuto) GetCaption() string  { return m.Data2.Caption }

func (m *TLInputBotInlineMessageMediaAuto) SetReplyMarkup(v *ReplyMarkup) { m.Data2.ReplyMarkup = v }
func (m *TLInputBotInlineMessageMediaAuto) GetReplyMarkup() *ReplyMarkup  { return m.Data2.ReplyMarkup }

func (m *TLInputBotInlineMessageMediaAuto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputBotInlineMessageMediaAuto))

	// flags
	var flags uint32 = 0
	if m.GetReplyMarkup() != nil {
		flags |= 1 << 2
	}
	x.UInt(flags)

	x.StringBytes(m.Data2.Caption)
	if m.GetReplyMarkup() != 0 {
		x.Bytes(m.Data2ReplyMarkup.Encode())
	}

	return x.buf
}

func (m *TLInputBotInlineMessageMediaAuto) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	m.Data2.Caption = x.StringBytes()
	if (flags & (1 << 3)) != 0 {
		m3 := &ReplyMarkup{}
		m3.Decode(dbuf)
		m.SetReplyMarkup(m3)
	}

	return dbuf.err
}

// inputBotInlineMessageText#3dcd7a87 flags:# no_webpage:flags.0?true message:string entities:flags.1?Vector<MessageEntity> reply_markup:flags.2?ReplyMarkup = InputBotInlineMessage;
func (m *TLInputBotInlineMessageText) To_InputBotInlineMessage() *InputBotInlineMessage {
	return &InputBotInlineMessage{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputBotInlineMessageText) SetNoWebpage(v bool) { m.Data2.NoWebpage = v }
func (m *TLInputBotInlineMessageText) GetNoWebpage() bool  { return m.Data2.NoWebpage }

func (m *TLInputBotInlineMessageText) SetMessage(v string) { m.Data2.Message = v }
func (m *TLInputBotInlineMessageText) GetMessage() string  { return m.Data2.Message }

func (m *TLInputBotInlineMessageText) SetEntities(v []*MessageEntity) { m.Data2.Entities = v }
func (m *TLInputBotInlineMessageText) GetEntities() []*MessageEntity  { return m.Data2.Entities }

func (m *TLInputBotInlineMessageText) SetReplyMarkup(v *ReplyMarkup) { m.Data2.ReplyMarkup = v }
func (m *TLInputBotInlineMessageText) GetReplyMarkup() *ReplyMarkup  { return m.Data2.ReplyMarkup }

func (m *TLInputBotInlineMessageText) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputBotInlineMessageText))

	// flags
	var flags uint32 = 0
	if m.GetNoWebpage() == true {
		flags |= 1 << 1
	}
	if m.GetEntities() != nil {
		flags |= 1 << 3
	}
	if m.GetReplyMarkup() != nil {
		flags |= 1 << 4
	}
	x.UInt(flags)

	x.StringBytes(m.Data2.Message)
	if m.Data2.Entities != 0 {

	}
	if m.GetReplyMarkup() != 0 {
		x.Bytes(m.Data2ReplyMarkup.Encode())
	}

	return x.buf
}

func (m *TLInputBotInlineMessageText) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m.Data2.NoWebpage = true
	}
	m.Data2.Message = x.StringBytes()
	if (flags & (1 << 4)) != 0 {

	}
	if (flags & (1 << 5)) != 0 {
		m5 := &ReplyMarkup{}
		m5.Decode(dbuf)
		m.SetReplyMarkup(m5)
	}

	return dbuf.err
}

// inputBotInlineMessageMediaGeo#f4a59de1 flags:# geo_point:InputGeoPoint reply_markup:flags.2?ReplyMarkup = InputBotInlineMessage;
func (m *TLInputBotInlineMessageMediaGeo) To_InputBotInlineMessage() *InputBotInlineMessage {
	return &InputBotInlineMessage{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputBotInlineMessageMediaGeo) SetGeoPoint(v *InputGeoPoint) { m.Data2.GeoPoint = v }
func (m *TLInputBotInlineMessageMediaGeo) GetGeoPoint() *InputGeoPoint  { return m.Data2.GeoPoint }

func (m *TLInputBotInlineMessageMediaGeo) SetReplyMarkup(v *ReplyMarkup) { m.Data2.ReplyMarkup = v }
func (m *TLInputBotInlineMessageMediaGeo) GetReplyMarkup() *ReplyMarkup  { return m.Data2.ReplyMarkup }

func (m *TLInputBotInlineMessageMediaGeo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputBotInlineMessageMediaGeo))

	// flags
	var flags uint32 = 0
	if m.GetReplyMarkup() != nil {
		flags |= 1 << 2
	}
	x.UInt(flags)

	x.Bytes(m.Data2GeoPoint.Encode())
	if m.GetReplyMarkup() != 0 {
		x.Bytes(m.Data2ReplyMarkup.Encode())
	}

	return x.buf
}

func (m *TLInputBotInlineMessageMediaGeo) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	m2 := &GeoPoint{}
	m2.Decode(dbuf)
	m.SetGeoPoint(m2)
	if (flags & (1 << 3)) != 0 {
		m3 := &ReplyMarkup{}
		m3.Decode(dbuf)
		m.SetReplyMarkup(m3)
	}

	return dbuf.err
}

// inputBotInlineMessageMediaVenue#aaafadc8 flags:# geo_point:InputGeoPoint title:string address:string provider:string venue_id:string reply_markup:flags.2?ReplyMarkup = InputBotInlineMessage;
func (m *TLInputBotInlineMessageMediaVenue) To_InputBotInlineMessage() *InputBotInlineMessage {
	return &InputBotInlineMessage{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputBotInlineMessageMediaVenue) SetGeoPoint(v *InputGeoPoint) { m.Data2.GeoPoint = v }
func (m *TLInputBotInlineMessageMediaVenue) GetGeoPoint() *InputGeoPoint  { return m.Data2.GeoPoint }

func (m *TLInputBotInlineMessageMediaVenue) SetTitle(v string) { m.Data2.Title = v }
func (m *TLInputBotInlineMessageMediaVenue) GetTitle() string  { return m.Data2.Title }

func (m *TLInputBotInlineMessageMediaVenue) SetAddress(v string) { m.Data2.Address = v }
func (m *TLInputBotInlineMessageMediaVenue) GetAddress() string  { return m.Data2.Address }

func (m *TLInputBotInlineMessageMediaVenue) SetProvider(v string) { m.Data2.Provider = v }
func (m *TLInputBotInlineMessageMediaVenue) GetProvider() string  { return m.Data2.Provider }

func (m *TLInputBotInlineMessageMediaVenue) SetVenueId(v string) { m.Data2.VenueId = v }
func (m *TLInputBotInlineMessageMediaVenue) GetVenueId() string  { return m.Data2.VenueId }

func (m *TLInputBotInlineMessageMediaVenue) SetReplyMarkup(v *ReplyMarkup) { m.Data2.ReplyMarkup = v }
func (m *TLInputBotInlineMessageMediaVenue) GetReplyMarkup() *ReplyMarkup  { return m.Data2.ReplyMarkup }

func (m *TLInputBotInlineMessageMediaVenue) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputBotInlineMessageMediaVenue))

	// flags
	var flags uint32 = 0
	if m.GetReplyMarkup() != nil {
		flags |= 1 << 6
	}
	x.UInt(flags)

	x.Bytes(m.Data2GeoPoint.Encode())
	x.StringBytes(m.Data2.Title)
	x.StringBytes(m.Data2.Address)
	x.StringBytes(m.Data2.Provider)
	x.StringBytes(m.Data2.VenueId)
	if m.GetReplyMarkup() != 0 {
		x.Bytes(m.Data2ReplyMarkup.Encode())
	}

	return x.buf
}

func (m *TLInputBotInlineMessageMediaVenue) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	m2 := &GeoPoint{}
	m2.Decode(dbuf)
	m.SetGeoPoint(m2)
	m.Data2.Title = x.StringBytes()
	m.Data2.Address = x.StringBytes()
	m.Data2.Provider = x.StringBytes()
	m.Data2.VenueId = x.StringBytes()
	if (flags & (1 << 7)) != 0 {
		m7 := &ReplyMarkup{}
		m7.Decode(dbuf)
		m.SetReplyMarkup(m7)
	}

	return dbuf.err
}

// inputBotInlineMessageMediaContact#2daf01a7 flags:# phone_number:string first_name:string last_name:string reply_markup:flags.2?ReplyMarkup = InputBotInlineMessage;
func (m *TLInputBotInlineMessageMediaContact) To_InputBotInlineMessage() *InputBotInlineMessage {
	return &InputBotInlineMessage{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputBotInlineMessageMediaContact) SetPhoneNumber(v string) { m.Data2.PhoneNumber = v }
func (m *TLInputBotInlineMessageMediaContact) GetPhoneNumber() string  { return m.Data2.PhoneNumber }

func (m *TLInputBotInlineMessageMediaContact) SetFirstName(v string) { m.Data2.FirstName = v }
func (m *TLInputBotInlineMessageMediaContact) GetFirstName() string  { return m.Data2.FirstName }

func (m *TLInputBotInlineMessageMediaContact) SetLastName(v string) { m.Data2.LastName = v }
func (m *TLInputBotInlineMessageMediaContact) GetLastName() string  { return m.Data2.LastName }

func (m *TLInputBotInlineMessageMediaContact) SetReplyMarkup(v *ReplyMarkup) { m.Data2.ReplyMarkup = v }
func (m *TLInputBotInlineMessageMediaContact) GetReplyMarkup() *ReplyMarkup {
	return m.Data2.ReplyMarkup
}

func (m *TLInputBotInlineMessageMediaContact) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputBotInlineMessageMediaContact))

	// flags
	var flags uint32 = 0
	if m.GetReplyMarkup() != nil {
		flags |= 1 << 4
	}
	x.UInt(flags)

	x.StringBytes(m.Data2.PhoneNumber)
	x.StringBytes(m.Data2.FirstName)
	x.StringBytes(m.Data2.LastName)
	if m.GetReplyMarkup() != 0 {
		x.Bytes(m.Data2ReplyMarkup.Encode())
	}

	return x.buf
}

func (m *TLInputBotInlineMessageMediaContact) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	m.Data2.PhoneNumber = x.StringBytes()
	m.Data2.FirstName = x.StringBytes()
	m.Data2.LastName = x.StringBytes()
	if (flags & (1 << 5)) != 0 {
		m5 := &ReplyMarkup{}
		m5.Decode(dbuf)
		m.SetReplyMarkup(m5)
	}

	return dbuf.err
}

// inputBotInlineMessageGame#4b425864 flags:# reply_markup:flags.2?ReplyMarkup = InputBotInlineMessage;
func (m *TLInputBotInlineMessageGame) To_InputBotInlineMessage() *InputBotInlineMessage {
	return &InputBotInlineMessage{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputBotInlineMessageGame) SetReplyMarkup(v *ReplyMarkup) { m.Data2.ReplyMarkup = v }
func (m *TLInputBotInlineMessageGame) GetReplyMarkup() *ReplyMarkup  { return m.Data2.ReplyMarkup }

func (m *TLInputBotInlineMessageGame) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputBotInlineMessageGame))

	// flags
	var flags uint32 = 0
	if m.GetReplyMarkup() != nil {
		flags |= 1 << 1
	}
	x.UInt(flags)

	if m.GetReplyMarkup() != 0 {
		x.Bytes(m.Data2ReplyMarkup.Encode())
	}

	return x.buf
}

func (m *TLInputBotInlineMessageGame) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m2 := &ReplyMarkup{}
		m2.Decode(dbuf)
		m.SetReplyMarkup(m2)
	}

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Upload_WebFile <--
//  + TL_UploadWebFile
//

func (m *Upload_WebFile) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_upload_webFile:
		t := m.To_UploadWebFile()
		return t.Encode()

	default:
		return []byte{}
	}
}

// upload.webFile#21e753bc size:int mime_type:string file_type:storage.FileType mtime:int bytes:bytes = upload.WebFile;
func (m *Upload_WebFile) To_UploadWebFile() *TLUploadWebFile {
	return &TLUploadWebFile{
		Data2: m.Data2,
	}
}

// upload.webFile#21e753bc size:int mime_type:string file_type:storage.FileType mtime:int bytes:bytes = upload.WebFile;
func (m *TLUploadWebFile) To_Upload_WebFile() *Upload_WebFile {
	return &Upload_WebFile{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUploadWebFile) SetSize(v int32) { m.Data2.Size = v }
func (m *TLUploadWebFile) GetSize() int32  { return m.Data2.Size }

func (m *TLUploadWebFile) SetMimeType(v string) { m.Data2.MimeType = v }
func (m *TLUploadWebFile) GetMimeType() string  { return m.Data2.MimeType }

func (m *TLUploadWebFile) SetFileType(v *Storage_FileType) { m.Data2.FileType = v }
func (m *TLUploadWebFile) GetFileType() *Storage_FileType  { return m.Data2.FileType }

func (m *TLUploadWebFile) SetMtime(v int32) { m.Data2.Mtime = v }
func (m *TLUploadWebFile) GetMtime() int32  { return m.Data2.Mtime }

func (m *TLUploadWebFile) SetBytes(v []byte) { m.Data2.Bytes = v }
func (m *TLUploadWebFile) GetBytes() []byte  { return m.Data2.Bytes }

func (m *TLUploadWebFile) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_upload_webFile))

	x.Int(m.Data2.Size)
	x.StringBytes(m.Data2.MimeType)
	x.Bytes(m.Data2FileType.Encode())
	x.Int(m.Data2.Mtime)
	x.StringBytes(m.Data2.Bytes)

	return x.buf
}

func (m *TLUploadWebFile) Decode(dbuf *DecodeBuf) error {
	m.Data2.Size = x.Int()
	m.Data2.MimeType = x.StringBytes()
	m3 := &FileType{}
	m3.Decode(dbuf)
	m.SetFileType(m3)
	m.Data2.Mtime = x.Int()
	m.Data2.Bytes = x.StringBytes()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// InputPaymentCredentials <--
//  + TL_InputPaymentCredentialsSaved
//  + TL_InputPaymentCredentials
//

func (m *InputPaymentCredentials) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_inputPaymentCredentialsSaved:
		t := m.To_InputPaymentCredentialsSaved()
		return t.Encode()
	case TLConstructor_CRC32_inputPaymentCredentials:
		t := m.To_InputPaymentCredentials()
		return t.Encode()

	default:
		return []byte{}
	}
}

// inputPaymentCredentialsSaved#c10eb2cf id:string tmp_password:bytes = InputPaymentCredentials;
func (m *InputPaymentCredentials) To_InputPaymentCredentialsSaved() *TLInputPaymentCredentialsSaved {
	return &TLInputPaymentCredentialsSaved{
		Data2: m.Data2,
	}
}

// inputPaymentCredentials#3417d728 flags:# save:flags.0?true data:DataJSON = InputPaymentCredentials;
func (m *InputPaymentCredentials) To_InputPaymentCredentials() *TLInputPaymentCredentials {
	return &TLInputPaymentCredentials{
		Data2: m.Data2,
	}
}

// inputPaymentCredentialsSaved#c10eb2cf id:string tmp_password:bytes = InputPaymentCredentials;
func (m *TLInputPaymentCredentialsSaved) To_InputPaymentCredentials() *InputPaymentCredentials {
	return &InputPaymentCredentials{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputPaymentCredentialsSaved) SetId(v string) { m.Data2.Id = v }
func (m *TLInputPaymentCredentialsSaved) GetId() string  { return m.Data2.Id }

func (m *TLInputPaymentCredentialsSaved) SetTmpPassword(v []byte) { m.Data2.TmpPassword = v }
func (m *TLInputPaymentCredentialsSaved) GetTmpPassword() []byte  { return m.Data2.TmpPassword }

func (m *TLInputPaymentCredentialsSaved) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputPaymentCredentialsSaved))

	x.StringBytes(m.Data2.Id)
	x.StringBytes(m.Data2.TmpPassword)

	return x.buf
}

func (m *TLInputPaymentCredentialsSaved) Decode(dbuf *DecodeBuf) error {
	m.Data2.Id = x.StringBytes()
	m.Data2.TmpPassword = x.StringBytes()

	return dbuf.err
}

// inputPaymentCredentials#3417d728 flags:# save:flags.0?true data:DataJSON = InputPaymentCredentials;
func (m *TLInputPaymentCredentials) To_InputPaymentCredentials() *InputPaymentCredentials {
	return &InputPaymentCredentials{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputPaymentCredentials) SetSave(v bool) { m.Data2.Save = v }
func (m *TLInputPaymentCredentials) GetSave() bool  { return m.Data2.Save }

func (m *TLInputPaymentCredentials) SetData(v *DataJSON) { m.Data2.Data = v }
func (m *TLInputPaymentCredentials) GetData() *DataJSON  { return m.Data2.Data }

func (m *TLInputPaymentCredentials) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputPaymentCredentials))

	// flags
	var flags uint32 = 0
	if m.GetSave() == true {
		flags |= 1 << 1
	}
	x.UInt(flags)

	x.Bytes(m.Data2Data.Encode())

	return x.buf
}

func (m *TLInputPaymentCredentials) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m.Data2.Save = true
	}
	m3 := &Data{}
	m3.Decode(dbuf)
	m.SetData(m3)

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// FileLocation <--
//  + TL_FileLocationUnavailable
//  + TL_FileLocation
//

func (m *FileLocation) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_fileLocationUnavailable:
		t := m.To_FileLocationUnavailable()
		return t.Encode()
	case TLConstructor_CRC32_fileLocation:
		t := m.To_FileLocation()
		return t.Encode()

	default:
		return []byte{}
	}
}

// fileLocationUnavailable#7c596b46 volume_id:long local_id:int secret:long = FileLocation;
func (m *FileLocation) To_FileLocationUnavailable() *TLFileLocationUnavailable {
	return &TLFileLocationUnavailable{
		Data2: m.Data2,
	}
}

// fileLocation#53d69076 dc_id:int volume_id:long local_id:int secret:long = FileLocation;
func (m *FileLocation) To_FileLocation() *TLFileLocation {
	return &TLFileLocation{
		Data2: m.Data2,
	}
}

// fileLocationUnavailable#7c596b46 volume_id:long local_id:int secret:long = FileLocation;
func (m *TLFileLocationUnavailable) To_FileLocation() *FileLocation {
	return &FileLocation{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLFileLocationUnavailable) SetVolumeId(v int64) { m.Data2.VolumeId = v }
func (m *TLFileLocationUnavailable) GetVolumeId() int64  { return m.Data2.VolumeId }

func (m *TLFileLocationUnavailable) SetLocalId(v int32) { m.Data2.LocalId = v }
func (m *TLFileLocationUnavailable) GetLocalId() int32  { return m.Data2.LocalId }

func (m *TLFileLocationUnavailable) SetSecret(v int64) { m.Data2.Secret = v }
func (m *TLFileLocationUnavailable) GetSecret() int64  { return m.Data2.Secret }

func (m *TLFileLocationUnavailable) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_fileLocationUnavailable))

	x.Long(m.Data2.VolumeId)
	x.Int(m.Data2.LocalId)
	x.Long(m.Data2.Secret)

	return x.buf
}

func (m *TLFileLocationUnavailable) Decode(dbuf *DecodeBuf) error {
	m.Data2.VolumeId = x.Long()
	m.Data2.LocalId = x.Int()
	m.Data2.Secret = x.Long()

	return dbuf.err
}

// fileLocation#53d69076 dc_id:int volume_id:long local_id:int secret:long = FileLocation;
func (m *TLFileLocation) To_FileLocation() *FileLocation {
	return &FileLocation{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLFileLocation) SetDcId(v int32) { m.Data2.DcId = v }
func (m *TLFileLocation) GetDcId() int32  { return m.Data2.DcId }

func (m *TLFileLocation) SetVolumeId(v int64) { m.Data2.VolumeId = v }
func (m *TLFileLocation) GetVolumeId() int64  { return m.Data2.VolumeId }

func (m *TLFileLocation) SetLocalId(v int32) { m.Data2.LocalId = v }
func (m *TLFileLocation) GetLocalId() int32  { return m.Data2.LocalId }

func (m *TLFileLocation) SetSecret(v int64) { m.Data2.Secret = v }
func (m *TLFileLocation) GetSecret() int64  { return m.Data2.Secret }

func (m *TLFileLocation) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_fileLocation))

	x.Int(m.Data2.DcId)
	x.Long(m.Data2.VolumeId)
	x.Int(m.Data2.LocalId)
	x.Long(m.Data2.Secret)

	return x.buf
}

func (m *TLFileLocation) Decode(dbuf *DecodeBuf) error {
	m.Data2.DcId = x.Int()
	m.Data2.VolumeId = x.Long()
	m.Data2.LocalId = x.Int()
	m.Data2.Secret = x.Long()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// EncryptedMessage <--
//  + TL_EncryptedMessage
//  + TL_EncryptedMessageService
//

func (m *EncryptedMessage) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_encryptedMessage:
		t := m.To_EncryptedMessage()
		return t.Encode()
	case TLConstructor_CRC32_encryptedMessageService:
		t := m.To_EncryptedMessageService()
		return t.Encode()

	default:
		return []byte{}
	}
}

// encryptedMessage#ed18c118 random_id:long chat_id:int date:int bytes:bytes file:EncryptedFile = EncryptedMessage;
func (m *EncryptedMessage) To_EncryptedMessage() *TLEncryptedMessage {
	return &TLEncryptedMessage{
		Data2: m.Data2,
	}
}

// encryptedMessageService#23734b06 random_id:long chat_id:int date:int bytes:bytes = EncryptedMessage;
func (m *EncryptedMessage) To_EncryptedMessageService() *TLEncryptedMessageService {
	return &TLEncryptedMessageService{
		Data2: m.Data2,
	}
}

// encryptedMessage#ed18c118 random_id:long chat_id:int date:int bytes:bytes file:EncryptedFile = EncryptedMessage;
func (m *TLEncryptedMessage) To_EncryptedMessage() *EncryptedMessage {
	return &EncryptedMessage{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLEncryptedMessage) SetRandomId(v int64) { m.Data2.RandomId = v }
func (m *TLEncryptedMessage) GetRandomId() int64  { return m.Data2.RandomId }

func (m *TLEncryptedMessage) SetChatId(v int32) { m.Data2.ChatId = v }
func (m *TLEncryptedMessage) GetChatId() int32  { return m.Data2.ChatId }

func (m *TLEncryptedMessage) SetDate(v int32) { m.Data2.Date = v }
func (m *TLEncryptedMessage) GetDate() int32  { return m.Data2.Date }

func (m *TLEncryptedMessage) SetBytes(v []byte) { m.Data2.Bytes = v }
func (m *TLEncryptedMessage) GetBytes() []byte  { return m.Data2.Bytes }

func (m *TLEncryptedMessage) SetFile(v *EncryptedFile) { m.Data2.File = v }
func (m *TLEncryptedMessage) GetFile() *EncryptedFile  { return m.Data2.File }

func (m *TLEncryptedMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_encryptedMessage))

	x.Long(m.Data2.RandomId)
	x.Int(m.Data2.ChatId)
	x.Int(m.Data2.Date)
	x.StringBytes(m.Data2.Bytes)
	x.Bytes(m.Data2File.Encode())

	return x.buf
}

func (m *TLEncryptedMessage) Decode(dbuf *DecodeBuf) error {
	m.Data2.RandomId = x.Long()
	m.Data2.ChatId = x.Int()
	m.Data2.Date = x.Int()
	m.Data2.Bytes = x.StringBytes()
	m5 := &File{}
	m5.Decode(dbuf)
	m.SetFile(m5)

	return dbuf.err
}

// encryptedMessageService#23734b06 random_id:long chat_id:int date:int bytes:bytes = EncryptedMessage;
func (m *TLEncryptedMessageService) To_EncryptedMessage() *EncryptedMessage {
	return &EncryptedMessage{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLEncryptedMessageService) SetRandomId(v int64) { m.Data2.RandomId = v }
func (m *TLEncryptedMessageService) GetRandomId() int64  { return m.Data2.RandomId }

func (m *TLEncryptedMessageService) SetChatId(v int32) { m.Data2.ChatId = v }
func (m *TLEncryptedMessageService) GetChatId() int32  { return m.Data2.ChatId }

func (m *TLEncryptedMessageService) SetDate(v int32) { m.Data2.Date = v }
func (m *TLEncryptedMessageService) GetDate() int32  { return m.Data2.Date }

func (m *TLEncryptedMessageService) SetBytes(v []byte) { m.Data2.Bytes = v }
func (m *TLEncryptedMessageService) GetBytes() []byte  { return m.Data2.Bytes }

func (m *TLEncryptedMessageService) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_encryptedMessageService))

	x.Long(m.Data2.RandomId)
	x.Int(m.Data2.ChatId)
	x.Int(m.Data2.Date)
	x.StringBytes(m.Data2.Bytes)

	return x.buf
}

func (m *TLEncryptedMessageService) Decode(dbuf *DecodeBuf) error {
	m.Data2.RandomId = x.Long()
	m.Data2.ChatId = x.Int()
	m.Data2.Date = x.Int()
	m.Data2.Bytes = x.StringBytes()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// SendMessageAction <--
//  + TL_SendMessageTypingAction
//  + TL_SendMessageCancelAction
//  + TL_SendMessageRecordVideoAction
//  + TL_SendMessageUploadVideoAction
//  + TL_SendMessageRecordAudioAction
//  + TL_SendMessageUploadAudioAction
//  + TL_SendMessageUploadPhotoAction
//  + TL_SendMessageUploadDocumentAction
//  + TL_SendMessageGeoLocationAction
//  + TL_SendMessageChooseContactAction
//  + TL_SendMessageGamePlayAction
//  + TL_SendMessageRecordRoundAction
//  + TL_SendMessageUploadRoundAction
//

func (m *SendMessageAction) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_sendMessageTypingAction:
		t := m.To_SendMessageTypingAction()
		return t.Encode()
	case TLConstructor_CRC32_sendMessageCancelAction:
		t := m.To_SendMessageCancelAction()
		return t.Encode()
	case TLConstructor_CRC32_sendMessageRecordVideoAction:
		t := m.To_SendMessageRecordVideoAction()
		return t.Encode()
	case TLConstructor_CRC32_sendMessageUploadVideoAction:
		t := m.To_SendMessageUploadVideoAction()
		return t.Encode()
	case TLConstructor_CRC32_sendMessageRecordAudioAction:
		t := m.To_SendMessageRecordAudioAction()
		return t.Encode()
	case TLConstructor_CRC32_sendMessageUploadAudioAction:
		t := m.To_SendMessageUploadAudioAction()
		return t.Encode()
	case TLConstructor_CRC32_sendMessageUploadPhotoAction:
		t := m.To_SendMessageUploadPhotoAction()
		return t.Encode()
	case TLConstructor_CRC32_sendMessageUploadDocumentAction:
		t := m.To_SendMessageUploadDocumentAction()
		return t.Encode()
	case TLConstructor_CRC32_sendMessageGeoLocationAction:
		t := m.To_SendMessageGeoLocationAction()
		return t.Encode()
	case TLConstructor_CRC32_sendMessageChooseContactAction:
		t := m.To_SendMessageChooseContactAction()
		return t.Encode()
	case TLConstructor_CRC32_sendMessageGamePlayAction:
		t := m.To_SendMessageGamePlayAction()
		return t.Encode()
	case TLConstructor_CRC32_sendMessageRecordRoundAction:
		t := m.To_SendMessageRecordRoundAction()
		return t.Encode()
	case TLConstructor_CRC32_sendMessageUploadRoundAction:
		t := m.To_SendMessageUploadRoundAction()
		return t.Encode()

	default:
		return []byte{}
	}
}

// sendMessageTypingAction#16bf744e = SendMessageAction;
func (m *SendMessageAction) To_SendMessageTypingAction() *TLSendMessageTypingAction {
	return &TLSendMessageTypingAction{
		Data2: m.Data2,
	}
}

// sendMessageCancelAction#fd5ec8f5 = SendMessageAction;
func (m *SendMessageAction) To_SendMessageCancelAction() *TLSendMessageCancelAction {
	return &TLSendMessageCancelAction{
		Data2: m.Data2,
	}
}

// sendMessageRecordVideoAction#a187d66f = SendMessageAction;
func (m *SendMessageAction) To_SendMessageRecordVideoAction() *TLSendMessageRecordVideoAction {
	return &TLSendMessageRecordVideoAction{
		Data2: m.Data2,
	}
}

// sendMessageUploadVideoAction#e9763aec progress:int = SendMessageAction;
func (m *SendMessageAction) To_SendMessageUploadVideoAction() *TLSendMessageUploadVideoAction {
	return &TLSendMessageUploadVideoAction{
		Data2: m.Data2,
	}
}

// sendMessageRecordAudioAction#d52f73f7 = SendMessageAction;
func (m *SendMessageAction) To_SendMessageRecordAudioAction() *TLSendMessageRecordAudioAction {
	return &TLSendMessageRecordAudioAction{
		Data2: m.Data2,
	}
}

// sendMessageUploadAudioAction#f351d7ab progress:int = SendMessageAction;
func (m *SendMessageAction) To_SendMessageUploadAudioAction() *TLSendMessageUploadAudioAction {
	return &TLSendMessageUploadAudioAction{
		Data2: m.Data2,
	}
}

// sendMessageUploadPhotoAction#d1d34a26 progress:int = SendMessageAction;
func (m *SendMessageAction) To_SendMessageUploadPhotoAction() *TLSendMessageUploadPhotoAction {
	return &TLSendMessageUploadPhotoAction{
		Data2: m.Data2,
	}
}

// sendMessageUploadDocumentAction#aa0cd9e4 progress:int = SendMessageAction;
func (m *SendMessageAction) To_SendMessageUploadDocumentAction() *TLSendMessageUploadDocumentAction {
	return &TLSendMessageUploadDocumentAction{
		Data2: m.Data2,
	}
}

// sendMessageGeoLocationAction#176f8ba1 = SendMessageAction;
func (m *SendMessageAction) To_SendMessageGeoLocationAction() *TLSendMessageGeoLocationAction {
	return &TLSendMessageGeoLocationAction{
		Data2: m.Data2,
	}
}

// sendMessageChooseContactAction#628cbc6f = SendMessageAction;
func (m *SendMessageAction) To_SendMessageChooseContactAction() *TLSendMessageChooseContactAction {
	return &TLSendMessageChooseContactAction{
		Data2: m.Data2,
	}
}

// sendMessageGamePlayAction#dd6a8f48 = SendMessageAction;
func (m *SendMessageAction) To_SendMessageGamePlayAction() *TLSendMessageGamePlayAction {
	return &TLSendMessageGamePlayAction{
		Data2: m.Data2,
	}
}

// sendMessageRecordRoundAction#88f27fbc = SendMessageAction;
func (m *SendMessageAction) To_SendMessageRecordRoundAction() *TLSendMessageRecordRoundAction {
	return &TLSendMessageRecordRoundAction{
		Data2: m.Data2,
	}
}

// sendMessageUploadRoundAction#243e1c66 progress:int = SendMessageAction;
func (m *SendMessageAction) To_SendMessageUploadRoundAction() *TLSendMessageUploadRoundAction {
	return &TLSendMessageUploadRoundAction{
		Data2: m.Data2,
	}
}

// sendMessageTypingAction#16bf744e = SendMessageAction;
func (m *TLSendMessageTypingAction) To_SendMessageAction() *SendMessageAction {
	return &SendMessageAction{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLSendMessageTypingAction) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_sendMessageTypingAction))

	return x.buf
}

func (m *TLSendMessageTypingAction) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// sendMessageCancelAction#fd5ec8f5 = SendMessageAction;
func (m *TLSendMessageCancelAction) To_SendMessageAction() *SendMessageAction {
	return &SendMessageAction{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLSendMessageCancelAction) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_sendMessageCancelAction))

	return x.buf
}

func (m *TLSendMessageCancelAction) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// sendMessageRecordVideoAction#a187d66f = SendMessageAction;
func (m *TLSendMessageRecordVideoAction) To_SendMessageAction() *SendMessageAction {
	return &SendMessageAction{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLSendMessageRecordVideoAction) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_sendMessageRecordVideoAction))

	return x.buf
}

func (m *TLSendMessageRecordVideoAction) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// sendMessageUploadVideoAction#e9763aec progress:int = SendMessageAction;
func (m *TLSendMessageUploadVideoAction) To_SendMessageAction() *SendMessageAction {
	return &SendMessageAction{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLSendMessageUploadVideoAction) SetProgress(v int32) { m.Data2.Progress = v }
func (m *TLSendMessageUploadVideoAction) GetProgress() int32  { return m.Data2.Progress }

func (m *TLSendMessageUploadVideoAction) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_sendMessageUploadVideoAction))

	x.Int(m.Data2.Progress)

	return x.buf
}

func (m *TLSendMessageUploadVideoAction) Decode(dbuf *DecodeBuf) error {
	m.Data2.Progress = x.Int()

	return dbuf.err
}

// sendMessageRecordAudioAction#d52f73f7 = SendMessageAction;
func (m *TLSendMessageRecordAudioAction) To_SendMessageAction() *SendMessageAction {
	return &SendMessageAction{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLSendMessageRecordAudioAction) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_sendMessageRecordAudioAction))

	return x.buf
}

func (m *TLSendMessageRecordAudioAction) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// sendMessageUploadAudioAction#f351d7ab progress:int = SendMessageAction;
func (m *TLSendMessageUploadAudioAction) To_SendMessageAction() *SendMessageAction {
	return &SendMessageAction{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLSendMessageUploadAudioAction) SetProgress(v int32) { m.Data2.Progress = v }
func (m *TLSendMessageUploadAudioAction) GetProgress() int32  { return m.Data2.Progress }

func (m *TLSendMessageUploadAudioAction) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_sendMessageUploadAudioAction))

	x.Int(m.Data2.Progress)

	return x.buf
}

func (m *TLSendMessageUploadAudioAction) Decode(dbuf *DecodeBuf) error {
	m.Data2.Progress = x.Int()

	return dbuf.err
}

// sendMessageUploadPhotoAction#d1d34a26 progress:int = SendMessageAction;
func (m *TLSendMessageUploadPhotoAction) To_SendMessageAction() *SendMessageAction {
	return &SendMessageAction{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLSendMessageUploadPhotoAction) SetProgress(v int32) { m.Data2.Progress = v }
func (m *TLSendMessageUploadPhotoAction) GetProgress() int32  { return m.Data2.Progress }

func (m *TLSendMessageUploadPhotoAction) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_sendMessageUploadPhotoAction))

	x.Int(m.Data2.Progress)

	return x.buf
}

func (m *TLSendMessageUploadPhotoAction) Decode(dbuf *DecodeBuf) error {
	m.Data2.Progress = x.Int()

	return dbuf.err
}

// sendMessageUploadDocumentAction#aa0cd9e4 progress:int = SendMessageAction;
func (m *TLSendMessageUploadDocumentAction) To_SendMessageAction() *SendMessageAction {
	return &SendMessageAction{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLSendMessageUploadDocumentAction) SetProgress(v int32) { m.Data2.Progress = v }
func (m *TLSendMessageUploadDocumentAction) GetProgress() int32  { return m.Data2.Progress }

func (m *TLSendMessageUploadDocumentAction) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_sendMessageUploadDocumentAction))

	x.Int(m.Data2.Progress)

	return x.buf
}

func (m *TLSendMessageUploadDocumentAction) Decode(dbuf *DecodeBuf) error {
	m.Data2.Progress = x.Int()

	return dbuf.err
}

// sendMessageGeoLocationAction#176f8ba1 = SendMessageAction;
func (m *TLSendMessageGeoLocationAction) To_SendMessageAction() *SendMessageAction {
	return &SendMessageAction{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLSendMessageGeoLocationAction) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_sendMessageGeoLocationAction))

	return x.buf
}

func (m *TLSendMessageGeoLocationAction) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// sendMessageChooseContactAction#628cbc6f = SendMessageAction;
func (m *TLSendMessageChooseContactAction) To_SendMessageAction() *SendMessageAction {
	return &SendMessageAction{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLSendMessageChooseContactAction) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_sendMessageChooseContactAction))

	return x.buf
}

func (m *TLSendMessageChooseContactAction) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// sendMessageGamePlayAction#dd6a8f48 = SendMessageAction;
func (m *TLSendMessageGamePlayAction) To_SendMessageAction() *SendMessageAction {
	return &SendMessageAction{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLSendMessageGamePlayAction) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_sendMessageGamePlayAction))

	return x.buf
}

func (m *TLSendMessageGamePlayAction) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// sendMessageRecordRoundAction#88f27fbc = SendMessageAction;
func (m *TLSendMessageRecordRoundAction) To_SendMessageAction() *SendMessageAction {
	return &SendMessageAction{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLSendMessageRecordRoundAction) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_sendMessageRecordRoundAction))

	return x.buf
}

func (m *TLSendMessageRecordRoundAction) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// sendMessageUploadRoundAction#243e1c66 progress:int = SendMessageAction;
func (m *TLSendMessageUploadRoundAction) To_SendMessageAction() *SendMessageAction {
	return &SendMessageAction{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLSendMessageUploadRoundAction) SetProgress(v int32) { m.Data2.Progress = v }
func (m *TLSendMessageUploadRoundAction) GetProgress() int32  { return m.Data2.Progress }

func (m *TLSendMessageUploadRoundAction) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_sendMessageUploadRoundAction))

	x.Int(m.Data2.Progress)

	return x.buf
}

func (m *TLSendMessageUploadRoundAction) Decode(dbuf *DecodeBuf) error {
	m.Data2.Progress = x.Int()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// ContactLink <--
//  + TL_ContactLinkUnknown
//  + TL_ContactLinkNone
//  + TL_ContactLinkHasPhone
//  + TL_ContactLinkContact
//

func (m *ContactLink) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_contactLinkUnknown:
		t := m.To_ContactLinkUnknown()
		return t.Encode()
	case TLConstructor_CRC32_contactLinkNone:
		t := m.To_ContactLinkNone()
		return t.Encode()
	case TLConstructor_CRC32_contactLinkHasPhone:
		t := m.To_ContactLinkHasPhone()
		return t.Encode()
	case TLConstructor_CRC32_contactLinkContact:
		t := m.To_ContactLinkContact()
		return t.Encode()

	default:
		return []byte{}
	}
}

// contactLinkUnknown#5f4f9247 = ContactLink;
func (m *ContactLink) To_ContactLinkUnknown() *TLContactLinkUnknown {
	return &TLContactLinkUnknown{
		Data2: m.Data2,
	}
}

// contactLinkNone#feedd3ad = ContactLink;
func (m *ContactLink) To_ContactLinkNone() *TLContactLinkNone {
	return &TLContactLinkNone{
		Data2: m.Data2,
	}
}

// contactLinkHasPhone#268f3f59 = ContactLink;
func (m *ContactLink) To_ContactLinkHasPhone() *TLContactLinkHasPhone {
	return &TLContactLinkHasPhone{
		Data2: m.Data2,
	}
}

// contactLinkContact#d502c2d0 = ContactLink;
func (m *ContactLink) To_ContactLinkContact() *TLContactLinkContact {
	return &TLContactLinkContact{
		Data2: m.Data2,
	}
}

// contactLinkUnknown#5f4f9247 = ContactLink;
func (m *TLContactLinkUnknown) To_ContactLink() *ContactLink {
	return &ContactLink{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLContactLinkUnknown) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_contactLinkUnknown))

	return x.buf
}

func (m *TLContactLinkUnknown) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// contactLinkNone#feedd3ad = ContactLink;
func (m *TLContactLinkNone) To_ContactLink() *ContactLink {
	return &ContactLink{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLContactLinkNone) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_contactLinkNone))

	return x.buf
}

func (m *TLContactLinkNone) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// contactLinkHasPhone#268f3f59 = ContactLink;
func (m *TLContactLinkHasPhone) To_ContactLink() *ContactLink {
	return &ContactLink{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLContactLinkHasPhone) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_contactLinkHasPhone))

	return x.buf
}

func (m *TLContactLinkHasPhone) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// contactLinkContact#d502c2d0 = ContactLink;
func (m *TLContactLinkContact) To_ContactLink() *ContactLink {
	return &ContactLink{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLContactLinkContact) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_contactLinkContact))

	return x.buf
}

func (m *TLContactLinkContact) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// LangPackDifference <--
//  + TL_LangPackDifference
//

func (m *LangPackDifference) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_langPackDifference:
		t := m.To_LangPackDifference()
		return t.Encode()

	default:
		return []byte{}
	}
}

// langPackDifference#f385c1f6 lang_code:string from_version:int version:int strings:Vector<LangPackString> = LangPackDifference;
func (m *LangPackDifference) To_LangPackDifference() *TLLangPackDifference {
	return &TLLangPackDifference{
		Data2: m.Data2,
	}
}

// langPackDifference#f385c1f6 lang_code:string from_version:int version:int strings:Vector<LangPackString> = LangPackDifference;
func (m *TLLangPackDifference) To_LangPackDifference() *LangPackDifference {
	return &LangPackDifference{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLLangPackDifference) SetLangCode(v string) { m.Data2.LangCode = v }
func (m *TLLangPackDifference) GetLangCode() string  { return m.Data2.LangCode }

func (m *TLLangPackDifference) SetFromVersion(v int32) { m.Data2.FromVersion = v }
func (m *TLLangPackDifference) GetFromVersion() int32  { return m.Data2.FromVersion }

func (m *TLLangPackDifference) SetVersion(v int32) { m.Data2.Version = v }
func (m *TLLangPackDifference) GetVersion() int32  { return m.Data2.Version }

func (m *TLLangPackDifference) SetStrings(v []*LangPackString) { m.Data2.Strings = v }
func (m *TLLangPackDifference) GetStrings() []*LangPackString  { return m.Data2.Strings }

func (m *TLLangPackDifference) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_langPackDifference))

	x.StringBytes(m.Data2.LangCode)
	x.Int(m.Data2.FromVersion)
	x.Int(m.Data2.Version)

	return x.buf
}

func (m *TLLangPackDifference) Decode(dbuf *DecodeBuf) error {
	m.Data2.LangCode = x.StringBytes()
	m.Data2.FromVersion = x.Int()
	m.Data2.Version = x.Int()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// PaymentSavedCredentials <--
//  + TL_PaymentSavedCredentialsCard
//

func (m *PaymentSavedCredentials) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_paymentSavedCredentialsCard:
		t := m.To_PaymentSavedCredentialsCard()
		return t.Encode()

	default:
		return []byte{}
	}
}

// paymentSavedCredentialsCard#cdc27a1f id:string title:string = PaymentSavedCredentials;
func (m *PaymentSavedCredentials) To_PaymentSavedCredentialsCard() *TLPaymentSavedCredentialsCard {
	return &TLPaymentSavedCredentialsCard{
		Data2: m.Data2,
	}
}

// paymentSavedCredentialsCard#cdc27a1f id:string title:string = PaymentSavedCredentials;
func (m *TLPaymentSavedCredentialsCard) To_PaymentSavedCredentials() *PaymentSavedCredentials {
	return &PaymentSavedCredentials{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPaymentSavedCredentialsCard) SetId(v string) { m.Data2.Id = v }
func (m *TLPaymentSavedCredentialsCard) GetId() string  { return m.Data2.Id }

func (m *TLPaymentSavedCredentialsCard) SetTitle(v string) { m.Data2.Title = v }
func (m *TLPaymentSavedCredentialsCard) GetTitle() string  { return m.Data2.Title }

func (m *TLPaymentSavedCredentialsCard) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_paymentSavedCredentialsCard))

	x.StringBytes(m.Data2.Id)
	x.StringBytes(m.Data2.Title)

	return x.buf
}

func (m *TLPaymentSavedCredentialsCard) Decode(dbuf *DecodeBuf) error {
	m.Data2.Id = x.StringBytes()
	m.Data2.Title = x.StringBytes()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// PhoneCall <--
//  + TL_PhoneCallEmpty
//  + TL_PhoneCallWaiting
//  + TL_PhoneCallRequested
//  + TL_PhoneCallAccepted
//  + TL_PhoneCall
//  + TL_PhoneCallDiscarded
//

func (m *PhoneCall) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_phoneCallEmpty:
		t := m.To_PhoneCallEmpty()
		return t.Encode()
	case TLConstructor_CRC32_phoneCallWaiting:
		t := m.To_PhoneCallWaiting()
		return t.Encode()
	case TLConstructor_CRC32_phoneCallRequested:
		t := m.To_PhoneCallRequested()
		return t.Encode()
	case TLConstructor_CRC32_phoneCallAccepted:
		t := m.To_PhoneCallAccepted()
		return t.Encode()
	case TLConstructor_CRC32_phoneCall:
		t := m.To_PhoneCall()
		return t.Encode()
	case TLConstructor_CRC32_phoneCallDiscarded:
		t := m.To_PhoneCallDiscarded()
		return t.Encode()

	default:
		return []byte{}
	}
}

// phoneCallEmpty#5366c915 id:long = PhoneCall;
func (m *PhoneCall) To_PhoneCallEmpty() *TLPhoneCallEmpty {
	return &TLPhoneCallEmpty{
		Data2: m.Data2,
	}
}

// phoneCallWaiting#1b8f4ad1 flags:# id:long access_hash:long date:int admin_id:int participant_id:int protocol:PhoneCallProtocol receive_date:flags.0?int = PhoneCall;
func (m *PhoneCall) To_PhoneCallWaiting() *TLPhoneCallWaiting {
	return &TLPhoneCallWaiting{
		Data2: m.Data2,
	}
}

// phoneCallRequested#83761ce4 id:long access_hash:long date:int admin_id:int participant_id:int g_a_hash:bytes protocol:PhoneCallProtocol = PhoneCall;
func (m *PhoneCall) To_PhoneCallRequested() *TLPhoneCallRequested {
	return &TLPhoneCallRequested{
		Data2: m.Data2,
	}
}

// phoneCallAccepted#6d003d3f id:long access_hash:long date:int admin_id:int participant_id:int g_b:bytes protocol:PhoneCallProtocol = PhoneCall;
func (m *PhoneCall) To_PhoneCallAccepted() *TLPhoneCallAccepted {
	return &TLPhoneCallAccepted{
		Data2: m.Data2,
	}
}

// phoneCall#ffe6ab67 id:long access_hash:long date:int admin_id:int participant_id:int g_a_or_b:bytes key_fingerprint:long protocol:PhoneCallProtocol connection:PhoneConnection alternative_connections:Vector<PhoneConnection> start_date:int = PhoneCall;
func (m *PhoneCall) To_PhoneCall() *TLPhoneCall {
	return &TLPhoneCall{
		Data2: m.Data2,
	}
}

// phoneCallDiscarded#50ca4de1 flags:# need_rating:flags.2?true need_debug:flags.3?true id:long reason:flags.0?PhoneCallDiscardReason duration:flags.1?int = PhoneCall;
func (m *PhoneCall) To_PhoneCallDiscarded() *TLPhoneCallDiscarded {
	return &TLPhoneCallDiscarded{
		Data2: m.Data2,
	}
}

// phoneCallEmpty#5366c915 id:long = PhoneCall;
func (m *TLPhoneCallEmpty) To_PhoneCall() *PhoneCall {
	return &PhoneCall{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPhoneCallEmpty) SetId(v int64) { m.Data2.Id = v }
func (m *TLPhoneCallEmpty) GetId() int64  { return m.Data2.Id }

func (m *TLPhoneCallEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_phoneCallEmpty))

	x.Long(m.Data2.Id)

	return x.buf
}

func (m *TLPhoneCallEmpty) Decode(dbuf *DecodeBuf) error {
	m.Data2.Id = x.Long()

	return dbuf.err
}

// phoneCallWaiting#1b8f4ad1 flags:# id:long access_hash:long date:int admin_id:int participant_id:int protocol:PhoneCallProtocol receive_date:flags.0?int = PhoneCall;
func (m *TLPhoneCallWaiting) To_PhoneCall() *PhoneCall {
	return &PhoneCall{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPhoneCallWaiting) SetId(v int64) { m.Data2.Id = v }
func (m *TLPhoneCallWaiting) GetId() int64  { return m.Data2.Id }

func (m *TLPhoneCallWaiting) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLPhoneCallWaiting) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLPhoneCallWaiting) SetDate(v int32) { m.Data2.Date = v }
func (m *TLPhoneCallWaiting) GetDate() int32  { return m.Data2.Date }

func (m *TLPhoneCallWaiting) SetAdminId(v int32) { m.Data2.AdminId = v }
func (m *TLPhoneCallWaiting) GetAdminId() int32  { return m.Data2.AdminId }

func (m *TLPhoneCallWaiting) SetParticipantId(v int32) { m.Data2.ParticipantId = v }
func (m *TLPhoneCallWaiting) GetParticipantId() int32  { return m.Data2.ParticipantId }

func (m *TLPhoneCallWaiting) SetProtocol(v *PhoneCallProtocol) { m.Data2.Protocol = v }
func (m *TLPhoneCallWaiting) GetProtocol() *PhoneCallProtocol  { return m.Data2.Protocol }

func (m *TLPhoneCallWaiting) SetReceiveDate(v int32) { m.Data2.ReceiveDate = v }
func (m *TLPhoneCallWaiting) GetReceiveDate() int32  { return m.Data2.ReceiveDate }

func (m *TLPhoneCallWaiting) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_phoneCallWaiting))

	// flags
	var flags uint32 = 0
	if m.GetReceiveDate() != 0 {
		flags |= 1 << 7
	}
	x.UInt(flags)

	x.Long(m.Data2.Id)
	x.Long(m.Data2.AccessHash)
	x.Int(m.Data2.Date)
	x.Int(m.Data2.AdminId)
	x.Int(m.Data2.ParticipantId)
	x.Bytes(m.Data2Protocol.Encode())
	if m.GetReceiveDate() != 0 {
		x.Int(m.Data2.ReceiveDate)
	}

	return x.buf
}

func (m *TLPhoneCallWaiting) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	m.Data2.Id = x.Long()
	m.Data2.AccessHash = x.Long()
	m.Data2.Date = x.Int()
	m.Data2.AdminId = x.Int()
	m.Data2.ParticipantId = x.Int()
	m7 := &Protocol{}
	m7.Decode(dbuf)
	m.SetProtocol(m7)
	if (flags & (1 << 8)) != 0 {
		m.Data2.ReceiveDate = x.Int()
	}

	return dbuf.err
}

// phoneCallRequested#83761ce4 id:long access_hash:long date:int admin_id:int participant_id:int g_a_hash:bytes protocol:PhoneCallProtocol = PhoneCall;
func (m *TLPhoneCallRequested) To_PhoneCall() *PhoneCall {
	return &PhoneCall{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPhoneCallRequested) SetId(v int64) { m.Data2.Id = v }
func (m *TLPhoneCallRequested) GetId() int64  { return m.Data2.Id }

func (m *TLPhoneCallRequested) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLPhoneCallRequested) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLPhoneCallRequested) SetDate(v int32) { m.Data2.Date = v }
func (m *TLPhoneCallRequested) GetDate() int32  { return m.Data2.Date }

func (m *TLPhoneCallRequested) SetAdminId(v int32) { m.Data2.AdminId = v }
func (m *TLPhoneCallRequested) GetAdminId() int32  { return m.Data2.AdminId }

func (m *TLPhoneCallRequested) SetParticipantId(v int32) { m.Data2.ParticipantId = v }
func (m *TLPhoneCallRequested) GetParticipantId() int32  { return m.Data2.ParticipantId }

func (m *TLPhoneCallRequested) SetGAHash(v []byte) { m.Data2.GAHash = v }
func (m *TLPhoneCallRequested) GetGAHash() []byte  { return m.Data2.GAHash }

func (m *TLPhoneCallRequested) SetProtocol(v *PhoneCallProtocol) { m.Data2.Protocol = v }
func (m *TLPhoneCallRequested) GetProtocol() *PhoneCallProtocol  { return m.Data2.Protocol }

func (m *TLPhoneCallRequested) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_phoneCallRequested))

	x.Long(m.Data2.Id)
	x.Long(m.Data2.AccessHash)
	x.Int(m.Data2.Date)
	x.Int(m.Data2.AdminId)
	x.Int(m.Data2.ParticipantId)
	x.StringBytes(m.Data2.GAHash)
	x.Bytes(m.Data2Protocol.Encode())

	return x.buf
}

func (m *TLPhoneCallRequested) Decode(dbuf *DecodeBuf) error {
	m.Data2.Id = x.Long()
	m.Data2.AccessHash = x.Long()
	m.Data2.Date = x.Int()
	m.Data2.AdminId = x.Int()
	m.Data2.ParticipantId = x.Int()
	m.Data2.GAHash = x.StringBytes()
	m7 := &Protocol{}
	m7.Decode(dbuf)
	m.SetProtocol(m7)

	return dbuf.err
}

// phoneCallAccepted#6d003d3f id:long access_hash:long date:int admin_id:int participant_id:int g_b:bytes protocol:PhoneCallProtocol = PhoneCall;
func (m *TLPhoneCallAccepted) To_PhoneCall() *PhoneCall {
	return &PhoneCall{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPhoneCallAccepted) SetId(v int64) { m.Data2.Id = v }
func (m *TLPhoneCallAccepted) GetId() int64  { return m.Data2.Id }

func (m *TLPhoneCallAccepted) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLPhoneCallAccepted) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLPhoneCallAccepted) SetDate(v int32) { m.Data2.Date = v }
func (m *TLPhoneCallAccepted) GetDate() int32  { return m.Data2.Date }

func (m *TLPhoneCallAccepted) SetAdminId(v int32) { m.Data2.AdminId = v }
func (m *TLPhoneCallAccepted) GetAdminId() int32  { return m.Data2.AdminId }

func (m *TLPhoneCallAccepted) SetParticipantId(v int32) { m.Data2.ParticipantId = v }
func (m *TLPhoneCallAccepted) GetParticipantId() int32  { return m.Data2.ParticipantId }

func (m *TLPhoneCallAccepted) SetGB(v []byte) { m.Data2.GB = v }
func (m *TLPhoneCallAccepted) GetGB() []byte  { return m.Data2.GB }

func (m *TLPhoneCallAccepted) SetProtocol(v *PhoneCallProtocol) { m.Data2.Protocol = v }
func (m *TLPhoneCallAccepted) GetProtocol() *PhoneCallProtocol  { return m.Data2.Protocol }

func (m *TLPhoneCallAccepted) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_phoneCallAccepted))

	x.Long(m.Data2.Id)
	x.Long(m.Data2.AccessHash)
	x.Int(m.Data2.Date)
	x.Int(m.Data2.AdminId)
	x.Int(m.Data2.ParticipantId)
	x.StringBytes(m.Data2.GB)
	x.Bytes(m.Data2Protocol.Encode())

	return x.buf
}

func (m *TLPhoneCallAccepted) Decode(dbuf *DecodeBuf) error {
	m.Data2.Id = x.Long()
	m.Data2.AccessHash = x.Long()
	m.Data2.Date = x.Int()
	m.Data2.AdminId = x.Int()
	m.Data2.ParticipantId = x.Int()
	m.Data2.GB = x.StringBytes()
	m7 := &Protocol{}
	m7.Decode(dbuf)
	m.SetProtocol(m7)

	return dbuf.err
}

// phoneCall#ffe6ab67 id:long access_hash:long date:int admin_id:int participant_id:int g_a_or_b:bytes key_fingerprint:long protocol:PhoneCallProtocol connection:PhoneConnection alternative_connections:Vector<PhoneConnection> start_date:int = PhoneCall;
func (m *TLPhoneCall) To_PhoneCall() *PhoneCall {
	return &PhoneCall{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPhoneCall) SetId(v int64) { m.Data2.Id = v }
func (m *TLPhoneCall) GetId() int64  { return m.Data2.Id }

func (m *TLPhoneCall) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLPhoneCall) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLPhoneCall) SetDate(v int32) { m.Data2.Date = v }
func (m *TLPhoneCall) GetDate() int32  { return m.Data2.Date }

func (m *TLPhoneCall) SetAdminId(v int32) { m.Data2.AdminId = v }
func (m *TLPhoneCall) GetAdminId() int32  { return m.Data2.AdminId }

func (m *TLPhoneCall) SetParticipantId(v int32) { m.Data2.ParticipantId = v }
func (m *TLPhoneCall) GetParticipantId() int32  { return m.Data2.ParticipantId }

func (m *TLPhoneCall) SetGAOrB(v []byte) { m.Data2.GAOrB = v }
func (m *TLPhoneCall) GetGAOrB() []byte  { return m.Data2.GAOrB }

func (m *TLPhoneCall) SetKeyFingerprint(v int64) { m.Data2.KeyFingerprint = v }
func (m *TLPhoneCall) GetKeyFingerprint() int64  { return m.Data2.KeyFingerprint }

func (m *TLPhoneCall) SetProtocol(v *PhoneCallProtocol) { m.Data2.Protocol = v }
func (m *TLPhoneCall) GetProtocol() *PhoneCallProtocol  { return m.Data2.Protocol }

func (m *TLPhoneCall) SetConnection(v *PhoneConnection) { m.Data2.Connection = v }
func (m *TLPhoneCall) GetConnection() *PhoneConnection  { return m.Data2.Connection }

func (m *TLPhoneCall) SetAlternativeConnections(v []*PhoneConnection) {
	m.Data2.AlternativeConnections = v
}
func (m *TLPhoneCall) GetAlternativeConnections() []*PhoneConnection {
	return m.Data2.AlternativeConnections
}

func (m *TLPhoneCall) SetStartDate(v int32) { m.Data2.StartDate = v }
func (m *TLPhoneCall) GetStartDate() int32  { return m.Data2.StartDate }

func (m *TLPhoneCall) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_phoneCall))

	x.Long(m.Data2.Id)
	x.Long(m.Data2.AccessHash)
	x.Int(m.Data2.Date)
	x.Int(m.Data2.AdminId)
	x.Int(m.Data2.ParticipantId)
	x.StringBytes(m.Data2.GAOrB)
	x.Long(m.Data2.KeyFingerprint)
	x.Bytes(m.Data2Protocol.Encode())
	x.Bytes(m.Data2Connection.Encode())

	x.Int(m.Data2.StartDate)

	return x.buf
}

func (m *TLPhoneCall) Decode(dbuf *DecodeBuf) error {
	m.Data2.Id = x.Long()
	m.Data2.AccessHash = x.Long()
	m.Data2.Date = x.Int()
	m.Data2.AdminId = x.Int()
	m.Data2.ParticipantId = x.Int()
	m.Data2.GAOrB = x.StringBytes()
	m.Data2.KeyFingerprint = x.Long()
	m8 := &Protocol{}
	m8.Decode(dbuf)
	m.SetProtocol(m8)
	m9 := &Connection{}
	m9.Decode(dbuf)
	m.SetConnection(m9)

	m.Data2.StartDate = x.Int()

	return dbuf.err
}

// phoneCallDiscarded#50ca4de1 flags:# need_rating:flags.2?true need_debug:flags.3?true id:long reason:flags.0?PhoneCallDiscardReason duration:flags.1?int = PhoneCall;
func (m *TLPhoneCallDiscarded) To_PhoneCall() *PhoneCall {
	return &PhoneCall{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPhoneCallDiscarded) SetNeedRating(v bool) { m.Data2.NeedRating = v }
func (m *TLPhoneCallDiscarded) GetNeedRating() bool  { return m.Data2.NeedRating }

func (m *TLPhoneCallDiscarded) SetNeedDebug(v bool) { m.Data2.NeedDebug = v }
func (m *TLPhoneCallDiscarded) GetNeedDebug() bool  { return m.Data2.NeedDebug }

func (m *TLPhoneCallDiscarded) SetId(v int64) { m.Data2.Id = v }
func (m *TLPhoneCallDiscarded) GetId() int64  { return m.Data2.Id }

func (m *TLPhoneCallDiscarded) SetReason(v *PhoneCallDiscardReason) { m.Data2.Reason = v }
func (m *TLPhoneCallDiscarded) GetReason() *PhoneCallDiscardReason  { return m.Data2.Reason }

func (m *TLPhoneCallDiscarded) SetDuration(v int32) { m.Data2.Duration = v }
func (m *TLPhoneCallDiscarded) GetDuration() int32  { return m.Data2.Duration }

func (m *TLPhoneCallDiscarded) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_phoneCallDiscarded))

	// flags
	var flags uint32 = 0
	if m.GetNeedRating() == true {
		flags |= 1 << 1
	}
	if m.GetNeedDebug() == true {
		flags |= 1 << 2
	}
	if m.GetReason() != nil {
		flags |= 1 << 4
	}
	if m.GetDuration() != 0 {
		flags |= 1 << 5
	}
	x.UInt(flags)

	x.Long(m.Data2.Id)
	if m.GetReason() != 0 {
		x.Bytes(m.Data2Reason.Encode())
	}
	if m.GetDuration() != 0 {
		x.Int(m.Data2.Duration)
	}

	return x.buf
}

func (m *TLPhoneCallDiscarded) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m.Data2.NeedRating = true
	}
	if (flags & (1 << 3)) != 0 {
		m.Data2.NeedDebug = true
	}
	m.Data2.Id = x.Long()
	if (flags & (1 << 5)) != 0 {
		m5 := &Reason{}
		m5.Decode(dbuf)
		m.SetReason(m5)
	}
	if (flags & (1 << 6)) != 0 {
		m.Data2.Duration = x.Int()
	}

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Contacts_ImportedContacts <--
//  + TL_ContactsImportedContacts
//

func (m *Contacts_ImportedContacts) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_contacts_importedContacts:
		t := m.To_ContactsImportedContacts()
		return t.Encode()

	default:
		return []byte{}
	}
}

// contacts.importedContacts#77d01c3b imported:Vector<ImportedContact> popular_invites:Vector<PopularContact> retry_contacts:Vector<long> users:Vector<User> = contacts.ImportedContacts;
func (m *Contacts_ImportedContacts) To_ContactsImportedContacts() *TLContactsImportedContacts {
	return &TLContactsImportedContacts{
		Data2: m.Data2,
	}
}

// contacts.importedContacts#77d01c3b imported:Vector<ImportedContact> popular_invites:Vector<PopularContact> retry_contacts:Vector<long> users:Vector<User> = contacts.ImportedContacts;
func (m *TLContactsImportedContacts) To_Contacts_ImportedContacts() *Contacts_ImportedContacts {
	return &Contacts_ImportedContacts{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLContactsImportedContacts) SetImported(v []*ImportedContact) { m.Data2.Imported = v }
func (m *TLContactsImportedContacts) GetImported() []*ImportedContact  { return m.Data2.Imported }

func (m *TLContactsImportedContacts) SetPopularInvites(v []*PopularContact) {
	m.Data2.PopularInvites = v
}
func (m *TLContactsImportedContacts) GetPopularInvites() []*PopularContact {
	return m.Data2.PopularInvites
}

func (m *TLContactsImportedContacts) SetRetryContacts(v []int64) { m.Data2.RetryContacts = v }
func (m *TLContactsImportedContacts) GetRetryContacts() []int64  { return m.Data2.RetryContacts }

func (m *TLContactsImportedContacts) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLContactsImportedContacts) GetUsers() []*User  { return m.Data2.Users }

func (m *TLContactsImportedContacts) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_contacts_importedContacts))

	return x.buf
}

func (m *TLContactsImportedContacts) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Messages_SavedGifs <--
//  + TL_MessagesSavedGifsNotModified
//  + TL_MessagesSavedGifs
//

func (m *Messages_SavedGifs) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_messages_savedGifsNotModified:
		t := m.To_MessagesSavedGifsNotModified()
		return t.Encode()
	case TLConstructor_CRC32_messages_savedGifs:
		t := m.To_MessagesSavedGifs()
		return t.Encode()

	default:
		return []byte{}
	}
}

// messages.savedGifsNotModified#e8025ca2 = messages.SavedGifs;
func (m *Messages_SavedGifs) To_MessagesSavedGifsNotModified() *TLMessagesSavedGifsNotModified {
	return &TLMessagesSavedGifsNotModified{
		Data2: m.Data2,
	}
}

// messages.savedGifs#2e0709a5 hash:int gifs:Vector<Document> = messages.SavedGifs;
func (m *Messages_SavedGifs) To_MessagesSavedGifs() *TLMessagesSavedGifs {
	return &TLMessagesSavedGifs{
		Data2: m.Data2,
	}
}

// messages.savedGifsNotModified#e8025ca2 = messages.SavedGifs;
func (m *TLMessagesSavedGifsNotModified) To_Messages_SavedGifs() *Messages_SavedGifs {
	return &Messages_SavedGifs{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessagesSavedGifsNotModified) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_savedGifsNotModified))

	return x.buf
}

func (m *TLMessagesSavedGifsNotModified) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// messages.savedGifs#2e0709a5 hash:int gifs:Vector<Document> = messages.SavedGifs;
func (m *TLMessagesSavedGifs) To_Messages_SavedGifs() *Messages_SavedGifs {
	return &Messages_SavedGifs{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessagesSavedGifs) SetHash(v int32) { m.Data2.Hash = v }
func (m *TLMessagesSavedGifs) GetHash() int32  { return m.Data2.Hash }

func (m *TLMessagesSavedGifs) SetGifs(v []*Document) { m.Data2.Gifs = v }
func (m *TLMessagesSavedGifs) GetGifs() []*Document  { return m.Data2.Gifs }

func (m *TLMessagesSavedGifs) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_savedGifs))

	x.Int(m.Data2.Hash)

	return x.buf
}

func (m *TLMessagesSavedGifs) Decode(dbuf *DecodeBuf) error {
	m.Data2.Hash = x.Int()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// InputBotInlineResult <--
//  + TL_InputBotInlineResult
//  + TL_InputBotInlineResultPhoto
//  + TL_InputBotInlineResultDocument
//  + TL_InputBotInlineResultGame
//

func (m *InputBotInlineResult) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_inputBotInlineResult:
		t := m.To_InputBotInlineResult()
		return t.Encode()
	case TLConstructor_CRC32_inputBotInlineResultPhoto:
		t := m.To_InputBotInlineResultPhoto()
		return t.Encode()
	case TLConstructor_CRC32_inputBotInlineResultDocument:
		t := m.To_InputBotInlineResultDocument()
		return t.Encode()
	case TLConstructor_CRC32_inputBotInlineResultGame:
		t := m.To_InputBotInlineResultGame()
		return t.Encode()

	default:
		return []byte{}
	}
}

// inputBotInlineResult#2cbbe15a flags:# id:string type:string title:flags.1?string description:flags.2?string url:flags.3?string thumb_url:flags.4?string content_url:flags.5?string content_type:flags.5?string w:flags.6?int h:flags.6?int duration:flags.7?int send_message:InputBotInlineMessage = InputBotInlineResult;
func (m *InputBotInlineResult) To_InputBotInlineResult() *TLInputBotInlineResult {
	return &TLInputBotInlineResult{
		Data2: m.Data2,
	}
}

// inputBotInlineResultPhoto#a8d864a7 id:string type:string photo:InputPhoto send_message:InputBotInlineMessage = InputBotInlineResult;
func (m *InputBotInlineResult) To_InputBotInlineResultPhoto() *TLInputBotInlineResultPhoto {
	return &TLInputBotInlineResultPhoto{
		Data2: m.Data2,
	}
}

// inputBotInlineResultDocument#fff8fdc4 flags:# id:string type:string title:flags.1?string description:flags.2?string document:InputDocument send_message:InputBotInlineMessage = InputBotInlineResult;
func (m *InputBotInlineResult) To_InputBotInlineResultDocument() *TLInputBotInlineResultDocument {
	return &TLInputBotInlineResultDocument{
		Data2: m.Data2,
	}
}

// inputBotInlineResultGame#4fa417f2 id:string short_name:string send_message:InputBotInlineMessage = InputBotInlineResult;
func (m *InputBotInlineResult) To_InputBotInlineResultGame() *TLInputBotInlineResultGame {
	return &TLInputBotInlineResultGame{
		Data2: m.Data2,
	}
}

// inputBotInlineResult#2cbbe15a flags:# id:string type:string title:flags.1?string description:flags.2?string url:flags.3?string thumb_url:flags.4?string content_url:flags.5?string content_type:flags.5?string w:flags.6?int h:flags.6?int duration:flags.7?int send_message:InputBotInlineMessage = InputBotInlineResult;
func (m *TLInputBotInlineResult) To_InputBotInlineResult() *InputBotInlineResult {
	return &InputBotInlineResult{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputBotInlineResult) SetId(v string) { m.Data2.Id = v }
func (m *TLInputBotInlineResult) GetId() string  { return m.Data2.Id }

func (m *TLInputBotInlineResult) SetType(v string) { m.Data2.Type = v }
func (m *TLInputBotInlineResult) GetType() string  { return m.Data2.Type }

func (m *TLInputBotInlineResult) SetTitle(v string) { m.Data2.Title = v }
func (m *TLInputBotInlineResult) GetTitle() string  { return m.Data2.Title }

func (m *TLInputBotInlineResult) SetDescription(v string) { m.Data2.Description = v }
func (m *TLInputBotInlineResult) GetDescription() string  { return m.Data2.Description }

func (m *TLInputBotInlineResult) SetUrl(v string) { m.Data2.Url = v }
func (m *TLInputBotInlineResult) GetUrl() string  { return m.Data2.Url }

func (m *TLInputBotInlineResult) SetThumbUrl(v string) { m.Data2.ThumbUrl = v }
func (m *TLInputBotInlineResult) GetThumbUrl() string  { return m.Data2.ThumbUrl }

func (m *TLInputBotInlineResult) SetContentUrl(v string) { m.Data2.ContentUrl = v }
func (m *TLInputBotInlineResult) GetContentUrl() string  { return m.Data2.ContentUrl }

func (m *TLInputBotInlineResult) SetContentType(v string) { m.Data2.ContentType = v }
func (m *TLInputBotInlineResult) GetContentType() string  { return m.Data2.ContentType }

func (m *TLInputBotInlineResult) SetW(v int32) { m.Data2.W = v }
func (m *TLInputBotInlineResult) GetW() int32  { return m.Data2.W }

func (m *TLInputBotInlineResult) SetH(v int32) { m.Data2.H = v }
func (m *TLInputBotInlineResult) GetH() int32  { return m.Data2.H }

func (m *TLInputBotInlineResult) SetDuration(v int32) { m.Data2.Duration = v }
func (m *TLInputBotInlineResult) GetDuration() int32  { return m.Data2.Duration }

func (m *TLInputBotInlineResult) SetSendMessage(v *InputBotInlineMessage) { m.Data2.SendMessage = v }
func (m *TLInputBotInlineResult) GetSendMessage() *InputBotInlineMessage  { return m.Data2.SendMessage }

func (m *TLInputBotInlineResult) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputBotInlineResult))

	// flags
	var flags uint32 = 0
	if m.GetTitle() != "" {
		flags |= 1 << 3
	}
	if m.GetDescription() != "" {
		flags |= 1 << 4
	}
	if m.GetUrl() != "" {
		flags |= 1 << 5
	}
	if m.GetThumbUrl() != "" {
		flags |= 1 << 6
	}
	if m.GetContentUrl() != "" {
		flags |= 1 << 7
	}
	if m.GetContentType() != "" {
		flags |= 1 << 8
	}
	if m.GetW() != 0 {
		flags |= 1 << 9
	}
	if m.GetH() != 0 {
		flags |= 1 << 10
	}
	if m.GetDuration() != 0 {
		flags |= 1 << 11
	}
	x.UInt(flags)

	x.StringBytes(m.Data2.Id)
	x.StringBytes(m.Data2.Type)
	if m.GetTitle() != 0 {
		x.StringBytes(m.Data2.Title)
	}
	if m.GetDescription() != 0 {
		x.StringBytes(m.Data2.Description)
	}
	if m.GetUrl() != 0 {
		x.StringBytes(m.Data2.Url)
	}
	if m.GetThumbUrl() != 0 {
		x.StringBytes(m.Data2.ThumbUrl)
	}
	if m.GetContentUrl() != 0 {
		x.StringBytes(m.Data2.ContentUrl)
	}
	if m.GetContentType() != 0 {
		x.StringBytes(m.Data2.ContentType)
	}
	if m.GetW() != 0 {
		x.Int(m.Data2.W)
	}
	if m.GetH() != 0 {
		x.Int(m.Data2.H)
	}
	if m.GetDuration() != 0 {
		x.Int(m.Data2.Duration)
	}
	x.Bytes(m.Data2SendMessage.Encode())

	return x.buf
}

func (m *TLInputBotInlineResult) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	m.Data2.Id = x.StringBytes()
	m.Data2.Type = x.StringBytes()
	if (flags & (1 << 4)) != 0 {
		m.Data2.Title = x.StringBytes()
	}
	if (flags & (1 << 5)) != 0 {
		m.Data2.Description = x.StringBytes()
	}
	if (flags & (1 << 6)) != 0 {
		m.Data2.Url = x.StringBytes()
	}
	if (flags & (1 << 7)) != 0 {
		m.Data2.ThumbUrl = x.StringBytes()
	}
	if (flags & (1 << 8)) != 0 {
		m.Data2.ContentUrl = x.StringBytes()
	}
	if (flags & (1 << 9)) != 0 {
		m.Data2.ContentType = x.StringBytes()
	}
	if (flags & (1 << 10)) != 0 {
		m.Data2.W = x.Int()
	}
	if (flags & (1 << 11)) != 0 {
		m.Data2.H = x.Int()
	}
	if (flags & (1 << 12)) != 0 {
		m.Data2.Duration = x.Int()
	}
	m13 := &SendMessage{}
	m13.Decode(dbuf)
	m.SetSendMessage(m13)

	return dbuf.err
}

// inputBotInlineResultPhoto#a8d864a7 id:string type:string photo:InputPhoto send_message:InputBotInlineMessage = InputBotInlineResult;
func (m *TLInputBotInlineResultPhoto) To_InputBotInlineResult() *InputBotInlineResult {
	return &InputBotInlineResult{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputBotInlineResultPhoto) SetId(v string) { m.Data2.Id = v }
func (m *TLInputBotInlineResultPhoto) GetId() string  { return m.Data2.Id }

func (m *TLInputBotInlineResultPhoto) SetType(v string) { m.Data2.Type = v }
func (m *TLInputBotInlineResultPhoto) GetType() string  { return m.Data2.Type }

func (m *TLInputBotInlineResultPhoto) SetPhoto(v *InputPhoto) { m.Data2.Photo = v }
func (m *TLInputBotInlineResultPhoto) GetPhoto() *InputPhoto  { return m.Data2.Photo }

func (m *TLInputBotInlineResultPhoto) SetSendMessage(v *InputBotInlineMessage) {
	m.Data2.SendMessage = v
}
func (m *TLInputBotInlineResultPhoto) GetSendMessage() *InputBotInlineMessage {
	return m.Data2.SendMessage
}

func (m *TLInputBotInlineResultPhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputBotInlineResultPhoto))

	x.StringBytes(m.Data2.Id)
	x.StringBytes(m.Data2.Type)
	x.Bytes(m.Data2Photo.Encode())
	x.Bytes(m.Data2SendMessage.Encode())

	return x.buf
}

func (m *TLInputBotInlineResultPhoto) Decode(dbuf *DecodeBuf) error {
	m.Data2.Id = x.StringBytes()
	m.Data2.Type = x.StringBytes()
	m3 := &Photo{}
	m3.Decode(dbuf)
	m.SetPhoto(m3)
	m4 := &SendMessage{}
	m4.Decode(dbuf)
	m.SetSendMessage(m4)

	return dbuf.err
}

// inputBotInlineResultDocument#fff8fdc4 flags:# id:string type:string title:flags.1?string description:flags.2?string document:InputDocument send_message:InputBotInlineMessage = InputBotInlineResult;
func (m *TLInputBotInlineResultDocument) To_InputBotInlineResult() *InputBotInlineResult {
	return &InputBotInlineResult{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputBotInlineResultDocument) SetId(v string) { m.Data2.Id = v }
func (m *TLInputBotInlineResultDocument) GetId() string  { return m.Data2.Id }

func (m *TLInputBotInlineResultDocument) SetType(v string) { m.Data2.Type = v }
func (m *TLInputBotInlineResultDocument) GetType() string  { return m.Data2.Type }

func (m *TLInputBotInlineResultDocument) SetTitle(v string) { m.Data2.Title = v }
func (m *TLInputBotInlineResultDocument) GetTitle() string  { return m.Data2.Title }

func (m *TLInputBotInlineResultDocument) SetDescription(v string) { m.Data2.Description = v }
func (m *TLInputBotInlineResultDocument) GetDescription() string  { return m.Data2.Description }

func (m *TLInputBotInlineResultDocument) SetDocument(v *InputDocument) { m.Data2.Document = v }
func (m *TLInputBotInlineResultDocument) GetDocument() *InputDocument  { return m.Data2.Document }

func (m *TLInputBotInlineResultDocument) SetSendMessage(v *InputBotInlineMessage) {
	m.Data2.SendMessage = v
}
func (m *TLInputBotInlineResultDocument) GetSendMessage() *InputBotInlineMessage {
	return m.Data2.SendMessage
}

func (m *TLInputBotInlineResultDocument) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputBotInlineResultDocument))

	// flags
	var flags uint32 = 0
	if m.GetTitle() != "" {
		flags |= 1 << 3
	}
	if m.GetDescription() != "" {
		flags |= 1 << 4
	}
	x.UInt(flags)

	x.StringBytes(m.Data2.Id)
	x.StringBytes(m.Data2.Type)
	if m.GetTitle() != 0 {
		x.StringBytes(m.Data2.Title)
	}
	if m.GetDescription() != 0 {
		x.StringBytes(m.Data2.Description)
	}
	x.Bytes(m.Data2Document.Encode())
	x.Bytes(m.Data2SendMessage.Encode())

	return x.buf
}

func (m *TLInputBotInlineResultDocument) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	m.Data2.Id = x.StringBytes()
	m.Data2.Type = x.StringBytes()
	if (flags & (1 << 4)) != 0 {
		m.Data2.Title = x.StringBytes()
	}
	if (flags & (1 << 5)) != 0 {
		m.Data2.Description = x.StringBytes()
	}
	m6 := &Document{}
	m6.Decode(dbuf)
	m.SetDocument(m6)
	m7 := &SendMessage{}
	m7.Decode(dbuf)
	m.SetSendMessage(m7)

	return dbuf.err
}

// inputBotInlineResultGame#4fa417f2 id:string short_name:string send_message:InputBotInlineMessage = InputBotInlineResult;
func (m *TLInputBotInlineResultGame) To_InputBotInlineResult() *InputBotInlineResult {
	return &InputBotInlineResult{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputBotInlineResultGame) SetId(v string) { m.Data2.Id = v }
func (m *TLInputBotInlineResultGame) GetId() string  { return m.Data2.Id }

func (m *TLInputBotInlineResultGame) SetShortName(v string) { m.Data2.ShortName = v }
func (m *TLInputBotInlineResultGame) GetShortName() string  { return m.Data2.ShortName }

func (m *TLInputBotInlineResultGame) SetSendMessage(v *InputBotInlineMessage) { m.Data2.SendMessage = v }
func (m *TLInputBotInlineResultGame) GetSendMessage() *InputBotInlineMessage {
	return m.Data2.SendMessage
}

func (m *TLInputBotInlineResultGame) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputBotInlineResultGame))

	x.StringBytes(m.Data2.Id)
	x.StringBytes(m.Data2.ShortName)
	x.Bytes(m.Data2SendMessage.Encode())

	return x.buf
}

func (m *TLInputBotInlineResultGame) Decode(dbuf *DecodeBuf) error {
	m.Data2.Id = x.StringBytes()
	m.Data2.ShortName = x.StringBytes()
	m3 := &SendMessage{}
	m3.Decode(dbuf)
	m.SetSendMessage(m3)

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Messages_StickerSetInstallResult <--
//  + TL_MessagesStickerSetInstallResultSuccess
//  + TL_MessagesStickerSetInstallResultArchive
//

func (m *Messages_StickerSetInstallResult) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_messages_stickerSetInstallResultSuccess:
		t := m.To_MessagesStickerSetInstallResultSuccess()
		return t.Encode()
	case TLConstructor_CRC32_messages_stickerSetInstallResultArchive:
		t := m.To_MessagesStickerSetInstallResultArchive()
		return t.Encode()

	default:
		return []byte{}
	}
}

// messages.stickerSetInstallResultSuccess#38641628 = messages.StickerSetInstallResult;
func (m *Messages_StickerSetInstallResult) To_MessagesStickerSetInstallResultSuccess() *TLMessagesStickerSetInstallResultSuccess {
	return &TLMessagesStickerSetInstallResultSuccess{
		Data2: m.Data2,
	}
}

// messages.stickerSetInstallResultArchive#35e410a8 sets:Vector<StickerSetCovered> = messages.StickerSetInstallResult;
func (m *Messages_StickerSetInstallResult) To_MessagesStickerSetInstallResultArchive() *TLMessagesStickerSetInstallResultArchive {
	return &TLMessagesStickerSetInstallResultArchive{
		Data2: m.Data2,
	}
}

// messages.stickerSetInstallResultSuccess#38641628 = messages.StickerSetInstallResult;
func (m *TLMessagesStickerSetInstallResultSuccess) To_Messages_StickerSetInstallResult() *Messages_StickerSetInstallResult {
	return &Messages_StickerSetInstallResult{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessagesStickerSetInstallResultSuccess) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_stickerSetInstallResultSuccess))

	return x.buf
}

func (m *TLMessagesStickerSetInstallResultSuccess) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// messages.stickerSetInstallResultArchive#35e410a8 sets:Vector<StickerSetCovered> = messages.StickerSetInstallResult;
func (m *TLMessagesStickerSetInstallResultArchive) To_Messages_StickerSetInstallResult() *Messages_StickerSetInstallResult {
	return &Messages_StickerSetInstallResult{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessagesStickerSetInstallResultArchive) SetSets(v []*StickerSetCovered) { m.Data2.Sets = v }
func (m *TLMessagesStickerSetInstallResultArchive) GetSets() []*StickerSetCovered  { return m.Data2.Sets }

func (m *TLMessagesStickerSetInstallResultArchive) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_stickerSetInstallResultArchive))

	return x.buf
}

func (m *TLMessagesStickerSetInstallResultArchive) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// PrivacyRule <--
//  + TL_PrivacyValueAllowContacts
//  + TL_PrivacyValueAllowAll
//  + TL_PrivacyValueAllowUsers
//  + TL_PrivacyValueDisallowContacts
//  + TL_PrivacyValueDisallowAll
//  + TL_PrivacyValueDisallowUsers
//

func (m *PrivacyRule) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_privacyValueAllowContacts:
		t := m.To_PrivacyValueAllowContacts()
		return t.Encode()
	case TLConstructor_CRC32_privacyValueAllowAll:
		t := m.To_PrivacyValueAllowAll()
		return t.Encode()
	case TLConstructor_CRC32_privacyValueAllowUsers:
		t := m.To_PrivacyValueAllowUsers()
		return t.Encode()
	case TLConstructor_CRC32_privacyValueDisallowContacts:
		t := m.To_PrivacyValueDisallowContacts()
		return t.Encode()
	case TLConstructor_CRC32_privacyValueDisallowAll:
		t := m.To_PrivacyValueDisallowAll()
		return t.Encode()
	case TLConstructor_CRC32_privacyValueDisallowUsers:
		t := m.To_PrivacyValueDisallowUsers()
		return t.Encode()

	default:
		return []byte{}
	}
}

// privacyValueAllowContacts#fffe1bac = PrivacyRule;
func (m *PrivacyRule) To_PrivacyValueAllowContacts() *TLPrivacyValueAllowContacts {
	return &TLPrivacyValueAllowContacts{
		Data2: m.Data2,
	}
}

// privacyValueAllowAll#65427b82 = PrivacyRule;
func (m *PrivacyRule) To_PrivacyValueAllowAll() *TLPrivacyValueAllowAll {
	return &TLPrivacyValueAllowAll{
		Data2: m.Data2,
	}
}

// privacyValueAllowUsers#4d5bbe0c users:Vector<int> = PrivacyRule;
func (m *PrivacyRule) To_PrivacyValueAllowUsers() *TLPrivacyValueAllowUsers {
	return &TLPrivacyValueAllowUsers{
		Data2: m.Data2,
	}
}

// privacyValueDisallowContacts#f888fa1a = PrivacyRule;
func (m *PrivacyRule) To_PrivacyValueDisallowContacts() *TLPrivacyValueDisallowContacts {
	return &TLPrivacyValueDisallowContacts{
		Data2: m.Data2,
	}
}

// privacyValueDisallowAll#8b73e763 = PrivacyRule;
func (m *PrivacyRule) To_PrivacyValueDisallowAll() *TLPrivacyValueDisallowAll {
	return &TLPrivacyValueDisallowAll{
		Data2: m.Data2,
	}
}

// privacyValueDisallowUsers#c7f49b7 users:Vector<int> = PrivacyRule;
func (m *PrivacyRule) To_PrivacyValueDisallowUsers() *TLPrivacyValueDisallowUsers {
	return &TLPrivacyValueDisallowUsers{
		Data2: m.Data2,
	}
}

// privacyValueAllowContacts#fffe1bac = PrivacyRule;
func (m *TLPrivacyValueAllowContacts) To_PrivacyRule() *PrivacyRule {
	return &PrivacyRule{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPrivacyValueAllowContacts) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_privacyValueAllowContacts))

	return x.buf
}

func (m *TLPrivacyValueAllowContacts) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// privacyValueAllowAll#65427b82 = PrivacyRule;
func (m *TLPrivacyValueAllowAll) To_PrivacyRule() *PrivacyRule {
	return &PrivacyRule{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPrivacyValueAllowAll) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_privacyValueAllowAll))

	return x.buf
}

func (m *TLPrivacyValueAllowAll) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// privacyValueAllowUsers#4d5bbe0c users:Vector<int> = PrivacyRule;
func (m *TLPrivacyValueAllowUsers) To_PrivacyRule() *PrivacyRule {
	return &PrivacyRule{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPrivacyValueAllowUsers) SetUsers(v []int32) { m.Data2.Users = v }
func (m *TLPrivacyValueAllowUsers) GetUsers() []int32  { return m.Data2.Users }

func (m *TLPrivacyValueAllowUsers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_privacyValueAllowUsers))

	return x.buf
}

func (m *TLPrivacyValueAllowUsers) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// privacyValueDisallowContacts#f888fa1a = PrivacyRule;
func (m *TLPrivacyValueDisallowContacts) To_PrivacyRule() *PrivacyRule {
	return &PrivacyRule{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPrivacyValueDisallowContacts) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_privacyValueDisallowContacts))

	return x.buf
}

func (m *TLPrivacyValueDisallowContacts) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// privacyValueDisallowAll#8b73e763 = PrivacyRule;
func (m *TLPrivacyValueDisallowAll) To_PrivacyRule() *PrivacyRule {
	return &PrivacyRule{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPrivacyValueDisallowAll) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_privacyValueDisallowAll))

	return x.buf
}

func (m *TLPrivacyValueDisallowAll) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// privacyValueDisallowUsers#c7f49b7 users:Vector<int> = PrivacyRule;
func (m *TLPrivacyValueDisallowUsers) To_PrivacyRule() *PrivacyRule {
	return &PrivacyRule{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPrivacyValueDisallowUsers) SetUsers(v []int32) { m.Data2.Users = v }
func (m *TLPrivacyValueDisallowUsers) GetUsers() []int32  { return m.Data2.Users }

func (m *TLPrivacyValueDisallowUsers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_privacyValueDisallowUsers))

	return x.buf
}

func (m *TLPrivacyValueDisallowUsers) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// ExportedMessageLink <--
//  + TL_ExportedMessageLink
//

func (m *ExportedMessageLink) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_exportedMessageLink:
		t := m.To_ExportedMessageLink()
		return t.Encode()

	default:
		return []byte{}
	}
}

// exportedMessageLink#1f486803 link:string = ExportedMessageLink;
func (m *ExportedMessageLink) To_ExportedMessageLink() *TLExportedMessageLink {
	return &TLExportedMessageLink{
		Data2: m.Data2,
	}
}

// exportedMessageLink#1f486803 link:string = ExportedMessageLink;
func (m *TLExportedMessageLink) To_ExportedMessageLink() *ExportedMessageLink {
	return &ExportedMessageLink{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLExportedMessageLink) SetLink(v string) { m.Data2.Link = v }
func (m *TLExportedMessageLink) GetLink() string  { return m.Data2.Link }

func (m *TLExportedMessageLink) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_exportedMessageLink))

	x.StringBytes(m.Data2.Link)

	return x.buf
}

func (m *TLExportedMessageLink) Decode(dbuf *DecodeBuf) error {
	m.Data2.Link = x.StringBytes()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// InputAppEvent <--
//  + TL_InputAppEvent
//

func (m *InputAppEvent) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_inputAppEvent:
		t := m.To_InputAppEvent()
		return t.Encode()

	default:
		return []byte{}
	}
}

// inputAppEvent#770656a8 time:double type:string peer:long data:string = InputAppEvent;
func (m *InputAppEvent) To_InputAppEvent() *TLInputAppEvent {
	return &TLInputAppEvent{
		Data2: m.Data2,
	}
}

// inputAppEvent#770656a8 time:double type:string peer:long data:string = InputAppEvent;
func (m *TLInputAppEvent) To_InputAppEvent() *InputAppEvent {
	return &InputAppEvent{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputAppEvent) SetTime(v float64) { m.Data2.Time = v }
func (m *TLInputAppEvent) GetTime() float64  { return m.Data2.Time }

func (m *TLInputAppEvent) SetType(v string) { m.Data2.Type = v }
func (m *TLInputAppEvent) GetType() string  { return m.Data2.Type }

func (m *TLInputAppEvent) SetPeer(v int64) { m.Data2.Peer = v }
func (m *TLInputAppEvent) GetPeer() int64  { return m.Data2.Peer }

func (m *TLInputAppEvent) SetData(v string) { m.Data2.Data = v }
func (m *TLInputAppEvent) GetData() string  { return m.Data2.Data }

func (m *TLInputAppEvent) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputAppEvent))

	x.Double(m.Data2.Time)
	x.StringBytes(m.Data2.Type)
	x.Long(m.Data2.Peer)
	x.StringBytes(m.Data2.Data)

	return x.buf
}

func (m *TLInputAppEvent) Decode(dbuf *DecodeBuf) error {
	m.Data2.Time = x.Double()
	m.Data2.Type = x.StringBytes()
	m.Data2.Peer = x.Long()
	m.Data2.Data = x.StringBytes()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Auth_CheckedPhone <--
//  + TL_AuthCheckedPhone
//

func (m *Auth_CheckedPhone) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_auth_checkedPhone:
		t := m.To_AuthCheckedPhone()
		return t.Encode()

	default:
		return []byte{}
	}
}

// auth.checkedPhone#811ea28e phone_registered:Bool = auth.CheckedPhone;
func (m *Auth_CheckedPhone) To_AuthCheckedPhone() *TLAuthCheckedPhone {
	return &TLAuthCheckedPhone{
		Data2: m.Data2,
	}
}

// auth.checkedPhone#811ea28e phone_registered:Bool = auth.CheckedPhone;
func (m *TLAuthCheckedPhone) To_Auth_CheckedPhone() *Auth_CheckedPhone {
	return &Auth_CheckedPhone{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLAuthCheckedPhone) SetPhoneRegistered(v *Bool) { m.Data2.PhoneRegistered = v }
func (m *TLAuthCheckedPhone) GetPhoneRegistered() *Bool  { return m.Data2.PhoneRegistered }

func (m *TLAuthCheckedPhone) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_auth_checkedPhone))

	x.Bytes(m.Data2PhoneRegistered.Encode())

	return x.buf
}

func (m *TLAuthCheckedPhone) Decode(dbuf *DecodeBuf) error {
	m1 := &PhoneRegistered{}
	m1.Decode(dbuf)
	m.SetPhoneRegistered(m1)

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// InputEncryptedChat <--
//  + TL_InputEncryptedChat
//

func (m *InputEncryptedChat) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_inputEncryptedChat:
		t := m.To_InputEncryptedChat()
		return t.Encode()

	default:
		return []byte{}
	}
}

// inputEncryptedChat#f141b5e1 chat_id:int access_hash:long = InputEncryptedChat;
func (m *InputEncryptedChat) To_InputEncryptedChat() *TLInputEncryptedChat {
	return &TLInputEncryptedChat{
		Data2: m.Data2,
	}
}

// inputEncryptedChat#f141b5e1 chat_id:int access_hash:long = InputEncryptedChat;
func (m *TLInputEncryptedChat) To_InputEncryptedChat() *InputEncryptedChat {
	return &InputEncryptedChat{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputEncryptedChat) SetChatId(v int32) { m.Data2.ChatId = v }
func (m *TLInputEncryptedChat) GetChatId() int32  { return m.Data2.ChatId }

func (m *TLInputEncryptedChat) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLInputEncryptedChat) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLInputEncryptedChat) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputEncryptedChat))

	x.Int(m.Data2.ChatId)
	x.Long(m.Data2.AccessHash)

	return x.buf
}

func (m *TLInputEncryptedChat) Decode(dbuf *DecodeBuf) error {
	m.Data2.ChatId = x.Int()
	m.Data2.AccessHash = x.Long()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Document <--
//  + TL_DocumentEmpty
//  + TL_Document
//

func (m *Document) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_documentEmpty:
		t := m.To_DocumentEmpty()
		return t.Encode()
	case TLConstructor_CRC32_document:
		t := m.To_Document()
		return t.Encode()

	default:
		return []byte{}
	}
}

// documentEmpty#36f8c871 id:long = Document;
func (m *Document) To_DocumentEmpty() *TLDocumentEmpty {
	return &TLDocumentEmpty{
		Data2: m.Data2,
	}
}

// document#87232bc7 id:long access_hash:long date:int mime_type:string size:int thumb:PhotoSize dc_id:int version:int attributes:Vector<DocumentAttribute> = Document;
func (m *Document) To_Document() *TLDocument {
	return &TLDocument{
		Data2: m.Data2,
	}
}

// documentEmpty#36f8c871 id:long = Document;
func (m *TLDocumentEmpty) To_Document() *Document {
	return &Document{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLDocumentEmpty) SetId(v int64) { m.Data2.Id = v }
func (m *TLDocumentEmpty) GetId() int64  { return m.Data2.Id }

func (m *TLDocumentEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_documentEmpty))

	x.Long(m.Data2.Id)

	return x.buf
}

func (m *TLDocumentEmpty) Decode(dbuf *DecodeBuf) error {
	m.Data2.Id = x.Long()

	return dbuf.err
}

// document#87232bc7 id:long access_hash:long date:int mime_type:string size:int thumb:PhotoSize dc_id:int version:int attributes:Vector<DocumentAttribute> = Document;
func (m *TLDocument) To_Document() *Document {
	return &Document{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLDocument) SetId(v int64) { m.Data2.Id = v }
func (m *TLDocument) GetId() int64  { return m.Data2.Id }

func (m *TLDocument) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLDocument) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLDocument) SetDate(v int32) { m.Data2.Date = v }
func (m *TLDocument) GetDate() int32  { return m.Data2.Date }

func (m *TLDocument) SetMimeType(v string) { m.Data2.MimeType = v }
func (m *TLDocument) GetMimeType() string  { return m.Data2.MimeType }

func (m *TLDocument) SetSize(v int32) { m.Data2.Size = v }
func (m *TLDocument) GetSize() int32  { return m.Data2.Size }

func (m *TLDocument) SetThumb(v *PhotoSize) { m.Data2.Thumb = v }
func (m *TLDocument) GetThumb() *PhotoSize  { return m.Data2.Thumb }

func (m *TLDocument) SetDcId(v int32) { m.Data2.DcId = v }
func (m *TLDocument) GetDcId() int32  { return m.Data2.DcId }

func (m *TLDocument) SetVersion(v int32) { m.Data2.Version = v }
func (m *TLDocument) GetVersion() int32  { return m.Data2.Version }

func (m *TLDocument) SetAttributes(v []*DocumentAttribute) { m.Data2.Attributes = v }
func (m *TLDocument) GetAttributes() []*DocumentAttribute  { return m.Data2.Attributes }

func (m *TLDocument) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_document))

	x.Long(m.Data2.Id)
	x.Long(m.Data2.AccessHash)
	x.Int(m.Data2.Date)
	x.StringBytes(m.Data2.MimeType)
	x.Int(m.Data2.Size)
	x.Bytes(m.Data2Thumb.Encode())
	x.Int(m.Data2.DcId)
	x.Int(m.Data2.Version)

	return x.buf
}

func (m *TLDocument) Decode(dbuf *DecodeBuf) error {
	m.Data2.Id = x.Long()
	m.Data2.AccessHash = x.Long()
	m.Data2.Date = x.Int()
	m.Data2.MimeType = x.StringBytes()
	m.Data2.Size = x.Int()
	m6 := &Thumb{}
	m6.Decode(dbuf)
	m.SetThumb(m6)
	m.Data2.DcId = x.Int()
	m.Data2.Version = x.Int()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Messages_PeerDialogs <--
//  + TL_MessagesPeerDialogs
//

func (m *Messages_PeerDialogs) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_messages_peerDialogs:
		t := m.To_MessagesPeerDialogs()
		return t.Encode()

	default:
		return []byte{}
	}
}

// messages.peerDialogs#3371c354 dialogs:Vector<Dialog> messages:Vector<Message> chats:Vector<Chat> users:Vector<User> state:updates.State = messages.PeerDialogs;
func (m *Messages_PeerDialogs) To_MessagesPeerDialogs() *TLMessagesPeerDialogs {
	return &TLMessagesPeerDialogs{
		Data2: m.Data2,
	}
}

// messages.peerDialogs#3371c354 dialogs:Vector<Dialog> messages:Vector<Message> chats:Vector<Chat> users:Vector<User> state:updates.State = messages.PeerDialogs;
func (m *TLMessagesPeerDialogs) To_Messages_PeerDialogs() *Messages_PeerDialogs {
	return &Messages_PeerDialogs{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessagesPeerDialogs) SetDialogs(v []*Dialog) { m.Data2.Dialogs = v }
func (m *TLMessagesPeerDialogs) GetDialogs() []*Dialog  { return m.Data2.Dialogs }

func (m *TLMessagesPeerDialogs) SetMessages(v []*Message) { m.Data2.Messages = v }
func (m *TLMessagesPeerDialogs) GetMessages() []*Message  { return m.Data2.Messages }

func (m *TLMessagesPeerDialogs) SetChats(v []*Chat) { m.Data2.Chats = v }
func (m *TLMessagesPeerDialogs) GetChats() []*Chat  { return m.Data2.Chats }

func (m *TLMessagesPeerDialogs) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLMessagesPeerDialogs) GetUsers() []*User  { return m.Data2.Users }

func (m *TLMessagesPeerDialogs) SetState(v *Updates_State) { m.Data2.State = v }
func (m *TLMessagesPeerDialogs) GetState() *Updates_State  { return m.Data2.State }

func (m *TLMessagesPeerDialogs) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_peerDialogs))

	x.Bytes(m.Data2State.Encode())

	return x.buf
}

func (m *TLMessagesPeerDialogs) Decode(dbuf *DecodeBuf) error {

	m5 := &State{}
	m5.Decode(dbuf)
	m.SetState(m5)

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Messages_Chats <--
//  + TL_MessagesChats
//  + TL_MessagesChatsSlice
//

func (m *Messages_Chats) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_messages_chats:
		t := m.To_MessagesChats()
		return t.Encode()
	case TLConstructor_CRC32_messages_chatsSlice:
		t := m.To_MessagesChatsSlice()
		return t.Encode()

	default:
		return []byte{}
	}
}

// messages.chats#64ff9fd5 chats:Vector<Chat> = messages.Chats;
func (m *Messages_Chats) To_MessagesChats() *TLMessagesChats {
	return &TLMessagesChats{
		Data2: m.Data2,
	}
}

// messages.chatsSlice#9cd81144 count:int chats:Vector<Chat> = messages.Chats;
func (m *Messages_Chats) To_MessagesChatsSlice() *TLMessagesChatsSlice {
	return &TLMessagesChatsSlice{
		Data2: m.Data2,
	}
}

// messages.chats#64ff9fd5 chats:Vector<Chat> = messages.Chats;
func (m *TLMessagesChats) To_Messages_Chats() *Messages_Chats {
	return &Messages_Chats{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessagesChats) SetChats(v []*Chat) { m.Data2.Chats = v }
func (m *TLMessagesChats) GetChats() []*Chat  { return m.Data2.Chats }

func (m *TLMessagesChats) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_chats))

	return x.buf
}

func (m *TLMessagesChats) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// messages.chatsSlice#9cd81144 count:int chats:Vector<Chat> = messages.Chats;
func (m *TLMessagesChatsSlice) To_Messages_Chats() *Messages_Chats {
	return &Messages_Chats{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessagesChatsSlice) SetCount(v int32) { m.Data2.Count = v }
func (m *TLMessagesChatsSlice) GetCount() int32  { return m.Data2.Count }

func (m *TLMessagesChatsSlice) SetChats(v []*Chat) { m.Data2.Chats = v }
func (m *TLMessagesChatsSlice) GetChats() []*Chat  { return m.Data2.Chats }

func (m *TLMessagesChatsSlice) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_chatsSlice))

	x.Int(m.Data2.Count)

	return x.buf
}

func (m *TLMessagesChatsSlice) Decode(dbuf *DecodeBuf) error {
	m.Data2.Count = x.Int()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// PrivacyKey <--
//  + TL_PrivacyKeyStatusTimestamp
//  + TL_PrivacyKeyChatInvite
//  + TL_PrivacyKeyPhoneCall
//

func (m *PrivacyKey) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_privacyKeyStatusTimestamp:
		t := m.To_PrivacyKeyStatusTimestamp()
		return t.Encode()
	case TLConstructor_CRC32_privacyKeyChatInvite:
		t := m.To_PrivacyKeyChatInvite()
		return t.Encode()
	case TLConstructor_CRC32_privacyKeyPhoneCall:
		t := m.To_PrivacyKeyPhoneCall()
		return t.Encode()

	default:
		return []byte{}
	}
}

// privacyKeyStatusTimestamp#bc2eab30 = PrivacyKey;
func (m *PrivacyKey) To_PrivacyKeyStatusTimestamp() *TLPrivacyKeyStatusTimestamp {
	return &TLPrivacyKeyStatusTimestamp{
		Data2: m.Data2,
	}
}

// privacyKeyChatInvite#500e6dfa = PrivacyKey;
func (m *PrivacyKey) To_PrivacyKeyChatInvite() *TLPrivacyKeyChatInvite {
	return &TLPrivacyKeyChatInvite{
		Data2: m.Data2,
	}
}

// privacyKeyPhoneCall#3d662b7b = PrivacyKey;
func (m *PrivacyKey) To_PrivacyKeyPhoneCall() *TLPrivacyKeyPhoneCall {
	return &TLPrivacyKeyPhoneCall{
		Data2: m.Data2,
	}
}

// privacyKeyStatusTimestamp#bc2eab30 = PrivacyKey;
func (m *TLPrivacyKeyStatusTimestamp) To_PrivacyKey() *PrivacyKey {
	return &PrivacyKey{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPrivacyKeyStatusTimestamp) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_privacyKeyStatusTimestamp))

	return x.buf
}

func (m *TLPrivacyKeyStatusTimestamp) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// privacyKeyChatInvite#500e6dfa = PrivacyKey;
func (m *TLPrivacyKeyChatInvite) To_PrivacyKey() *PrivacyKey {
	return &PrivacyKey{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPrivacyKeyChatInvite) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_privacyKeyChatInvite))

	return x.buf
}

func (m *TLPrivacyKeyChatInvite) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// privacyKeyPhoneCall#3d662b7b = PrivacyKey;
func (m *TLPrivacyKeyPhoneCall) To_PrivacyKey() *PrivacyKey {
	return &PrivacyKey{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPrivacyKeyPhoneCall) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_privacyKeyPhoneCall))

	return x.buf
}

func (m *TLPrivacyKeyPhoneCall) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// DocumentAttribute <--
//  + TL_DocumentAttributeImageSize
//  + TL_DocumentAttributeAnimated
//  + TL_DocumentAttributeSticker
//  + TL_DocumentAttributeVideo
//  + TL_DocumentAttributeAudio
//  + TL_DocumentAttributeFilename
//  + TL_DocumentAttributeHasStickers
//

func (m *DocumentAttribute) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_documentAttributeImageSize:
		t := m.To_DocumentAttributeImageSize()
		return t.Encode()
	case TLConstructor_CRC32_documentAttributeAnimated:
		t := m.To_DocumentAttributeAnimated()
		return t.Encode()
	case TLConstructor_CRC32_documentAttributeSticker:
		t := m.To_DocumentAttributeSticker()
		return t.Encode()
	case TLConstructor_CRC32_documentAttributeVideo:
		t := m.To_DocumentAttributeVideo()
		return t.Encode()
	case TLConstructor_CRC32_documentAttributeAudio:
		t := m.To_DocumentAttributeAudio()
		return t.Encode()
	case TLConstructor_CRC32_documentAttributeFilename:
		t := m.To_DocumentAttributeFilename()
		return t.Encode()
	case TLConstructor_CRC32_documentAttributeHasStickers:
		t := m.To_DocumentAttributeHasStickers()
		return t.Encode()

	default:
		return []byte{}
	}
}

// documentAttributeImageSize#6c37c15c w:int h:int = DocumentAttribute;
func (m *DocumentAttribute) To_DocumentAttributeImageSize() *TLDocumentAttributeImageSize {
	return &TLDocumentAttributeImageSize{
		Data2: m.Data2,
	}
}

// documentAttributeAnimated#11b58939 = DocumentAttribute;
func (m *DocumentAttribute) To_DocumentAttributeAnimated() *TLDocumentAttributeAnimated {
	return &TLDocumentAttributeAnimated{
		Data2: m.Data2,
	}
}

// documentAttributeSticker#6319d612 flags:# mask:flags.1?true alt:string stickerset:InputStickerSet mask_coords:flags.0?MaskCoords = DocumentAttribute;
func (m *DocumentAttribute) To_DocumentAttributeSticker() *TLDocumentAttributeSticker {
	return &TLDocumentAttributeSticker{
		Data2: m.Data2,
	}
}

// documentAttributeVideo#ef02ce6 flags:# round_message:flags.0?true duration:int w:int h:int = DocumentAttribute;
func (m *DocumentAttribute) To_DocumentAttributeVideo() *TLDocumentAttributeVideo {
	return &TLDocumentAttributeVideo{
		Data2: m.Data2,
	}
}

// documentAttributeAudio#9852f9c6 flags:# voice:flags.10?true duration:int title:flags.0?string performer:flags.1?string waveform:flags.2?bytes = DocumentAttribute;
func (m *DocumentAttribute) To_DocumentAttributeAudio() *TLDocumentAttributeAudio {
	return &TLDocumentAttributeAudio{
		Data2: m.Data2,
	}
}

// documentAttributeFilename#15590068 file_name:string = DocumentAttribute;
func (m *DocumentAttribute) To_DocumentAttributeFilename() *TLDocumentAttributeFilename {
	return &TLDocumentAttributeFilename{
		Data2: m.Data2,
	}
}

// documentAttributeHasStickers#9801d2f7 = DocumentAttribute;
func (m *DocumentAttribute) To_DocumentAttributeHasStickers() *TLDocumentAttributeHasStickers {
	return &TLDocumentAttributeHasStickers{
		Data2: m.Data2,
	}
}

// documentAttributeImageSize#6c37c15c w:int h:int = DocumentAttribute;
func (m *TLDocumentAttributeImageSize) To_DocumentAttribute() *DocumentAttribute {
	return &DocumentAttribute{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLDocumentAttributeImageSize) SetW(v int32) { m.Data2.W = v }
func (m *TLDocumentAttributeImageSize) GetW() int32  { return m.Data2.W }

func (m *TLDocumentAttributeImageSize) SetH(v int32) { m.Data2.H = v }
func (m *TLDocumentAttributeImageSize) GetH() int32  { return m.Data2.H }

func (m *TLDocumentAttributeImageSize) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_documentAttributeImageSize))

	x.Int(m.Data2.W)
	x.Int(m.Data2.H)

	return x.buf
}

func (m *TLDocumentAttributeImageSize) Decode(dbuf *DecodeBuf) error {
	m.Data2.W = x.Int()
	m.Data2.H = x.Int()

	return dbuf.err
}

// documentAttributeAnimated#11b58939 = DocumentAttribute;
func (m *TLDocumentAttributeAnimated) To_DocumentAttribute() *DocumentAttribute {
	return &DocumentAttribute{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLDocumentAttributeAnimated) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_documentAttributeAnimated))

	return x.buf
}

func (m *TLDocumentAttributeAnimated) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// documentAttributeSticker#6319d612 flags:# mask:flags.1?true alt:string stickerset:InputStickerSet mask_coords:flags.0?MaskCoords = DocumentAttribute;
func (m *TLDocumentAttributeSticker) To_DocumentAttribute() *DocumentAttribute {
	return &DocumentAttribute{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLDocumentAttributeSticker) SetMask(v bool) { m.Data2.Mask = v }
func (m *TLDocumentAttributeSticker) GetMask() bool  { return m.Data2.Mask }

func (m *TLDocumentAttributeSticker) SetAlt(v string) { m.Data2.Alt = v }
func (m *TLDocumentAttributeSticker) GetAlt() string  { return m.Data2.Alt }

func (m *TLDocumentAttributeSticker) SetStickerset(v *InputStickerSet) { m.Data2.Stickerset = v }
func (m *TLDocumentAttributeSticker) GetStickerset() *InputStickerSet  { return m.Data2.Stickerset }

func (m *TLDocumentAttributeSticker) SetMaskCoords(v *MaskCoords) { m.Data2.MaskCoords = v }
func (m *TLDocumentAttributeSticker) GetMaskCoords() *MaskCoords  { return m.Data2.MaskCoords }

func (m *TLDocumentAttributeSticker) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_documentAttributeSticker))

	// flags
	var flags uint32 = 0
	if m.GetMask() == true {
		flags |= 1 << 1
	}
	if m.GetMaskCoords() != nil {
		flags |= 1 << 4
	}
	x.UInt(flags)

	x.StringBytes(m.Data2.Alt)
	x.Bytes(m.Data2Stickerset.Encode())
	if m.GetMaskCoords() != 0 {
		x.Bytes(m.Data2MaskCoords.Encode())
	}

	return x.buf
}

func (m *TLDocumentAttributeSticker) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m.Data2.Mask = true
	}
	m.Data2.Alt = x.StringBytes()
	m4 := &Stickerset{}
	m4.Decode(dbuf)
	m.SetStickerset(m4)
	if (flags & (1 << 5)) != 0 {
		m5 := &MaskCoords{}
		m5.Decode(dbuf)
		m.SetMaskCoords(m5)
	}

	return dbuf.err
}

// documentAttributeVideo#ef02ce6 flags:# round_message:flags.0?true duration:int w:int h:int = DocumentAttribute;
func (m *TLDocumentAttributeVideo) To_DocumentAttribute() *DocumentAttribute {
	return &DocumentAttribute{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLDocumentAttributeVideo) SetRoundMessage(v bool) { m.Data2.RoundMessage = v }
func (m *TLDocumentAttributeVideo) GetRoundMessage() bool  { return m.Data2.RoundMessage }

func (m *TLDocumentAttributeVideo) SetDuration(v int32) { m.Data2.Duration = v }
func (m *TLDocumentAttributeVideo) GetDuration() int32  { return m.Data2.Duration }

func (m *TLDocumentAttributeVideo) SetW(v int32) { m.Data2.W = v }
func (m *TLDocumentAttributeVideo) GetW() int32  { return m.Data2.W }

func (m *TLDocumentAttributeVideo) SetH(v int32) { m.Data2.H = v }
func (m *TLDocumentAttributeVideo) GetH() int32  { return m.Data2.H }

func (m *TLDocumentAttributeVideo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_documentAttributeVideo))

	// flags
	var flags uint32 = 0
	if m.GetRoundMessage() == true {
		flags |= 1 << 1
	}
	x.UInt(flags)

	x.Int(m.Data2.Duration)
	x.Int(m.Data2.W)
	x.Int(m.Data2.H)

	return x.buf
}

func (m *TLDocumentAttributeVideo) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m.Data2.RoundMessage = true
	}
	m.Data2.Duration = x.Int()
	m.Data2.W = x.Int()
	m.Data2.H = x.Int()

	return dbuf.err
}

// documentAttributeAudio#9852f9c6 flags:# voice:flags.10?true duration:int title:flags.0?string performer:flags.1?string waveform:flags.2?bytes = DocumentAttribute;
func (m *TLDocumentAttributeAudio) To_DocumentAttribute() *DocumentAttribute {
	return &DocumentAttribute{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLDocumentAttributeAudio) SetVoice(v bool) { m.Data2.Voice = v }
func (m *TLDocumentAttributeAudio) GetVoice() bool  { return m.Data2.Voice }

func (m *TLDocumentAttributeAudio) SetDuration(v int32) { m.Data2.Duration = v }
func (m *TLDocumentAttributeAudio) GetDuration() int32  { return m.Data2.Duration }

func (m *TLDocumentAttributeAudio) SetTitle(v string) { m.Data2.Title = v }
func (m *TLDocumentAttributeAudio) GetTitle() string  { return m.Data2.Title }

func (m *TLDocumentAttributeAudio) SetPerformer(v string) { m.Data2.Performer = v }
func (m *TLDocumentAttributeAudio) GetPerformer() string  { return m.Data2.Performer }

func (m *TLDocumentAttributeAudio) SetWaveform(v []byte) { m.Data2.Waveform = v }
func (m *TLDocumentAttributeAudio) GetWaveform() []byte  { return m.Data2.Waveform }

func (m *TLDocumentAttributeAudio) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_documentAttributeAudio))

	// flags
	var flags uint32 = 0
	if m.GetVoice() == true {
		flags |= 1 << 1
	}
	if m.GetTitle() != "" {
		flags |= 1 << 3
	}
	if m.GetPerformer() != "" {
		flags |= 1 << 4
	}
	if m.GetWaveform() != nil {
		flags |= 1 << 5
	}
	x.UInt(flags)

	x.Int(m.Data2.Duration)
	if m.GetTitle() != 0 {
		x.StringBytes(m.Data2.Title)
	}
	if m.GetPerformer() != 0 {
		x.StringBytes(m.Data2.Performer)
	}
	if m.GetWaveform() != 0 {
		x.StringBytes(m.Data2.Waveform)
	}

	return x.buf
}

func (m *TLDocumentAttributeAudio) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m.Data2.Voice = true
	}
	m.Data2.Duration = x.Int()
	if (flags & (1 << 4)) != 0 {
		m.Data2.Title = x.StringBytes()
	}
	if (flags & (1 << 5)) != 0 {
		m.Data2.Performer = x.StringBytes()
	}
	if (flags & (1 << 6)) != 0 {
		m.Data2.Waveform = x.StringBytes()
	}

	return dbuf.err
}

// documentAttributeFilename#15590068 file_name:string = DocumentAttribute;
func (m *TLDocumentAttributeFilename) To_DocumentAttribute() *DocumentAttribute {
	return &DocumentAttribute{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLDocumentAttributeFilename) SetFileName(v string) { m.Data2.FileName = v }
func (m *TLDocumentAttributeFilename) GetFileName() string  { return m.Data2.FileName }

func (m *TLDocumentAttributeFilename) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_documentAttributeFilename))

	x.StringBytes(m.Data2.FileName)

	return x.buf
}

func (m *TLDocumentAttributeFilename) Decode(dbuf *DecodeBuf) error {
	m.Data2.FileName = x.StringBytes()

	return dbuf.err
}

// documentAttributeHasStickers#9801d2f7 = DocumentAttribute;
func (m *TLDocumentAttributeHasStickers) To_DocumentAttribute() *DocumentAttribute {
	return &DocumentAttribute{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLDocumentAttributeHasStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_documentAttributeHasStickers))

	return x.buf
}

func (m *TLDocumentAttributeHasStickers) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// ChannelMessagesFilter <--
//  + TL_ChannelMessagesFilterEmpty
//  + TL_ChannelMessagesFilter
//

func (m *ChannelMessagesFilter) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_channelMessagesFilterEmpty:
		t := m.To_ChannelMessagesFilterEmpty()
		return t.Encode()
	case TLConstructor_CRC32_channelMessagesFilter:
		t := m.To_ChannelMessagesFilter()
		return t.Encode()

	default:
		return []byte{}
	}
}

// channelMessagesFilterEmpty#94d42ee7 = ChannelMessagesFilter;
func (m *ChannelMessagesFilter) To_ChannelMessagesFilterEmpty() *TLChannelMessagesFilterEmpty {
	return &TLChannelMessagesFilterEmpty{
		Data2: m.Data2,
	}
}

// channelMessagesFilter#cd77d957 flags:# exclude_new_messages:flags.1?true ranges:Vector<MessageRange> = ChannelMessagesFilter;
func (m *ChannelMessagesFilter) To_ChannelMessagesFilter() *TLChannelMessagesFilter {
	return &TLChannelMessagesFilter{
		Data2: m.Data2,
	}
}

// channelMessagesFilterEmpty#94d42ee7 = ChannelMessagesFilter;
func (m *TLChannelMessagesFilterEmpty) To_ChannelMessagesFilter() *ChannelMessagesFilter {
	return &ChannelMessagesFilter{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLChannelMessagesFilterEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelMessagesFilterEmpty))

	return x.buf
}

func (m *TLChannelMessagesFilterEmpty) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// channelMessagesFilter#cd77d957 flags:# exclude_new_messages:flags.1?true ranges:Vector<MessageRange> = ChannelMessagesFilter;
func (m *TLChannelMessagesFilter) To_ChannelMessagesFilter() *ChannelMessagesFilter {
	return &ChannelMessagesFilter{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLChannelMessagesFilter) SetExcludeNewMessages(v bool) { m.Data2.ExcludeNewMessages = v }
func (m *TLChannelMessagesFilter) GetExcludeNewMessages() bool  { return m.Data2.ExcludeNewMessages }

func (m *TLChannelMessagesFilter) SetRanges(v []*MessageRange) { m.Data2.Ranges = v }
func (m *TLChannelMessagesFilter) GetRanges() []*MessageRange  { return m.Data2.Ranges }

func (m *TLChannelMessagesFilter) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelMessagesFilter))

	// flags
	var flags uint32 = 0
	if m.GetExcludeNewMessages() == true {
		flags |= 1 << 1
	}
	x.UInt(flags)

	return x.buf
}

func (m *TLChannelMessagesFilter) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m.Data2.ExcludeNewMessages = true
	}

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// WallPaper <--
//  + TL_WallPaper
//  + TL_WallPaperSolid
//

func (m *WallPaper) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_wallPaper:
		t := m.To_WallPaper()
		return t.Encode()
	case TLConstructor_CRC32_wallPaperSolid:
		t := m.To_WallPaperSolid()
		return t.Encode()

	default:
		return []byte{}
	}
}

// wallPaper#ccb03657 id:int title:string sizes:Vector<PhotoSize> color:int = WallPaper;
func (m *WallPaper) To_WallPaper() *TLWallPaper {
	return &TLWallPaper{
		Data2: m.Data2,
	}
}

// wallPaperSolid#63117f24 id:int title:string bg_color:int color:int = WallPaper;
func (m *WallPaper) To_WallPaperSolid() *TLWallPaperSolid {
	return &TLWallPaperSolid{
		Data2: m.Data2,
	}
}

// wallPaper#ccb03657 id:int title:string sizes:Vector<PhotoSize> color:int = WallPaper;
func (m *TLWallPaper) To_WallPaper() *WallPaper {
	return &WallPaper{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLWallPaper) SetId(v int32) { m.Data2.Id = v }
func (m *TLWallPaper) GetId() int32  { return m.Data2.Id }

func (m *TLWallPaper) SetTitle(v string) { m.Data2.Title = v }
func (m *TLWallPaper) GetTitle() string  { return m.Data2.Title }

func (m *TLWallPaper) SetSizes(v []*PhotoSize) { m.Data2.Sizes = v }
func (m *TLWallPaper) GetSizes() []*PhotoSize  { return m.Data2.Sizes }

func (m *TLWallPaper) SetColor(v int32) { m.Data2.Color = v }
func (m *TLWallPaper) GetColor() int32  { return m.Data2.Color }

func (m *TLWallPaper) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_wallPaper))

	x.Int(m.Data2.Id)
	x.StringBytes(m.Data2.Title)

	x.Int(m.Data2.Color)

	return x.buf
}

func (m *TLWallPaper) Decode(dbuf *DecodeBuf) error {
	m.Data2.Id = x.Int()
	m.Data2.Title = x.StringBytes()

	m.Data2.Color = x.Int()

	return dbuf.err
}

// wallPaperSolid#63117f24 id:int title:string bg_color:int color:int = WallPaper;
func (m *TLWallPaperSolid) To_WallPaper() *WallPaper {
	return &WallPaper{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLWallPaperSolid) SetId(v int32) { m.Data2.Id = v }
func (m *TLWallPaperSolid) GetId() int32  { return m.Data2.Id }

func (m *TLWallPaperSolid) SetTitle(v string) { m.Data2.Title = v }
func (m *TLWallPaperSolid) GetTitle() string  { return m.Data2.Title }

func (m *TLWallPaperSolid) SetBgColor(v int32) { m.Data2.BgColor = v }
func (m *TLWallPaperSolid) GetBgColor() int32  { return m.Data2.BgColor }

func (m *TLWallPaperSolid) SetColor(v int32) { m.Data2.Color = v }
func (m *TLWallPaperSolid) GetColor() int32  { return m.Data2.Color }

func (m *TLWallPaperSolid) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_wallPaperSolid))

	x.Int(m.Data2.Id)
	x.StringBytes(m.Data2.Title)
	x.Int(m.Data2.BgColor)
	x.Int(m.Data2.Color)

	return x.buf
}

func (m *TLWallPaperSolid) Decode(dbuf *DecodeBuf) error {
	m.Data2.Id = x.Int()
	m.Data2.Title = x.StringBytes()
	m.Data2.BgColor = x.Int()
	m.Data2.Color = x.Int()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Updates_Difference <--
//  + TL_UpdatesDifferenceEmpty
//  + TL_UpdatesDifference
//  + TL_UpdatesDifferenceSlice
//  + TL_UpdatesDifferenceTooLong
//

func (m *Updates_Difference) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_updates_differenceEmpty:
		t := m.To_UpdatesDifferenceEmpty()
		return t.Encode()
	case TLConstructor_CRC32_updates_difference:
		t := m.To_UpdatesDifference()
		return t.Encode()
	case TLConstructor_CRC32_updates_differenceSlice:
		t := m.To_UpdatesDifferenceSlice()
		return t.Encode()
	case TLConstructor_CRC32_updates_differenceTooLong:
		t := m.To_UpdatesDifferenceTooLong()
		return t.Encode()

	default:
		return []byte{}
	}
}

// updates.differenceEmpty#5d75a138 date:int seq:int = updates.Difference;
func (m *Updates_Difference) To_UpdatesDifferenceEmpty() *TLUpdatesDifferenceEmpty {
	return &TLUpdatesDifferenceEmpty{
		Data2: m.Data2,
	}
}

// updates.difference#f49ca0 new_messages:Vector<Message> new_encrypted_messages:Vector<EncryptedMessage> other_updates:Vector<Update> chats:Vector<Chat> users:Vector<User> state:updates.State = updates.Difference;
func (m *Updates_Difference) To_UpdatesDifference() *TLUpdatesDifference {
	return &TLUpdatesDifference{
		Data2: m.Data2,
	}
}

// updates.differenceSlice#a8fb1981 new_messages:Vector<Message> new_encrypted_messages:Vector<EncryptedMessage> other_updates:Vector<Update> chats:Vector<Chat> users:Vector<User> intermediate_state:updates.State = updates.Difference;
func (m *Updates_Difference) To_UpdatesDifferenceSlice() *TLUpdatesDifferenceSlice {
	return &TLUpdatesDifferenceSlice{
		Data2: m.Data2,
	}
}

// updates.differenceTooLong#4afe8f6d pts:int = updates.Difference;
func (m *Updates_Difference) To_UpdatesDifferenceTooLong() *TLUpdatesDifferenceTooLong {
	return &TLUpdatesDifferenceTooLong{
		Data2: m.Data2,
	}
}

// updates.differenceEmpty#5d75a138 date:int seq:int = updates.Difference;
func (m *TLUpdatesDifferenceEmpty) To_Updates_Difference() *Updates_Difference {
	return &Updates_Difference{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdatesDifferenceEmpty) SetDate(v int32) { m.Data2.Date = v }
func (m *TLUpdatesDifferenceEmpty) GetDate() int32  { return m.Data2.Date }

func (m *TLUpdatesDifferenceEmpty) SetSeq(v int32) { m.Data2.Seq = v }
func (m *TLUpdatesDifferenceEmpty) GetSeq() int32  { return m.Data2.Seq }

func (m *TLUpdatesDifferenceEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updates_differenceEmpty))

	x.Int(m.Data2.Date)
	x.Int(m.Data2.Seq)

	return x.buf
}

func (m *TLUpdatesDifferenceEmpty) Decode(dbuf *DecodeBuf) error {
	m.Data2.Date = x.Int()
	m.Data2.Seq = x.Int()

	return dbuf.err
}

// updates.difference#f49ca0 new_messages:Vector<Message> new_encrypted_messages:Vector<EncryptedMessage> other_updates:Vector<Update> chats:Vector<Chat> users:Vector<User> state:updates.State = updates.Difference;
func (m *TLUpdatesDifference) To_Updates_Difference() *Updates_Difference {
	return &Updates_Difference{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdatesDifference) SetNewMessages(v []*Message) { m.Data2.NewMessages = v }
func (m *TLUpdatesDifference) GetNewMessages() []*Message  { return m.Data2.NewMessages }

func (m *TLUpdatesDifference) SetNewEncryptedMessages(v []*EncryptedMessage) {
	m.Data2.NewEncryptedMessages = v
}
func (m *TLUpdatesDifference) GetNewEncryptedMessages() []*EncryptedMessage {
	return m.Data2.NewEncryptedMessages
}

func (m *TLUpdatesDifference) SetOtherUpdates(v []*Update) { m.Data2.OtherUpdates = v }
func (m *TLUpdatesDifference) GetOtherUpdates() []*Update  { return m.Data2.OtherUpdates }

func (m *TLUpdatesDifference) SetChats(v []*Chat) { m.Data2.Chats = v }
func (m *TLUpdatesDifference) GetChats() []*Chat  { return m.Data2.Chats }

func (m *TLUpdatesDifference) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLUpdatesDifference) GetUsers() []*User  { return m.Data2.Users }

func (m *TLUpdatesDifference) SetState(v *Updates_State) { m.Data2.State = v }
func (m *TLUpdatesDifference) GetState() *Updates_State  { return m.Data2.State }

func (m *TLUpdatesDifference) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updates_difference))

	x.Bytes(m.Data2State.Encode())

	return x.buf
}

func (m *TLUpdatesDifference) Decode(dbuf *DecodeBuf) error {

	m6 := &State{}
	m6.Decode(dbuf)
	m.SetState(m6)

	return dbuf.err
}

// updates.differenceSlice#a8fb1981 new_messages:Vector<Message> new_encrypted_messages:Vector<EncryptedMessage> other_updates:Vector<Update> chats:Vector<Chat> users:Vector<User> intermediate_state:updates.State = updates.Difference;
func (m *TLUpdatesDifferenceSlice) To_Updates_Difference() *Updates_Difference {
	return &Updates_Difference{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdatesDifferenceSlice) SetNewMessages(v []*Message) { m.Data2.NewMessages = v }
func (m *TLUpdatesDifferenceSlice) GetNewMessages() []*Message  { return m.Data2.NewMessages }

func (m *TLUpdatesDifferenceSlice) SetNewEncryptedMessages(v []*EncryptedMessage) {
	m.Data2.NewEncryptedMessages = v
}
func (m *TLUpdatesDifferenceSlice) GetNewEncryptedMessages() []*EncryptedMessage {
	return m.Data2.NewEncryptedMessages
}

func (m *TLUpdatesDifferenceSlice) SetOtherUpdates(v []*Update) { m.Data2.OtherUpdates = v }
func (m *TLUpdatesDifferenceSlice) GetOtherUpdates() []*Update  { return m.Data2.OtherUpdates }

func (m *TLUpdatesDifferenceSlice) SetChats(v []*Chat) { m.Data2.Chats = v }
func (m *TLUpdatesDifferenceSlice) GetChats() []*Chat  { return m.Data2.Chats }

func (m *TLUpdatesDifferenceSlice) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLUpdatesDifferenceSlice) GetUsers() []*User  { return m.Data2.Users }

func (m *TLUpdatesDifferenceSlice) SetIntermediateState(v *Updates_State) {
	m.Data2.IntermediateState = v
}
func (m *TLUpdatesDifferenceSlice) GetIntermediateState() *Updates_State {
	return m.Data2.IntermediateState
}

func (m *TLUpdatesDifferenceSlice) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updates_differenceSlice))

	x.Bytes(m.Data2IntermediateState.Encode())

	return x.buf
}

func (m *TLUpdatesDifferenceSlice) Decode(dbuf *DecodeBuf) error {

	m6 := &IntermediateState{}
	m6.Decode(dbuf)
	m.SetIntermediateState(m6)

	return dbuf.err
}

// updates.differenceTooLong#4afe8f6d pts:int = updates.Difference;
func (m *TLUpdatesDifferenceTooLong) To_Updates_Difference() *Updates_Difference {
	return &Updates_Difference{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdatesDifferenceTooLong) SetPts(v int32) { m.Data2.Pts = v }
func (m *TLUpdatesDifferenceTooLong) GetPts() int32  { return m.Data2.Pts }

func (m *TLUpdatesDifferenceTooLong) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updates_differenceTooLong))

	x.Int(m.Data2.Pts)

	return x.buf
}

func (m *TLUpdatesDifferenceTooLong) Decode(dbuf *DecodeBuf) error {
	m.Data2.Pts = x.Int()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Updates <--
//  + TL_UpdatesTooLong
//  + TL_UpdateShortMessage
//  + TL_UpdateShortChatMessage
//  + TL_UpdateShort
//  + TL_UpdatesCombined
//  + TL_Updates
//  + TL_UpdateShortSentMessage
//

func (m *Updates) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_updatesTooLong:
		t := m.To_UpdatesTooLong()
		return t.Encode()
	case TLConstructor_CRC32_updateShortMessage:
		t := m.To_UpdateShortMessage()
		return t.Encode()
	case TLConstructor_CRC32_updateShortChatMessage:
		t := m.To_UpdateShortChatMessage()
		return t.Encode()
	case TLConstructor_CRC32_updateShort:
		t := m.To_UpdateShort()
		return t.Encode()
	case TLConstructor_CRC32_updatesCombined:
		t := m.To_UpdatesCombined()
		return t.Encode()
	case TLConstructor_CRC32_updates:
		t := m.To_Updates()
		return t.Encode()
	case TLConstructor_CRC32_updateShortSentMessage:
		t := m.To_UpdateShortSentMessage()
		return t.Encode()

	default:
		return []byte{}
	}
}

// updatesTooLong#e317af7e = Updates;
func (m *Updates) To_UpdatesTooLong() *TLUpdatesTooLong {
	return &TLUpdatesTooLong{
		Data2: m.Data2,
	}
}

// updateShortMessage#914fbf11 flags:# out:flags.1?true mentioned:flags.4?true media_unread:flags.5?true silent:flags.13?true id:int user_id:int message:string pts:int pts_count:int date:int fwd_from:flags.2?MessageFwdHeader via_bot_id:flags.11?int reply_to_msg_id:flags.3?int entities:flags.7?Vector<MessageEntity> = Updates;
func (m *Updates) To_UpdateShortMessage() *TLUpdateShortMessage {
	return &TLUpdateShortMessage{
		Data2: m.Data2,
	}
}

// updateShortChatMessage#16812688 flags:# out:flags.1?true mentioned:flags.4?true media_unread:flags.5?true silent:flags.13?true id:int from_id:int chat_id:int message:string pts:int pts_count:int date:int fwd_from:flags.2?MessageFwdHeader via_bot_id:flags.11?int reply_to_msg_id:flags.3?int entities:flags.7?Vector<MessageEntity> = Updates;
func (m *Updates) To_UpdateShortChatMessage() *TLUpdateShortChatMessage {
	return &TLUpdateShortChatMessage{
		Data2: m.Data2,
	}
}

// updateShort#78d4dec1 update:Update date:int = Updates;
func (m *Updates) To_UpdateShort() *TLUpdateShort {
	return &TLUpdateShort{
		Data2: m.Data2,
	}
}

// updatesCombined#725b04c3 updates:Vector<Update> users:Vector<User> chats:Vector<Chat> date:int seq_start:int seq:int = Updates;
func (m *Updates) To_UpdatesCombined() *TLUpdatesCombined {
	return &TLUpdatesCombined{
		Data2: m.Data2,
	}
}

// updates#74ae4240 updates:Vector<Update> users:Vector<User> chats:Vector<Chat> date:int seq:int = Updates;
func (m *Updates) To_Updates() *TLUpdates {
	return &TLUpdates{
		Data2: m.Data2,
	}
}

// updateShortSentMessage#11f1331c flags:# out:flags.1?true id:int pts:int pts_count:int date:int media:flags.9?MessageMedia entities:flags.7?Vector<MessageEntity> = Updates;
func (m *Updates) To_UpdateShortSentMessage() *TLUpdateShortSentMessage {
	return &TLUpdateShortSentMessage{
		Data2: m.Data2,
	}
}

// updatesTooLong#e317af7e = Updates;
func (m *TLUpdatesTooLong) To_Updates() *Updates {
	return &Updates{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdatesTooLong) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updatesTooLong))

	return x.buf
}

func (m *TLUpdatesTooLong) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// updateShortMessage#914fbf11 flags:# out:flags.1?true mentioned:flags.4?true media_unread:flags.5?true silent:flags.13?true id:int user_id:int message:string pts:int pts_count:int date:int fwd_from:flags.2?MessageFwdHeader via_bot_id:flags.11?int reply_to_msg_id:flags.3?int entities:flags.7?Vector<MessageEntity> = Updates;
func (m *TLUpdateShortMessage) To_Updates() *Updates {
	return &Updates{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateShortMessage) SetOut(v bool) { m.Data2.Out = v }
func (m *TLUpdateShortMessage) GetOut() bool  { return m.Data2.Out }

func (m *TLUpdateShortMessage) SetMentioned(v bool) { m.Data2.Mentioned = v }
func (m *TLUpdateShortMessage) GetMentioned() bool  { return m.Data2.Mentioned }

func (m *TLUpdateShortMessage) SetMediaUnread(v bool) { m.Data2.MediaUnread = v }
func (m *TLUpdateShortMessage) GetMediaUnread() bool  { return m.Data2.MediaUnread }

func (m *TLUpdateShortMessage) SetSilent(v bool) { m.Data2.Silent = v }
func (m *TLUpdateShortMessage) GetSilent() bool  { return m.Data2.Silent }

func (m *TLUpdateShortMessage) SetId(v int32) { m.Data2.Id = v }
func (m *TLUpdateShortMessage) GetId() int32  { return m.Data2.Id }

func (m *TLUpdateShortMessage) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLUpdateShortMessage) GetUserId() int32  { return m.Data2.UserId }

func (m *TLUpdateShortMessage) SetMessage(v string) { m.Data2.Message = v }
func (m *TLUpdateShortMessage) GetMessage() string  { return m.Data2.Message }

func (m *TLUpdateShortMessage) SetPts(v int32) { m.Data2.Pts = v }
func (m *TLUpdateShortMessage) GetPts() int32  { return m.Data2.Pts }

func (m *TLUpdateShortMessage) SetPtsCount(v int32) { m.Data2.PtsCount = v }
func (m *TLUpdateShortMessage) GetPtsCount() int32  { return m.Data2.PtsCount }

func (m *TLUpdateShortMessage) SetDate(v int32) { m.Data2.Date = v }
func (m *TLUpdateShortMessage) GetDate() int32  { return m.Data2.Date }

func (m *TLUpdateShortMessage) SetFwdFrom(v *MessageFwdHeader) { m.Data2.FwdFrom = v }
func (m *TLUpdateShortMessage) GetFwdFrom() *MessageFwdHeader  { return m.Data2.FwdFrom }

func (m *TLUpdateShortMessage) SetViaBotId(v int32) { m.Data2.ViaBotId = v }
func (m *TLUpdateShortMessage) GetViaBotId() int32  { return m.Data2.ViaBotId }

func (m *TLUpdateShortMessage) SetReplyToMsgId(v int32) { m.Data2.ReplyToMsgId = v }
func (m *TLUpdateShortMessage) GetReplyToMsgId() int32  { return m.Data2.ReplyToMsgId }

func (m *TLUpdateShortMessage) SetEntities(v []*MessageEntity) { m.Data2.Entities = v }
func (m *TLUpdateShortMessage) GetEntities() []*MessageEntity  { return m.Data2.Entities }

func (m *TLUpdateShortMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateShortMessage))

	// flags
	var flags uint32 = 0
	if m.GetOut() == true {
		flags |= 1 << 1
	}
	if m.GetMentioned() == true {
		flags |= 1 << 2
	}
	if m.GetMediaUnread() == true {
		flags |= 1 << 3
	}
	if m.GetSilent() == true {
		flags |= 1 << 4
	}
	if m.GetFwdFrom() != nil {
		flags |= 1 << 11
	}
	if m.GetViaBotId() != 0 {
		flags |= 1 << 12
	}
	if m.GetReplyToMsgId() != 0 {
		flags |= 1 << 13
	}
	if m.GetEntities() != nil {
		flags |= 1 << 14
	}
	x.UInt(flags)

	x.Int(m.Data2.Id)
	x.Int(m.Data2.UserId)
	x.StringBytes(m.Data2.Message)
	x.Int(m.Data2.Pts)
	x.Int(m.Data2.PtsCount)
	x.Int(m.Data2.Date)
	if m.GetFwdFrom() != 0 {
		x.Bytes(m.Data2FwdFrom.Encode())
	}
	if m.GetViaBotId() != 0 {
		x.Int(m.Data2.ViaBotId)
	}
	if m.GetReplyToMsgId() != 0 {
		x.Int(m.Data2.ReplyToMsgId)
	}
	if m.Data2.Entities != 0 {

	}

	return x.buf
}

func (m *TLUpdateShortMessage) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m.Data2.Out = true
	}
	if (flags & (1 << 3)) != 0 {
		m.Data2.Mentioned = true
	}
	if (flags & (1 << 4)) != 0 {
		m.Data2.MediaUnread = true
	}
	if (flags & (1 << 5)) != 0 {
		m.Data2.Silent = true
	}
	m.Data2.Id = x.Int()
	m.Data2.UserId = x.Int()
	m.Data2.Message = x.StringBytes()
	m.Data2.Pts = x.Int()
	m.Data2.PtsCount = x.Int()
	m.Data2.Date = x.Int()
	if (flags & (1 << 12)) != 0 {
		m12 := &FwdFrom{}
		m12.Decode(dbuf)
		m.SetFwdFrom(m12)
	}
	if (flags & (1 << 13)) != 0 {
		m.Data2.ViaBotId = x.Int()
	}
	if (flags & (1 << 14)) != 0 {
		m.Data2.ReplyToMsgId = x.Int()
	}
	if (flags & (1 << 15)) != 0 {

	}

	return dbuf.err
}

// updateShortChatMessage#16812688 flags:# out:flags.1?true mentioned:flags.4?true media_unread:flags.5?true silent:flags.13?true id:int from_id:int chat_id:int message:string pts:int pts_count:int date:int fwd_from:flags.2?MessageFwdHeader via_bot_id:flags.11?int reply_to_msg_id:flags.3?int entities:flags.7?Vector<MessageEntity> = Updates;
func (m *TLUpdateShortChatMessage) To_Updates() *Updates {
	return &Updates{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateShortChatMessage) SetOut(v bool) { m.Data2.Out = v }
func (m *TLUpdateShortChatMessage) GetOut() bool  { return m.Data2.Out }

func (m *TLUpdateShortChatMessage) SetMentioned(v bool) { m.Data2.Mentioned = v }
func (m *TLUpdateShortChatMessage) GetMentioned() bool  { return m.Data2.Mentioned }

func (m *TLUpdateShortChatMessage) SetMediaUnread(v bool) { m.Data2.MediaUnread = v }
func (m *TLUpdateShortChatMessage) GetMediaUnread() bool  { return m.Data2.MediaUnread }

func (m *TLUpdateShortChatMessage) SetSilent(v bool) { m.Data2.Silent = v }
func (m *TLUpdateShortChatMessage) GetSilent() bool  { return m.Data2.Silent }

func (m *TLUpdateShortChatMessage) SetId(v int32) { m.Data2.Id = v }
func (m *TLUpdateShortChatMessage) GetId() int32  { return m.Data2.Id }

func (m *TLUpdateShortChatMessage) SetFromId(v int32) { m.Data2.FromId = v }
func (m *TLUpdateShortChatMessage) GetFromId() int32  { return m.Data2.FromId }

func (m *TLUpdateShortChatMessage) SetChatId(v int32) { m.Data2.ChatId = v }
func (m *TLUpdateShortChatMessage) GetChatId() int32  { return m.Data2.ChatId }

func (m *TLUpdateShortChatMessage) SetMessage(v string) { m.Data2.Message = v }
func (m *TLUpdateShortChatMessage) GetMessage() string  { return m.Data2.Message }

func (m *TLUpdateShortChatMessage) SetPts(v int32) { m.Data2.Pts = v }
func (m *TLUpdateShortChatMessage) GetPts() int32  { return m.Data2.Pts }

func (m *TLUpdateShortChatMessage) SetPtsCount(v int32) { m.Data2.PtsCount = v }
func (m *TLUpdateShortChatMessage) GetPtsCount() int32  { return m.Data2.PtsCount }

func (m *TLUpdateShortChatMessage) SetDate(v int32) { m.Data2.Date = v }
func (m *TLUpdateShortChatMessage) GetDate() int32  { return m.Data2.Date }

func (m *TLUpdateShortChatMessage) SetFwdFrom(v *MessageFwdHeader) { m.Data2.FwdFrom = v }
func (m *TLUpdateShortChatMessage) GetFwdFrom() *MessageFwdHeader  { return m.Data2.FwdFrom }

func (m *TLUpdateShortChatMessage) SetViaBotId(v int32) { m.Data2.ViaBotId = v }
func (m *TLUpdateShortChatMessage) GetViaBotId() int32  { return m.Data2.ViaBotId }

func (m *TLUpdateShortChatMessage) SetReplyToMsgId(v int32) { m.Data2.ReplyToMsgId = v }
func (m *TLUpdateShortChatMessage) GetReplyToMsgId() int32  { return m.Data2.ReplyToMsgId }

func (m *TLUpdateShortChatMessage) SetEntities(v []*MessageEntity) { m.Data2.Entities = v }
func (m *TLUpdateShortChatMessage) GetEntities() []*MessageEntity  { return m.Data2.Entities }

func (m *TLUpdateShortChatMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateShortChatMessage))

	// flags
	var flags uint32 = 0
	if m.GetOut() == true {
		flags |= 1 << 1
	}
	if m.GetMentioned() == true {
		flags |= 1 << 2
	}
	if m.GetMediaUnread() == true {
		flags |= 1 << 3
	}
	if m.GetSilent() == true {
		flags |= 1 << 4
	}
	if m.GetFwdFrom() != nil {
		flags |= 1 << 12
	}
	if m.GetViaBotId() != 0 {
		flags |= 1 << 13
	}
	if m.GetReplyToMsgId() != 0 {
		flags |= 1 << 14
	}
	if m.GetEntities() != nil {
		flags |= 1 << 15
	}
	x.UInt(flags)

	x.Int(m.Data2.Id)
	x.Int(m.Data2.FromId)
	x.Int(m.Data2.ChatId)
	x.StringBytes(m.Data2.Message)
	x.Int(m.Data2.Pts)
	x.Int(m.Data2.PtsCount)
	x.Int(m.Data2.Date)
	if m.GetFwdFrom() != 0 {
		x.Bytes(m.Data2FwdFrom.Encode())
	}
	if m.GetViaBotId() != 0 {
		x.Int(m.Data2.ViaBotId)
	}
	if m.GetReplyToMsgId() != 0 {
		x.Int(m.Data2.ReplyToMsgId)
	}
	if m.Data2.Entities != 0 {

	}

	return x.buf
}

func (m *TLUpdateShortChatMessage) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m.Data2.Out = true
	}
	if (flags & (1 << 3)) != 0 {
		m.Data2.Mentioned = true
	}
	if (flags & (1 << 4)) != 0 {
		m.Data2.MediaUnread = true
	}
	if (flags & (1 << 5)) != 0 {
		m.Data2.Silent = true
	}
	m.Data2.Id = x.Int()
	m.Data2.FromId = x.Int()
	m.Data2.ChatId = x.Int()
	m.Data2.Message = x.StringBytes()
	m.Data2.Pts = x.Int()
	m.Data2.PtsCount = x.Int()
	m.Data2.Date = x.Int()
	if (flags & (1 << 13)) != 0 {
		m13 := &FwdFrom{}
		m13.Decode(dbuf)
		m.SetFwdFrom(m13)
	}
	if (flags & (1 << 14)) != 0 {
		m.Data2.ViaBotId = x.Int()
	}
	if (flags & (1 << 15)) != 0 {
		m.Data2.ReplyToMsgId = x.Int()
	}
	if (flags & (1 << 16)) != 0 {

	}

	return dbuf.err
}

// updateShort#78d4dec1 update:Update date:int = Updates;
func (m *TLUpdateShort) To_Updates() *Updates {
	return &Updates{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateShort) SetUpdate(v *Update) { m.Data2.Update = v }
func (m *TLUpdateShort) GetUpdate() *Update  { return m.Data2.Update }

func (m *TLUpdateShort) SetDate(v int32) { m.Data2.Date = v }
func (m *TLUpdateShort) GetDate() int32  { return m.Data2.Date }

func (m *TLUpdateShort) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateShort))

	x.Bytes(m.Data2Update.Encode())
	x.Int(m.Data2.Date)

	return x.buf
}

func (m *TLUpdateShort) Decode(dbuf *DecodeBuf) error {
	m1 := &Update{}
	m1.Decode(dbuf)
	m.SetUpdate(m1)
	m.Data2.Date = x.Int()

	return dbuf.err
}

// updatesCombined#725b04c3 updates:Vector<Update> users:Vector<User> chats:Vector<Chat> date:int seq_start:int seq:int = Updates;
func (m *TLUpdatesCombined) To_Updates() *Updates {
	return &Updates{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdatesCombined) SetUpdates(v []*Update) { m.Data2.Updates = v }
func (m *TLUpdatesCombined) GetUpdates() []*Update  { return m.Data2.Updates }

func (m *TLUpdatesCombined) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLUpdatesCombined) GetUsers() []*User  { return m.Data2.Users }

func (m *TLUpdatesCombined) SetChats(v []*Chat) { m.Data2.Chats = v }
func (m *TLUpdatesCombined) GetChats() []*Chat  { return m.Data2.Chats }

func (m *TLUpdatesCombined) SetDate(v int32) { m.Data2.Date = v }
func (m *TLUpdatesCombined) GetDate() int32  { return m.Data2.Date }

func (m *TLUpdatesCombined) SetSeqStart(v int32) { m.Data2.SeqStart = v }
func (m *TLUpdatesCombined) GetSeqStart() int32  { return m.Data2.SeqStart }

func (m *TLUpdatesCombined) SetSeq(v int32) { m.Data2.Seq = v }
func (m *TLUpdatesCombined) GetSeq() int32  { return m.Data2.Seq }

func (m *TLUpdatesCombined) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updatesCombined))

	x.Int(m.Data2.Date)
	x.Int(m.Data2.SeqStart)
	x.Int(m.Data2.Seq)

	return x.buf
}

func (m *TLUpdatesCombined) Decode(dbuf *DecodeBuf) error {

	m.Data2.Date = x.Int()
	m.Data2.SeqStart = x.Int()
	m.Data2.Seq = x.Int()

	return dbuf.err
}

// updates#74ae4240 updates:Vector<Update> users:Vector<User> chats:Vector<Chat> date:int seq:int = Updates;
func (m *TLUpdates) To_Updates() *Updates {
	return &Updates{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdates) SetUpdates(v []*Update) { m.Data2.Updates = v }
func (m *TLUpdates) GetUpdates() []*Update  { return m.Data2.Updates }

func (m *TLUpdates) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLUpdates) GetUsers() []*User  { return m.Data2.Users }

func (m *TLUpdates) SetChats(v []*Chat) { m.Data2.Chats = v }
func (m *TLUpdates) GetChats() []*Chat  { return m.Data2.Chats }

func (m *TLUpdates) SetDate(v int32) { m.Data2.Date = v }
func (m *TLUpdates) GetDate() int32  { return m.Data2.Date }

func (m *TLUpdates) SetSeq(v int32) { m.Data2.Seq = v }
func (m *TLUpdates) GetSeq() int32  { return m.Data2.Seq }

func (m *TLUpdates) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updates))

	x.Int(m.Data2.Date)
	x.Int(m.Data2.Seq)

	return x.buf
}

func (m *TLUpdates) Decode(dbuf *DecodeBuf) error {

	m.Data2.Date = x.Int()
	m.Data2.Seq = x.Int()

	return dbuf.err
}

// updateShortSentMessage#11f1331c flags:# out:flags.1?true id:int pts:int pts_count:int date:int media:flags.9?MessageMedia entities:flags.7?Vector<MessageEntity> = Updates;
func (m *TLUpdateShortSentMessage) To_Updates() *Updates {
	return &Updates{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateShortSentMessage) SetOut(v bool) { m.Data2.Out = v }
func (m *TLUpdateShortSentMessage) GetOut() bool  { return m.Data2.Out }

func (m *TLUpdateShortSentMessage) SetId(v int32) { m.Data2.Id = v }
func (m *TLUpdateShortSentMessage) GetId() int32  { return m.Data2.Id }

func (m *TLUpdateShortSentMessage) SetPts(v int32) { m.Data2.Pts = v }
func (m *TLUpdateShortSentMessage) GetPts() int32  { return m.Data2.Pts }

func (m *TLUpdateShortSentMessage) SetPtsCount(v int32) { m.Data2.PtsCount = v }
func (m *TLUpdateShortSentMessage) GetPtsCount() int32  { return m.Data2.PtsCount }

func (m *TLUpdateShortSentMessage) SetDate(v int32) { m.Data2.Date = v }
func (m *TLUpdateShortSentMessage) GetDate() int32  { return m.Data2.Date }

func (m *TLUpdateShortSentMessage) SetMedia(v *MessageMedia) { m.Data2.Media = v }
func (m *TLUpdateShortSentMessage) GetMedia() *MessageMedia  { return m.Data2.Media }

func (m *TLUpdateShortSentMessage) SetEntities(v []*MessageEntity) { m.Data2.Entities = v }
func (m *TLUpdateShortSentMessage) GetEntities() []*MessageEntity  { return m.Data2.Entities }

func (m *TLUpdateShortSentMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateShortSentMessage))

	// flags
	var flags uint32 = 0
	if m.GetOut() == true {
		flags |= 1 << 1
	}
	if m.GetMedia() != nil {
		flags |= 1 << 6
	}
	if m.GetEntities() != nil {
		flags |= 1 << 7
	}
	x.UInt(flags)

	x.Int(m.Data2.Id)
	x.Int(m.Data2.Pts)
	x.Int(m.Data2.PtsCount)
	x.Int(m.Data2.Date)
	if m.GetMedia() != 0 {
		x.Bytes(m.Data2Media.Encode())
	}
	if m.Data2.Entities != 0 {

	}

	return x.buf
}

func (m *TLUpdateShortSentMessage) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m.Data2.Out = true
	}
	m.Data2.Id = x.Int()
	m.Data2.Pts = x.Int()
	m.Data2.PtsCount = x.Int()
	m.Data2.Date = x.Int()
	if (flags & (1 << 7)) != 0 {
		m7 := &Media{}
		m7.Decode(dbuf)
		m.SetMedia(m7)
	}
	if (flags & (1 << 8)) != 0 {

	}

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Account_Password <--
//  + TL_AccountNoPassword
//  + TL_AccountPassword
//

func (m *Account_Password) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_account_noPassword:
		t := m.To_AccountNoPassword()
		return t.Encode()
	case TLConstructor_CRC32_account_password:
		t := m.To_AccountPassword()
		return t.Encode()

	default:
		return []byte{}
	}
}

// account.noPassword#96dabc18 new_salt:bytes email_unconfirmed_pattern:string = account.Password;
func (m *Account_Password) To_AccountNoPassword() *TLAccountNoPassword {
	return &TLAccountNoPassword{
		Data2: m.Data2,
	}
}

// account.password#7c18141c current_salt:bytes new_salt:bytes hint:string has_recovery:Bool email_unconfirmed_pattern:string = account.Password;
func (m *Account_Password) To_AccountPassword() *TLAccountPassword {
	return &TLAccountPassword{
		Data2: m.Data2,
	}
}

// account.noPassword#96dabc18 new_salt:bytes email_unconfirmed_pattern:string = account.Password;
func (m *TLAccountNoPassword) To_Account_Password() *Account_Password {
	return &Account_Password{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLAccountNoPassword) SetNewSalt(v []byte) { m.Data2.NewSalt = v }
func (m *TLAccountNoPassword) GetNewSalt() []byte  { return m.Data2.NewSalt }

func (m *TLAccountNoPassword) SetEmailUnconfirmedPattern(v string) {
	m.Data2.EmailUnconfirmedPattern = v
}
func (m *TLAccountNoPassword) GetEmailUnconfirmedPattern() string {
	return m.Data2.EmailUnconfirmedPattern
}

func (m *TLAccountNoPassword) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_account_noPassword))

	x.StringBytes(m.Data2.NewSalt)
	x.StringBytes(m.Data2.EmailUnconfirmedPattern)

	return x.buf
}

func (m *TLAccountNoPassword) Decode(dbuf *DecodeBuf) error {
	m.Data2.NewSalt = x.StringBytes()
	m.Data2.EmailUnconfirmedPattern = x.StringBytes()

	return dbuf.err
}

// account.password#7c18141c current_salt:bytes new_salt:bytes hint:string has_recovery:Bool email_unconfirmed_pattern:string = account.Password;
func (m *TLAccountPassword) To_Account_Password() *Account_Password {
	return &Account_Password{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLAccountPassword) SetCurrentSalt(v []byte) { m.Data2.CurrentSalt = v }
func (m *TLAccountPassword) GetCurrentSalt() []byte  { return m.Data2.CurrentSalt }

func (m *TLAccountPassword) SetNewSalt(v []byte) { m.Data2.NewSalt = v }
func (m *TLAccountPassword) GetNewSalt() []byte  { return m.Data2.NewSalt }

func (m *TLAccountPassword) SetHint(v string) { m.Data2.Hint = v }
func (m *TLAccountPassword) GetHint() string  { return m.Data2.Hint }

func (m *TLAccountPassword) SetHasRecovery(v *Bool) { m.Data2.HasRecovery = v }
func (m *TLAccountPassword) GetHasRecovery() *Bool  { return m.Data2.HasRecovery }

func (m *TLAccountPassword) SetEmailUnconfirmedPattern(v string) { m.Data2.EmailUnconfirmedPattern = v }
func (m *TLAccountPassword) GetEmailUnconfirmedPattern() string {
	return m.Data2.EmailUnconfirmedPattern
}

func (m *TLAccountPassword) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_account_password))

	x.StringBytes(m.Data2.CurrentSalt)
	x.StringBytes(m.Data2.NewSalt)
	x.StringBytes(m.Data2.Hint)
	x.Bytes(m.Data2HasRecovery.Encode())
	x.StringBytes(m.Data2.EmailUnconfirmedPattern)

	return x.buf
}

func (m *TLAccountPassword) Decode(dbuf *DecodeBuf) error {
	m.Data2.CurrentSalt = x.StringBytes()
	m.Data2.NewSalt = x.StringBytes()
	m.Data2.Hint = x.StringBytes()
	m4 := &HasRecovery{}
	m4.Decode(dbuf)
	m.SetHasRecovery(m4)
	m.Data2.EmailUnconfirmedPattern = x.StringBytes()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// ChatFull <--
//  + TL_ChatFull
//  + TL_ChannelFull
//

func (m *ChatFull) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_chatFull:
		t := m.To_ChatFull()
		return t.Encode()
	case TLConstructor_CRC32_channelFull:
		t := m.To_ChannelFull()
		return t.Encode()

	default:
		return []byte{}
	}
}

// chatFull#2e02a614 id:int participants:ChatParticipants chat_photo:Photo notify_settings:PeerNotifySettings exported_invite:ExportedChatInvite bot_info:Vector<BotInfo> = ChatFull;
func (m *ChatFull) To_ChatFull() *TLChatFull {
	return &TLChatFull{
		Data2: m.Data2,
	}
}

// channelFull#17f45fcf flags:# can_view_participants:flags.3?true can_set_username:flags.6?true can_set_stickers:flags.7?true id:int about:string participants_count:flags.0?int admins_count:flags.1?int kicked_count:flags.2?int banned_count:flags.2?int read_inbox_max_id:int read_outbox_max_id:int unread_count:int chat_photo:Photo notify_settings:PeerNotifySettings exported_invite:ExportedChatInvite bot_info:Vector<BotInfo> migrated_from_chat_id:flags.4?int migrated_from_max_id:flags.4?int pinned_msg_id:flags.5?int stickerset:flags.8?StickerSet = ChatFull;
func (m *ChatFull) To_ChannelFull() *TLChannelFull {
	return &TLChannelFull{
		Data2: m.Data2,
	}
}

// chatFull#2e02a614 id:int participants:ChatParticipants chat_photo:Photo notify_settings:PeerNotifySettings exported_invite:ExportedChatInvite bot_info:Vector<BotInfo> = ChatFull;
func (m *TLChatFull) To_ChatFull() *ChatFull {
	return &ChatFull{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLChatFull) SetId(v int32) { m.Data2.Id = v }
func (m *TLChatFull) GetId() int32  { return m.Data2.Id }

func (m *TLChatFull) SetParticipants(v *ChatParticipants) { m.Data2.Participants = v }
func (m *TLChatFull) GetParticipants() *ChatParticipants  { return m.Data2.Participants }

func (m *TLChatFull) SetChatPhoto(v *Photo) { m.Data2.ChatPhoto = v }
func (m *TLChatFull) GetChatPhoto() *Photo  { return m.Data2.ChatPhoto }

func (m *TLChatFull) SetNotifySettings(v *PeerNotifySettings) { m.Data2.NotifySettings = v }
func (m *TLChatFull) GetNotifySettings() *PeerNotifySettings  { return m.Data2.NotifySettings }

func (m *TLChatFull) SetExportedInvite(v *ExportedChatInvite) { m.Data2.ExportedInvite = v }
func (m *TLChatFull) GetExportedInvite() *ExportedChatInvite  { return m.Data2.ExportedInvite }

func (m *TLChatFull) SetBotInfo(v []*BotInfo) { m.Data2.BotInfo = v }
func (m *TLChatFull) GetBotInfo() []*BotInfo  { return m.Data2.BotInfo }

func (m *TLChatFull) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_chatFull))

	x.Int(m.Data2.Id)
	x.Bytes(m.Data2Participants.Encode())
	x.Bytes(m.Data2ChatPhoto.Encode())
	x.Bytes(m.Data2NotifySettings.Encode())
	x.Bytes(m.Data2ExportedInvite.Encode())

	return x.buf
}

func (m *TLChatFull) Decode(dbuf *DecodeBuf) error {
	m.Data2.Id = x.Int()
	m2 := &Participants{}
	m2.Decode(dbuf)
	m.SetParticipants(m2)
	m3 := &ChatPhoto{}
	m3.Decode(dbuf)
	m.SetChatPhoto(m3)
	m4 := &NotifySettings{}
	m4.Decode(dbuf)
	m.SetNotifySettings(m4)
	m5 := &ExportedInvite{}
	m5.Decode(dbuf)
	m.SetExportedInvite(m5)

	return dbuf.err
}

// channelFull#17f45fcf flags:# can_view_participants:flags.3?true can_set_username:flags.6?true can_set_stickers:flags.7?true id:int about:string participants_count:flags.0?int admins_count:flags.1?int kicked_count:flags.2?int banned_count:flags.2?int read_inbox_max_id:int read_outbox_max_id:int unread_count:int chat_photo:Photo notify_settings:PeerNotifySettings exported_invite:ExportedChatInvite bot_info:Vector<BotInfo> migrated_from_chat_id:flags.4?int migrated_from_max_id:flags.4?int pinned_msg_id:flags.5?int stickerset:flags.8?StickerSet = ChatFull;
func (m *TLChannelFull) To_ChatFull() *ChatFull {
	return &ChatFull{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLChannelFull) SetCanViewParticipants(v bool) { m.Data2.CanViewParticipants = v }
func (m *TLChannelFull) GetCanViewParticipants() bool  { return m.Data2.CanViewParticipants }

func (m *TLChannelFull) SetCanSetUsername(v bool) { m.Data2.CanSetUsername = v }
func (m *TLChannelFull) GetCanSetUsername() bool  { return m.Data2.CanSetUsername }

func (m *TLChannelFull) SetCanSetStickers(v bool) { m.Data2.CanSetStickers = v }
func (m *TLChannelFull) GetCanSetStickers() bool  { return m.Data2.CanSetStickers }

func (m *TLChannelFull) SetId(v int32) { m.Data2.Id = v }
func (m *TLChannelFull) GetId() int32  { return m.Data2.Id }

func (m *TLChannelFull) SetAbout(v string) { m.Data2.About = v }
func (m *TLChannelFull) GetAbout() string  { return m.Data2.About }

func (m *TLChannelFull) SetParticipantsCount(v int32) { m.Data2.ParticipantsCount = v }
func (m *TLChannelFull) GetParticipantsCount() int32  { return m.Data2.ParticipantsCount }

func (m *TLChannelFull) SetAdminsCount(v int32) { m.Data2.AdminsCount = v }
func (m *TLChannelFull) GetAdminsCount() int32  { return m.Data2.AdminsCount }

func (m *TLChannelFull) SetKickedCount(v int32) { m.Data2.KickedCount = v }
func (m *TLChannelFull) GetKickedCount() int32  { return m.Data2.KickedCount }

func (m *TLChannelFull) SetBannedCount(v int32) { m.Data2.BannedCount = v }
func (m *TLChannelFull) GetBannedCount() int32  { return m.Data2.BannedCount }

func (m *TLChannelFull) SetReadInboxMaxId(v int32) { m.Data2.ReadInboxMaxId = v }
func (m *TLChannelFull) GetReadInboxMaxId() int32  { return m.Data2.ReadInboxMaxId }

func (m *TLChannelFull) SetReadOutboxMaxId(v int32) { m.Data2.ReadOutboxMaxId = v }
func (m *TLChannelFull) GetReadOutboxMaxId() int32  { return m.Data2.ReadOutboxMaxId }

func (m *TLChannelFull) SetUnreadCount(v int32) { m.Data2.UnreadCount = v }
func (m *TLChannelFull) GetUnreadCount() int32  { return m.Data2.UnreadCount }

func (m *TLChannelFull) SetChatPhoto(v *Photo) { m.Data2.ChatPhoto = v }
func (m *TLChannelFull) GetChatPhoto() *Photo  { return m.Data2.ChatPhoto }

func (m *TLChannelFull) SetNotifySettings(v *PeerNotifySettings) { m.Data2.NotifySettings = v }
func (m *TLChannelFull) GetNotifySettings() *PeerNotifySettings  { return m.Data2.NotifySettings }

func (m *TLChannelFull) SetExportedInvite(v *ExportedChatInvite) { m.Data2.ExportedInvite = v }
func (m *TLChannelFull) GetExportedInvite() *ExportedChatInvite  { return m.Data2.ExportedInvite }

func (m *TLChannelFull) SetBotInfo(v []*BotInfo) { m.Data2.BotInfo = v }
func (m *TLChannelFull) GetBotInfo() []*BotInfo  { return m.Data2.BotInfo }

func (m *TLChannelFull) SetMigratedFromChatId(v int32) { m.Data2.MigratedFromChatId = v }
func (m *TLChannelFull) GetMigratedFromChatId() int32  { return m.Data2.MigratedFromChatId }

func (m *TLChannelFull) SetMigratedFromMaxId(v int32) { m.Data2.MigratedFromMaxId = v }
func (m *TLChannelFull) GetMigratedFromMaxId() int32  { return m.Data2.MigratedFromMaxId }

func (m *TLChannelFull) SetPinnedMsgId(v int32) { m.Data2.PinnedMsgId = v }
func (m *TLChannelFull) GetPinnedMsgId() int32  { return m.Data2.PinnedMsgId }

func (m *TLChannelFull) SetStickerset(v *StickerSet) { m.Data2.Stickerset = v }
func (m *TLChannelFull) GetStickerset() *StickerSet  { return m.Data2.Stickerset }

func (m *TLChannelFull) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelFull))

	// flags
	var flags uint32 = 0
	if m.GetCanViewParticipants() == true {
		flags |= 1 << 1
	}
	if m.GetCanSetUsername() == true {
		flags |= 1 << 2
	}
	if m.GetCanSetStickers() == true {
		flags |= 1 << 3
	}
	if m.GetParticipantsCount() != 0 {
		flags |= 1 << 6
	}
	if m.GetAdminsCount() != 0 {
		flags |= 1 << 7
	}
	if m.GetKickedCount() != 0 {
		flags |= 1 << 8
	}
	if m.GetBannedCount() != 0 {
		flags |= 1 << 9
	}
	if m.GetMigratedFromChatId() != 0 {
		flags |= 1 << 17
	}
	if m.GetMigratedFromMaxId() != 0 {
		flags |= 1 << 18
	}
	if m.GetPinnedMsgId() != 0 {
		flags |= 1 << 19
	}
	if m.GetStickerset() != nil {
		flags |= 1 << 20
	}
	x.UInt(flags)

	x.Int(m.Data2.Id)
	x.StringBytes(m.Data2.About)
	if m.GetParticipantsCount() != 0 {
		x.Int(m.Data2.ParticipantsCount)
	}
	if m.GetAdminsCount() != 0 {
		x.Int(m.Data2.AdminsCount)
	}
	if m.GetKickedCount() != 0 {
		x.Int(m.Data2.KickedCount)
	}
	if m.GetBannedCount() != 0 {
		x.Int(m.Data2.BannedCount)
	}
	x.Int(m.Data2.ReadInboxMaxId)
	x.Int(m.Data2.ReadOutboxMaxId)
	x.Int(m.Data2.UnreadCount)
	x.Bytes(m.Data2ChatPhoto.Encode())
	x.Bytes(m.Data2NotifySettings.Encode())
	x.Bytes(m.Data2ExportedInvite.Encode())

	if m.GetMigratedFromChatId() != 0 {
		x.Int(m.Data2.MigratedFromChatId)
	}
	if m.GetMigratedFromMaxId() != 0 {
		x.Int(m.Data2.MigratedFromMaxId)
	}
	if m.GetPinnedMsgId() != 0 {
		x.Int(m.Data2.PinnedMsgId)
	}
	if m.GetStickerset() != 0 {
		x.Bytes(m.Data2Stickerset.Encode())
	}

	return x.buf
}

func (m *TLChannelFull) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m.Data2.CanViewParticipants = true
	}
	if (flags & (1 << 3)) != 0 {
		m.Data2.CanSetUsername = true
	}
	if (flags & (1 << 4)) != 0 {
		m.Data2.CanSetStickers = true
	}
	m.Data2.Id = x.Int()
	m.Data2.About = x.StringBytes()
	if (flags & (1 << 7)) != 0 {
		m.Data2.ParticipantsCount = x.Int()
	}
	if (flags & (1 << 8)) != 0 {
		m.Data2.AdminsCount = x.Int()
	}
	if (flags & (1 << 9)) != 0 {
		m.Data2.KickedCount = x.Int()
	}
	if (flags & (1 << 10)) != 0 {
		m.Data2.BannedCount = x.Int()
	}
	m.Data2.ReadInboxMaxId = x.Int()
	m.Data2.ReadOutboxMaxId = x.Int()
	m.Data2.UnreadCount = x.Int()
	m14 := &ChatPhoto{}
	m14.Decode(dbuf)
	m.SetChatPhoto(m14)
	m15 := &NotifySettings{}
	m15.Decode(dbuf)
	m.SetNotifySettings(m15)
	m16 := &ExportedInvite{}
	m16.Decode(dbuf)
	m.SetExportedInvite(m16)

	if (flags & (1 << 18)) != 0 {
		m.Data2.MigratedFromChatId = x.Int()
	}
	if (flags & (1 << 19)) != 0 {
		m.Data2.MigratedFromMaxId = x.Int()
	}
	if (flags & (1 << 20)) != 0 {
		m.Data2.PinnedMsgId = x.Int()
	}
	if (flags & (1 << 21)) != 0 {
		m21 := &Stickerset{}
		m21.Decode(dbuf)
		m.SetStickerset(m21)
	}

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// ChatPhoto <--
//  + TL_ChatPhotoEmpty
//  + TL_ChatPhoto
//

func (m *ChatPhoto) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_chatPhotoEmpty:
		t := m.To_ChatPhotoEmpty()
		return t.Encode()
	case TLConstructor_CRC32_chatPhoto:
		t := m.To_ChatPhoto()
		return t.Encode()

	default:
		return []byte{}
	}
}

// chatPhotoEmpty#37c1011c = ChatPhoto;
func (m *ChatPhoto) To_ChatPhotoEmpty() *TLChatPhotoEmpty {
	return &TLChatPhotoEmpty{
		Data2: m.Data2,
	}
}

// chatPhoto#6153276a photo_small:FileLocation photo_big:FileLocation = ChatPhoto;
func (m *ChatPhoto) To_ChatPhoto() *TLChatPhoto {
	return &TLChatPhoto{
		Data2: m.Data2,
	}
}

// chatPhotoEmpty#37c1011c = ChatPhoto;
func (m *TLChatPhotoEmpty) To_ChatPhoto() *ChatPhoto {
	return &ChatPhoto{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLChatPhotoEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_chatPhotoEmpty))

	return x.buf
}

func (m *TLChatPhotoEmpty) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// chatPhoto#6153276a photo_small:FileLocation photo_big:FileLocation = ChatPhoto;
func (m *TLChatPhoto) To_ChatPhoto() *ChatPhoto {
	return &ChatPhoto{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLChatPhoto) SetPhotoSmall(v *FileLocation) { m.Data2.PhotoSmall = v }
func (m *TLChatPhoto) GetPhotoSmall() *FileLocation  { return m.Data2.PhotoSmall }

func (m *TLChatPhoto) SetPhotoBig(v *FileLocation) { m.Data2.PhotoBig = v }
func (m *TLChatPhoto) GetPhotoBig() *FileLocation  { return m.Data2.PhotoBig }

func (m *TLChatPhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_chatPhoto))

	x.Bytes(m.Data2PhotoSmall.Encode())
	x.Bytes(m.Data2PhotoBig.Encode())

	return x.buf
}

func (m *TLChatPhoto) Decode(dbuf *DecodeBuf) error {
	m1 := &PhotoSmall{}
	m1.Decode(dbuf)
	m.SetPhotoSmall(m1)
	m2 := &PhotoBig{}
	m2.Decode(dbuf)
	m.SetPhotoBig(m2)

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// MessageMedia <--
//  + TL_MessageMediaEmpty
//  + TL_MessageMediaPhoto
//  + TL_MessageMediaGeo
//  + TL_MessageMediaContact
//  + TL_MessageMediaUnsupported
//  + TL_MessageMediaDocument
//  + TL_MessageMediaWebPage
//  + TL_MessageMediaVenue
//  + TL_MessageMediaGame
//  + TL_MessageMediaInvoice
//

func (m *MessageMedia) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_messageMediaEmpty:
		t := m.To_MessageMediaEmpty()
		return t.Encode()
	case TLConstructor_CRC32_messageMediaPhoto:
		t := m.To_MessageMediaPhoto()
		return t.Encode()
	case TLConstructor_CRC32_messageMediaGeo:
		t := m.To_MessageMediaGeo()
		return t.Encode()
	case TLConstructor_CRC32_messageMediaContact:
		t := m.To_MessageMediaContact()
		return t.Encode()
	case TLConstructor_CRC32_messageMediaUnsupported:
		t := m.To_MessageMediaUnsupported()
		return t.Encode()
	case TLConstructor_CRC32_messageMediaDocument:
		t := m.To_MessageMediaDocument()
		return t.Encode()
	case TLConstructor_CRC32_messageMediaWebPage:
		t := m.To_MessageMediaWebPage()
		return t.Encode()
	case TLConstructor_CRC32_messageMediaVenue:
		t := m.To_MessageMediaVenue()
		return t.Encode()
	case TLConstructor_CRC32_messageMediaGame:
		t := m.To_MessageMediaGame()
		return t.Encode()
	case TLConstructor_CRC32_messageMediaInvoice:
		t := m.To_MessageMediaInvoice()
		return t.Encode()

	default:
		return []byte{}
	}
}

// messageMediaEmpty#3ded6320 = MessageMedia;
func (m *MessageMedia) To_MessageMediaEmpty() *TLMessageMediaEmpty {
	return &TLMessageMediaEmpty{
		Data2: m.Data2,
	}
}

// messageMediaPhoto#b5223b0f flags:# photo:flags.0?Photo caption:flags.1?string ttl_seconds:flags.2?int = MessageMedia;
func (m *MessageMedia) To_MessageMediaPhoto() *TLMessageMediaPhoto {
	return &TLMessageMediaPhoto{
		Data2: m.Data2,
	}
}

// messageMediaGeo#56e0d474 geo:GeoPoint = MessageMedia;
func (m *MessageMedia) To_MessageMediaGeo() *TLMessageMediaGeo {
	return &TLMessageMediaGeo{
		Data2: m.Data2,
	}
}

// messageMediaContact#5e7d2f39 phone_number:string first_name:string last_name:string user_id:int = MessageMedia;
func (m *MessageMedia) To_MessageMediaContact() *TLMessageMediaContact {
	return &TLMessageMediaContact{
		Data2: m.Data2,
	}
}

// messageMediaUnsupported#9f84f49e = MessageMedia;
func (m *MessageMedia) To_MessageMediaUnsupported() *TLMessageMediaUnsupported {
	return &TLMessageMediaUnsupported{
		Data2: m.Data2,
	}
}

// messageMediaDocument#7c4414d3 flags:# document:flags.0?Document caption:flags.1?string ttl_seconds:flags.2?int = MessageMedia;
func (m *MessageMedia) To_MessageMediaDocument() *TLMessageMediaDocument {
	return &TLMessageMediaDocument{
		Data2: m.Data2,
	}
}

// messageMediaWebPage#a32dd600 webpage:WebPage = MessageMedia;
func (m *MessageMedia) To_MessageMediaWebPage() *TLMessageMediaWebPage {
	return &TLMessageMediaWebPage{
		Data2: m.Data2,
	}
}

// messageMediaVenue#7912b71f geo:GeoPoint title:string address:string provider:string venue_id:string = MessageMedia;
func (m *MessageMedia) To_MessageMediaVenue() *TLMessageMediaVenue {
	return &TLMessageMediaVenue{
		Data2: m.Data2,
	}
}

// messageMediaGame#fdb19008 game:Game = MessageMedia;
func (m *MessageMedia) To_MessageMediaGame() *TLMessageMediaGame {
	return &TLMessageMediaGame{
		Data2: m.Data2,
	}
}

// messageMediaInvoice#84551347 flags:# shipping_address_requested:flags.1?true test:flags.3?true title:string description:string photo:flags.0?WebDocument receipt_msg_id:flags.2?int currency:string total_amount:long start_param:string = MessageMedia;
func (m *MessageMedia) To_MessageMediaInvoice() *TLMessageMediaInvoice {
	return &TLMessageMediaInvoice{
		Data2: m.Data2,
	}
}

// messageMediaEmpty#3ded6320 = MessageMedia;
func (m *TLMessageMediaEmpty) To_MessageMedia() *MessageMedia {
	return &MessageMedia{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessageMediaEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageMediaEmpty))

	return x.buf
}

func (m *TLMessageMediaEmpty) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// messageMediaPhoto#b5223b0f flags:# photo:flags.0?Photo caption:flags.1?string ttl_seconds:flags.2?int = MessageMedia;
func (m *TLMessageMediaPhoto) To_MessageMedia() *MessageMedia {
	return &MessageMedia{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessageMediaPhoto) SetPhoto(v *Photo) { m.Data2.Photo_1 = v }
func (m *TLMessageMediaPhoto) GetPhoto() *Photo  { return m.Data2.Photo_1 }

func (m *TLMessageMediaPhoto) SetCaption(v string) { m.Data2.Caption = v }
func (m *TLMessageMediaPhoto) GetCaption() string  { return m.Data2.Caption }

func (m *TLMessageMediaPhoto) SetTtlSeconds(v int32) { m.Data2.TtlSeconds = v }
func (m *TLMessageMediaPhoto) GetTtlSeconds() int32  { return m.Data2.TtlSeconds }

func (m *TLMessageMediaPhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageMediaPhoto))

	// flags
	var flags uint32 = 0
	if m.GetPhoto() != nil {
		flags |= 1 << 1
	}
	if m.GetCaption() != "" {
		flags |= 1 << 2
	}
	if m.GetTtlSeconds() != 0 {
		flags |= 1 << 3
	}
	x.UInt(flags)

	if m.GetPhoto() != 0 {
		x.Bytes(m.Data2Photo.Encode())
	}
	if m.GetCaption() != 0 {
		x.StringBytes(m.Data2.Caption)
	}
	if m.GetTtlSeconds() != 0 {
		x.Int(m.Data2.TtlSeconds)
	}

	return x.buf
}

func (m *TLMessageMediaPhoto) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m2 := &Photo{}
		m2.Decode(dbuf)
		m.SetPhoto(m2)
	}
	if (flags & (1 << 3)) != 0 {
		m.Data2.Caption = x.StringBytes()
	}
	if (flags & (1 << 4)) != 0 {
		m.Data2.TtlSeconds = x.Int()
	}

	return dbuf.err
}

// messageMediaGeo#56e0d474 geo:GeoPoint = MessageMedia;
func (m *TLMessageMediaGeo) To_MessageMedia() *MessageMedia {
	return &MessageMedia{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessageMediaGeo) SetGeo(v *GeoPoint) { m.Data2.Geo = v }
func (m *TLMessageMediaGeo) GetGeo() *GeoPoint  { return m.Data2.Geo }

func (m *TLMessageMediaGeo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageMediaGeo))

	x.Bytes(m.Data2Geo.Encode())

	return x.buf
}

func (m *TLMessageMediaGeo) Decode(dbuf *DecodeBuf) error {
	m1 := &Geo{}
	m1.Decode(dbuf)
	m.SetGeo(m1)

	return dbuf.err
}

// messageMediaContact#5e7d2f39 phone_number:string first_name:string last_name:string user_id:int = MessageMedia;
func (m *TLMessageMediaContact) To_MessageMedia() *MessageMedia {
	return &MessageMedia{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessageMediaContact) SetPhoneNumber(v string) { m.Data2.PhoneNumber = v }
func (m *TLMessageMediaContact) GetPhoneNumber() string  { return m.Data2.PhoneNumber }

func (m *TLMessageMediaContact) SetFirstName(v string) { m.Data2.FirstName = v }
func (m *TLMessageMediaContact) GetFirstName() string  { return m.Data2.FirstName }

func (m *TLMessageMediaContact) SetLastName(v string) { m.Data2.LastName = v }
func (m *TLMessageMediaContact) GetLastName() string  { return m.Data2.LastName }

func (m *TLMessageMediaContact) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLMessageMediaContact) GetUserId() int32  { return m.Data2.UserId }

func (m *TLMessageMediaContact) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageMediaContact))

	x.StringBytes(m.Data2.PhoneNumber)
	x.StringBytes(m.Data2.FirstName)
	x.StringBytes(m.Data2.LastName)
	x.Int(m.Data2.UserId)

	return x.buf
}

func (m *TLMessageMediaContact) Decode(dbuf *DecodeBuf) error {
	m.Data2.PhoneNumber = x.StringBytes()
	m.Data2.FirstName = x.StringBytes()
	m.Data2.LastName = x.StringBytes()
	m.Data2.UserId = x.Int()

	return dbuf.err
}

// messageMediaUnsupported#9f84f49e = MessageMedia;
func (m *TLMessageMediaUnsupported) To_MessageMedia() *MessageMedia {
	return &MessageMedia{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessageMediaUnsupported) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageMediaUnsupported))

	return x.buf
}

func (m *TLMessageMediaUnsupported) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// messageMediaDocument#7c4414d3 flags:# document:flags.0?Document caption:flags.1?string ttl_seconds:flags.2?int = MessageMedia;
func (m *TLMessageMediaDocument) To_MessageMedia() *MessageMedia {
	return &MessageMedia{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessageMediaDocument) SetDocument(v *Document) { m.Data2.Document = v }
func (m *TLMessageMediaDocument) GetDocument() *Document  { return m.Data2.Document }

func (m *TLMessageMediaDocument) SetCaption(v string) { m.Data2.Caption = v }
func (m *TLMessageMediaDocument) GetCaption() string  { return m.Data2.Caption }

func (m *TLMessageMediaDocument) SetTtlSeconds(v int32) { m.Data2.TtlSeconds = v }
func (m *TLMessageMediaDocument) GetTtlSeconds() int32  { return m.Data2.TtlSeconds }

func (m *TLMessageMediaDocument) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageMediaDocument))

	// flags
	var flags uint32 = 0
	if m.GetDocument() != nil {
		flags |= 1 << 1
	}
	if m.GetCaption() != "" {
		flags |= 1 << 2
	}
	if m.GetTtlSeconds() != 0 {
		flags |= 1 << 3
	}
	x.UInt(flags)

	if m.GetDocument() != 0 {
		x.Bytes(m.Data2Document.Encode())
	}
	if m.GetCaption() != 0 {
		x.StringBytes(m.Data2.Caption)
	}
	if m.GetTtlSeconds() != 0 {
		x.Int(m.Data2.TtlSeconds)
	}

	return x.buf
}

func (m *TLMessageMediaDocument) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m2 := &Document{}
		m2.Decode(dbuf)
		m.SetDocument(m2)
	}
	if (flags & (1 << 3)) != 0 {
		m.Data2.Caption = x.StringBytes()
	}
	if (flags & (1 << 4)) != 0 {
		m.Data2.TtlSeconds = x.Int()
	}

	return dbuf.err
}

// messageMediaWebPage#a32dd600 webpage:WebPage = MessageMedia;
func (m *TLMessageMediaWebPage) To_MessageMedia() *MessageMedia {
	return &MessageMedia{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessageMediaWebPage) SetWebpage(v *WebPage) { m.Data2.Webpage = v }
func (m *TLMessageMediaWebPage) GetWebpage() *WebPage  { return m.Data2.Webpage }

func (m *TLMessageMediaWebPage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageMediaWebPage))

	x.Bytes(m.Data2Webpage.Encode())

	return x.buf
}

func (m *TLMessageMediaWebPage) Decode(dbuf *DecodeBuf) error {
	m1 := &Webpage{}
	m1.Decode(dbuf)
	m.SetWebpage(m1)

	return dbuf.err
}

// messageMediaVenue#7912b71f geo:GeoPoint title:string address:string provider:string venue_id:string = MessageMedia;
func (m *TLMessageMediaVenue) To_MessageMedia() *MessageMedia {
	return &MessageMedia{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessageMediaVenue) SetGeo(v *GeoPoint) { m.Data2.Geo = v }
func (m *TLMessageMediaVenue) GetGeo() *GeoPoint  { return m.Data2.Geo }

func (m *TLMessageMediaVenue) SetTitle(v string) { m.Data2.Title = v }
func (m *TLMessageMediaVenue) GetTitle() string  { return m.Data2.Title }

func (m *TLMessageMediaVenue) SetAddress(v string) { m.Data2.Address = v }
func (m *TLMessageMediaVenue) GetAddress() string  { return m.Data2.Address }

func (m *TLMessageMediaVenue) SetProvider(v string) { m.Data2.Provider = v }
func (m *TLMessageMediaVenue) GetProvider() string  { return m.Data2.Provider }

func (m *TLMessageMediaVenue) SetVenueId(v string) { m.Data2.VenueId = v }
func (m *TLMessageMediaVenue) GetVenueId() string  { return m.Data2.VenueId }

func (m *TLMessageMediaVenue) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageMediaVenue))

	x.Bytes(m.Data2Geo.Encode())
	x.StringBytes(m.Data2.Title)
	x.StringBytes(m.Data2.Address)
	x.StringBytes(m.Data2.Provider)
	x.StringBytes(m.Data2.VenueId)

	return x.buf
}

func (m *TLMessageMediaVenue) Decode(dbuf *DecodeBuf) error {
	m1 := &Geo{}
	m1.Decode(dbuf)
	m.SetGeo(m1)
	m.Data2.Title = x.StringBytes()
	m.Data2.Address = x.StringBytes()
	m.Data2.Provider = x.StringBytes()
	m.Data2.VenueId = x.StringBytes()

	return dbuf.err
}

// messageMediaGame#fdb19008 game:Game = MessageMedia;
func (m *TLMessageMediaGame) To_MessageMedia() *MessageMedia {
	return &MessageMedia{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessageMediaGame) SetGame(v *Game) { m.Data2.Game = v }
func (m *TLMessageMediaGame) GetGame() *Game  { return m.Data2.Game }

func (m *TLMessageMediaGame) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageMediaGame))

	x.Bytes(m.Data2Game.Encode())

	return x.buf
}

func (m *TLMessageMediaGame) Decode(dbuf *DecodeBuf) error {
	m1 := &Game{}
	m1.Decode(dbuf)
	m.SetGame(m1)

	return dbuf.err
}

// messageMediaInvoice#84551347 flags:# shipping_address_requested:flags.1?true test:flags.3?true title:string description:string photo:flags.0?WebDocument receipt_msg_id:flags.2?int currency:string total_amount:long start_param:string = MessageMedia;
func (m *TLMessageMediaInvoice) To_MessageMedia() *MessageMedia {
	return &MessageMedia{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessageMediaInvoice) SetShippingAddressRequested(v bool) {
	m.Data2.ShippingAddressRequested = v
}
func (m *TLMessageMediaInvoice) GetShippingAddressRequested() bool {
	return m.Data2.ShippingAddressRequested
}

func (m *TLMessageMediaInvoice) SetTest(v bool) { m.Data2.Test = v }
func (m *TLMessageMediaInvoice) GetTest() bool  { return m.Data2.Test }

func (m *TLMessageMediaInvoice) SetTitle(v string) { m.Data2.Title = v }
func (m *TLMessageMediaInvoice) GetTitle() string  { return m.Data2.Title }

func (m *TLMessageMediaInvoice) SetDescription(v string) { m.Data2.Description = v }
func (m *TLMessageMediaInvoice) GetDescription() string  { return m.Data2.Description }

func (m *TLMessageMediaInvoice) SetPhoto(v *WebDocument) { m.Data2.Photo_19 = v }
func (m *TLMessageMediaInvoice) GetPhoto() *WebDocument  { return m.Data2.Photo_19 }

func (m *TLMessageMediaInvoice) SetReceiptMsgId(v int32) { m.Data2.ReceiptMsgId = v }
func (m *TLMessageMediaInvoice) GetReceiptMsgId() int32  { return m.Data2.ReceiptMsgId }

func (m *TLMessageMediaInvoice) SetCurrency(v string) { m.Data2.Currency = v }
func (m *TLMessageMediaInvoice) GetCurrency() string  { return m.Data2.Currency }

func (m *TLMessageMediaInvoice) SetTotalAmount(v int64) { m.Data2.TotalAmount = v }
func (m *TLMessageMediaInvoice) GetTotalAmount() int64  { return m.Data2.TotalAmount }

func (m *TLMessageMediaInvoice) SetStartParam(v string) { m.Data2.StartParam = v }
func (m *TLMessageMediaInvoice) GetStartParam() string  { return m.Data2.StartParam }

func (m *TLMessageMediaInvoice) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageMediaInvoice))

	// flags
	var flags uint32 = 0
	if m.GetShippingAddressRequested() == true {
		flags |= 1 << 1
	}
	if m.GetTest() == true {
		flags |= 1 << 2
	}
	if m.GetPhoto() != nil {
		flags |= 1 << 5
	}
	if m.GetReceiptMsgId() != 0 {
		flags |= 1 << 6
	}
	x.UInt(flags)

	x.StringBytes(m.Data2.Title)
	x.StringBytes(m.Data2.Description)
	if m.GetPhoto() != 0 {
		x.Bytes(m.Data2Photo.Encode())
	}
	if m.GetReceiptMsgId() != 0 {
		x.Int(m.Data2.ReceiptMsgId)
	}
	x.StringBytes(m.Data2.Currency)
	x.Long(m.Data2.TotalAmount)
	x.StringBytes(m.Data2.StartParam)

	return x.buf
}

func (m *TLMessageMediaInvoice) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m.Data2.ShippingAddressRequested = true
	}
	if (flags & (1 << 3)) != 0 {
		m.Data2.Test = true
	}
	m.Data2.Title = x.StringBytes()
	m.Data2.Description = x.StringBytes()
	if (flags & (1 << 6)) != 0 {
		m6 := &Photo{}
		m6.Decode(dbuf)
		m.SetPhoto(m6)
	}
	if (flags & (1 << 7)) != 0 {
		m.Data2.ReceiptMsgId = x.Int()
	}
	m.Data2.Currency = x.StringBytes()
	m.Data2.TotalAmount = x.Long()
	m.Data2.StartParam = x.StringBytes()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Photo <--
//  + TL_PhotoEmpty
//  + TL_Photo
//

func (m *Photo) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_photoEmpty:
		t := m.To_PhotoEmpty()
		return t.Encode()
	case TLConstructor_CRC32_photo:
		t := m.To_Photo()
		return t.Encode()

	default:
		return []byte{}
	}
}

// photoEmpty#2331b22d id:long = Photo;
func (m *Photo) To_PhotoEmpty() *TLPhotoEmpty {
	return &TLPhotoEmpty{
		Data2: m.Data2,
	}
}

// photo#9288dd29 flags:# has_stickers:flags.0?true id:long access_hash:long date:int sizes:Vector<PhotoSize> = Photo;
func (m *Photo) To_Photo() *TLPhoto {
	return &TLPhoto{
		Data2: m.Data2,
	}
}

// photoEmpty#2331b22d id:long = Photo;
func (m *TLPhotoEmpty) To_Photo() *Photo {
	return &Photo{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPhotoEmpty) SetId(v int64) { m.Data2.Id = v }
func (m *TLPhotoEmpty) GetId() int64  { return m.Data2.Id }

func (m *TLPhotoEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_photoEmpty))

	x.Long(m.Data2.Id)

	return x.buf
}

func (m *TLPhotoEmpty) Decode(dbuf *DecodeBuf) error {
	m.Data2.Id = x.Long()

	return dbuf.err
}

// photo#9288dd29 flags:# has_stickers:flags.0?true id:long access_hash:long date:int sizes:Vector<PhotoSize> = Photo;
func (m *TLPhoto) To_Photo() *Photo {
	return &Photo{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPhoto) SetHasStickers(v bool) { m.Data2.HasStickers = v }
func (m *TLPhoto) GetHasStickers() bool  { return m.Data2.HasStickers }

func (m *TLPhoto) SetId(v int64) { m.Data2.Id = v }
func (m *TLPhoto) GetId() int64  { return m.Data2.Id }

func (m *TLPhoto) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLPhoto) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLPhoto) SetDate(v int32) { m.Data2.Date = v }
func (m *TLPhoto) GetDate() int32  { return m.Data2.Date }

func (m *TLPhoto) SetSizes(v []*PhotoSize) { m.Data2.Sizes = v }
func (m *TLPhoto) GetSizes() []*PhotoSize  { return m.Data2.Sizes }

func (m *TLPhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_photo))

	// flags
	var flags uint32 = 0
	if m.GetHasStickers() == true {
		flags |= 1 << 1
	}
	x.UInt(flags)

	x.Long(m.Data2.Id)
	x.Long(m.Data2.AccessHash)
	x.Int(m.Data2.Date)

	return x.buf
}

func (m *TLPhoto) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m.Data2.HasStickers = true
	}
	m.Data2.Id = x.Long()
	m.Data2.AccessHash = x.Long()
	m.Data2.Date = x.Int()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Updates_ChannelDifference <--
//  + TL_UpdatesChannelDifferenceEmpty
//  + TL_UpdatesChannelDifferenceTooLong
//  + TL_UpdatesChannelDifference
//

func (m *Updates_ChannelDifference) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_updates_channelDifferenceEmpty:
		t := m.To_UpdatesChannelDifferenceEmpty()
		return t.Encode()
	case TLConstructor_CRC32_updates_channelDifferenceTooLong:
		t := m.To_UpdatesChannelDifferenceTooLong()
		return t.Encode()
	case TLConstructor_CRC32_updates_channelDifference:
		t := m.To_UpdatesChannelDifference()
		return t.Encode()

	default:
		return []byte{}
	}
}

// updates.channelDifferenceEmpty#3e11affb flags:# final:flags.0?true pts:int timeout:flags.1?int = updates.ChannelDifference;
func (m *Updates_ChannelDifference) To_UpdatesChannelDifferenceEmpty() *TLUpdatesChannelDifferenceEmpty {
	return &TLUpdatesChannelDifferenceEmpty{
		Data2: m.Data2,
	}
}

// updates.channelDifferenceTooLong#6a9d7b35 flags:# final:flags.0?true pts:int timeout:flags.1?int top_message:int read_inbox_max_id:int read_outbox_max_id:int unread_count:int unread_mentions_count:int messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = updates.ChannelDifference;
func (m *Updates_ChannelDifference) To_UpdatesChannelDifferenceTooLong() *TLUpdatesChannelDifferenceTooLong {
	return &TLUpdatesChannelDifferenceTooLong{
		Data2: m.Data2,
	}
}

// updates.channelDifference#2064674e flags:# final:flags.0?true pts:int timeout:flags.1?int new_messages:Vector<Message> other_updates:Vector<Update> chats:Vector<Chat> users:Vector<User> = updates.ChannelDifference;
func (m *Updates_ChannelDifference) To_UpdatesChannelDifference() *TLUpdatesChannelDifference {
	return &TLUpdatesChannelDifference{
		Data2: m.Data2,
	}
}

// updates.channelDifferenceEmpty#3e11affb flags:# final:flags.0?true pts:int timeout:flags.1?int = updates.ChannelDifference;
func (m *TLUpdatesChannelDifferenceEmpty) To_Updates_ChannelDifference() *Updates_ChannelDifference {
	return &Updates_ChannelDifference{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdatesChannelDifferenceEmpty) SetFinal(v bool) { m.Data2.Final = v }
func (m *TLUpdatesChannelDifferenceEmpty) GetFinal() bool  { return m.Data2.Final }

func (m *TLUpdatesChannelDifferenceEmpty) SetPts(v int32) { m.Data2.Pts = v }
func (m *TLUpdatesChannelDifferenceEmpty) GetPts() int32  { return m.Data2.Pts }

func (m *TLUpdatesChannelDifferenceEmpty) SetTimeout(v int32) { m.Data2.Timeout = v }
func (m *TLUpdatesChannelDifferenceEmpty) GetTimeout() int32  { return m.Data2.Timeout }

func (m *TLUpdatesChannelDifferenceEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updates_channelDifferenceEmpty))

	// flags
	var flags uint32 = 0
	if m.GetFinal() == true {
		flags |= 1 << 1
	}
	if m.GetTimeout() != 0 {
		flags |= 1 << 3
	}
	x.UInt(flags)

	x.Int(m.Data2.Pts)
	if m.GetTimeout() != 0 {
		x.Int(m.Data2.Timeout)
	}

	return x.buf
}

func (m *TLUpdatesChannelDifferenceEmpty) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m.Data2.Final = true
	}
	m.Data2.Pts = x.Int()
	if (flags & (1 << 4)) != 0 {
		m.Data2.Timeout = x.Int()
	}

	return dbuf.err
}

// updates.channelDifferenceTooLong#6a9d7b35 flags:# final:flags.0?true pts:int timeout:flags.1?int top_message:int read_inbox_max_id:int read_outbox_max_id:int unread_count:int unread_mentions_count:int messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = updates.ChannelDifference;
func (m *TLUpdatesChannelDifferenceTooLong) To_Updates_ChannelDifference() *Updates_ChannelDifference {
	return &Updates_ChannelDifference{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdatesChannelDifferenceTooLong) SetFinal(v bool) { m.Data2.Final = v }
func (m *TLUpdatesChannelDifferenceTooLong) GetFinal() bool  { return m.Data2.Final }

func (m *TLUpdatesChannelDifferenceTooLong) SetPts(v int32) { m.Data2.Pts = v }
func (m *TLUpdatesChannelDifferenceTooLong) GetPts() int32  { return m.Data2.Pts }

func (m *TLUpdatesChannelDifferenceTooLong) SetTimeout(v int32) { m.Data2.Timeout = v }
func (m *TLUpdatesChannelDifferenceTooLong) GetTimeout() int32  { return m.Data2.Timeout }

func (m *TLUpdatesChannelDifferenceTooLong) SetTopMessage(v int32) { m.Data2.TopMessage = v }
func (m *TLUpdatesChannelDifferenceTooLong) GetTopMessage() int32  { return m.Data2.TopMessage }

func (m *TLUpdatesChannelDifferenceTooLong) SetReadInboxMaxId(v int32) { m.Data2.ReadInboxMaxId = v }
func (m *TLUpdatesChannelDifferenceTooLong) GetReadInboxMaxId() int32  { return m.Data2.ReadInboxMaxId }

func (m *TLUpdatesChannelDifferenceTooLong) SetReadOutboxMaxId(v int32) { m.Data2.ReadOutboxMaxId = v }
func (m *TLUpdatesChannelDifferenceTooLong) GetReadOutboxMaxId() int32  { return m.Data2.ReadOutboxMaxId }

func (m *TLUpdatesChannelDifferenceTooLong) SetUnreadCount(v int32) { m.Data2.UnreadCount = v }
func (m *TLUpdatesChannelDifferenceTooLong) GetUnreadCount() int32  { return m.Data2.UnreadCount }

func (m *TLUpdatesChannelDifferenceTooLong) SetUnreadMentionsCount(v int32) {
	m.Data2.UnreadMentionsCount = v
}
func (m *TLUpdatesChannelDifferenceTooLong) GetUnreadMentionsCount() int32 {
	return m.Data2.UnreadMentionsCount
}

func (m *TLUpdatesChannelDifferenceTooLong) SetMessages(v []*Message) { m.Data2.Messages = v }
func (m *TLUpdatesChannelDifferenceTooLong) GetMessages() []*Message  { return m.Data2.Messages }

func (m *TLUpdatesChannelDifferenceTooLong) SetChats(v []*Chat) { m.Data2.Chats = v }
func (m *TLUpdatesChannelDifferenceTooLong) GetChats() []*Chat  { return m.Data2.Chats }

func (m *TLUpdatesChannelDifferenceTooLong) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLUpdatesChannelDifferenceTooLong) GetUsers() []*User  { return m.Data2.Users }

func (m *TLUpdatesChannelDifferenceTooLong) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updates_channelDifferenceTooLong))

	// flags
	var flags uint32 = 0
	if m.GetFinal() == true {
		flags |= 1 << 1
	}
	if m.GetTimeout() != 0 {
		flags |= 1 << 3
	}
	x.UInt(flags)

	x.Int(m.Data2.Pts)
	if m.GetTimeout() != 0 {
		x.Int(m.Data2.Timeout)
	}
	x.Int(m.Data2.TopMessage)
	x.Int(m.Data2.ReadInboxMaxId)
	x.Int(m.Data2.ReadOutboxMaxId)
	x.Int(m.Data2.UnreadCount)
	x.Int(m.Data2.UnreadMentionsCount)

	return x.buf
}

func (m *TLUpdatesChannelDifferenceTooLong) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m.Data2.Final = true
	}
	m.Data2.Pts = x.Int()
	if (flags & (1 << 4)) != 0 {
		m.Data2.Timeout = x.Int()
	}
	m.Data2.TopMessage = x.Int()
	m.Data2.ReadInboxMaxId = x.Int()
	m.Data2.ReadOutboxMaxId = x.Int()
	m.Data2.UnreadCount = x.Int()
	m.Data2.UnreadMentionsCount = x.Int()

	return dbuf.err
}

// updates.channelDifference#2064674e flags:# final:flags.0?true pts:int timeout:flags.1?int new_messages:Vector<Message> other_updates:Vector<Update> chats:Vector<Chat> users:Vector<User> = updates.ChannelDifference;
func (m *TLUpdatesChannelDifference) To_Updates_ChannelDifference() *Updates_ChannelDifference {
	return &Updates_ChannelDifference{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdatesChannelDifference) SetFinal(v bool) { m.Data2.Final = v }
func (m *TLUpdatesChannelDifference) GetFinal() bool  { return m.Data2.Final }

func (m *TLUpdatesChannelDifference) SetPts(v int32) { m.Data2.Pts = v }
func (m *TLUpdatesChannelDifference) GetPts() int32  { return m.Data2.Pts }

func (m *TLUpdatesChannelDifference) SetTimeout(v int32) { m.Data2.Timeout = v }
func (m *TLUpdatesChannelDifference) GetTimeout() int32  { return m.Data2.Timeout }

func (m *TLUpdatesChannelDifference) SetNewMessages(v []*Message) { m.Data2.NewMessages = v }
func (m *TLUpdatesChannelDifference) GetNewMessages() []*Message  { return m.Data2.NewMessages }

func (m *TLUpdatesChannelDifference) SetOtherUpdates(v []*Update) { m.Data2.OtherUpdates = v }
func (m *TLUpdatesChannelDifference) GetOtherUpdates() []*Update  { return m.Data2.OtherUpdates }

func (m *TLUpdatesChannelDifference) SetChats(v []*Chat) { m.Data2.Chats = v }
func (m *TLUpdatesChannelDifference) GetChats() []*Chat  { return m.Data2.Chats }

func (m *TLUpdatesChannelDifference) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLUpdatesChannelDifference) GetUsers() []*User  { return m.Data2.Users }

func (m *TLUpdatesChannelDifference) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updates_channelDifference))

	// flags
	var flags uint32 = 0
	if m.GetFinal() == true {
		flags |= 1 << 1
	}
	if m.GetTimeout() != 0 {
		flags |= 1 << 3
	}
	x.UInt(flags)

	x.Int(m.Data2.Pts)
	if m.GetTimeout() != 0 {
		x.Int(m.Data2.Timeout)
	}

	return x.buf
}

func (m *TLUpdatesChannelDifference) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m.Data2.Final = true
	}
	m.Data2.Pts = x.Int()
	if (flags & (1 << 4)) != 0 {
		m.Data2.Timeout = x.Int()
	}

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Channels_ChannelParticipants <--
//  + TL_ChannelsChannelParticipants
//

func (m *Channels_ChannelParticipants) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_channels_channelParticipants:
		t := m.To_ChannelsChannelParticipants()
		return t.Encode()

	default:
		return []byte{}
	}
}

// channels.channelParticipants#f56ee2a8 count:int participants:Vector<ChannelParticipant> users:Vector<User> = channels.ChannelParticipants;
func (m *Channels_ChannelParticipants) To_ChannelsChannelParticipants() *TLChannelsChannelParticipants {
	return &TLChannelsChannelParticipants{
		Data2: m.Data2,
	}
}

// channels.channelParticipants#f56ee2a8 count:int participants:Vector<ChannelParticipant> users:Vector<User> = channels.ChannelParticipants;
func (m *TLChannelsChannelParticipants) To_Channels_ChannelParticipants() *Channels_ChannelParticipants {
	return &Channels_ChannelParticipants{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLChannelsChannelParticipants) SetCount(v int32) { m.Data2.Count = v }
func (m *TLChannelsChannelParticipants) GetCount() int32  { return m.Data2.Count }

func (m *TLChannelsChannelParticipants) SetParticipants(v []*ChannelParticipant) {
	m.Data2.Participants = v
}
func (m *TLChannelsChannelParticipants) GetParticipants() []*ChannelParticipant {
	return m.Data2.Participants
}

func (m *TLChannelsChannelParticipants) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLChannelsChannelParticipants) GetUsers() []*User  { return m.Data2.Users }

func (m *TLChannelsChannelParticipants) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channels_channelParticipants))

	x.Int(m.Data2.Count)

	return x.buf
}

func (m *TLChannelsChannelParticipants) Decode(dbuf *DecodeBuf) error {
	m.Data2.Count = x.Int()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// WebDocument <--
//  + TL_WebDocument
//

func (m *WebDocument) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_webDocument:
		t := m.To_WebDocument()
		return t.Encode()

	default:
		return []byte{}
	}
}

// webDocument#c61acbd8 url:string access_hash:long size:int mime_type:string attributes:Vector<DocumentAttribute> dc_id:int = WebDocument;
func (m *WebDocument) To_WebDocument() *TLWebDocument {
	return &TLWebDocument{
		Data2: m.Data2,
	}
}

// webDocument#c61acbd8 url:string access_hash:long size:int mime_type:string attributes:Vector<DocumentAttribute> dc_id:int = WebDocument;
func (m *TLWebDocument) To_WebDocument() *WebDocument {
	return &WebDocument{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLWebDocument) SetUrl(v string) { m.Data2.Url = v }
func (m *TLWebDocument) GetUrl() string  { return m.Data2.Url }

func (m *TLWebDocument) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLWebDocument) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLWebDocument) SetSize(v int32) { m.Data2.Size = v }
func (m *TLWebDocument) GetSize() int32  { return m.Data2.Size }

func (m *TLWebDocument) SetMimeType(v string) { m.Data2.MimeType = v }
func (m *TLWebDocument) GetMimeType() string  { return m.Data2.MimeType }

func (m *TLWebDocument) SetAttributes(v []*DocumentAttribute) { m.Data2.Attributes = v }
func (m *TLWebDocument) GetAttributes() []*DocumentAttribute  { return m.Data2.Attributes }

func (m *TLWebDocument) SetDcId(v int32) { m.Data2.DcId = v }
func (m *TLWebDocument) GetDcId() int32  { return m.Data2.DcId }

func (m *TLWebDocument) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_webDocument))

	x.StringBytes(m.Data2.Url)
	x.Long(m.Data2.AccessHash)
	x.Int(m.Data2.Size)
	x.StringBytes(m.Data2.MimeType)

	x.Int(m.Data2.DcId)

	return x.buf
}

func (m *TLWebDocument) Decode(dbuf *DecodeBuf) error {
	m.Data2.Url = x.StringBytes()
	m.Data2.AccessHash = x.Long()
	m.Data2.Size = x.Int()
	m.Data2.MimeType = x.StringBytes()

	m.Data2.DcId = x.Int()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Upload_CdnFile <--
//  + TL_UploadCdnFileReuploadNeeded
//  + TL_UploadCdnFile
//

func (m *Upload_CdnFile) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_upload_cdnFileReuploadNeeded:
		t := m.To_UploadCdnFileReuploadNeeded()
		return t.Encode()
	case TLConstructor_CRC32_upload_cdnFile:
		t := m.To_UploadCdnFile()
		return t.Encode()

	default:
		return []byte{}
	}
}

// upload.cdnFileReuploadNeeded#eea8e46e request_token:bytes = upload.CdnFile;
func (m *Upload_CdnFile) To_UploadCdnFileReuploadNeeded() *TLUploadCdnFileReuploadNeeded {
	return &TLUploadCdnFileReuploadNeeded{
		Data2: m.Data2,
	}
}

// upload.cdnFile#a99fca4f bytes:bytes = upload.CdnFile;
func (m *Upload_CdnFile) To_UploadCdnFile() *TLUploadCdnFile {
	return &TLUploadCdnFile{
		Data2: m.Data2,
	}
}

// upload.cdnFileReuploadNeeded#eea8e46e request_token:bytes = upload.CdnFile;
func (m *TLUploadCdnFileReuploadNeeded) To_Upload_CdnFile() *Upload_CdnFile {
	return &Upload_CdnFile{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUploadCdnFileReuploadNeeded) SetRequestToken(v []byte) { m.Data2.RequestToken = v }
func (m *TLUploadCdnFileReuploadNeeded) GetRequestToken() []byte  { return m.Data2.RequestToken }

func (m *TLUploadCdnFileReuploadNeeded) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_upload_cdnFileReuploadNeeded))

	x.StringBytes(m.Data2.RequestToken)

	return x.buf
}

func (m *TLUploadCdnFileReuploadNeeded) Decode(dbuf *DecodeBuf) error {
	m.Data2.RequestToken = x.StringBytes()

	return dbuf.err
}

// upload.cdnFile#a99fca4f bytes:bytes = upload.CdnFile;
func (m *TLUploadCdnFile) To_Upload_CdnFile() *Upload_CdnFile {
	return &Upload_CdnFile{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUploadCdnFile) SetBytes(v []byte) { m.Data2.Bytes = v }
func (m *TLUploadCdnFile) GetBytes() []byte  { return m.Data2.Bytes }

func (m *TLUploadCdnFile) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_upload_cdnFile))

	x.StringBytes(m.Data2.Bytes)

	return x.buf
}

func (m *TLUploadCdnFile) Decode(dbuf *DecodeBuf) error {
	m.Data2.Bytes = x.StringBytes()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Null <--
//  + TL_Null
//

func (m *Null) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_null:
		t := m.To_Null()
		return t.Encode()

	default:
		return []byte{}
	}
}

// null#56730bcc = Null;
func (m *Null) To_Null() *TLNull {
	return &TLNull{
		Data2: m.Data2,
	}
}

// null#56730bcc = Null;
func (m *TLNull) To_Null() *Null {
	return &Null{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLNull) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_null))

	return x.buf
}

func (m *TLNull) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// InputFileLocation <--
//  + TL_InputFileLocation
//  + TL_InputEncryptedFileLocation
//  + TL_InputDocumentFileLocation
//

func (m *InputFileLocation) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_inputFileLocation:
		t := m.To_InputFileLocation()
		return t.Encode()
	case TLConstructor_CRC32_inputEncryptedFileLocation:
		t := m.To_InputEncryptedFileLocation()
		return t.Encode()
	case TLConstructor_CRC32_inputDocumentFileLocation:
		t := m.To_InputDocumentFileLocation()
		return t.Encode()

	default:
		return []byte{}
	}
}

// inputFileLocation#14637196 volume_id:long local_id:int secret:long = InputFileLocation;
func (m *InputFileLocation) To_InputFileLocation() *TLInputFileLocation {
	return &TLInputFileLocation{
		Data2: m.Data2,
	}
}

// inputEncryptedFileLocation#f5235d55 id:long access_hash:long = InputFileLocation;
func (m *InputFileLocation) To_InputEncryptedFileLocation() *TLInputEncryptedFileLocation {
	return &TLInputEncryptedFileLocation{
		Data2: m.Data2,
	}
}

// inputDocumentFileLocation#430f0724 id:long access_hash:long version:int = InputFileLocation;
func (m *InputFileLocation) To_InputDocumentFileLocation() *TLInputDocumentFileLocation {
	return &TLInputDocumentFileLocation{
		Data2: m.Data2,
	}
}

// inputFileLocation#14637196 volume_id:long local_id:int secret:long = InputFileLocation;
func (m *TLInputFileLocation) To_InputFileLocation() *InputFileLocation {
	return &InputFileLocation{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputFileLocation) SetVolumeId(v int64) { m.Data2.VolumeId = v }
func (m *TLInputFileLocation) GetVolumeId() int64  { return m.Data2.VolumeId }

func (m *TLInputFileLocation) SetLocalId(v int32) { m.Data2.LocalId = v }
func (m *TLInputFileLocation) GetLocalId() int32  { return m.Data2.LocalId }

func (m *TLInputFileLocation) SetSecret(v int64) { m.Data2.Secret = v }
func (m *TLInputFileLocation) GetSecret() int64  { return m.Data2.Secret }

func (m *TLInputFileLocation) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputFileLocation))

	x.Long(m.Data2.VolumeId)
	x.Int(m.Data2.LocalId)
	x.Long(m.Data2.Secret)

	return x.buf
}

func (m *TLInputFileLocation) Decode(dbuf *DecodeBuf) error {
	m.Data2.VolumeId = x.Long()
	m.Data2.LocalId = x.Int()
	m.Data2.Secret = x.Long()

	return dbuf.err
}

// inputEncryptedFileLocation#f5235d55 id:long access_hash:long = InputFileLocation;
func (m *TLInputEncryptedFileLocation) To_InputFileLocation() *InputFileLocation {
	return &InputFileLocation{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputEncryptedFileLocation) SetId(v int64) { m.Data2.Id = v }
func (m *TLInputEncryptedFileLocation) GetId() int64  { return m.Data2.Id }

func (m *TLInputEncryptedFileLocation) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLInputEncryptedFileLocation) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLInputEncryptedFileLocation) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputEncryptedFileLocation))

	x.Long(m.Data2.Id)
	x.Long(m.Data2.AccessHash)

	return x.buf
}

func (m *TLInputEncryptedFileLocation) Decode(dbuf *DecodeBuf) error {
	m.Data2.Id = x.Long()
	m.Data2.AccessHash = x.Long()

	return dbuf.err
}

// inputDocumentFileLocation#430f0724 id:long access_hash:long version:int = InputFileLocation;
func (m *TLInputDocumentFileLocation) To_InputFileLocation() *InputFileLocation {
	return &InputFileLocation{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputDocumentFileLocation) SetId(v int64) { m.Data2.Id = v }
func (m *TLInputDocumentFileLocation) GetId() int64  { return m.Data2.Id }

func (m *TLInputDocumentFileLocation) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLInputDocumentFileLocation) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLInputDocumentFileLocation) SetVersion(v int32) { m.Data2.Version = v }
func (m *TLInputDocumentFileLocation) GetVersion() int32  { return m.Data2.Version }

func (m *TLInputDocumentFileLocation) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputDocumentFileLocation))

	x.Long(m.Data2.Id)
	x.Long(m.Data2.AccessHash)
	x.Int(m.Data2.Version)

	return x.buf
}

func (m *TLInputDocumentFileLocation) Decode(dbuf *DecodeBuf) error {
	m.Data2.Id = x.Long()
	m.Data2.AccessHash = x.Long()
	m.Data2.Version = x.Int()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Storage_FileType <--
//  + TL_StorageFileUnknown
//  + TL_StorageFilePartial
//  + TL_StorageFileJpeg
//  + TL_StorageFileGif
//  + TL_StorageFilePng
//  + TL_StorageFilePdf
//  + TL_StorageFileMp3
//  + TL_StorageFileMov
//  + TL_StorageFileMp4
//  + TL_StorageFileWebp
//

func (m *Storage_FileType) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_storage_fileUnknown:
		t := m.To_StorageFileUnknown()
		return t.Encode()
	case TLConstructor_CRC32_storage_filePartial:
		t := m.To_StorageFilePartial()
		return t.Encode()
	case TLConstructor_CRC32_storage_fileJpeg:
		t := m.To_StorageFileJpeg()
		return t.Encode()
	case TLConstructor_CRC32_storage_fileGif:
		t := m.To_StorageFileGif()
		return t.Encode()
	case TLConstructor_CRC32_storage_filePng:
		t := m.To_StorageFilePng()
		return t.Encode()
	case TLConstructor_CRC32_storage_filePdf:
		t := m.To_StorageFilePdf()
		return t.Encode()
	case TLConstructor_CRC32_storage_fileMp3:
		t := m.To_StorageFileMp3()
		return t.Encode()
	case TLConstructor_CRC32_storage_fileMov:
		t := m.To_StorageFileMov()
		return t.Encode()
	case TLConstructor_CRC32_storage_fileMp4:
		t := m.To_StorageFileMp4()
		return t.Encode()
	case TLConstructor_CRC32_storage_fileWebp:
		t := m.To_StorageFileWebp()
		return t.Encode()

	default:
		return []byte{}
	}
}

// storage.fileUnknown#aa963b05 = storage.FileType;
func (m *Storage_FileType) To_StorageFileUnknown() *TLStorageFileUnknown {
	return &TLStorageFileUnknown{
		Data2: m.Data2,
	}
}

// storage.filePartial#40bc6f52 = storage.FileType;
func (m *Storage_FileType) To_StorageFilePartial() *TLStorageFilePartial {
	return &TLStorageFilePartial{
		Data2: m.Data2,
	}
}

// storage.fileJpeg#7efe0e = storage.FileType;
func (m *Storage_FileType) To_StorageFileJpeg() *TLStorageFileJpeg {
	return &TLStorageFileJpeg{
		Data2: m.Data2,
	}
}

// storage.fileGif#cae1aadf = storage.FileType;
func (m *Storage_FileType) To_StorageFileGif() *TLStorageFileGif {
	return &TLStorageFileGif{
		Data2: m.Data2,
	}
}

// storage.filePng#a4f63c0 = storage.FileType;
func (m *Storage_FileType) To_StorageFilePng() *TLStorageFilePng {
	return &TLStorageFilePng{
		Data2: m.Data2,
	}
}

// storage.filePdf#ae1e508d = storage.FileType;
func (m *Storage_FileType) To_StorageFilePdf() *TLStorageFilePdf {
	return &TLStorageFilePdf{
		Data2: m.Data2,
	}
}

// storage.fileMp3#528a0677 = storage.FileType;
func (m *Storage_FileType) To_StorageFileMp3() *TLStorageFileMp3 {
	return &TLStorageFileMp3{
		Data2: m.Data2,
	}
}

// storage.fileMov#4b09ebbc = storage.FileType;
func (m *Storage_FileType) To_StorageFileMov() *TLStorageFileMov {
	return &TLStorageFileMov{
		Data2: m.Data2,
	}
}

// storage.fileMp4#b3cea0e4 = storage.FileType;
func (m *Storage_FileType) To_StorageFileMp4() *TLStorageFileMp4 {
	return &TLStorageFileMp4{
		Data2: m.Data2,
	}
}

// storage.fileWebp#1081464c = storage.FileType;
func (m *Storage_FileType) To_StorageFileWebp() *TLStorageFileWebp {
	return &TLStorageFileWebp{
		Data2: m.Data2,
	}
}

// storage.fileUnknown#aa963b05 = storage.FileType;
func (m *TLStorageFileUnknown) To_Storage_FileType() *Storage_FileType {
	return &Storage_FileType{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLStorageFileUnknown) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_storage_fileUnknown))

	return x.buf
}

func (m *TLStorageFileUnknown) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// storage.filePartial#40bc6f52 = storage.FileType;
func (m *TLStorageFilePartial) To_Storage_FileType() *Storage_FileType {
	return &Storage_FileType{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLStorageFilePartial) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_storage_filePartial))

	return x.buf
}

func (m *TLStorageFilePartial) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// storage.fileJpeg#7efe0e = storage.FileType;
func (m *TLStorageFileJpeg) To_Storage_FileType() *Storage_FileType {
	return &Storage_FileType{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLStorageFileJpeg) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_storage_fileJpeg))

	return x.buf
}

func (m *TLStorageFileJpeg) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// storage.fileGif#cae1aadf = storage.FileType;
func (m *TLStorageFileGif) To_Storage_FileType() *Storage_FileType {
	return &Storage_FileType{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLStorageFileGif) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_storage_fileGif))

	return x.buf
}

func (m *TLStorageFileGif) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// storage.filePng#a4f63c0 = storage.FileType;
func (m *TLStorageFilePng) To_Storage_FileType() *Storage_FileType {
	return &Storage_FileType{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLStorageFilePng) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_storage_filePng))

	return x.buf
}

func (m *TLStorageFilePng) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// storage.filePdf#ae1e508d = storage.FileType;
func (m *TLStorageFilePdf) To_Storage_FileType() *Storage_FileType {
	return &Storage_FileType{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLStorageFilePdf) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_storage_filePdf))

	return x.buf
}

func (m *TLStorageFilePdf) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// storage.fileMp3#528a0677 = storage.FileType;
func (m *TLStorageFileMp3) To_Storage_FileType() *Storage_FileType {
	return &Storage_FileType{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLStorageFileMp3) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_storage_fileMp3))

	return x.buf
}

func (m *TLStorageFileMp3) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// storage.fileMov#4b09ebbc = storage.FileType;
func (m *TLStorageFileMov) To_Storage_FileType() *Storage_FileType {
	return &Storage_FileType{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLStorageFileMov) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_storage_fileMov))

	return x.buf
}

func (m *TLStorageFileMov) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// storage.fileMp4#b3cea0e4 = storage.FileType;
func (m *TLStorageFileMp4) To_Storage_FileType() *Storage_FileType {
	return &Storage_FileType{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLStorageFileMp4) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_storage_fileMp4))

	return x.buf
}

func (m *TLStorageFileMp4) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// storage.fileWebp#1081464c = storage.FileType;
func (m *TLStorageFileWebp) To_Storage_FileType() *Storage_FileType {
	return &Storage_FileType{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLStorageFileWebp) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_storage_fileWebp))

	return x.buf
}

func (m *TLStorageFileWebp) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Page <--
//  + TL_PagePart
//  + TL_PageFull
//

func (m *Page) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_pagePart:
		t := m.To_PagePart()
		return t.Encode()
	case TLConstructor_CRC32_pageFull:
		t := m.To_PageFull()
		return t.Encode()

	default:
		return []byte{}
	}
}

// pagePart#8e3f9ebe blocks:Vector<PageBlock> photos:Vector<Photo> documents:Vector<Document> = Page;
func (m *Page) To_PagePart() *TLPagePart {
	return &TLPagePart{
		Data2: m.Data2,
	}
}

// pageFull#556ec7aa blocks:Vector<PageBlock> photos:Vector<Photo> documents:Vector<Document> = Page;
func (m *Page) To_PageFull() *TLPageFull {
	return &TLPageFull{
		Data2: m.Data2,
	}
}

// pagePart#8e3f9ebe blocks:Vector<PageBlock> photos:Vector<Photo> documents:Vector<Document> = Page;
func (m *TLPagePart) To_Page() *Page {
	return &Page{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPagePart) SetBlocks(v []*PageBlock) { m.Data2.Blocks = v }
func (m *TLPagePart) GetBlocks() []*PageBlock  { return m.Data2.Blocks }

func (m *TLPagePart) SetPhotos(v []*Photo) { m.Data2.Photos = v }
func (m *TLPagePart) GetPhotos() []*Photo  { return m.Data2.Photos }

func (m *TLPagePart) SetDocuments(v []*Document) { m.Data2.Documents = v }
func (m *TLPagePart) GetDocuments() []*Document  { return m.Data2.Documents }

func (m *TLPagePart) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_pagePart))

	return x.buf
}

func (m *TLPagePart) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// pageFull#556ec7aa blocks:Vector<PageBlock> photos:Vector<Photo> documents:Vector<Document> = Page;
func (m *TLPageFull) To_Page() *Page {
	return &Page{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPageFull) SetBlocks(v []*PageBlock) { m.Data2.Blocks = v }
func (m *TLPageFull) GetBlocks() []*PageBlock  { return m.Data2.Blocks }

func (m *TLPageFull) SetPhotos(v []*Photo) { m.Data2.Photos = v }
func (m *TLPageFull) GetPhotos() []*Photo  { return m.Data2.Photos }

func (m *TLPageFull) SetDocuments(v []*Document) { m.Data2.Documents = v }
func (m *TLPageFull) GetDocuments() []*Document  { return m.Data2.Documents }

func (m *TLPageFull) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_pageFull))

	return x.buf
}

func (m *TLPageFull) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// KeyboardButton <--
//  + TL_KeyboardButton
//  + TL_KeyboardButtonUrl
//  + TL_KeyboardButtonCallback
//  + TL_KeyboardButtonRequestPhone
//  + TL_KeyboardButtonRequestGeoLocation
//  + TL_KeyboardButtonSwitchInline
//  + TL_KeyboardButtonGame
//  + TL_KeyboardButtonBuy
//

func (m *KeyboardButton) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_keyboardButton:
		t := m.To_KeyboardButton()
		return t.Encode()
	case TLConstructor_CRC32_keyboardButtonUrl:
		t := m.To_KeyboardButtonUrl()
		return t.Encode()
	case TLConstructor_CRC32_keyboardButtonCallback:
		t := m.To_KeyboardButtonCallback()
		return t.Encode()
	case TLConstructor_CRC32_keyboardButtonRequestPhone:
		t := m.To_KeyboardButtonRequestPhone()
		return t.Encode()
	case TLConstructor_CRC32_keyboardButtonRequestGeoLocation:
		t := m.To_KeyboardButtonRequestGeoLocation()
		return t.Encode()
	case TLConstructor_CRC32_keyboardButtonSwitchInline:
		t := m.To_KeyboardButtonSwitchInline()
		return t.Encode()
	case TLConstructor_CRC32_keyboardButtonGame:
		t := m.To_KeyboardButtonGame()
		return t.Encode()
	case TLConstructor_CRC32_keyboardButtonBuy:
		t := m.To_KeyboardButtonBuy()
		return t.Encode()

	default:
		return []byte{}
	}
}

// keyboardButton#a2fa4880 text:string = KeyboardButton;
func (m *KeyboardButton) To_KeyboardButton() *TLKeyboardButton {
	return &TLKeyboardButton{
		Data2: m.Data2,
	}
}

// keyboardButtonUrl#258aff05 text:string url:string = KeyboardButton;
func (m *KeyboardButton) To_KeyboardButtonUrl() *TLKeyboardButtonUrl {
	return &TLKeyboardButtonUrl{
		Data2: m.Data2,
	}
}

// keyboardButtonCallback#683a5e46 text:string data:bytes = KeyboardButton;
func (m *KeyboardButton) To_KeyboardButtonCallback() *TLKeyboardButtonCallback {
	return &TLKeyboardButtonCallback{
		Data2: m.Data2,
	}
}

// keyboardButtonRequestPhone#b16a6c29 text:string = KeyboardButton;
func (m *KeyboardButton) To_KeyboardButtonRequestPhone() *TLKeyboardButtonRequestPhone {
	return &TLKeyboardButtonRequestPhone{
		Data2: m.Data2,
	}
}

// keyboardButtonRequestGeoLocation#fc796b3f text:string = KeyboardButton;
func (m *KeyboardButton) To_KeyboardButtonRequestGeoLocation() *TLKeyboardButtonRequestGeoLocation {
	return &TLKeyboardButtonRequestGeoLocation{
		Data2: m.Data2,
	}
}

// keyboardButtonSwitchInline#568a748 flags:# same_peer:flags.0?true text:string query:string = KeyboardButton;
func (m *KeyboardButton) To_KeyboardButtonSwitchInline() *TLKeyboardButtonSwitchInline {
	return &TLKeyboardButtonSwitchInline{
		Data2: m.Data2,
	}
}

// keyboardButtonGame#50f41ccf text:string = KeyboardButton;
func (m *KeyboardButton) To_KeyboardButtonGame() *TLKeyboardButtonGame {
	return &TLKeyboardButtonGame{
		Data2: m.Data2,
	}
}

// keyboardButtonBuy#afd93fbb text:string = KeyboardButton;
func (m *KeyboardButton) To_KeyboardButtonBuy() *TLKeyboardButtonBuy {
	return &TLKeyboardButtonBuy{
		Data2: m.Data2,
	}
}

// keyboardButton#a2fa4880 text:string = KeyboardButton;
func (m *TLKeyboardButton) To_KeyboardButton() *KeyboardButton {
	return &KeyboardButton{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLKeyboardButton) SetText(v string) { m.Data2.Text = v }
func (m *TLKeyboardButton) GetText() string  { return m.Data2.Text }

func (m *TLKeyboardButton) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_keyboardButton))

	x.StringBytes(m.Data2.Text)

	return x.buf
}

func (m *TLKeyboardButton) Decode(dbuf *DecodeBuf) error {
	m.Data2.Text = x.StringBytes()

	return dbuf.err
}

// keyboardButtonUrl#258aff05 text:string url:string = KeyboardButton;
func (m *TLKeyboardButtonUrl) To_KeyboardButton() *KeyboardButton {
	return &KeyboardButton{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLKeyboardButtonUrl) SetText(v string) { m.Data2.Text = v }
func (m *TLKeyboardButtonUrl) GetText() string  { return m.Data2.Text }

func (m *TLKeyboardButtonUrl) SetUrl(v string) { m.Data2.Url = v }
func (m *TLKeyboardButtonUrl) GetUrl() string  { return m.Data2.Url }

func (m *TLKeyboardButtonUrl) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_keyboardButtonUrl))

	x.StringBytes(m.Data2.Text)
	x.StringBytes(m.Data2.Url)

	return x.buf
}

func (m *TLKeyboardButtonUrl) Decode(dbuf *DecodeBuf) error {
	m.Data2.Text = x.StringBytes()
	m.Data2.Url = x.StringBytes()

	return dbuf.err
}

// keyboardButtonCallback#683a5e46 text:string data:bytes = KeyboardButton;
func (m *TLKeyboardButtonCallback) To_KeyboardButton() *KeyboardButton {
	return &KeyboardButton{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLKeyboardButtonCallback) SetText(v string) { m.Data2.Text = v }
func (m *TLKeyboardButtonCallback) GetText() string  { return m.Data2.Text }

func (m *TLKeyboardButtonCallback) SetData(v []byte) { m.Data2.Data = v }
func (m *TLKeyboardButtonCallback) GetData() []byte  { return m.Data2.Data }

func (m *TLKeyboardButtonCallback) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_keyboardButtonCallback))

	x.StringBytes(m.Data2.Text)
	x.StringBytes(m.Data2.Data)

	return x.buf
}

func (m *TLKeyboardButtonCallback) Decode(dbuf *DecodeBuf) error {
	m.Data2.Text = x.StringBytes()
	m.Data2.Data = x.StringBytes()

	return dbuf.err
}

// keyboardButtonRequestPhone#b16a6c29 text:string = KeyboardButton;
func (m *TLKeyboardButtonRequestPhone) To_KeyboardButton() *KeyboardButton {
	return &KeyboardButton{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLKeyboardButtonRequestPhone) SetText(v string) { m.Data2.Text = v }
func (m *TLKeyboardButtonRequestPhone) GetText() string  { return m.Data2.Text }

func (m *TLKeyboardButtonRequestPhone) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_keyboardButtonRequestPhone))

	x.StringBytes(m.Data2.Text)

	return x.buf
}

func (m *TLKeyboardButtonRequestPhone) Decode(dbuf *DecodeBuf) error {
	m.Data2.Text = x.StringBytes()

	return dbuf.err
}

// keyboardButtonRequestGeoLocation#fc796b3f text:string = KeyboardButton;
func (m *TLKeyboardButtonRequestGeoLocation) To_KeyboardButton() *KeyboardButton {
	return &KeyboardButton{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLKeyboardButtonRequestGeoLocation) SetText(v string) { m.Data2.Text = v }
func (m *TLKeyboardButtonRequestGeoLocation) GetText() string  { return m.Data2.Text }

func (m *TLKeyboardButtonRequestGeoLocation) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_keyboardButtonRequestGeoLocation))

	x.StringBytes(m.Data2.Text)

	return x.buf
}

func (m *TLKeyboardButtonRequestGeoLocation) Decode(dbuf *DecodeBuf) error {
	m.Data2.Text = x.StringBytes()

	return dbuf.err
}

// keyboardButtonSwitchInline#568a748 flags:# same_peer:flags.0?true text:string query:string = KeyboardButton;
func (m *TLKeyboardButtonSwitchInline) To_KeyboardButton() *KeyboardButton {
	return &KeyboardButton{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLKeyboardButtonSwitchInline) SetSamePeer(v bool) { m.Data2.SamePeer = v }
func (m *TLKeyboardButtonSwitchInline) GetSamePeer() bool  { return m.Data2.SamePeer }

func (m *TLKeyboardButtonSwitchInline) SetText(v string) { m.Data2.Text = v }
func (m *TLKeyboardButtonSwitchInline) GetText() string  { return m.Data2.Text }

func (m *TLKeyboardButtonSwitchInline) SetQuery(v string) { m.Data2.Query = v }
func (m *TLKeyboardButtonSwitchInline) GetQuery() string  { return m.Data2.Query }

func (m *TLKeyboardButtonSwitchInline) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_keyboardButtonSwitchInline))

	// flags
	var flags uint32 = 0
	if m.GetSamePeer() == true {
		flags |= 1 << 1
	}
	x.UInt(flags)

	x.StringBytes(m.Data2.Text)
	x.StringBytes(m.Data2.Query)

	return x.buf
}

func (m *TLKeyboardButtonSwitchInline) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m.Data2.SamePeer = true
	}
	m.Data2.Text = x.StringBytes()
	m.Data2.Query = x.StringBytes()

	return dbuf.err
}

// keyboardButtonGame#50f41ccf text:string = KeyboardButton;
func (m *TLKeyboardButtonGame) To_KeyboardButton() *KeyboardButton {
	return &KeyboardButton{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLKeyboardButtonGame) SetText(v string) { m.Data2.Text = v }
func (m *TLKeyboardButtonGame) GetText() string  { return m.Data2.Text }

func (m *TLKeyboardButtonGame) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_keyboardButtonGame))

	x.StringBytes(m.Data2.Text)

	return x.buf
}

func (m *TLKeyboardButtonGame) Decode(dbuf *DecodeBuf) error {
	m.Data2.Text = x.StringBytes()

	return dbuf.err
}

// keyboardButtonBuy#afd93fbb text:string = KeyboardButton;
func (m *TLKeyboardButtonBuy) To_KeyboardButton() *KeyboardButton {
	return &KeyboardButton{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLKeyboardButtonBuy) SetText(v string) { m.Data2.Text = v }
func (m *TLKeyboardButtonBuy) GetText() string  { return m.Data2.Text }

func (m *TLKeyboardButtonBuy) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_keyboardButtonBuy))

	x.StringBytes(m.Data2.Text)

	return x.buf
}

func (m *TLKeyboardButtonBuy) Decode(dbuf *DecodeBuf) error {
	m.Data2.Text = x.StringBytes()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// InputChannel <--
//  + TL_InputChannelEmpty
//  + TL_InputChannel
//

func (m *InputChannel) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_inputChannelEmpty:
		t := m.To_InputChannelEmpty()
		return t.Encode()
	case TLConstructor_CRC32_inputChannel:
		t := m.To_InputChannel()
		return t.Encode()

	default:
		return []byte{}
	}
}

// inputChannelEmpty#ee8c1e86 = InputChannel;
func (m *InputChannel) To_InputChannelEmpty() *TLInputChannelEmpty {
	return &TLInputChannelEmpty{
		Data2: m.Data2,
	}
}

// inputChannel#afeb712e channel_id:int access_hash:long = InputChannel;
func (m *InputChannel) To_InputChannel() *TLInputChannel {
	return &TLInputChannel{
		Data2: m.Data2,
	}
}

// inputChannelEmpty#ee8c1e86 = InputChannel;
func (m *TLInputChannelEmpty) To_InputChannel() *InputChannel {
	return &InputChannel{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputChannelEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputChannelEmpty))

	return x.buf
}

func (m *TLInputChannelEmpty) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// inputChannel#afeb712e channel_id:int access_hash:long = InputChannel;
func (m *TLInputChannel) To_InputChannel() *InputChannel {
	return &InputChannel{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputChannel) SetChannelId(v int32) { m.Data2.ChannelId = v }
func (m *TLInputChannel) GetChannelId() int32  { return m.Data2.ChannelId }

func (m *TLInputChannel) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLInputChannel) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLInputChannel) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputChannel))

	x.Int(m.Data2.ChannelId)
	x.Long(m.Data2.AccessHash)

	return x.buf
}

func (m *TLInputChannel) Decode(dbuf *DecodeBuf) error {
	m.Data2.ChannelId = x.Int()
	m.Data2.AccessHash = x.Long()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Phone_PhoneCall <--
//  + TL_PhonePhoneCall
//

func (m *Phone_PhoneCall) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_phone_phoneCall:
		t := m.To_PhonePhoneCall()
		return t.Encode()

	default:
		return []byte{}
	}
}

// phone.phoneCall#ec82e140 phone_call:PhoneCall users:Vector<User> = phone.PhoneCall;
func (m *Phone_PhoneCall) To_PhonePhoneCall() *TLPhonePhoneCall {
	return &TLPhonePhoneCall{
		Data2: m.Data2,
	}
}

// phone.phoneCall#ec82e140 phone_call:PhoneCall users:Vector<User> = phone.PhoneCall;
func (m *TLPhonePhoneCall) To_Phone_PhoneCall() *Phone_PhoneCall {
	return &Phone_PhoneCall{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPhonePhoneCall) SetPhoneCall(v *PhoneCall) { m.Data2.PhoneCall = v }
func (m *TLPhonePhoneCall) GetPhoneCall() *PhoneCall  { return m.Data2.PhoneCall }

func (m *TLPhonePhoneCall) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLPhonePhoneCall) GetUsers() []*User  { return m.Data2.Users }

func (m *TLPhonePhoneCall) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_phone_phoneCall))

	x.Bytes(m.Data2PhoneCall.Encode())

	return x.buf
}

func (m *TLPhonePhoneCall) Decode(dbuf *DecodeBuf) error {
	m1 := &PhoneCall{}
	m1.Decode(dbuf)
	m.SetPhoneCall(m1)

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// InputMedia <--
//  + TL_InputMediaEmpty
//  + TL_InputMediaUploadedPhoto
//  + TL_InputMediaPhoto
//  + TL_InputMediaGeoPoint
//  + TL_InputMediaContact
//  + TL_InputMediaUploadedDocument
//  + TL_InputMediaDocument
//  + TL_InputMediaVenue
//  + TL_InputMediaGifExternal
//  + TL_InputMediaPhotoExternal
//  + TL_InputMediaDocumentExternal
//  + TL_InputMediaGame
//  + TL_InputMediaInvoice
//

func (m *InputMedia) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_inputMediaEmpty:
		t := m.To_InputMediaEmpty()
		return t.Encode()
	case TLConstructor_CRC32_inputMediaUploadedPhoto:
		t := m.To_InputMediaUploadedPhoto()
		return t.Encode()
	case TLConstructor_CRC32_inputMediaPhoto:
		t := m.To_InputMediaPhoto()
		return t.Encode()
	case TLConstructor_CRC32_inputMediaGeoPoint:
		t := m.To_InputMediaGeoPoint()
		return t.Encode()
	case TLConstructor_CRC32_inputMediaContact:
		t := m.To_InputMediaContact()
		return t.Encode()
	case TLConstructor_CRC32_inputMediaUploadedDocument:
		t := m.To_InputMediaUploadedDocument()
		return t.Encode()
	case TLConstructor_CRC32_inputMediaDocument:
		t := m.To_InputMediaDocument()
		return t.Encode()
	case TLConstructor_CRC32_inputMediaVenue:
		t := m.To_InputMediaVenue()
		return t.Encode()
	case TLConstructor_CRC32_inputMediaGifExternal:
		t := m.To_InputMediaGifExternal()
		return t.Encode()
	case TLConstructor_CRC32_inputMediaPhotoExternal:
		t := m.To_InputMediaPhotoExternal()
		return t.Encode()
	case TLConstructor_CRC32_inputMediaDocumentExternal:
		t := m.To_InputMediaDocumentExternal()
		return t.Encode()
	case TLConstructor_CRC32_inputMediaGame:
		t := m.To_InputMediaGame()
		return t.Encode()
	case TLConstructor_CRC32_inputMediaInvoice:
		t := m.To_InputMediaInvoice()
		return t.Encode()

	default:
		return []byte{}
	}
}

// inputMediaEmpty#9664f57f = InputMedia;
func (m *InputMedia) To_InputMediaEmpty() *TLInputMediaEmpty {
	return &TLInputMediaEmpty{
		Data2: m.Data2,
	}
}

// inputMediaUploadedPhoto#2f37e231 flags:# file:InputFile caption:string stickers:flags.0?Vector<InputDocument> ttl_seconds:flags.1?int = InputMedia;
func (m *InputMedia) To_InputMediaUploadedPhoto() *TLInputMediaUploadedPhoto {
	return &TLInputMediaUploadedPhoto{
		Data2: m.Data2,
	}
}

// inputMediaPhoto#81fa373a flags:# id:InputPhoto caption:string ttl_seconds:flags.0?int = InputMedia;
func (m *InputMedia) To_InputMediaPhoto() *TLInputMediaPhoto {
	return &TLInputMediaPhoto{
		Data2: m.Data2,
	}
}

// inputMediaGeoPoint#f9c44144 geo_point:InputGeoPoint = InputMedia;
func (m *InputMedia) To_InputMediaGeoPoint() *TLInputMediaGeoPoint {
	return &TLInputMediaGeoPoint{
		Data2: m.Data2,
	}
}

// inputMediaContact#a6e45987 phone_number:string first_name:string last_name:string = InputMedia;
func (m *InputMedia) To_InputMediaContact() *TLInputMediaContact {
	return &TLInputMediaContact{
		Data2: m.Data2,
	}
}

// inputMediaUploadedDocument#e39621fd flags:# file:InputFile thumb:flags.2?InputFile mime_type:string attributes:Vector<DocumentAttribute> caption:string stickers:flags.0?Vector<InputDocument> ttl_seconds:flags.1?int = InputMedia;
func (m *InputMedia) To_InputMediaUploadedDocument() *TLInputMediaUploadedDocument {
	return &TLInputMediaUploadedDocument{
		Data2: m.Data2,
	}
}

// inputMediaDocument#5acb668e flags:# id:InputDocument caption:string ttl_seconds:flags.0?int = InputMedia;
func (m *InputMedia) To_InputMediaDocument() *TLInputMediaDocument {
	return &TLInputMediaDocument{
		Data2: m.Data2,
	}
}

// inputMediaVenue#2827a81a geo_point:InputGeoPoint title:string address:string provider:string venue_id:string = InputMedia;
func (m *InputMedia) To_InputMediaVenue() *TLInputMediaVenue {
	return &TLInputMediaVenue{
		Data2: m.Data2,
	}
}

// inputMediaGifExternal#4843b0fd url:string q:string = InputMedia;
func (m *InputMedia) To_InputMediaGifExternal() *TLInputMediaGifExternal {
	return &TLInputMediaGifExternal{
		Data2: m.Data2,
	}
}

// inputMediaPhotoExternal#922aec1 flags:# url:string caption:string ttl_seconds:flags.0?int = InputMedia;
func (m *InputMedia) To_InputMediaPhotoExternal() *TLInputMediaPhotoExternal {
	return &TLInputMediaPhotoExternal{
		Data2: m.Data2,
	}
}

// inputMediaDocumentExternal#b6f74335 flags:# url:string caption:string ttl_seconds:flags.0?int = InputMedia;
func (m *InputMedia) To_InputMediaDocumentExternal() *TLInputMediaDocumentExternal {
	return &TLInputMediaDocumentExternal{
		Data2: m.Data2,
	}
}

// inputMediaGame#d33f43f3 id:InputGame = InputMedia;
func (m *InputMedia) To_InputMediaGame() *TLInputMediaGame {
	return &TLInputMediaGame{
		Data2: m.Data2,
	}
}

// inputMediaInvoice#92153685 flags:# title:string description:string photo:flags.0?InputWebDocument invoice:Invoice payload:bytes provider:string start_param:string = InputMedia;
func (m *InputMedia) To_InputMediaInvoice() *TLInputMediaInvoice {
	return &TLInputMediaInvoice{
		Data2: m.Data2,
	}
}

// inputMediaEmpty#9664f57f = InputMedia;
func (m *TLInputMediaEmpty) To_InputMedia() *InputMedia {
	return &InputMedia{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputMediaEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputMediaEmpty))

	return x.buf
}

func (m *TLInputMediaEmpty) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// inputMediaUploadedPhoto#2f37e231 flags:# file:InputFile caption:string stickers:flags.0?Vector<InputDocument> ttl_seconds:flags.1?int = InputMedia;
func (m *TLInputMediaUploadedPhoto) To_InputMedia() *InputMedia {
	return &InputMedia{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputMediaUploadedPhoto) SetFile(v *InputFile) { m.Data2.File = v }
func (m *TLInputMediaUploadedPhoto) GetFile() *InputFile  { return m.Data2.File }

func (m *TLInputMediaUploadedPhoto) SetCaption(v string) { m.Data2.Caption = v }
func (m *TLInputMediaUploadedPhoto) GetCaption() string  { return m.Data2.Caption }

func (m *TLInputMediaUploadedPhoto) SetStickers(v []*InputDocument) { m.Data2.Stickers = v }
func (m *TLInputMediaUploadedPhoto) GetStickers() []*InputDocument  { return m.Data2.Stickers }

func (m *TLInputMediaUploadedPhoto) SetTtlSeconds(v int32) { m.Data2.TtlSeconds = v }
func (m *TLInputMediaUploadedPhoto) GetTtlSeconds() int32  { return m.Data2.TtlSeconds }

func (m *TLInputMediaUploadedPhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputMediaUploadedPhoto))

	// flags
	var flags uint32 = 0
	if m.GetStickers() != nil {
		flags |= 1 << 3
	}
	if m.GetTtlSeconds() != 0 {
		flags |= 1 << 4
	}
	x.UInt(flags)

	x.Bytes(m.Data2File.Encode())
	x.StringBytes(m.Data2.Caption)
	if m.Data2.Stickers != 0 {

	}
	if m.GetTtlSeconds() != 0 {
		x.Int(m.Data2.TtlSeconds)
	}

	return x.buf
}

func (m *TLInputMediaUploadedPhoto) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	m2 := &File{}
	m2.Decode(dbuf)
	m.SetFile(m2)
	m.Data2.Caption = x.StringBytes()
	if (flags & (1 << 4)) != 0 {

	}
	if (flags & (1 << 5)) != 0 {
		m.Data2.TtlSeconds = x.Int()
	}

	return dbuf.err
}

// inputMediaPhoto#81fa373a flags:# id:InputPhoto caption:string ttl_seconds:flags.0?int = InputMedia;
func (m *TLInputMediaPhoto) To_InputMedia() *InputMedia {
	return &InputMedia{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputMediaPhoto) SetId(v *InputPhoto) { m.Data2.Id_5 = v }
func (m *TLInputMediaPhoto) GetId() *InputPhoto  { return m.Data2.Id_5 }

func (m *TLInputMediaPhoto) SetCaption(v string) { m.Data2.Caption = v }
func (m *TLInputMediaPhoto) GetCaption() string  { return m.Data2.Caption }

func (m *TLInputMediaPhoto) SetTtlSeconds(v int32) { m.Data2.TtlSeconds = v }
func (m *TLInputMediaPhoto) GetTtlSeconds() int32  { return m.Data2.TtlSeconds }

func (m *TLInputMediaPhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputMediaPhoto))

	// flags
	var flags uint32 = 0
	if m.GetTtlSeconds() != 0 {
		flags |= 1 << 3
	}
	x.UInt(flags)

	x.Bytes(m.Data2Id.Encode())
	x.StringBytes(m.Data2.Caption)
	if m.GetTtlSeconds() != 0 {
		x.Int(m.Data2.TtlSeconds)
	}

	return x.buf
}

func (m *TLInputMediaPhoto) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	m2 := &Id{}
	m2.Decode(dbuf)
	m.SetId(m2)
	m.Data2.Caption = x.StringBytes()
	if (flags & (1 << 4)) != 0 {
		m.Data2.TtlSeconds = x.Int()
	}

	return dbuf.err
}

// inputMediaGeoPoint#f9c44144 geo_point:InputGeoPoint = InputMedia;
func (m *TLInputMediaGeoPoint) To_InputMedia() *InputMedia {
	return &InputMedia{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputMediaGeoPoint) SetGeoPoint(v *InputGeoPoint) { m.Data2.GeoPoint = v }
func (m *TLInputMediaGeoPoint) GetGeoPoint() *InputGeoPoint  { return m.Data2.GeoPoint }

func (m *TLInputMediaGeoPoint) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputMediaGeoPoint))

	x.Bytes(m.Data2GeoPoint.Encode())

	return x.buf
}

func (m *TLInputMediaGeoPoint) Decode(dbuf *DecodeBuf) error {
	m1 := &GeoPoint{}
	m1.Decode(dbuf)
	m.SetGeoPoint(m1)

	return dbuf.err
}

// inputMediaContact#a6e45987 phone_number:string first_name:string last_name:string = InputMedia;
func (m *TLInputMediaContact) To_InputMedia() *InputMedia {
	return &InputMedia{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputMediaContact) SetPhoneNumber(v string) { m.Data2.PhoneNumber = v }
func (m *TLInputMediaContact) GetPhoneNumber() string  { return m.Data2.PhoneNumber }

func (m *TLInputMediaContact) SetFirstName(v string) { m.Data2.FirstName = v }
func (m *TLInputMediaContact) GetFirstName() string  { return m.Data2.FirstName }

func (m *TLInputMediaContact) SetLastName(v string) { m.Data2.LastName = v }
func (m *TLInputMediaContact) GetLastName() string  { return m.Data2.LastName }

func (m *TLInputMediaContact) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputMediaContact))

	x.StringBytes(m.Data2.PhoneNumber)
	x.StringBytes(m.Data2.FirstName)
	x.StringBytes(m.Data2.LastName)

	return x.buf
}

func (m *TLInputMediaContact) Decode(dbuf *DecodeBuf) error {
	m.Data2.PhoneNumber = x.StringBytes()
	m.Data2.FirstName = x.StringBytes()
	m.Data2.LastName = x.StringBytes()

	return dbuf.err
}

// inputMediaUploadedDocument#e39621fd flags:# file:InputFile thumb:flags.2?InputFile mime_type:string attributes:Vector<DocumentAttribute> caption:string stickers:flags.0?Vector<InputDocument> ttl_seconds:flags.1?int = InputMedia;
func (m *TLInputMediaUploadedDocument) To_InputMedia() *InputMedia {
	return &InputMedia{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputMediaUploadedDocument) SetFile(v *InputFile) { m.Data2.File = v }
func (m *TLInputMediaUploadedDocument) GetFile() *InputFile  { return m.Data2.File }

func (m *TLInputMediaUploadedDocument) SetThumb(v *InputFile) { m.Data2.Thumb = v }
func (m *TLInputMediaUploadedDocument) GetThumb() *InputFile  { return m.Data2.Thumb }

func (m *TLInputMediaUploadedDocument) SetMimeType(v string) { m.Data2.MimeType = v }
func (m *TLInputMediaUploadedDocument) GetMimeType() string  { return m.Data2.MimeType }

func (m *TLInputMediaUploadedDocument) SetAttributes(v []*DocumentAttribute) { m.Data2.Attributes = v }
func (m *TLInputMediaUploadedDocument) GetAttributes() []*DocumentAttribute  { return m.Data2.Attributes }

func (m *TLInputMediaUploadedDocument) SetCaption(v string) { m.Data2.Caption = v }
func (m *TLInputMediaUploadedDocument) GetCaption() string  { return m.Data2.Caption }

func (m *TLInputMediaUploadedDocument) SetStickers(v []*InputDocument) { m.Data2.Stickers = v }
func (m *TLInputMediaUploadedDocument) GetStickers() []*InputDocument  { return m.Data2.Stickers }

func (m *TLInputMediaUploadedDocument) SetTtlSeconds(v int32) { m.Data2.TtlSeconds = v }
func (m *TLInputMediaUploadedDocument) GetTtlSeconds() int32  { return m.Data2.TtlSeconds }

func (m *TLInputMediaUploadedDocument) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputMediaUploadedDocument))

	// flags
	var flags uint32 = 0
	if m.GetThumb() != nil {
		flags |= 1 << 2
	}
	if m.GetStickers() != nil {
		flags |= 1 << 6
	}
	if m.GetTtlSeconds() != 0 {
		flags |= 1 << 7
	}
	x.UInt(flags)

	x.Bytes(m.Data2File.Encode())
	if m.GetThumb() != 0 {
		x.Bytes(m.Data2Thumb.Encode())
	}
	x.StringBytes(m.Data2.MimeType)

	x.StringBytes(m.Data2.Caption)
	if m.Data2.Stickers != 0 {

	}
	if m.GetTtlSeconds() != 0 {
		x.Int(m.Data2.TtlSeconds)
	}

	return x.buf
}

func (m *TLInputMediaUploadedDocument) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	m2 := &File{}
	m2.Decode(dbuf)
	m.SetFile(m2)
	if (flags & (1 << 3)) != 0 {
		m3 := &Thumb{}
		m3.Decode(dbuf)
		m.SetThumb(m3)
	}
	m.Data2.MimeType = x.StringBytes()

	m.Data2.Caption = x.StringBytes()
	if (flags & (1 << 7)) != 0 {

	}
	if (flags & (1 << 8)) != 0 {
		m.Data2.TtlSeconds = x.Int()
	}

	return dbuf.err
}

// inputMediaDocument#5acb668e flags:# id:InputDocument caption:string ttl_seconds:flags.0?int = InputMedia;
func (m *TLInputMediaDocument) To_InputMedia() *InputMedia {
	return &InputMedia{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputMediaDocument) SetId(v *InputDocument) { m.Data2.Id_13 = v }
func (m *TLInputMediaDocument) GetId() *InputDocument  { return m.Data2.Id_13 }

func (m *TLInputMediaDocument) SetCaption(v string) { m.Data2.Caption = v }
func (m *TLInputMediaDocument) GetCaption() string  { return m.Data2.Caption }

func (m *TLInputMediaDocument) SetTtlSeconds(v int32) { m.Data2.TtlSeconds = v }
func (m *TLInputMediaDocument) GetTtlSeconds() int32  { return m.Data2.TtlSeconds }

func (m *TLInputMediaDocument) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputMediaDocument))

	// flags
	var flags uint32 = 0
	if m.GetTtlSeconds() != 0 {
		flags |= 1 << 3
	}
	x.UInt(flags)

	x.Bytes(m.Data2Id.Encode())
	x.StringBytes(m.Data2.Caption)
	if m.GetTtlSeconds() != 0 {
		x.Int(m.Data2.TtlSeconds)
	}

	return x.buf
}

func (m *TLInputMediaDocument) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	m2 := &Id{}
	m2.Decode(dbuf)
	m.SetId(m2)
	m.Data2.Caption = x.StringBytes()
	if (flags & (1 << 4)) != 0 {
		m.Data2.TtlSeconds = x.Int()
	}

	return dbuf.err
}

// inputMediaVenue#2827a81a geo_point:InputGeoPoint title:string address:string provider:string venue_id:string = InputMedia;
func (m *TLInputMediaVenue) To_InputMedia() *InputMedia {
	return &InputMedia{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputMediaVenue) SetGeoPoint(v *InputGeoPoint) { m.Data2.GeoPoint = v }
func (m *TLInputMediaVenue) GetGeoPoint() *InputGeoPoint  { return m.Data2.GeoPoint }

func (m *TLInputMediaVenue) SetTitle(v string) { m.Data2.Title = v }
func (m *TLInputMediaVenue) GetTitle() string  { return m.Data2.Title }

func (m *TLInputMediaVenue) SetAddress(v string) { m.Data2.Address = v }
func (m *TLInputMediaVenue) GetAddress() string  { return m.Data2.Address }

func (m *TLInputMediaVenue) SetProvider(v string) { m.Data2.Provider = v }
func (m *TLInputMediaVenue) GetProvider() string  { return m.Data2.Provider }

func (m *TLInputMediaVenue) SetVenueId(v string) { m.Data2.VenueId = v }
func (m *TLInputMediaVenue) GetVenueId() string  { return m.Data2.VenueId }

func (m *TLInputMediaVenue) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputMediaVenue))

	x.Bytes(m.Data2GeoPoint.Encode())
	x.StringBytes(m.Data2.Title)
	x.StringBytes(m.Data2.Address)
	x.StringBytes(m.Data2.Provider)
	x.StringBytes(m.Data2.VenueId)

	return x.buf
}

func (m *TLInputMediaVenue) Decode(dbuf *DecodeBuf) error {
	m1 := &GeoPoint{}
	m1.Decode(dbuf)
	m.SetGeoPoint(m1)
	m.Data2.Title = x.StringBytes()
	m.Data2.Address = x.StringBytes()
	m.Data2.Provider = x.StringBytes()
	m.Data2.VenueId = x.StringBytes()

	return dbuf.err
}

// inputMediaGifExternal#4843b0fd url:string q:string = InputMedia;
func (m *TLInputMediaGifExternal) To_InputMedia() *InputMedia {
	return &InputMedia{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputMediaGifExternal) SetUrl(v string) { m.Data2.Url = v }
func (m *TLInputMediaGifExternal) GetUrl() string  { return m.Data2.Url }

func (m *TLInputMediaGifExternal) SetQ(v string) { m.Data2.Q = v }
func (m *TLInputMediaGifExternal) GetQ() string  { return m.Data2.Q }

func (m *TLInputMediaGifExternal) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputMediaGifExternal))

	x.StringBytes(m.Data2.Url)
	x.StringBytes(m.Data2.Q)

	return x.buf
}

func (m *TLInputMediaGifExternal) Decode(dbuf *DecodeBuf) error {
	m.Data2.Url = x.StringBytes()
	m.Data2.Q = x.StringBytes()

	return dbuf.err
}

// inputMediaPhotoExternal#922aec1 flags:# url:string caption:string ttl_seconds:flags.0?int = InputMedia;
func (m *TLInputMediaPhotoExternal) To_InputMedia() *InputMedia {
	return &InputMedia{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputMediaPhotoExternal) SetUrl(v string) { m.Data2.Url = v }
func (m *TLInputMediaPhotoExternal) GetUrl() string  { return m.Data2.Url }

func (m *TLInputMediaPhotoExternal) SetCaption(v string) { m.Data2.Caption = v }
func (m *TLInputMediaPhotoExternal) GetCaption() string  { return m.Data2.Caption }

func (m *TLInputMediaPhotoExternal) SetTtlSeconds(v int32) { m.Data2.TtlSeconds = v }
func (m *TLInputMediaPhotoExternal) GetTtlSeconds() int32  { return m.Data2.TtlSeconds }

func (m *TLInputMediaPhotoExternal) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputMediaPhotoExternal))

	// flags
	var flags uint32 = 0
	if m.GetTtlSeconds() != 0 {
		flags |= 1 << 3
	}
	x.UInt(flags)

	x.StringBytes(m.Data2.Url)
	x.StringBytes(m.Data2.Caption)
	if m.GetTtlSeconds() != 0 {
		x.Int(m.Data2.TtlSeconds)
	}

	return x.buf
}

func (m *TLInputMediaPhotoExternal) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	m.Data2.Url = x.StringBytes()
	m.Data2.Caption = x.StringBytes()
	if (flags & (1 << 4)) != 0 {
		m.Data2.TtlSeconds = x.Int()
	}

	return dbuf.err
}

// inputMediaDocumentExternal#b6f74335 flags:# url:string caption:string ttl_seconds:flags.0?int = InputMedia;
func (m *TLInputMediaDocumentExternal) To_InputMedia() *InputMedia {
	return &InputMedia{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputMediaDocumentExternal) SetUrl(v string) { m.Data2.Url = v }
func (m *TLInputMediaDocumentExternal) GetUrl() string  { return m.Data2.Url }

func (m *TLInputMediaDocumentExternal) SetCaption(v string) { m.Data2.Caption = v }
func (m *TLInputMediaDocumentExternal) GetCaption() string  { return m.Data2.Caption }

func (m *TLInputMediaDocumentExternal) SetTtlSeconds(v int32) { m.Data2.TtlSeconds = v }
func (m *TLInputMediaDocumentExternal) GetTtlSeconds() int32  { return m.Data2.TtlSeconds }

func (m *TLInputMediaDocumentExternal) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputMediaDocumentExternal))

	// flags
	var flags uint32 = 0
	if m.GetTtlSeconds() != 0 {
		flags |= 1 << 3
	}
	x.UInt(flags)

	x.StringBytes(m.Data2.Url)
	x.StringBytes(m.Data2.Caption)
	if m.GetTtlSeconds() != 0 {
		x.Int(m.Data2.TtlSeconds)
	}

	return x.buf
}

func (m *TLInputMediaDocumentExternal) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	m.Data2.Url = x.StringBytes()
	m.Data2.Caption = x.StringBytes()
	if (flags & (1 << 4)) != 0 {
		m.Data2.TtlSeconds = x.Int()
	}

	return dbuf.err
}

// inputMediaGame#d33f43f3 id:InputGame = InputMedia;
func (m *TLInputMediaGame) To_InputMedia() *InputMedia {
	return &InputMedia{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputMediaGame) SetId(v *InputGame) { m.Data2.Id_20 = v }
func (m *TLInputMediaGame) GetId() *InputGame  { return m.Data2.Id_20 }

func (m *TLInputMediaGame) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputMediaGame))

	x.Bytes(m.Data2Id.Encode())

	return x.buf
}

func (m *TLInputMediaGame) Decode(dbuf *DecodeBuf) error {
	m1 := &Id{}
	m1.Decode(dbuf)
	m.SetId(m1)

	return dbuf.err
}

// inputMediaInvoice#92153685 flags:# title:string description:string photo:flags.0?InputWebDocument invoice:Invoice payload:bytes provider:string start_param:string = InputMedia;
func (m *TLInputMediaInvoice) To_InputMedia() *InputMedia {
	return &InputMedia{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputMediaInvoice) SetTitle(v string) { m.Data2.Title = v }
func (m *TLInputMediaInvoice) GetTitle() string  { return m.Data2.Title }

func (m *TLInputMediaInvoice) SetDescription(v string) { m.Data2.Description = v }
func (m *TLInputMediaInvoice) GetDescription() string  { return m.Data2.Description }

func (m *TLInputMediaInvoice) SetPhoto(v *InputWebDocument) { m.Data2.Photo = v }
func (m *TLInputMediaInvoice) GetPhoto() *InputWebDocument  { return m.Data2.Photo }

func (m *TLInputMediaInvoice) SetInvoice(v *Invoice) { m.Data2.Invoice = v }
func (m *TLInputMediaInvoice) GetInvoice() *Invoice  { return m.Data2.Invoice }

func (m *TLInputMediaInvoice) SetPayload(v []byte) { m.Data2.Payload = v }
func (m *TLInputMediaInvoice) GetPayload() []byte  { return m.Data2.Payload }

func (m *TLInputMediaInvoice) SetProvider(v string) { m.Data2.Provider = v }
func (m *TLInputMediaInvoice) GetProvider() string  { return m.Data2.Provider }

func (m *TLInputMediaInvoice) SetStartParam(v string) { m.Data2.StartParam = v }
func (m *TLInputMediaInvoice) GetStartParam() string  { return m.Data2.StartParam }

func (m *TLInputMediaInvoice) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputMediaInvoice))

	// flags
	var flags uint32 = 0
	if m.GetPhoto() != nil {
		flags |= 1 << 3
	}
	x.UInt(flags)

	x.StringBytes(m.Data2.Title)
	x.StringBytes(m.Data2.Description)
	if m.GetPhoto() != 0 {
		x.Bytes(m.Data2Photo.Encode())
	}
	x.Bytes(m.Data2Invoice.Encode())
	x.StringBytes(m.Data2.Payload)
	x.StringBytes(m.Data2.Provider)
	x.StringBytes(m.Data2.StartParam)

	return x.buf
}

func (m *TLInputMediaInvoice) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	m.Data2.Title = x.StringBytes()
	m.Data2.Description = x.StringBytes()
	if (flags & (1 << 4)) != 0 {
		m4 := &Photo{}
		m4.Decode(dbuf)
		m.SetPhoto(m4)
	}
	m5 := &Invoice{}
	m5.Decode(dbuf)
	m.SetInvoice(m5)
	m.Data2.Payload = x.StringBytes()
	m.Data2.Provider = x.StringBytes()
	m.Data2.StartParam = x.StringBytes()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// User <--
//  + TL_UserEmpty
//  + TL_User
//

func (m *User) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_userEmpty:
		t := m.To_UserEmpty()
		return t.Encode()
	case TLConstructor_CRC32_user:
		t := m.To_User()
		return t.Encode()

	default:
		return []byte{}
	}
}

// userEmpty#200250ba id:int = User;
func (m *User) To_UserEmpty() *TLUserEmpty {
	return &TLUserEmpty{
		Data2: m.Data2,
	}
}

// user#2e13f4c3 flags:# self:flags.10?true contact:flags.11?true mutual_contact:flags.12?true deleted:flags.13?true bot:flags.14?true bot_chat_history:flags.15?true bot_nochats:flags.16?true verified:flags.17?true restricted:flags.18?true min:flags.20?true bot_inline_geo:flags.21?true id:int access_hash:flags.0?long first_name:flags.1?string last_name:flags.2?string username:flags.3?string phone:flags.4?string photo:flags.5?UserProfilePhoto status:flags.6?UserStatus bot_info_version:flags.14?int restriction_reason:flags.18?string bot_inline_placeholder:flags.19?string lang_code:flags.22?string = User;
func (m *User) To_User() *TLUser {
	return &TLUser{
		Data2: m.Data2,
	}
}

// userEmpty#200250ba id:int = User;
func (m *TLUserEmpty) To_User() *User {
	return &User{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUserEmpty) SetId(v int32) { m.Data2.Id = v }
func (m *TLUserEmpty) GetId() int32  { return m.Data2.Id }

func (m *TLUserEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_userEmpty))

	x.Int(m.Data2.Id)

	return x.buf
}

func (m *TLUserEmpty) Decode(dbuf *DecodeBuf) error {
	m.Data2.Id = x.Int()

	return dbuf.err
}

// user#2e13f4c3 flags:# self:flags.10?true contact:flags.11?true mutual_contact:flags.12?true deleted:flags.13?true bot:flags.14?true bot_chat_history:flags.15?true bot_nochats:flags.16?true verified:flags.17?true restricted:flags.18?true min:flags.20?true bot_inline_geo:flags.21?true id:int access_hash:flags.0?long first_name:flags.1?string last_name:flags.2?string username:flags.3?string phone:flags.4?string photo:flags.5?UserProfilePhoto status:flags.6?UserStatus bot_info_version:flags.14?int restriction_reason:flags.18?string bot_inline_placeholder:flags.19?string lang_code:flags.22?string = User;
func (m *TLUser) To_User() *User {
	return &User{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUser) SetSelf(v bool) { m.Data2.Self = v }
func (m *TLUser) GetSelf() bool  { return m.Data2.Self }

func (m *TLUser) SetContact(v bool) { m.Data2.Contact = v }
func (m *TLUser) GetContact() bool  { return m.Data2.Contact }

func (m *TLUser) SetMutualContact(v bool) { m.Data2.MutualContact = v }
func (m *TLUser) GetMutualContact() bool  { return m.Data2.MutualContact }

func (m *TLUser) SetDeleted(v bool) { m.Data2.Deleted = v }
func (m *TLUser) GetDeleted() bool  { return m.Data2.Deleted }

func (m *TLUser) SetBot(v bool) { m.Data2.Bot = v }
func (m *TLUser) GetBot() bool  { return m.Data2.Bot }

func (m *TLUser) SetBotChatHistory(v bool) { m.Data2.BotChatHistory = v }
func (m *TLUser) GetBotChatHistory() bool  { return m.Data2.BotChatHistory }

func (m *TLUser) SetBotNochats(v bool) { m.Data2.BotNochats = v }
func (m *TLUser) GetBotNochats() bool  { return m.Data2.BotNochats }

func (m *TLUser) SetVerified(v bool) { m.Data2.Verified = v }
func (m *TLUser) GetVerified() bool  { return m.Data2.Verified }

func (m *TLUser) SetRestricted(v bool) { m.Data2.Restricted = v }
func (m *TLUser) GetRestricted() bool  { return m.Data2.Restricted }

func (m *TLUser) SetMin(v bool) { m.Data2.Min = v }
func (m *TLUser) GetMin() bool  { return m.Data2.Min }

func (m *TLUser) SetBotInlineGeo(v bool) { m.Data2.BotInlineGeo = v }
func (m *TLUser) GetBotInlineGeo() bool  { return m.Data2.BotInlineGeo }

func (m *TLUser) SetId(v int32) { m.Data2.Id = v }
func (m *TLUser) GetId() int32  { return m.Data2.Id }

func (m *TLUser) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLUser) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLUser) SetFirstName(v string) { m.Data2.FirstName = v }
func (m *TLUser) GetFirstName() string  { return m.Data2.FirstName }

func (m *TLUser) SetLastName(v string) { m.Data2.LastName = v }
func (m *TLUser) GetLastName() string  { return m.Data2.LastName }

func (m *TLUser) SetUsername(v string) { m.Data2.Username = v }
func (m *TLUser) GetUsername() string  { return m.Data2.Username }

func (m *TLUser) SetPhone(v string) { m.Data2.Phone = v }
func (m *TLUser) GetPhone() string  { return m.Data2.Phone }

func (m *TLUser) SetPhoto(v *UserProfilePhoto) { m.Data2.Photo = v }
func (m *TLUser) GetPhoto() *UserProfilePhoto  { return m.Data2.Photo }

func (m *TLUser) SetStatus(v *UserStatus) { m.Data2.Status = v }
func (m *TLUser) GetStatus() *UserStatus  { return m.Data2.Status }

func (m *TLUser) SetBotInfoVersion(v int32) { m.Data2.BotInfoVersion = v }
func (m *TLUser) GetBotInfoVersion() int32  { return m.Data2.BotInfoVersion }

func (m *TLUser) SetRestrictionReason(v string) { m.Data2.RestrictionReason = v }
func (m *TLUser) GetRestrictionReason() string  { return m.Data2.RestrictionReason }

func (m *TLUser) SetBotInlinePlaceholder(v string) { m.Data2.BotInlinePlaceholder = v }
func (m *TLUser) GetBotInlinePlaceholder() string  { return m.Data2.BotInlinePlaceholder }

func (m *TLUser) SetLangCode(v string) { m.Data2.LangCode = v }
func (m *TLUser) GetLangCode() string  { return m.Data2.LangCode }

func (m *TLUser) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_user))

	// flags
	var flags uint32 = 0
	if m.GetSelf() == true {
		flags |= 1 << 1
	}
	if m.GetContact() == true {
		flags |= 1 << 2
	}
	if m.GetMutualContact() == true {
		flags |= 1 << 3
	}
	if m.GetDeleted() == true {
		flags |= 1 << 4
	}
	if m.GetBot() == true {
		flags |= 1 << 5
	}
	if m.GetBotChatHistory() == true {
		flags |= 1 << 6
	}
	if m.GetBotNochats() == true {
		flags |= 1 << 7
	}
	if m.GetVerified() == true {
		flags |= 1 << 8
	}
	if m.GetRestricted() == true {
		flags |= 1 << 9
	}
	if m.GetMin() == true {
		flags |= 1 << 10
	}
	if m.GetBotInlineGeo() == true {
		flags |= 1 << 11
	}
	if m.GetAccessHash() != 0 {
		flags |= 1 << 13
	}
	if m.GetFirstName() != "" {
		flags |= 1 << 14
	}
	if m.GetLastName() != "" {
		flags |= 1 << 15
	}
	if m.GetUsername() != "" {
		flags |= 1 << 16
	}
	if m.GetPhone() != "" {
		flags |= 1 << 17
	}
	if m.GetPhoto() != nil {
		flags |= 1 << 18
	}
	if m.GetStatus() != nil {
		flags |= 1 << 19
	}
	if m.GetBotInfoVersion() != 0 {
		flags |= 1 << 20
	}
	if m.GetRestrictionReason() != "" {
		flags |= 1 << 21
	}
	if m.GetBotInlinePlaceholder() != "" {
		flags |= 1 << 22
	}
	if m.GetLangCode() != "" {
		flags |= 1 << 23
	}
	x.UInt(flags)

	x.Int(m.Data2.Id)
	if m.GetAccessHash() != 0 {
		x.Long(m.Data2.AccessHash)
	}
	if m.GetFirstName() != 0 {
		x.StringBytes(m.Data2.FirstName)
	}
	if m.GetLastName() != 0 {
		x.StringBytes(m.Data2.LastName)
	}
	if m.GetUsername() != 0 {
		x.StringBytes(m.Data2.Username)
	}
	if m.GetPhone() != 0 {
		x.StringBytes(m.Data2.Phone)
	}
	if m.GetPhoto() != 0 {
		x.Bytes(m.Data2Photo.Encode())
	}
	if m.GetStatus() != 0 {
		x.Bytes(m.Data2Status.Encode())
	}
	if m.GetBotInfoVersion() != 0 {
		x.Int(m.Data2.BotInfoVersion)
	}
	if m.GetRestrictionReason() != 0 {
		x.StringBytes(m.Data2.RestrictionReason)
	}
	if m.GetBotInlinePlaceholder() != 0 {
		x.StringBytes(m.Data2.BotInlinePlaceholder)
	}
	if m.GetLangCode() != 0 {
		x.StringBytes(m.Data2.LangCode)
	}

	return x.buf
}

func (m *TLUser) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m.Data2.Self = true
	}
	if (flags & (1 << 3)) != 0 {
		m.Data2.Contact = true
	}
	if (flags & (1 << 4)) != 0 {
		m.Data2.MutualContact = true
	}
	if (flags & (1 << 5)) != 0 {
		m.Data2.Deleted = true
	}
	if (flags & (1 << 6)) != 0 {
		m.Data2.Bot = true
	}
	if (flags & (1 << 7)) != 0 {
		m.Data2.BotChatHistory = true
	}
	if (flags & (1 << 8)) != 0 {
		m.Data2.BotNochats = true
	}
	if (flags & (1 << 9)) != 0 {
		m.Data2.Verified = true
	}
	if (flags & (1 << 10)) != 0 {
		m.Data2.Restricted = true
	}
	if (flags & (1 << 11)) != 0 {
		m.Data2.Min = true
	}
	if (flags & (1 << 12)) != 0 {
		m.Data2.BotInlineGeo = true
	}
	m.Data2.Id = x.Int()
	if (flags & (1 << 14)) != 0 {
		m.Data2.AccessHash = x.Long()
	}
	if (flags & (1 << 15)) != 0 {
		m.Data2.FirstName = x.StringBytes()
	}
	if (flags & (1 << 16)) != 0 {
		m.Data2.LastName = x.StringBytes()
	}
	if (flags & (1 << 17)) != 0 {
		m.Data2.Username = x.StringBytes()
	}
	if (flags & (1 << 18)) != 0 {
		m.Data2.Phone = x.StringBytes()
	}
	if (flags & (1 << 19)) != 0 {
		m19 := &Photo{}
		m19.Decode(dbuf)
		m.SetPhoto(m19)
	}
	if (flags & (1 << 20)) != 0 {
		m20 := &Status{}
		m20.Decode(dbuf)
		m.SetStatus(m20)
	}
	if (flags & (1 << 21)) != 0 {
		m.Data2.BotInfoVersion = x.Int()
	}
	if (flags & (1 << 22)) != 0 {
		m.Data2.RestrictionReason = x.StringBytes()
	}
	if (flags & (1 << 23)) != 0 {
		m.Data2.BotInlinePlaceholder = x.StringBytes()
	}
	if (flags & (1 << 24)) != 0 {
		m.Data2.LangCode = x.StringBytes()
	}

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Chat <--
//  + TL_ChatEmpty
//  + TL_Chat
//  + TL_ChatForbidden
//  + TL_Channel
//  + TL_ChannelForbidden
//

func (m *Chat) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_chatEmpty:
		t := m.To_ChatEmpty()
		return t.Encode()
	case TLConstructor_CRC32_chat:
		t := m.To_Chat()
		return t.Encode()
	case TLConstructor_CRC32_chatForbidden:
		t := m.To_ChatForbidden()
		return t.Encode()
	case TLConstructor_CRC32_channel:
		t := m.To_Channel()
		return t.Encode()
	case TLConstructor_CRC32_channelForbidden:
		t := m.To_ChannelForbidden()
		return t.Encode()

	default:
		return []byte{}
	}
}

// chatEmpty#9ba2d800 id:int = Chat;
func (m *Chat) To_ChatEmpty() *TLChatEmpty {
	return &TLChatEmpty{
		Data2: m.Data2,
	}
}

// chat#d91cdd54 flags:# creator:flags.0?true kicked:flags.1?true left:flags.2?true admins_enabled:flags.3?true admin:flags.4?true deactivated:flags.5?true id:int title:string photo:ChatPhoto participants_count:int date:int version:int migrated_to:flags.6?InputChannel = Chat;
func (m *Chat) To_Chat() *TLChat {
	return &TLChat{
		Data2: m.Data2,
	}
}

// chatForbidden#7328bdb id:int title:string = Chat;
func (m *Chat) To_ChatForbidden() *TLChatForbidden {
	return &TLChatForbidden{
		Data2: m.Data2,
	}
}

// channel#cb44b1c flags:# creator:flags.0?true left:flags.2?true editor:flags.3?true broadcast:flags.5?true verified:flags.7?true megagroup:flags.8?true restricted:flags.9?true democracy:flags.10?true signatures:flags.11?true min:flags.12?true id:int access_hash:flags.13?long title:string username:flags.6?string photo:ChatPhoto date:int version:int restriction_reason:flags.9?string admin_rights:flags.14?ChannelAdminRights banned_rights:flags.15?ChannelBannedRights = Chat;
func (m *Chat) To_Channel() *TLChannel {
	return &TLChannel{
		Data2: m.Data2,
	}
}

// channelForbidden#289da732 flags:# broadcast:flags.5?true megagroup:flags.8?true id:int access_hash:long title:string until_date:flags.16?int = Chat;
func (m *Chat) To_ChannelForbidden() *TLChannelForbidden {
	return &TLChannelForbidden{
		Data2: m.Data2,
	}
}

// chatEmpty#9ba2d800 id:int = Chat;
func (m *TLChatEmpty) To_Chat() *Chat {
	return &Chat{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLChatEmpty) SetId(v int32) { m.Data2.Id = v }
func (m *TLChatEmpty) GetId() int32  { return m.Data2.Id }

func (m *TLChatEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_chatEmpty))

	x.Int(m.Data2.Id)

	return x.buf
}

func (m *TLChatEmpty) Decode(dbuf *DecodeBuf) error {
	m.Data2.Id = x.Int()

	return dbuf.err
}

// chat#d91cdd54 flags:# creator:flags.0?true kicked:flags.1?true left:flags.2?true admins_enabled:flags.3?true admin:flags.4?true deactivated:flags.5?true id:int title:string photo:ChatPhoto participants_count:int date:int version:int migrated_to:flags.6?InputChannel = Chat;
func (m *TLChat) To_Chat() *Chat {
	return &Chat{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLChat) SetCreator(v bool) { m.Data2.Creator = v }
func (m *TLChat) GetCreator() bool  { return m.Data2.Creator }

func (m *TLChat) SetKicked(v bool) { m.Data2.Kicked = v }
func (m *TLChat) GetKicked() bool  { return m.Data2.Kicked }

func (m *TLChat) SetLeft(v bool) { m.Data2.Left = v }
func (m *TLChat) GetLeft() bool  { return m.Data2.Left }

func (m *TLChat) SetAdminsEnabled(v bool) { m.Data2.AdminsEnabled = v }
func (m *TLChat) GetAdminsEnabled() bool  { return m.Data2.AdminsEnabled }

func (m *TLChat) SetAdmin(v bool) { m.Data2.Admin = v }
func (m *TLChat) GetAdmin() bool  { return m.Data2.Admin }

func (m *TLChat) SetDeactivated(v bool) { m.Data2.Deactivated = v }
func (m *TLChat) GetDeactivated() bool  { return m.Data2.Deactivated }

func (m *TLChat) SetId(v int32) { m.Data2.Id = v }
func (m *TLChat) GetId() int32  { return m.Data2.Id }

func (m *TLChat) SetTitle(v string) { m.Data2.Title = v }
func (m *TLChat) GetTitle() string  { return m.Data2.Title }

func (m *TLChat) SetPhoto(v *ChatPhoto) { m.Data2.Photo = v }
func (m *TLChat) GetPhoto() *ChatPhoto  { return m.Data2.Photo }

func (m *TLChat) SetParticipantsCount(v int32) { m.Data2.ParticipantsCount = v }
func (m *TLChat) GetParticipantsCount() int32  { return m.Data2.ParticipantsCount }

func (m *TLChat) SetDate(v int32) { m.Data2.Date = v }
func (m *TLChat) GetDate() int32  { return m.Data2.Date }

func (m *TLChat) SetVersion(v int32) { m.Data2.Version = v }
func (m *TLChat) GetVersion() int32  { return m.Data2.Version }

func (m *TLChat) SetMigratedTo(v *InputChannel) { m.Data2.MigratedTo = v }
func (m *TLChat) GetMigratedTo() *InputChannel  { return m.Data2.MigratedTo }

func (m *TLChat) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_chat))

	// flags
	var flags uint32 = 0
	if m.GetCreator() == true {
		flags |= 1 << 1
	}
	if m.GetKicked() == true {
		flags |= 1 << 2
	}
	if m.GetLeft() == true {
		flags |= 1 << 3
	}
	if m.GetAdminsEnabled() == true {
		flags |= 1 << 4
	}
	if m.GetAdmin() == true {
		flags |= 1 << 5
	}
	if m.GetDeactivated() == true {
		flags |= 1 << 6
	}
	if m.GetMigratedTo() != nil {
		flags |= 1 << 13
	}
	x.UInt(flags)

	x.Int(m.Data2.Id)
	x.StringBytes(m.Data2.Title)
	x.Bytes(m.Data2Photo.Encode())
	x.Int(m.Data2.ParticipantsCount)
	x.Int(m.Data2.Date)
	x.Int(m.Data2.Version)
	if m.GetMigratedTo() != 0 {
		x.Bytes(m.Data2MigratedTo.Encode())
	}

	return x.buf
}

func (m *TLChat) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m.Data2.Creator = true
	}
	if (flags & (1 << 3)) != 0 {
		m.Data2.Kicked = true
	}
	if (flags & (1 << 4)) != 0 {
		m.Data2.Left = true
	}
	if (flags & (1 << 5)) != 0 {
		m.Data2.AdminsEnabled = true
	}
	if (flags & (1 << 6)) != 0 {
		m.Data2.Admin = true
	}
	if (flags & (1 << 7)) != 0 {
		m.Data2.Deactivated = true
	}
	m.Data2.Id = x.Int()
	m.Data2.Title = x.StringBytes()
	m10 := &Photo{}
	m10.Decode(dbuf)
	m.SetPhoto(m10)
	m.Data2.ParticipantsCount = x.Int()
	m.Data2.Date = x.Int()
	m.Data2.Version = x.Int()
	if (flags & (1 << 14)) != 0 {
		m14 := &MigratedTo{}
		m14.Decode(dbuf)
		m.SetMigratedTo(m14)
	}

	return dbuf.err
}

// chatForbidden#7328bdb id:int title:string = Chat;
func (m *TLChatForbidden) To_Chat() *Chat {
	return &Chat{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLChatForbidden) SetId(v int32) { m.Data2.Id = v }
func (m *TLChatForbidden) GetId() int32  { return m.Data2.Id }

func (m *TLChatForbidden) SetTitle(v string) { m.Data2.Title = v }
func (m *TLChatForbidden) GetTitle() string  { return m.Data2.Title }

func (m *TLChatForbidden) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_chatForbidden))

	x.Int(m.Data2.Id)
	x.StringBytes(m.Data2.Title)

	return x.buf
}

func (m *TLChatForbidden) Decode(dbuf *DecodeBuf) error {
	m.Data2.Id = x.Int()
	m.Data2.Title = x.StringBytes()

	return dbuf.err
}

// channel#cb44b1c flags:# creator:flags.0?true left:flags.2?true editor:flags.3?true broadcast:flags.5?true verified:flags.7?true megagroup:flags.8?true restricted:flags.9?true democracy:flags.10?true signatures:flags.11?true min:flags.12?true id:int access_hash:flags.13?long title:string username:flags.6?string photo:ChatPhoto date:int version:int restriction_reason:flags.9?string admin_rights:flags.14?ChannelAdminRights banned_rights:flags.15?ChannelBannedRights = Chat;
func (m *TLChannel) To_Chat() *Chat {
	return &Chat{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLChannel) SetCreator(v bool) { m.Data2.Creator = v }
func (m *TLChannel) GetCreator() bool  { return m.Data2.Creator }

func (m *TLChannel) SetLeft(v bool) { m.Data2.Left = v }
func (m *TLChannel) GetLeft() bool  { return m.Data2.Left }

func (m *TLChannel) SetEditor(v bool) { m.Data2.Editor = v }
func (m *TLChannel) GetEditor() bool  { return m.Data2.Editor }

func (m *TLChannel) SetBroadcast(v bool) { m.Data2.Broadcast = v }
func (m *TLChannel) GetBroadcast() bool  { return m.Data2.Broadcast }

func (m *TLChannel) SetVerified(v bool) { m.Data2.Verified = v }
func (m *TLChannel) GetVerified() bool  { return m.Data2.Verified }

func (m *TLChannel) SetMegagroup(v bool) { m.Data2.Megagroup = v }
func (m *TLChannel) GetMegagroup() bool  { return m.Data2.Megagroup }

func (m *TLChannel) SetRestricted(v bool) { m.Data2.Restricted = v }
func (m *TLChannel) GetRestricted() bool  { return m.Data2.Restricted }

func (m *TLChannel) SetDemocracy(v bool) { m.Data2.Democracy = v }
func (m *TLChannel) GetDemocracy() bool  { return m.Data2.Democracy }

func (m *TLChannel) SetSignatures(v bool) { m.Data2.Signatures = v }
func (m *TLChannel) GetSignatures() bool  { return m.Data2.Signatures }

func (m *TLChannel) SetMin(v bool) { m.Data2.Min = v }
func (m *TLChannel) GetMin() bool  { return m.Data2.Min }

func (m *TLChannel) SetId(v int32) { m.Data2.Id = v }
func (m *TLChannel) GetId() int32  { return m.Data2.Id }

func (m *TLChannel) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLChannel) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLChannel) SetTitle(v string) { m.Data2.Title = v }
func (m *TLChannel) GetTitle() string  { return m.Data2.Title }

func (m *TLChannel) SetUsername(v string) { m.Data2.Username = v }
func (m *TLChannel) GetUsername() string  { return m.Data2.Username }

func (m *TLChannel) SetPhoto(v *ChatPhoto) { m.Data2.Photo = v }
func (m *TLChannel) GetPhoto() *ChatPhoto  { return m.Data2.Photo }

func (m *TLChannel) SetDate(v int32) { m.Data2.Date = v }
func (m *TLChannel) GetDate() int32  { return m.Data2.Date }

func (m *TLChannel) SetVersion(v int32) { m.Data2.Version = v }
func (m *TLChannel) GetVersion() int32  { return m.Data2.Version }

func (m *TLChannel) SetRestrictionReason(v string) { m.Data2.RestrictionReason = v }
func (m *TLChannel) GetRestrictionReason() string  { return m.Data2.RestrictionReason }

func (m *TLChannel) SetAdminRights(v *ChannelAdminRights) { m.Data2.AdminRights = v }
func (m *TLChannel) GetAdminRights() *ChannelAdminRights  { return m.Data2.AdminRights }

func (m *TLChannel) SetBannedRights(v *ChannelBannedRights) { m.Data2.BannedRights = v }
func (m *TLChannel) GetBannedRights() *ChannelBannedRights  { return m.Data2.BannedRights }

func (m *TLChannel) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channel))

	// flags
	var flags uint32 = 0
	if m.GetCreator() == true {
		flags |= 1 << 1
	}
	if m.GetLeft() == true {
		flags |= 1 << 2
	}
	if m.GetEditor() == true {
		flags |= 1 << 3
	}
	if m.GetBroadcast() == true {
		flags |= 1 << 4
	}
	if m.GetVerified() == true {
		flags |= 1 << 5
	}
	if m.GetMegagroup() == true {
		flags |= 1 << 6
	}
	if m.GetRestricted() == true {
		flags |= 1 << 7
	}
	if m.GetDemocracy() == true {
		flags |= 1 << 8
	}
	if m.GetSignatures() == true {
		flags |= 1 << 9
	}
	if m.GetMin() == true {
		flags |= 1 << 10
	}
	if m.GetAccessHash() != 0 {
		flags |= 1 << 12
	}
	if m.GetUsername() != "" {
		flags |= 1 << 14
	}
	if m.GetRestrictionReason() != "" {
		flags |= 1 << 18
	}
	if m.GetAdminRights() != nil {
		flags |= 1 << 19
	}
	if m.GetBannedRights() != nil {
		flags |= 1 << 20
	}
	x.UInt(flags)

	x.Int(m.Data2.Id)
	if m.GetAccessHash() != 0 {
		x.Long(m.Data2.AccessHash)
	}
	x.StringBytes(m.Data2.Title)
	if m.GetUsername() != 0 {
		x.StringBytes(m.Data2.Username)
	}
	x.Bytes(m.Data2Photo.Encode())
	x.Int(m.Data2.Date)
	x.Int(m.Data2.Version)
	if m.GetRestrictionReason() != 0 {
		x.StringBytes(m.Data2.RestrictionReason)
	}
	if m.GetAdminRights() != 0 {
		x.Bytes(m.Data2AdminRights.Encode())
	}
	if m.GetBannedRights() != 0 {
		x.Bytes(m.Data2BannedRights.Encode())
	}

	return x.buf
}

func (m *TLChannel) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m.Data2.Creator = true
	}
	if (flags & (1 << 3)) != 0 {
		m.Data2.Left = true
	}
	if (flags & (1 << 4)) != 0 {
		m.Data2.Editor = true
	}
	if (flags & (1 << 5)) != 0 {
		m.Data2.Broadcast = true
	}
	if (flags & (1 << 6)) != 0 {
		m.Data2.Verified = true
	}
	if (flags & (1 << 7)) != 0 {
		m.Data2.Megagroup = true
	}
	if (flags & (1 << 8)) != 0 {
		m.Data2.Restricted = true
	}
	if (flags & (1 << 9)) != 0 {
		m.Data2.Democracy = true
	}
	if (flags & (1 << 10)) != 0 {
		m.Data2.Signatures = true
	}
	if (flags & (1 << 11)) != 0 {
		m.Data2.Min = true
	}
	m.Data2.Id = x.Int()
	if (flags & (1 << 13)) != 0 {
		m.Data2.AccessHash = x.Long()
	}
	m.Data2.Title = x.StringBytes()
	if (flags & (1 << 15)) != 0 {
		m.Data2.Username = x.StringBytes()
	}
	m16 := &Photo{}
	m16.Decode(dbuf)
	m.SetPhoto(m16)
	m.Data2.Date = x.Int()
	m.Data2.Version = x.Int()
	if (flags & (1 << 19)) != 0 {
		m.Data2.RestrictionReason = x.StringBytes()
	}
	if (flags & (1 << 20)) != 0 {
		m20 := &AdminRights{}
		m20.Decode(dbuf)
		m.SetAdminRights(m20)
	}
	if (flags & (1 << 21)) != 0 {
		m21 := &BannedRights{}
		m21.Decode(dbuf)
		m.SetBannedRights(m21)
	}

	return dbuf.err
}

// channelForbidden#289da732 flags:# broadcast:flags.5?true megagroup:flags.8?true id:int access_hash:long title:string until_date:flags.16?int = Chat;
func (m *TLChannelForbidden) To_Chat() *Chat {
	return &Chat{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLChannelForbidden) SetBroadcast(v bool) { m.Data2.Broadcast = v }
func (m *TLChannelForbidden) GetBroadcast() bool  { return m.Data2.Broadcast }

func (m *TLChannelForbidden) SetMegagroup(v bool) { m.Data2.Megagroup = v }
func (m *TLChannelForbidden) GetMegagroup() bool  { return m.Data2.Megagroup }

func (m *TLChannelForbidden) SetId(v int32) { m.Data2.Id = v }
func (m *TLChannelForbidden) GetId() int32  { return m.Data2.Id }

func (m *TLChannelForbidden) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLChannelForbidden) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLChannelForbidden) SetTitle(v string) { m.Data2.Title = v }
func (m *TLChannelForbidden) GetTitle() string  { return m.Data2.Title }

func (m *TLChannelForbidden) SetUntilDate(v int32) { m.Data2.UntilDate = v }
func (m *TLChannelForbidden) GetUntilDate() int32  { return m.Data2.UntilDate }

func (m *TLChannelForbidden) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelForbidden))

	// flags
	var flags uint32 = 0
	if m.GetBroadcast() == true {
		flags |= 1 << 1
	}
	if m.GetMegagroup() == true {
		flags |= 1 << 2
	}
	if m.GetUntilDate() != 0 {
		flags |= 1 << 6
	}
	x.UInt(flags)

	x.Int(m.Data2.Id)
	x.Long(m.Data2.AccessHash)
	x.StringBytes(m.Data2.Title)
	if m.GetUntilDate() != 0 {
		x.Int(m.Data2.UntilDate)
	}

	return x.buf
}

func (m *TLChannelForbidden) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m.Data2.Broadcast = true
	}
	if (flags & (1 << 3)) != 0 {
		m.Data2.Megagroup = true
	}
	m.Data2.Id = x.Int()
	m.Data2.AccessHash = x.Long()
	m.Data2.Title = x.StringBytes()
	if (flags & (1 << 7)) != 0 {
		m.Data2.UntilDate = x.Int()
	}

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Updates_State <--
//  + TL_UpdatesState
//

func (m *Updates_State) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_updates_state:
		t := m.To_UpdatesState()
		return t.Encode()

	default:
		return []byte{}
	}
}

// updates.state#a56c2a3e pts:int qts:int date:int seq:int unread_count:int = updates.State;
func (m *Updates_State) To_UpdatesState() *TLUpdatesState {
	return &TLUpdatesState{
		Data2: m.Data2,
	}
}

// updates.state#a56c2a3e pts:int qts:int date:int seq:int unread_count:int = updates.State;
func (m *TLUpdatesState) To_Updates_State() *Updates_State {
	return &Updates_State{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdatesState) SetPts(v int32) { m.Data2.Pts = v }
func (m *TLUpdatesState) GetPts() int32  { return m.Data2.Pts }

func (m *TLUpdatesState) SetQts(v int32) { m.Data2.Qts = v }
func (m *TLUpdatesState) GetQts() int32  { return m.Data2.Qts }

func (m *TLUpdatesState) SetDate(v int32) { m.Data2.Date = v }
func (m *TLUpdatesState) GetDate() int32  { return m.Data2.Date }

func (m *TLUpdatesState) SetSeq(v int32) { m.Data2.Seq = v }
func (m *TLUpdatesState) GetSeq() int32  { return m.Data2.Seq }

func (m *TLUpdatesState) SetUnreadCount(v int32) { m.Data2.UnreadCount = v }
func (m *TLUpdatesState) GetUnreadCount() int32  { return m.Data2.UnreadCount }

func (m *TLUpdatesState) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updates_state))

	x.Int(m.Data2.Pts)
	x.Int(m.Data2.Qts)
	x.Int(m.Data2.Date)
	x.Int(m.Data2.Seq)
	x.Int(m.Data2.UnreadCount)

	return x.buf
}

func (m *TLUpdatesState) Decode(dbuf *DecodeBuf) error {
	m.Data2.Pts = x.Int()
	m.Data2.Qts = x.Int()
	m.Data2.Date = x.Int()
	m.Data2.Seq = x.Int()
	m.Data2.UnreadCount = x.Int()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// TopPeerCategoryPeers <--
//  + TL_TopPeerCategoryPeers
//

func (m *TopPeerCategoryPeers) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_topPeerCategoryPeers:
		t := m.To_TopPeerCategoryPeers()
		return t.Encode()

	default:
		return []byte{}
	}
}

// topPeerCategoryPeers#fb834291 category:TopPeerCategory count:int peers:Vector<TopPeer> = TopPeerCategoryPeers;
func (m *TopPeerCategoryPeers) To_TopPeerCategoryPeers() *TLTopPeerCategoryPeers {
	return &TLTopPeerCategoryPeers{
		Data2: m.Data2,
	}
}

// topPeerCategoryPeers#fb834291 category:TopPeerCategory count:int peers:Vector<TopPeer> = TopPeerCategoryPeers;
func (m *TLTopPeerCategoryPeers) To_TopPeerCategoryPeers() *TopPeerCategoryPeers {
	return &TopPeerCategoryPeers{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLTopPeerCategoryPeers) SetCategory(v *TopPeerCategory) { m.Data2.Category = v }
func (m *TLTopPeerCategoryPeers) GetCategory() *TopPeerCategory  { return m.Data2.Category }

func (m *TLTopPeerCategoryPeers) SetCount(v int32) { m.Data2.Count = v }
func (m *TLTopPeerCategoryPeers) GetCount() int32  { return m.Data2.Count }

func (m *TLTopPeerCategoryPeers) SetPeers(v []*TopPeer) { m.Data2.Peers = v }
func (m *TLTopPeerCategoryPeers) GetPeers() []*TopPeer  { return m.Data2.Peers }

func (m *TLTopPeerCategoryPeers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_topPeerCategoryPeers))

	x.Bytes(m.Data2Category.Encode())
	x.Int(m.Data2.Count)

	return x.buf
}

func (m *TLTopPeerCategoryPeers) Decode(dbuf *DecodeBuf) error {
	m1 := &Category{}
	m1.Decode(dbuf)
	m.SetCategory(m1)
	m.Data2.Count = x.Int()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// InputGeoPoint <--
//  + TL_InputGeoPointEmpty
//  + TL_InputGeoPoint
//

func (m *InputGeoPoint) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_inputGeoPointEmpty:
		t := m.To_InputGeoPointEmpty()
		return t.Encode()
	case TLConstructor_CRC32_inputGeoPoint:
		t := m.To_InputGeoPoint()
		return t.Encode()

	default:
		return []byte{}
	}
}

// inputGeoPointEmpty#e4c123d6 = InputGeoPoint;
func (m *InputGeoPoint) To_InputGeoPointEmpty() *TLInputGeoPointEmpty {
	return &TLInputGeoPointEmpty{
		Data2: m.Data2,
	}
}

// inputGeoPoint#f3b7acc9 lat:double long:double = InputGeoPoint;
func (m *InputGeoPoint) To_InputGeoPoint() *TLInputGeoPoint {
	return &TLInputGeoPoint{
		Data2: m.Data2,
	}
}

// inputGeoPointEmpty#e4c123d6 = InputGeoPoint;
func (m *TLInputGeoPointEmpty) To_InputGeoPoint() *InputGeoPoint {
	return &InputGeoPoint{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputGeoPointEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputGeoPointEmpty))

	return x.buf
}

func (m *TLInputGeoPointEmpty) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// inputGeoPoint#f3b7acc9 lat:double long:double = InputGeoPoint;
func (m *TLInputGeoPoint) To_InputGeoPoint() *InputGeoPoint {
	return &InputGeoPoint{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputGeoPoint) SetLat(v float64) { m.Data2.Lat = v }
func (m *TLInputGeoPoint) GetLat() float64  { return m.Data2.Lat }

func (m *TLInputGeoPoint) SetLong(v float64) { m.Data2.Long = v }
func (m *TLInputGeoPoint) GetLong() float64  { return m.Data2.Long }

func (m *TLInputGeoPoint) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputGeoPoint))

	x.Double(m.Data2.Lat)
	x.Double(m.Data2.Long)

	return x.buf
}

func (m *TLInputGeoPoint) Decode(dbuf *DecodeBuf) error {
	m.Data2.Lat = x.Double()
	m.Data2.Long = x.Double()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// ExportedChatInvite <--
//  + TL_ChatInviteEmpty
//  + TL_ChatInviteExported
//

func (m *ExportedChatInvite) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_chatInviteEmpty:
		t := m.To_ChatInviteEmpty()
		return t.Encode()
	case TLConstructor_CRC32_chatInviteExported:
		t := m.To_ChatInviteExported()
		return t.Encode()

	default:
		return []byte{}
	}
}

// chatInviteEmpty#69df3769 = ExportedChatInvite;
func (m *ExportedChatInvite) To_ChatInviteEmpty() *TLChatInviteEmpty {
	return &TLChatInviteEmpty{
		Data2: m.Data2,
	}
}

// chatInviteExported#fc2e05bc link:string = ExportedChatInvite;
func (m *ExportedChatInvite) To_ChatInviteExported() *TLChatInviteExported {
	return &TLChatInviteExported{
		Data2: m.Data2,
	}
}

// chatInviteEmpty#69df3769 = ExportedChatInvite;
func (m *TLChatInviteEmpty) To_ExportedChatInvite() *ExportedChatInvite {
	return &ExportedChatInvite{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLChatInviteEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_chatInviteEmpty))

	return x.buf
}

func (m *TLChatInviteEmpty) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// chatInviteExported#fc2e05bc link:string = ExportedChatInvite;
func (m *TLChatInviteExported) To_ExportedChatInvite() *ExportedChatInvite {
	return &ExportedChatInvite{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLChatInviteExported) SetLink(v string) { m.Data2.Link = v }
func (m *TLChatInviteExported) GetLink() string  { return m.Data2.Link }

func (m *TLChatInviteExported) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_chatInviteExported))

	x.StringBytes(m.Data2.Link)

	return x.buf
}

func (m *TLChatInviteExported) Decode(dbuf *DecodeBuf) error {
	m.Data2.Link = x.StringBytes()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// ChatInvite <--
//  + TL_ChatInviteAlready
//  + TL_ChatInvite
//

func (m *ChatInvite) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_chatInviteAlready:
		t := m.To_ChatInviteAlready()
		return t.Encode()
	case TLConstructor_CRC32_chatInvite:
		t := m.To_ChatInvite()
		return t.Encode()

	default:
		return []byte{}
	}
}

// chatInviteAlready#5a686d7c chat:Chat = ChatInvite;
func (m *ChatInvite) To_ChatInviteAlready() *TLChatInviteAlready {
	return &TLChatInviteAlready{
		Data2: m.Data2,
	}
}

// chatInvite#db74f558 flags:# channel:flags.0?true broadcast:flags.1?true public:flags.2?true megagroup:flags.3?true title:string photo:ChatPhoto participants_count:int participants:flags.4?Vector<User> = ChatInvite;
func (m *ChatInvite) To_ChatInvite() *TLChatInvite {
	return &TLChatInvite{
		Data2: m.Data2,
	}
}

// chatInviteAlready#5a686d7c chat:Chat = ChatInvite;
func (m *TLChatInviteAlready) To_ChatInvite() *ChatInvite {
	return &ChatInvite{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLChatInviteAlready) SetChat(v *Chat) { m.Data2.Chat = v }
func (m *TLChatInviteAlready) GetChat() *Chat  { return m.Data2.Chat }

func (m *TLChatInviteAlready) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_chatInviteAlready))

	x.Bytes(m.Data2Chat.Encode())

	return x.buf
}

func (m *TLChatInviteAlready) Decode(dbuf *DecodeBuf) error {
	m1 := &Chat{}
	m1.Decode(dbuf)
	m.SetChat(m1)

	return dbuf.err
}

// chatInvite#db74f558 flags:# channel:flags.0?true broadcast:flags.1?true public:flags.2?true megagroup:flags.3?true title:string photo:ChatPhoto participants_count:int participants:flags.4?Vector<User> = ChatInvite;
func (m *TLChatInvite) To_ChatInvite() *ChatInvite {
	return &ChatInvite{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLChatInvite) SetChannel(v bool) { m.Data2.Channel = v }
func (m *TLChatInvite) GetChannel() bool  { return m.Data2.Channel }

func (m *TLChatInvite) SetBroadcast(v bool) { m.Data2.Broadcast = v }
func (m *TLChatInvite) GetBroadcast() bool  { return m.Data2.Broadcast }

func (m *TLChatInvite) SetPublic(v bool) { m.Data2.Public = v }
func (m *TLChatInvite) GetPublic() bool  { return m.Data2.Public }

func (m *TLChatInvite) SetMegagroup(v bool) { m.Data2.Megagroup = v }
func (m *TLChatInvite) GetMegagroup() bool  { return m.Data2.Megagroup }

func (m *TLChatInvite) SetTitle(v string) { m.Data2.Title = v }
func (m *TLChatInvite) GetTitle() string  { return m.Data2.Title }

func (m *TLChatInvite) SetPhoto(v *ChatPhoto) { m.Data2.Photo = v }
func (m *TLChatInvite) GetPhoto() *ChatPhoto  { return m.Data2.Photo }

func (m *TLChatInvite) SetParticipantsCount(v int32) { m.Data2.ParticipantsCount = v }
func (m *TLChatInvite) GetParticipantsCount() int32  { return m.Data2.ParticipantsCount }

func (m *TLChatInvite) SetParticipants(v []*User) { m.Data2.Participants = v }
func (m *TLChatInvite) GetParticipants() []*User  { return m.Data2.Participants }

func (m *TLChatInvite) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_chatInvite))

	// flags
	var flags uint32 = 0
	if m.GetChannel() == true {
		flags |= 1 << 1
	}
	if m.GetBroadcast() == true {
		flags |= 1 << 2
	}
	if m.GetPublic() == true {
		flags |= 1 << 3
	}
	if m.GetMegagroup() == true {
		flags |= 1 << 4
	}
	if m.GetParticipants() != nil {
		flags |= 1 << 8
	}
	x.UInt(flags)

	x.StringBytes(m.Data2.Title)
	x.Bytes(m.Data2Photo.Encode())
	x.Int(m.Data2.ParticipantsCount)
	if m.Data2.Participants != 0 {

	}

	return x.buf
}

func (m *TLChatInvite) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m.Data2.Channel = true
	}
	if (flags & (1 << 3)) != 0 {
		m.Data2.Broadcast = true
	}
	if (flags & (1 << 4)) != 0 {
		m.Data2.Public = true
	}
	if (flags & (1 << 5)) != 0 {
		m.Data2.Megagroup = true
	}
	m.Data2.Title = x.StringBytes()
	m7 := &Photo{}
	m7.Decode(dbuf)
	m.SetPhoto(m7)
	m.Data2.ParticipantsCount = x.Int()
	if (flags & (1 << 9)) != 0 {

	}

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Messages_MessageEditData <--
//  + TL_MessagesMessageEditData
//

func (m *Messages_MessageEditData) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_messages_messageEditData:
		t := m.To_MessagesMessageEditData()
		return t.Encode()

	default:
		return []byte{}
	}
}

// messages.messageEditData#26b5dde6 flags:# caption:flags.0?true = messages.MessageEditData;
func (m *Messages_MessageEditData) To_MessagesMessageEditData() *TLMessagesMessageEditData {
	return &TLMessagesMessageEditData{
		Data2: m.Data2,
	}
}

// messages.messageEditData#26b5dde6 flags:# caption:flags.0?true = messages.MessageEditData;
func (m *TLMessagesMessageEditData) To_Messages_MessageEditData() *Messages_MessageEditData {
	return &Messages_MessageEditData{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessagesMessageEditData) SetCaption(v bool) { m.Data2.Caption = v }
func (m *TLMessagesMessageEditData) GetCaption() bool  { return m.Data2.Caption }

func (m *TLMessagesMessageEditData) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_messageEditData))

	// flags
	var flags uint32 = 0
	if m.GetCaption() == true {
		flags |= 1 << 1
	}
	x.UInt(flags)

	return x.buf
}

func (m *TLMessagesMessageEditData) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m.Data2.Caption = true
	}

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Auth_SentCodeType <--
//  + TL_AuthSentCodeTypeApp
//  + TL_AuthSentCodeTypeSms
//  + TL_AuthSentCodeTypeCall
//  + TL_AuthSentCodeTypeFlashCall
//

func (m *Auth_SentCodeType) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_auth_sentCodeTypeApp:
		t := m.To_AuthSentCodeTypeApp()
		return t.Encode()
	case TLConstructor_CRC32_auth_sentCodeTypeSms:
		t := m.To_AuthSentCodeTypeSms()
		return t.Encode()
	case TLConstructor_CRC32_auth_sentCodeTypeCall:
		t := m.To_AuthSentCodeTypeCall()
		return t.Encode()
	case TLConstructor_CRC32_auth_sentCodeTypeFlashCall:
		t := m.To_AuthSentCodeTypeFlashCall()
		return t.Encode()

	default:
		return []byte{}
	}
}

// auth.sentCodeTypeApp#3dbb5986 length:int = auth.SentCodeType;
func (m *Auth_SentCodeType) To_AuthSentCodeTypeApp() *TLAuthSentCodeTypeApp {
	return &TLAuthSentCodeTypeApp{
		Data2: m.Data2,
	}
}

// auth.sentCodeTypeSms#c000bba2 length:int = auth.SentCodeType;
func (m *Auth_SentCodeType) To_AuthSentCodeTypeSms() *TLAuthSentCodeTypeSms {
	return &TLAuthSentCodeTypeSms{
		Data2: m.Data2,
	}
}

// auth.sentCodeTypeCall#5353e5a7 length:int = auth.SentCodeType;
func (m *Auth_SentCodeType) To_AuthSentCodeTypeCall() *TLAuthSentCodeTypeCall {
	return &TLAuthSentCodeTypeCall{
		Data2: m.Data2,
	}
}

// auth.sentCodeTypeFlashCall#ab03c6d9 pattern:string = auth.SentCodeType;
func (m *Auth_SentCodeType) To_AuthSentCodeTypeFlashCall() *TLAuthSentCodeTypeFlashCall {
	return &TLAuthSentCodeTypeFlashCall{
		Data2: m.Data2,
	}
}

// auth.sentCodeTypeApp#3dbb5986 length:int = auth.SentCodeType;
func (m *TLAuthSentCodeTypeApp) To_Auth_SentCodeType() *Auth_SentCodeType {
	return &Auth_SentCodeType{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLAuthSentCodeTypeApp) SetLength(v int32) { m.Data2.Length = v }
func (m *TLAuthSentCodeTypeApp) GetLength() int32  { return m.Data2.Length }

func (m *TLAuthSentCodeTypeApp) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_auth_sentCodeTypeApp))

	x.Int(m.Data2.Length)

	return x.buf
}

func (m *TLAuthSentCodeTypeApp) Decode(dbuf *DecodeBuf) error {
	m.Data2.Length = x.Int()

	return dbuf.err
}

// auth.sentCodeTypeSms#c000bba2 length:int = auth.SentCodeType;
func (m *TLAuthSentCodeTypeSms) To_Auth_SentCodeType() *Auth_SentCodeType {
	return &Auth_SentCodeType{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLAuthSentCodeTypeSms) SetLength(v int32) { m.Data2.Length = v }
func (m *TLAuthSentCodeTypeSms) GetLength() int32  { return m.Data2.Length }

func (m *TLAuthSentCodeTypeSms) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_auth_sentCodeTypeSms))

	x.Int(m.Data2.Length)

	return x.buf
}

func (m *TLAuthSentCodeTypeSms) Decode(dbuf *DecodeBuf) error {
	m.Data2.Length = x.Int()

	return dbuf.err
}

// auth.sentCodeTypeCall#5353e5a7 length:int = auth.SentCodeType;
func (m *TLAuthSentCodeTypeCall) To_Auth_SentCodeType() *Auth_SentCodeType {
	return &Auth_SentCodeType{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLAuthSentCodeTypeCall) SetLength(v int32) { m.Data2.Length = v }
func (m *TLAuthSentCodeTypeCall) GetLength() int32  { return m.Data2.Length }

func (m *TLAuthSentCodeTypeCall) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_auth_sentCodeTypeCall))

	x.Int(m.Data2.Length)

	return x.buf
}

func (m *TLAuthSentCodeTypeCall) Decode(dbuf *DecodeBuf) error {
	m.Data2.Length = x.Int()

	return dbuf.err
}

// auth.sentCodeTypeFlashCall#ab03c6d9 pattern:string = auth.SentCodeType;
func (m *TLAuthSentCodeTypeFlashCall) To_Auth_SentCodeType() *Auth_SentCodeType {
	return &Auth_SentCodeType{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLAuthSentCodeTypeFlashCall) SetPattern(v string) { m.Data2.Pattern = v }
func (m *TLAuthSentCodeTypeFlashCall) GetPattern() string  { return m.Data2.Pattern }

func (m *TLAuthSentCodeTypeFlashCall) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_auth_sentCodeTypeFlashCall))

	x.StringBytes(m.Data2.Pattern)

	return x.buf
}

func (m *TLAuthSentCodeTypeFlashCall) Decode(dbuf *DecodeBuf) error {
	m.Data2.Pattern = x.StringBytes()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// InputBotInlineMessageID <--
//  + TL_InputBotInlineMessageID
//

func (m *InputBotInlineMessageID) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_inputBotInlineMessageID:
		t := m.To_InputBotInlineMessageID()
		return t.Encode()

	default:
		return []byte{}
	}
}

// inputBotInlineMessageID#890c3d89 dc_id:int id:long access_hash:long = InputBotInlineMessageID;
func (m *InputBotInlineMessageID) To_InputBotInlineMessageID() *TLInputBotInlineMessageID {
	return &TLInputBotInlineMessageID{
		Data2: m.Data2,
	}
}

// inputBotInlineMessageID#890c3d89 dc_id:int id:long access_hash:long = InputBotInlineMessageID;
func (m *TLInputBotInlineMessageID) To_InputBotInlineMessageID() *InputBotInlineMessageID {
	return &InputBotInlineMessageID{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputBotInlineMessageID) SetDcId(v int32) { m.Data2.DcId = v }
func (m *TLInputBotInlineMessageID) GetDcId() int32  { return m.Data2.DcId }

func (m *TLInputBotInlineMessageID) SetId(v int64) { m.Data2.Id = v }
func (m *TLInputBotInlineMessageID) GetId() int64  { return m.Data2.Id }

func (m *TLInputBotInlineMessageID) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLInputBotInlineMessageID) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLInputBotInlineMessageID) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputBotInlineMessageID))

	x.Int(m.Data2.DcId)
	x.Long(m.Data2.Id)
	x.Long(m.Data2.AccessHash)

	return x.buf
}

func (m *TLInputBotInlineMessageID) Decode(dbuf *DecodeBuf) error {
	m.Data2.DcId = x.Int()
	m.Data2.Id = x.Long()
	m.Data2.AccessHash = x.Long()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Payments_PaymentReceipt <--
//  + TL_PaymentsPaymentReceipt
//

func (m *Payments_PaymentReceipt) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_payments_paymentReceipt:
		t := m.To_PaymentsPaymentReceipt()
		return t.Encode()

	default:
		return []byte{}
	}
}

// payments.paymentReceipt#500911e1 flags:# date:int bot_id:int invoice:Invoice provider_id:int info:flags.0?PaymentRequestedInfo shipping:flags.1?ShippingOption currency:string total_amount:long credentials_title:string users:Vector<User> = payments.PaymentReceipt;
func (m *Payments_PaymentReceipt) To_PaymentsPaymentReceipt() *TLPaymentsPaymentReceipt {
	return &TLPaymentsPaymentReceipt{
		Data2: m.Data2,
	}
}

// payments.paymentReceipt#500911e1 flags:# date:int bot_id:int invoice:Invoice provider_id:int info:flags.0?PaymentRequestedInfo shipping:flags.1?ShippingOption currency:string total_amount:long credentials_title:string users:Vector<User> = payments.PaymentReceipt;
func (m *TLPaymentsPaymentReceipt) To_Payments_PaymentReceipt() *Payments_PaymentReceipt {
	return &Payments_PaymentReceipt{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPaymentsPaymentReceipt) SetDate(v int32) { m.Data2.Date = v }
func (m *TLPaymentsPaymentReceipt) GetDate() int32  { return m.Data2.Date }

func (m *TLPaymentsPaymentReceipt) SetBotId(v int32) { m.Data2.BotId = v }
func (m *TLPaymentsPaymentReceipt) GetBotId() int32  { return m.Data2.BotId }

func (m *TLPaymentsPaymentReceipt) SetInvoice(v *Invoice) { m.Data2.Invoice = v }
func (m *TLPaymentsPaymentReceipt) GetInvoice() *Invoice  { return m.Data2.Invoice }

func (m *TLPaymentsPaymentReceipt) SetProviderId(v int32) { m.Data2.ProviderId = v }
func (m *TLPaymentsPaymentReceipt) GetProviderId() int32  { return m.Data2.ProviderId }

func (m *TLPaymentsPaymentReceipt) SetInfo(v *PaymentRequestedInfo) { m.Data2.Info = v }
func (m *TLPaymentsPaymentReceipt) GetInfo() *PaymentRequestedInfo  { return m.Data2.Info }

func (m *TLPaymentsPaymentReceipt) SetShipping(v *ShippingOption) { m.Data2.Shipping = v }
func (m *TLPaymentsPaymentReceipt) GetShipping() *ShippingOption  { return m.Data2.Shipping }

func (m *TLPaymentsPaymentReceipt) SetCurrency(v string) { m.Data2.Currency = v }
func (m *TLPaymentsPaymentReceipt) GetCurrency() string  { return m.Data2.Currency }

func (m *TLPaymentsPaymentReceipt) SetTotalAmount(v int64) { m.Data2.TotalAmount = v }
func (m *TLPaymentsPaymentReceipt) GetTotalAmount() int64  { return m.Data2.TotalAmount }

func (m *TLPaymentsPaymentReceipt) SetCredentialsTitle(v string) { m.Data2.CredentialsTitle = v }
func (m *TLPaymentsPaymentReceipt) GetCredentialsTitle() string  { return m.Data2.CredentialsTitle }

func (m *TLPaymentsPaymentReceipt) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLPaymentsPaymentReceipt) GetUsers() []*User  { return m.Data2.Users }

func (m *TLPaymentsPaymentReceipt) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_payments_paymentReceipt))

	// flags
	var flags uint32 = 0
	if m.GetInfo() != nil {
		flags |= 1 << 5
	}
	if m.GetShipping() != nil {
		flags |= 1 << 6
	}
	x.UInt(flags)

	x.Int(m.Data2.Date)
	x.Int(m.Data2.BotId)
	x.Bytes(m.Data2Invoice.Encode())
	x.Int(m.Data2.ProviderId)
	if m.GetInfo() != 0 {
		x.Bytes(m.Data2Info.Encode())
	}
	if m.GetShipping() != 0 {
		x.Bytes(m.Data2Shipping.Encode())
	}
	x.StringBytes(m.Data2.Currency)
	x.Long(m.Data2.TotalAmount)
	x.StringBytes(m.Data2.CredentialsTitle)

	return x.buf
}

func (m *TLPaymentsPaymentReceipt) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	m.Data2.Date = x.Int()
	m.Data2.BotId = x.Int()
	m4 := &Invoice{}
	m4.Decode(dbuf)
	m.SetInvoice(m4)
	m.Data2.ProviderId = x.Int()
	if (flags & (1 << 6)) != 0 {
		m6 := &Info{}
		m6.Decode(dbuf)
		m.SetInfo(m6)
	}
	if (flags & (1 << 7)) != 0 {
		m7 := &Shipping{}
		m7.Decode(dbuf)
		m.SetShipping(m7)
	}
	m.Data2.Currency = x.StringBytes()
	m.Data2.TotalAmount = x.Long()
	m.Data2.CredentialsTitle = x.StringBytes()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// CdnConfig <--
//  + TL_CdnConfig
//

func (m *CdnConfig) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_cdnConfig:
		t := m.To_CdnConfig()
		return t.Encode()

	default:
		return []byte{}
	}
}

// cdnConfig#5725e40a public_keys:Vector<CdnPublicKey> = CdnConfig;
func (m *CdnConfig) To_CdnConfig() *TLCdnConfig {
	return &TLCdnConfig{
		Data2: m.Data2,
	}
}

// cdnConfig#5725e40a public_keys:Vector<CdnPublicKey> = CdnConfig;
func (m *TLCdnConfig) To_CdnConfig() *CdnConfig {
	return &CdnConfig{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLCdnConfig) SetPublicKeys(v []*CdnPublicKey) { m.Data2.PublicKeys = v }
func (m *TLCdnConfig) GetPublicKeys() []*CdnPublicKey  { return m.Data2.PublicKeys }

func (m *TLCdnConfig) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_cdnConfig))

	return x.buf
}

func (m *TLCdnConfig) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// UserFull <--
//  + TL_UserFull
//

func (m *UserFull) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_userFull:
		t := m.To_UserFull()
		return t.Encode()

	default:
		return []byte{}
	}
}

// userFull#f220f3f flags:# blocked:flags.0?true phone_calls_available:flags.4?true phone_calls_private:flags.5?true user:User about:flags.1?string link:contacts.Link profile_photo:flags.2?Photo notify_settings:PeerNotifySettings bot_info:flags.3?BotInfo common_chats_count:int = UserFull;
func (m *UserFull) To_UserFull() *TLUserFull {
	return &TLUserFull{
		Data2: m.Data2,
	}
}

// userFull#f220f3f flags:# blocked:flags.0?true phone_calls_available:flags.4?true phone_calls_private:flags.5?true user:User about:flags.1?string link:contacts.Link profile_photo:flags.2?Photo notify_settings:PeerNotifySettings bot_info:flags.3?BotInfo common_chats_count:int = UserFull;
func (m *TLUserFull) To_UserFull() *UserFull {
	return &UserFull{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUserFull) SetBlocked(v bool) { m.Data2.Blocked = v }
func (m *TLUserFull) GetBlocked() bool  { return m.Data2.Blocked }

func (m *TLUserFull) SetPhoneCallsAvailable(v bool) { m.Data2.PhoneCallsAvailable = v }
func (m *TLUserFull) GetPhoneCallsAvailable() bool  { return m.Data2.PhoneCallsAvailable }

func (m *TLUserFull) SetPhoneCallsPrivate(v bool) { m.Data2.PhoneCallsPrivate = v }
func (m *TLUserFull) GetPhoneCallsPrivate() bool  { return m.Data2.PhoneCallsPrivate }

func (m *TLUserFull) SetUser(v *User) { m.Data2.User = v }
func (m *TLUserFull) GetUser() *User  { return m.Data2.User }

func (m *TLUserFull) SetAbout(v string) { m.Data2.About = v }
func (m *TLUserFull) GetAbout() string  { return m.Data2.About }

func (m *TLUserFull) SetLink(v *Contacts_Link) { m.Data2.Link = v }
func (m *TLUserFull) GetLink() *Contacts_Link  { return m.Data2.Link }

func (m *TLUserFull) SetProfilePhoto(v *Photo) { m.Data2.ProfilePhoto = v }
func (m *TLUserFull) GetProfilePhoto() *Photo  { return m.Data2.ProfilePhoto }

func (m *TLUserFull) SetNotifySettings(v *PeerNotifySettings) { m.Data2.NotifySettings = v }
func (m *TLUserFull) GetNotifySettings() *PeerNotifySettings  { return m.Data2.NotifySettings }

func (m *TLUserFull) SetBotInfo(v *BotInfo) { m.Data2.BotInfo = v }
func (m *TLUserFull) GetBotInfo() *BotInfo  { return m.Data2.BotInfo }

func (m *TLUserFull) SetCommonChatsCount(v int32) { m.Data2.CommonChatsCount = v }
func (m *TLUserFull) GetCommonChatsCount() int32  { return m.Data2.CommonChatsCount }

func (m *TLUserFull) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_userFull))

	// flags
	var flags uint32 = 0
	if m.GetBlocked() == true {
		flags |= 1 << 1
	}
	if m.GetPhoneCallsAvailable() == true {
		flags |= 1 << 2
	}
	if m.GetPhoneCallsPrivate() == true {
		flags |= 1 << 3
	}
	if m.GetAbout() != "" {
		flags |= 1 << 5
	}
	if m.GetProfilePhoto() != nil {
		flags |= 1 << 7
	}
	if m.GetBotInfo() != nil {
		flags |= 1 << 9
	}
	x.UInt(flags)

	x.Bytes(m.Data2User.Encode())
	if m.GetAbout() != 0 {
		x.StringBytes(m.Data2.About)
	}
	x.Bytes(m.Data2Link.Encode())
	if m.GetProfilePhoto() != 0 {
		x.Bytes(m.Data2ProfilePhoto.Encode())
	}
	x.Bytes(m.Data2NotifySettings.Encode())
	if m.GetBotInfo() != 0 {
		x.Bytes(m.Data2BotInfo.Encode())
	}
	x.Int(m.Data2.CommonChatsCount)

	return x.buf
}

func (m *TLUserFull) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m.Data2.Blocked = true
	}
	if (flags & (1 << 3)) != 0 {
		m.Data2.PhoneCallsAvailable = true
	}
	if (flags & (1 << 4)) != 0 {
		m.Data2.PhoneCallsPrivate = true
	}
	m5 := &User{}
	m5.Decode(dbuf)
	m.SetUser(m5)
	if (flags & (1 << 6)) != 0 {
		m.Data2.About = x.StringBytes()
	}
	m7 := &Link{}
	m7.Decode(dbuf)
	m.SetLink(m7)
	if (flags & (1 << 8)) != 0 {
		m8 := &ProfilePhoto{}
		m8.Decode(dbuf)
		m.SetProfilePhoto(m8)
	}
	m9 := &NotifySettings{}
	m9.Decode(dbuf)
	m.SetNotifySettings(m9)
	if (flags & (1 << 10)) != 0 {
		m10 := &BotInfo{}
		m10.Decode(dbuf)
		m.SetBotInfo(m10)
	}
	m.Data2.CommonChatsCount = x.Int()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Messages_SentEncryptedMessage <--
//  + TL_MessagesSentEncryptedMessage
//  + TL_MessagesSentEncryptedFile
//

func (m *Messages_SentEncryptedMessage) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_messages_sentEncryptedMessage:
		t := m.To_MessagesSentEncryptedMessage()
		return t.Encode()
	case TLConstructor_CRC32_messages_sentEncryptedFile:
		t := m.To_MessagesSentEncryptedFile()
		return t.Encode()

	default:
		return []byte{}
	}
}

// messages.sentEncryptedMessage#560f8935 date:int = messages.SentEncryptedMessage;
func (m *Messages_SentEncryptedMessage) To_MessagesSentEncryptedMessage() *TLMessagesSentEncryptedMessage {
	return &TLMessagesSentEncryptedMessage{
		Data2: m.Data2,
	}
}

// messages.sentEncryptedFile#9493ff32 date:int file:EncryptedFile = messages.SentEncryptedMessage;
func (m *Messages_SentEncryptedMessage) To_MessagesSentEncryptedFile() *TLMessagesSentEncryptedFile {
	return &TLMessagesSentEncryptedFile{
		Data2: m.Data2,
	}
}

// messages.sentEncryptedMessage#560f8935 date:int = messages.SentEncryptedMessage;
func (m *TLMessagesSentEncryptedMessage) To_Messages_SentEncryptedMessage() *Messages_SentEncryptedMessage {
	return &Messages_SentEncryptedMessage{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessagesSentEncryptedMessage) SetDate(v int32) { m.Data2.Date = v }
func (m *TLMessagesSentEncryptedMessage) GetDate() int32  { return m.Data2.Date }

func (m *TLMessagesSentEncryptedMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_sentEncryptedMessage))

	x.Int(m.Data2.Date)

	return x.buf
}

func (m *TLMessagesSentEncryptedMessage) Decode(dbuf *DecodeBuf) error {
	m.Data2.Date = x.Int()

	return dbuf.err
}

// messages.sentEncryptedFile#9493ff32 date:int file:EncryptedFile = messages.SentEncryptedMessage;
func (m *TLMessagesSentEncryptedFile) To_Messages_SentEncryptedMessage() *Messages_SentEncryptedMessage {
	return &Messages_SentEncryptedMessage{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessagesSentEncryptedFile) SetDate(v int32) { m.Data2.Date = v }
func (m *TLMessagesSentEncryptedFile) GetDate() int32  { return m.Data2.Date }

func (m *TLMessagesSentEncryptedFile) SetFile(v *EncryptedFile) { m.Data2.File = v }
func (m *TLMessagesSentEncryptedFile) GetFile() *EncryptedFile  { return m.Data2.File }

func (m *TLMessagesSentEncryptedFile) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_sentEncryptedFile))

	x.Int(m.Data2.Date)
	x.Bytes(m.Data2File.Encode())

	return x.buf
}

func (m *TLMessagesSentEncryptedFile) Decode(dbuf *DecodeBuf) error {
	m.Data2.Date = x.Int()
	m2 := &File{}
	m2.Decode(dbuf)
	m.SetFile(m2)

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Messages_Stickers <--
//  + TL_MessagesStickersNotModified
//  + TL_MessagesStickers
//

func (m *Messages_Stickers) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_messages_stickersNotModified:
		t := m.To_MessagesStickersNotModified()
		return t.Encode()
	case TLConstructor_CRC32_messages_stickers:
		t := m.To_MessagesStickers()
		return t.Encode()

	default:
		return []byte{}
	}
}

// messages.stickersNotModified#f1749a22 = messages.Stickers;
func (m *Messages_Stickers) To_MessagesStickersNotModified() *TLMessagesStickersNotModified {
	return &TLMessagesStickersNotModified{
		Data2: m.Data2,
	}
}

// messages.stickers#8a8ecd32 hash:string stickers:Vector<Document> = messages.Stickers;
func (m *Messages_Stickers) To_MessagesStickers() *TLMessagesStickers {
	return &TLMessagesStickers{
		Data2: m.Data2,
	}
}

// messages.stickersNotModified#f1749a22 = messages.Stickers;
func (m *TLMessagesStickersNotModified) To_Messages_Stickers() *Messages_Stickers {
	return &Messages_Stickers{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessagesStickersNotModified) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_stickersNotModified))

	return x.buf
}

func (m *TLMessagesStickersNotModified) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// messages.stickers#8a8ecd32 hash:string stickers:Vector<Document> = messages.Stickers;
func (m *TLMessagesStickers) To_Messages_Stickers() *Messages_Stickers {
	return &Messages_Stickers{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessagesStickers) SetHash(v string) { m.Data2.Hash = v }
func (m *TLMessagesStickers) GetHash() string  { return m.Data2.Hash }

func (m *TLMessagesStickers) SetStickers(v []*Document) { m.Data2.Stickers = v }
func (m *TLMessagesStickers) GetStickers() []*Document  { return m.Data2.Stickers }

func (m *TLMessagesStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_stickers))

	x.StringBytes(m.Data2.Hash)

	return x.buf
}

func (m *TLMessagesStickers) Decode(dbuf *DecodeBuf) error {
	m.Data2.Hash = x.StringBytes()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// DisabledFeature <--
//  + TL_DisabledFeature
//

func (m *DisabledFeature) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_disabledFeature:
		t := m.To_DisabledFeature()
		return t.Encode()

	default:
		return []byte{}
	}
}

// disabledFeature#ae636f24 feature:string description:string = DisabledFeature;
func (m *DisabledFeature) To_DisabledFeature() *TLDisabledFeature {
	return &TLDisabledFeature{
		Data2: m.Data2,
	}
}

// disabledFeature#ae636f24 feature:string description:string = DisabledFeature;
func (m *TLDisabledFeature) To_DisabledFeature() *DisabledFeature {
	return &DisabledFeature{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLDisabledFeature) SetFeature(v string) { m.Data2.Feature = v }
func (m *TLDisabledFeature) GetFeature() string  { return m.Data2.Feature }

func (m *TLDisabledFeature) SetDescription(v string) { m.Data2.Description = v }
func (m *TLDisabledFeature) GetDescription() string  { return m.Data2.Description }

func (m *TLDisabledFeature) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_disabledFeature))

	x.StringBytes(m.Data2.Feature)
	x.StringBytes(m.Data2.Description)

	return x.buf
}

func (m *TLDisabledFeature) Decode(dbuf *DecodeBuf) error {
	m.Data2.Feature = x.StringBytes()
	m.Data2.Description = x.StringBytes()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// ChannelParticipantsFilter <--
//  + TL_ChannelParticipantsRecent
//  + TL_ChannelParticipantsAdmins
//  + TL_ChannelParticipantsKicked
//  + TL_ChannelParticipantsBots
//  + TL_ChannelParticipantsBanned
//  + TL_ChannelParticipantsSearch
//

func (m *ChannelParticipantsFilter) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_channelParticipantsRecent:
		t := m.To_ChannelParticipantsRecent()
		return t.Encode()
	case TLConstructor_CRC32_channelParticipantsAdmins:
		t := m.To_ChannelParticipantsAdmins()
		return t.Encode()
	case TLConstructor_CRC32_channelParticipantsKicked:
		t := m.To_ChannelParticipantsKicked()
		return t.Encode()
	case TLConstructor_CRC32_channelParticipantsBots:
		t := m.To_ChannelParticipantsBots()
		return t.Encode()
	case TLConstructor_CRC32_channelParticipantsBanned:
		t := m.To_ChannelParticipantsBanned()
		return t.Encode()
	case TLConstructor_CRC32_channelParticipantsSearch:
		t := m.To_ChannelParticipantsSearch()
		return t.Encode()

	default:
		return []byte{}
	}
}

// channelParticipantsRecent#de3f3c79 = ChannelParticipantsFilter;
func (m *ChannelParticipantsFilter) To_ChannelParticipantsRecent() *TLChannelParticipantsRecent {
	return &TLChannelParticipantsRecent{
		Data2: m.Data2,
	}
}

// channelParticipantsAdmins#b4608969 = ChannelParticipantsFilter;
func (m *ChannelParticipantsFilter) To_ChannelParticipantsAdmins() *TLChannelParticipantsAdmins {
	return &TLChannelParticipantsAdmins{
		Data2: m.Data2,
	}
}

// channelParticipantsKicked#a3b54985 q:string = ChannelParticipantsFilter;
func (m *ChannelParticipantsFilter) To_ChannelParticipantsKicked() *TLChannelParticipantsKicked {
	return &TLChannelParticipantsKicked{
		Data2: m.Data2,
	}
}

// channelParticipantsBots#b0d1865b = ChannelParticipantsFilter;
func (m *ChannelParticipantsFilter) To_ChannelParticipantsBots() *TLChannelParticipantsBots {
	return &TLChannelParticipantsBots{
		Data2: m.Data2,
	}
}

// channelParticipantsBanned#1427a5e1 q:string = ChannelParticipantsFilter;
func (m *ChannelParticipantsFilter) To_ChannelParticipantsBanned() *TLChannelParticipantsBanned {
	return &TLChannelParticipantsBanned{
		Data2: m.Data2,
	}
}

// channelParticipantsSearch#656ac4b q:string = ChannelParticipantsFilter;
func (m *ChannelParticipantsFilter) To_ChannelParticipantsSearch() *TLChannelParticipantsSearch {
	return &TLChannelParticipantsSearch{
		Data2: m.Data2,
	}
}

// channelParticipantsRecent#de3f3c79 = ChannelParticipantsFilter;
func (m *TLChannelParticipantsRecent) To_ChannelParticipantsFilter() *ChannelParticipantsFilter {
	return &ChannelParticipantsFilter{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLChannelParticipantsRecent) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelParticipantsRecent))

	return x.buf
}

func (m *TLChannelParticipantsRecent) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// channelParticipantsAdmins#b4608969 = ChannelParticipantsFilter;
func (m *TLChannelParticipantsAdmins) To_ChannelParticipantsFilter() *ChannelParticipantsFilter {
	return &ChannelParticipantsFilter{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLChannelParticipantsAdmins) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelParticipantsAdmins))

	return x.buf
}

func (m *TLChannelParticipantsAdmins) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// channelParticipantsKicked#a3b54985 q:string = ChannelParticipantsFilter;
func (m *TLChannelParticipantsKicked) To_ChannelParticipantsFilter() *ChannelParticipantsFilter {
	return &ChannelParticipantsFilter{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLChannelParticipantsKicked) SetQ(v string) { m.Data2.Q = v }
func (m *TLChannelParticipantsKicked) GetQ() string  { return m.Data2.Q }

func (m *TLChannelParticipantsKicked) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelParticipantsKicked))

	x.StringBytes(m.Data2.Q)

	return x.buf
}

func (m *TLChannelParticipantsKicked) Decode(dbuf *DecodeBuf) error {
	m.Data2.Q = x.StringBytes()

	return dbuf.err
}

// channelParticipantsBots#b0d1865b = ChannelParticipantsFilter;
func (m *TLChannelParticipantsBots) To_ChannelParticipantsFilter() *ChannelParticipantsFilter {
	return &ChannelParticipantsFilter{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLChannelParticipantsBots) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelParticipantsBots))

	return x.buf
}

func (m *TLChannelParticipantsBots) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// channelParticipantsBanned#1427a5e1 q:string = ChannelParticipantsFilter;
func (m *TLChannelParticipantsBanned) To_ChannelParticipantsFilter() *ChannelParticipantsFilter {
	return &ChannelParticipantsFilter{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLChannelParticipantsBanned) SetQ(v string) { m.Data2.Q = v }
func (m *TLChannelParticipantsBanned) GetQ() string  { return m.Data2.Q }

func (m *TLChannelParticipantsBanned) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelParticipantsBanned))

	x.StringBytes(m.Data2.Q)

	return x.buf
}

func (m *TLChannelParticipantsBanned) Decode(dbuf *DecodeBuf) error {
	m.Data2.Q = x.StringBytes()

	return dbuf.err
}

// channelParticipantsSearch#656ac4b q:string = ChannelParticipantsFilter;
func (m *TLChannelParticipantsSearch) To_ChannelParticipantsFilter() *ChannelParticipantsFilter {
	return &ChannelParticipantsFilter{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLChannelParticipantsSearch) SetQ(v string) { m.Data2.Q = v }
func (m *TLChannelParticipantsSearch) GetQ() string  { return m.Data2.Q }

func (m *TLChannelParticipantsSearch) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelParticipantsSearch))

	x.StringBytes(m.Data2.Q)

	return x.buf
}

func (m *TLChannelParticipantsSearch) Decode(dbuf *DecodeBuf) error {
	m.Data2.Q = x.StringBytes()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// UserProfilePhoto <--
//  + TL_UserProfilePhotoEmpty
//  + TL_UserProfilePhoto
//

func (m *UserProfilePhoto) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_userProfilePhotoEmpty:
		t := m.To_UserProfilePhotoEmpty()
		return t.Encode()
	case TLConstructor_CRC32_userProfilePhoto:
		t := m.To_UserProfilePhoto()
		return t.Encode()

	default:
		return []byte{}
	}
}

// userProfilePhotoEmpty#4f11bae1 = UserProfilePhoto;
func (m *UserProfilePhoto) To_UserProfilePhotoEmpty() *TLUserProfilePhotoEmpty {
	return &TLUserProfilePhotoEmpty{
		Data2: m.Data2,
	}
}

// userProfilePhoto#d559d8c8 photo_id:long photo_small:FileLocation photo_big:FileLocation = UserProfilePhoto;
func (m *UserProfilePhoto) To_UserProfilePhoto() *TLUserProfilePhoto {
	return &TLUserProfilePhoto{
		Data2: m.Data2,
	}
}

// userProfilePhotoEmpty#4f11bae1 = UserProfilePhoto;
func (m *TLUserProfilePhotoEmpty) To_UserProfilePhoto() *UserProfilePhoto {
	return &UserProfilePhoto{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUserProfilePhotoEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_userProfilePhotoEmpty))

	return x.buf
}

func (m *TLUserProfilePhotoEmpty) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// userProfilePhoto#d559d8c8 photo_id:long photo_small:FileLocation photo_big:FileLocation = UserProfilePhoto;
func (m *TLUserProfilePhoto) To_UserProfilePhoto() *UserProfilePhoto {
	return &UserProfilePhoto{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUserProfilePhoto) SetPhotoId(v int64) { m.Data2.PhotoId = v }
func (m *TLUserProfilePhoto) GetPhotoId() int64  { return m.Data2.PhotoId }

func (m *TLUserProfilePhoto) SetPhotoSmall(v *FileLocation) { m.Data2.PhotoSmall = v }
func (m *TLUserProfilePhoto) GetPhotoSmall() *FileLocation  { return m.Data2.PhotoSmall }

func (m *TLUserProfilePhoto) SetPhotoBig(v *FileLocation) { m.Data2.PhotoBig = v }
func (m *TLUserProfilePhoto) GetPhotoBig() *FileLocation  { return m.Data2.PhotoBig }

func (m *TLUserProfilePhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_userProfilePhoto))

	x.Long(m.Data2.PhotoId)
	x.Bytes(m.Data2PhotoSmall.Encode())
	x.Bytes(m.Data2PhotoBig.Encode())

	return x.buf
}

func (m *TLUserProfilePhoto) Decode(dbuf *DecodeBuf) error {
	m.Data2.PhotoId = x.Long()
	m2 := &PhotoSmall{}
	m2.Decode(dbuf)
	m.SetPhotoSmall(m2)
	m3 := &PhotoBig{}
	m3.Decode(dbuf)
	m.SetPhotoBig(m3)

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Contacts_Link <--
//  + TL_ContactsLink
//

func (m *Contacts_Link) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_contacts_link:
		t := m.To_ContactsLink()
		return t.Encode()

	default:
		return []byte{}
	}
}

// contacts.link#3ace484c my_link:ContactLink foreign_link:ContactLink user:User = contacts.Link;
func (m *Contacts_Link) To_ContactsLink() *TLContactsLink {
	return &TLContactsLink{
		Data2: m.Data2,
	}
}

// contacts.link#3ace484c my_link:ContactLink foreign_link:ContactLink user:User = contacts.Link;
func (m *TLContactsLink) To_Contacts_Link() *Contacts_Link {
	return &Contacts_Link{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLContactsLink) SetMyLink(v *ContactLink) { m.Data2.MyLink = v }
func (m *TLContactsLink) GetMyLink() *ContactLink  { return m.Data2.MyLink }

func (m *TLContactsLink) SetForeignLink(v *ContactLink) { m.Data2.ForeignLink = v }
func (m *TLContactsLink) GetForeignLink() *ContactLink  { return m.Data2.ForeignLink }

func (m *TLContactsLink) SetUser(v *User) { m.Data2.User = v }
func (m *TLContactsLink) GetUser() *User  { return m.Data2.User }

func (m *TLContactsLink) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_contacts_link))

	x.Bytes(m.Data2MyLink.Encode())
	x.Bytes(m.Data2ForeignLink.Encode())
	x.Bytes(m.Data2User.Encode())

	return x.buf
}

func (m *TLContactsLink) Decode(dbuf *DecodeBuf) error {
	m1 := &MyLink{}
	m1.Decode(dbuf)
	m.SetMyLink(m1)
	m2 := &ForeignLink{}
	m2.Decode(dbuf)
	m.SetForeignLink(m2)
	m3 := &User{}
	m3.Decode(dbuf)
	m.SetUser(m3)

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Contacts_Contacts <--
//  + TL_ContactsContactsNotModified
//  + TL_ContactsContacts
//

func (m *Contacts_Contacts) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_contacts_contactsNotModified:
		t := m.To_ContactsContactsNotModified()
		return t.Encode()
	case TLConstructor_CRC32_contacts_contacts:
		t := m.To_ContactsContacts()
		return t.Encode()

	default:
		return []byte{}
	}
}

// contacts.contactsNotModified#b74ba9d2 = contacts.Contacts;
func (m *Contacts_Contacts) To_ContactsContactsNotModified() *TLContactsContactsNotModified {
	return &TLContactsContactsNotModified{
		Data2: m.Data2,
	}
}

// contacts.contacts#eae87e42 contacts:Vector<Contact> saved_count:int users:Vector<User> = contacts.Contacts;
func (m *Contacts_Contacts) To_ContactsContacts() *TLContactsContacts {
	return &TLContactsContacts{
		Data2: m.Data2,
	}
}

// contacts.contactsNotModified#b74ba9d2 = contacts.Contacts;
func (m *TLContactsContactsNotModified) To_Contacts_Contacts() *Contacts_Contacts {
	return &Contacts_Contacts{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLContactsContactsNotModified) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_contacts_contactsNotModified))

	return x.buf
}

func (m *TLContactsContactsNotModified) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// contacts.contacts#eae87e42 contacts:Vector<Contact> saved_count:int users:Vector<User> = contacts.Contacts;
func (m *TLContactsContacts) To_Contacts_Contacts() *Contacts_Contacts {
	return &Contacts_Contacts{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLContactsContacts) SetContacts(v []*Contact) { m.Data2.Contacts = v }
func (m *TLContactsContacts) GetContacts() []*Contact  { return m.Data2.Contacts }

func (m *TLContactsContacts) SetSavedCount(v int32) { m.Data2.SavedCount = v }
func (m *TLContactsContacts) GetSavedCount() int32  { return m.Data2.SavedCount }

func (m *TLContactsContacts) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLContactsContacts) GetUsers() []*User  { return m.Data2.Users }

func (m *TLContactsContacts) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_contacts_contacts))

	x.Int(m.Data2.SavedCount)

	return x.buf
}

func (m *TLContactsContacts) Decode(dbuf *DecodeBuf) error {

	m.Data2.SavedCount = x.Int()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// NotifyPeer <--
//  + TL_NotifyPeer
//  + TL_NotifyUsers
//  + TL_NotifyChats
//  + TL_NotifyAll
//

func (m *NotifyPeer) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_notifyPeer:
		t := m.To_NotifyPeer()
		return t.Encode()
	case TLConstructor_CRC32_notifyUsers:
		t := m.To_NotifyUsers()
		return t.Encode()
	case TLConstructor_CRC32_notifyChats:
		t := m.To_NotifyChats()
		return t.Encode()
	case TLConstructor_CRC32_notifyAll:
		t := m.To_NotifyAll()
		return t.Encode()

	default:
		return []byte{}
	}
}

// notifyPeer#9fd40bd8 peer:Peer = NotifyPeer;
func (m *NotifyPeer) To_NotifyPeer() *TLNotifyPeer {
	return &TLNotifyPeer{
		Data2: m.Data2,
	}
}

// notifyUsers#b4c83b4c = NotifyPeer;
func (m *NotifyPeer) To_NotifyUsers() *TLNotifyUsers {
	return &TLNotifyUsers{
		Data2: m.Data2,
	}
}

// notifyChats#c007cec3 = NotifyPeer;
func (m *NotifyPeer) To_NotifyChats() *TLNotifyChats {
	return &TLNotifyChats{
		Data2: m.Data2,
	}
}

// notifyAll#74d07c60 = NotifyPeer;
func (m *NotifyPeer) To_NotifyAll() *TLNotifyAll {
	return &TLNotifyAll{
		Data2: m.Data2,
	}
}

// notifyPeer#9fd40bd8 peer:Peer = NotifyPeer;
func (m *TLNotifyPeer) To_NotifyPeer() *NotifyPeer {
	return &NotifyPeer{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLNotifyPeer) SetPeer(v *Peer) { m.Data2.Peer = v }
func (m *TLNotifyPeer) GetPeer() *Peer  { return m.Data2.Peer }

func (m *TLNotifyPeer) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_notifyPeer))

	x.Bytes(m.Data2Peer.Encode())

	return x.buf
}

func (m *TLNotifyPeer) Decode(dbuf *DecodeBuf) error {
	m1 := &Peer{}
	m1.Decode(dbuf)
	m.SetPeer(m1)

	return dbuf.err
}

// notifyUsers#b4c83b4c = NotifyPeer;
func (m *TLNotifyUsers) To_NotifyPeer() *NotifyPeer {
	return &NotifyPeer{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLNotifyUsers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_notifyUsers))

	return x.buf
}

func (m *TLNotifyUsers) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// notifyChats#c007cec3 = NotifyPeer;
func (m *TLNotifyChats) To_NotifyPeer() *NotifyPeer {
	return &NotifyPeer{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLNotifyChats) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_notifyChats))

	return x.buf
}

func (m *TLNotifyChats) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// notifyAll#74d07c60 = NotifyPeer;
func (m *TLNotifyAll) To_NotifyPeer() *NotifyPeer {
	return &NotifyPeer{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLNotifyAll) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_notifyAll))

	return x.buf
}

func (m *TLNotifyAll) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// ContactStatus <--
//  + TL_ContactStatus
//

func (m *ContactStatus) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_contactStatus:
		t := m.To_ContactStatus()
		return t.Encode()

	default:
		return []byte{}
	}
}

// contactStatus#d3680c61 user_id:int status:UserStatus = ContactStatus;
func (m *ContactStatus) To_ContactStatus() *TLContactStatus {
	return &TLContactStatus{
		Data2: m.Data2,
	}
}

// contactStatus#d3680c61 user_id:int status:UserStatus = ContactStatus;
func (m *TLContactStatus) To_ContactStatus() *ContactStatus {
	return &ContactStatus{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLContactStatus) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLContactStatus) GetUserId() int32  { return m.Data2.UserId }

func (m *TLContactStatus) SetStatus(v *UserStatus) { m.Data2.Status = v }
func (m *TLContactStatus) GetStatus() *UserStatus  { return m.Data2.Status }

func (m *TLContactStatus) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_contactStatus))

	x.Int(m.Data2.UserId)
	x.Bytes(m.Data2Status.Encode())

	return x.buf
}

func (m *TLContactStatus) Decode(dbuf *DecodeBuf) error {
	m.Data2.UserId = x.Int()
	m2 := &Status{}
	m2.Decode(dbuf)
	m.SetStatus(m2)

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Account_Authorizations <--
//  + TL_AccountAuthorizations
//

func (m *Account_Authorizations) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_account_authorizations:
		t := m.To_AccountAuthorizations()
		return t.Encode()

	default:
		return []byte{}
	}
}

// account.authorizations#1250abde authorizations:Vector<Authorization> = account.Authorizations;
func (m *Account_Authorizations) To_AccountAuthorizations() *TLAccountAuthorizations {
	return &TLAccountAuthorizations{
		Data2: m.Data2,
	}
}

// account.authorizations#1250abde authorizations:Vector<Authorization> = account.Authorizations;
func (m *TLAccountAuthorizations) To_Account_Authorizations() *Account_Authorizations {
	return &Account_Authorizations{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLAccountAuthorizations) SetAuthorizations(v []*Authorization) { m.Data2.Authorizations = v }
func (m *TLAccountAuthorizations) GetAuthorizations() []*Authorization  { return m.Data2.Authorizations }

func (m *TLAccountAuthorizations) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_account_authorizations))

	return x.buf
}

func (m *TLAccountAuthorizations) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// LangPackLanguage <--
//  + TL_LangPackLanguage
//

func (m *LangPackLanguage) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_langPackLanguage:
		t := m.To_LangPackLanguage()
		return t.Encode()

	default:
		return []byte{}
	}
}

// langPackLanguage#117698f1 name:string native_name:string lang_code:string = LangPackLanguage;
func (m *LangPackLanguage) To_LangPackLanguage() *TLLangPackLanguage {
	return &TLLangPackLanguage{
		Data2: m.Data2,
	}
}

// langPackLanguage#117698f1 name:string native_name:string lang_code:string = LangPackLanguage;
func (m *TLLangPackLanguage) To_LangPackLanguage() *LangPackLanguage {
	return &LangPackLanguage{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLLangPackLanguage) SetName(v string) { m.Data2.Name = v }
func (m *TLLangPackLanguage) GetName() string  { return m.Data2.Name }

func (m *TLLangPackLanguage) SetNativeName(v string) { m.Data2.NativeName = v }
func (m *TLLangPackLanguage) GetNativeName() string  { return m.Data2.NativeName }

func (m *TLLangPackLanguage) SetLangCode(v string) { m.Data2.LangCode = v }
func (m *TLLangPackLanguage) GetLangCode() string  { return m.Data2.LangCode }

func (m *TLLangPackLanguage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_langPackLanguage))

	x.StringBytes(m.Data2.Name)
	x.StringBytes(m.Data2.NativeName)
	x.StringBytes(m.Data2.LangCode)

	return x.buf
}

func (m *TLLangPackLanguage) Decode(dbuf *DecodeBuf) error {
	m.Data2.Name = x.StringBytes()
	m.Data2.NativeName = x.StringBytes()
	m.Data2.LangCode = x.StringBytes()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Messages_BotCallbackAnswer <--
//  + TL_MessagesBotCallbackAnswer
//

func (m *Messages_BotCallbackAnswer) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_messages_botCallbackAnswer:
		t := m.To_MessagesBotCallbackAnswer()
		return t.Encode()

	default:
		return []byte{}
	}
}

// messages.botCallbackAnswer#36585ea4 flags:# alert:flags.1?true has_url:flags.3?true message:flags.0?string url:flags.2?string cache_time:int = messages.BotCallbackAnswer;
func (m *Messages_BotCallbackAnswer) To_MessagesBotCallbackAnswer() *TLMessagesBotCallbackAnswer {
	return &TLMessagesBotCallbackAnswer{
		Data2: m.Data2,
	}
}

// messages.botCallbackAnswer#36585ea4 flags:# alert:flags.1?true has_url:flags.3?true message:flags.0?string url:flags.2?string cache_time:int = messages.BotCallbackAnswer;
func (m *TLMessagesBotCallbackAnswer) To_Messages_BotCallbackAnswer() *Messages_BotCallbackAnswer {
	return &Messages_BotCallbackAnswer{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessagesBotCallbackAnswer) SetAlert(v bool) { m.Data2.Alert = v }
func (m *TLMessagesBotCallbackAnswer) GetAlert() bool  { return m.Data2.Alert }

func (m *TLMessagesBotCallbackAnswer) SetHasUrl(v bool) { m.Data2.HasUrl = v }
func (m *TLMessagesBotCallbackAnswer) GetHasUrl() bool  { return m.Data2.HasUrl }

func (m *TLMessagesBotCallbackAnswer) SetMessage(v string) { m.Data2.Message = v }
func (m *TLMessagesBotCallbackAnswer) GetMessage() string  { return m.Data2.Message }

func (m *TLMessagesBotCallbackAnswer) SetUrl(v string) { m.Data2.Url = v }
func (m *TLMessagesBotCallbackAnswer) GetUrl() string  { return m.Data2.Url }

func (m *TLMessagesBotCallbackAnswer) SetCacheTime(v int32) { m.Data2.CacheTime = v }
func (m *TLMessagesBotCallbackAnswer) GetCacheTime() int32  { return m.Data2.CacheTime }

func (m *TLMessagesBotCallbackAnswer) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_botCallbackAnswer))

	// flags
	var flags uint32 = 0
	if m.GetAlert() == true {
		flags |= 1 << 1
	}
	if m.GetHasUrl() == true {
		flags |= 1 << 2
	}
	if m.GetMessage() != "" {
		flags |= 1 << 3
	}
	if m.GetUrl() != "" {
		flags |= 1 << 4
	}
	x.UInt(flags)

	if m.GetMessage() != 0 {
		x.StringBytes(m.Data2.Message)
	}
	if m.GetUrl() != 0 {
		x.StringBytes(m.Data2.Url)
	}
	x.Int(m.Data2.CacheTime)

	return x.buf
}

func (m *TLMessagesBotCallbackAnswer) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m.Data2.Alert = true
	}
	if (flags & (1 << 3)) != 0 {
		m.Data2.HasUrl = true
	}
	if (flags & (1 << 4)) != 0 {
		m.Data2.Message = x.StringBytes()
	}
	if (flags & (1 << 5)) != 0 {
		m.Data2.Url = x.StringBytes()
	}
	m.Data2.CacheTime = x.Int()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// InlineBotSwitchPM <--
//  + TL_InlineBotSwitchPM
//

func (m *InlineBotSwitchPM) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_inlineBotSwitchPM:
		t := m.To_InlineBotSwitchPM()
		return t.Encode()

	default:
		return []byte{}
	}
}

// inlineBotSwitchPM#3c20629f text:string start_param:string = InlineBotSwitchPM;
func (m *InlineBotSwitchPM) To_InlineBotSwitchPM() *TLInlineBotSwitchPM {
	return &TLInlineBotSwitchPM{
		Data2: m.Data2,
	}
}

// inlineBotSwitchPM#3c20629f text:string start_param:string = InlineBotSwitchPM;
func (m *TLInlineBotSwitchPM) To_InlineBotSwitchPM() *InlineBotSwitchPM {
	return &InlineBotSwitchPM{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInlineBotSwitchPM) SetText(v string) { m.Data2.Text = v }
func (m *TLInlineBotSwitchPM) GetText() string  { return m.Data2.Text }

func (m *TLInlineBotSwitchPM) SetStartParam(v string) { m.Data2.StartParam = v }
func (m *TLInlineBotSwitchPM) GetStartParam() string  { return m.Data2.StartParam }

func (m *TLInlineBotSwitchPM) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inlineBotSwitchPM))

	x.StringBytes(m.Data2.Text)
	x.StringBytes(m.Data2.StartParam)

	return x.buf
}

func (m *TLInlineBotSwitchPM) Decode(dbuf *DecodeBuf) error {
	m.Data2.Text = x.StringBytes()
	m.Data2.StartParam = x.StringBytes()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// PostAddress <--
//  + TL_PostAddress
//

func (m *PostAddress) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_postAddress:
		t := m.To_PostAddress()
		return t.Encode()

	default:
		return []byte{}
	}
}

// postAddress#1e8caaeb street_line1:string street_line2:string city:string state:string country_iso2:string post_code:string = PostAddress;
func (m *PostAddress) To_PostAddress() *TLPostAddress {
	return &TLPostAddress{
		Data2: m.Data2,
	}
}

// postAddress#1e8caaeb street_line1:string street_line2:string city:string state:string country_iso2:string post_code:string = PostAddress;
func (m *TLPostAddress) To_PostAddress() *PostAddress {
	return &PostAddress{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPostAddress) SetStreetLine1(v string) { m.Data2.StreetLine1 = v }
func (m *TLPostAddress) GetStreetLine1() string  { return m.Data2.StreetLine1 }

func (m *TLPostAddress) SetStreetLine2(v string) { m.Data2.StreetLine2 = v }
func (m *TLPostAddress) GetStreetLine2() string  { return m.Data2.StreetLine2 }

func (m *TLPostAddress) SetCity(v string) { m.Data2.City = v }
func (m *TLPostAddress) GetCity() string  { return m.Data2.City }

func (m *TLPostAddress) SetState(v string) { m.Data2.State = v }
func (m *TLPostAddress) GetState() string  { return m.Data2.State }

func (m *TLPostAddress) SetCountryIso2(v string) { m.Data2.CountryIso2 = v }
func (m *TLPostAddress) GetCountryIso2() string  { return m.Data2.CountryIso2 }

func (m *TLPostAddress) SetPostCode(v string) { m.Data2.PostCode = v }
func (m *TLPostAddress) GetPostCode() string  { return m.Data2.PostCode }

func (m *TLPostAddress) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_postAddress))

	x.StringBytes(m.Data2.StreetLine1)
	x.StringBytes(m.Data2.StreetLine2)
	x.StringBytes(m.Data2.City)
	x.StringBytes(m.Data2.State)
	x.StringBytes(m.Data2.CountryIso2)
	x.StringBytes(m.Data2.PostCode)

	return x.buf
}

func (m *TLPostAddress) Decode(dbuf *DecodeBuf) error {
	m.Data2.StreetLine1 = x.StringBytes()
	m.Data2.StreetLine2 = x.StringBytes()
	m.Data2.City = x.StringBytes()
	m.Data2.State = x.StringBytes()
	m.Data2.CountryIso2 = x.StringBytes()
	m.Data2.PostCode = x.StringBytes()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// ChannelAdminRights <--
//  + TL_ChannelAdminRights
//

func (m *ChannelAdminRights) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_channelAdminRights:
		t := m.To_ChannelAdminRights()
		return t.Encode()

	default:
		return []byte{}
	}
}

// channelAdminRights#5d7ceba5 flags:# change_info:flags.0?true post_messages:flags.1?true edit_messages:flags.2?true delete_messages:flags.3?true ban_users:flags.4?true invite_users:flags.5?true invite_link:flags.6?true pin_messages:flags.7?true add_admins:flags.9?true = ChannelAdminRights;
func (m *ChannelAdminRights) To_ChannelAdminRights() *TLChannelAdminRights {
	return &TLChannelAdminRights{
		Data2: m.Data2,
	}
}

// channelAdminRights#5d7ceba5 flags:# change_info:flags.0?true post_messages:flags.1?true edit_messages:flags.2?true delete_messages:flags.3?true ban_users:flags.4?true invite_users:flags.5?true invite_link:flags.6?true pin_messages:flags.7?true add_admins:flags.9?true = ChannelAdminRights;
func (m *TLChannelAdminRights) To_ChannelAdminRights() *ChannelAdminRights {
	return &ChannelAdminRights{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLChannelAdminRights) SetChangeInfo(v bool) { m.Data2.ChangeInfo = v }
func (m *TLChannelAdminRights) GetChangeInfo() bool  { return m.Data2.ChangeInfo }

func (m *TLChannelAdminRights) SetPostMessages(v bool) { m.Data2.PostMessages = v }
func (m *TLChannelAdminRights) GetPostMessages() bool  { return m.Data2.PostMessages }

func (m *TLChannelAdminRights) SetEditMessages(v bool) { m.Data2.EditMessages = v }
func (m *TLChannelAdminRights) GetEditMessages() bool  { return m.Data2.EditMessages }

func (m *TLChannelAdminRights) SetDeleteMessages(v bool) { m.Data2.DeleteMessages = v }
func (m *TLChannelAdminRights) GetDeleteMessages() bool  { return m.Data2.DeleteMessages }

func (m *TLChannelAdminRights) SetBanUsers(v bool) { m.Data2.BanUsers = v }
func (m *TLChannelAdminRights) GetBanUsers() bool  { return m.Data2.BanUsers }

func (m *TLChannelAdminRights) SetInviteUsers(v bool) { m.Data2.InviteUsers = v }
func (m *TLChannelAdminRights) GetInviteUsers() bool  { return m.Data2.InviteUsers }

func (m *TLChannelAdminRights) SetInviteLink(v bool) { m.Data2.InviteLink = v }
func (m *TLChannelAdminRights) GetInviteLink() bool  { return m.Data2.InviteLink }

func (m *TLChannelAdminRights) SetPinMessages(v bool) { m.Data2.PinMessages = v }
func (m *TLChannelAdminRights) GetPinMessages() bool  { return m.Data2.PinMessages }

func (m *TLChannelAdminRights) SetAddAdmins(v bool) { m.Data2.AddAdmins = v }
func (m *TLChannelAdminRights) GetAddAdmins() bool  { return m.Data2.AddAdmins }

func (m *TLChannelAdminRights) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelAdminRights))

	// flags
	var flags uint32 = 0
	if m.GetChangeInfo() == true {
		flags |= 1 << 1
	}
	if m.GetPostMessages() == true {
		flags |= 1 << 2
	}
	if m.GetEditMessages() == true {
		flags |= 1 << 3
	}
	if m.GetDeleteMessages() == true {
		flags |= 1 << 4
	}
	if m.GetBanUsers() == true {
		flags |= 1 << 5
	}
	if m.GetInviteUsers() == true {
		flags |= 1 << 6
	}
	if m.GetInviteLink() == true {
		flags |= 1 << 7
	}
	if m.GetPinMessages() == true {
		flags |= 1 << 8
	}
	if m.GetAddAdmins() == true {
		flags |= 1 << 9
	}
	x.UInt(flags)

	return x.buf
}

func (m *TLChannelAdminRights) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m.Data2.ChangeInfo = true
	}
	if (flags & (1 << 3)) != 0 {
		m.Data2.PostMessages = true
	}
	if (flags & (1 << 4)) != 0 {
		m.Data2.EditMessages = true
	}
	if (flags & (1 << 5)) != 0 {
		m.Data2.DeleteMessages = true
	}
	if (flags & (1 << 6)) != 0 {
		m.Data2.BanUsers = true
	}
	if (flags & (1 << 7)) != 0 {
		m.Data2.InviteUsers = true
	}
	if (flags & (1 << 8)) != 0 {
		m.Data2.InviteLink = true
	}
	if (flags & (1 << 9)) != 0 {
		m.Data2.PinMessages = true
	}
	if (flags & (1 << 10)) != 0 {
		m.Data2.AddAdmins = true
	}

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Messages_AffectedHistory <--
//  + TL_MessagesAffectedHistory
//

func (m *Messages_AffectedHistory) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_messages_affectedHistory:
		t := m.To_MessagesAffectedHistory()
		return t.Encode()

	default:
		return []byte{}
	}
}

// messages.affectedHistory#b45c69d1 pts:int pts_count:int offset:int = messages.AffectedHistory;
func (m *Messages_AffectedHistory) To_MessagesAffectedHistory() *TLMessagesAffectedHistory {
	return &TLMessagesAffectedHistory{
		Data2: m.Data2,
	}
}

// messages.affectedHistory#b45c69d1 pts:int pts_count:int offset:int = messages.AffectedHistory;
func (m *TLMessagesAffectedHistory) To_Messages_AffectedHistory() *Messages_AffectedHistory {
	return &Messages_AffectedHistory{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessagesAffectedHistory) SetPts(v int32) { m.Data2.Pts = v }
func (m *TLMessagesAffectedHistory) GetPts() int32  { return m.Data2.Pts }

func (m *TLMessagesAffectedHistory) SetPtsCount(v int32) { m.Data2.PtsCount = v }
func (m *TLMessagesAffectedHistory) GetPtsCount() int32  { return m.Data2.PtsCount }

func (m *TLMessagesAffectedHistory) SetOffset(v int32) { m.Data2.Offset = v }
func (m *TLMessagesAffectedHistory) GetOffset() int32  { return m.Data2.Offset }

func (m *TLMessagesAffectedHistory) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_affectedHistory))

	x.Int(m.Data2.Pts)
	x.Int(m.Data2.PtsCount)
	x.Int(m.Data2.Offset)

	return x.buf
}

func (m *TLMessagesAffectedHistory) Decode(dbuf *DecodeBuf) error {
	m.Data2.Pts = x.Int()
	m.Data2.PtsCount = x.Int()
	m.Data2.Offset = x.Int()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// EncryptedChat <--
//  + TL_EncryptedChatEmpty
//  + TL_EncryptedChatWaiting
//  + TL_EncryptedChatRequested
//  + TL_EncryptedChat
//  + TL_EncryptedChatDiscarded
//

func (m *EncryptedChat) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_encryptedChatEmpty:
		t := m.To_EncryptedChatEmpty()
		return t.Encode()
	case TLConstructor_CRC32_encryptedChatWaiting:
		t := m.To_EncryptedChatWaiting()
		return t.Encode()
	case TLConstructor_CRC32_encryptedChatRequested:
		t := m.To_EncryptedChatRequested()
		return t.Encode()
	case TLConstructor_CRC32_encryptedChat:
		t := m.To_EncryptedChat()
		return t.Encode()
	case TLConstructor_CRC32_encryptedChatDiscarded:
		t := m.To_EncryptedChatDiscarded()
		return t.Encode()

	default:
		return []byte{}
	}
}

// encryptedChatEmpty#ab7ec0a0 id:int = EncryptedChat;
func (m *EncryptedChat) To_EncryptedChatEmpty() *TLEncryptedChatEmpty {
	return &TLEncryptedChatEmpty{
		Data2: m.Data2,
	}
}

// encryptedChatWaiting#3bf703dc id:int access_hash:long date:int admin_id:int participant_id:int = EncryptedChat;
func (m *EncryptedChat) To_EncryptedChatWaiting() *TLEncryptedChatWaiting {
	return &TLEncryptedChatWaiting{
		Data2: m.Data2,
	}
}

// encryptedChatRequested#c878527e id:int access_hash:long date:int admin_id:int participant_id:int g_a:bytes = EncryptedChat;
func (m *EncryptedChat) To_EncryptedChatRequested() *TLEncryptedChatRequested {
	return &TLEncryptedChatRequested{
		Data2: m.Data2,
	}
}

// encryptedChat#fa56ce36 id:int access_hash:long date:int admin_id:int participant_id:int g_a_or_b:bytes key_fingerprint:long = EncryptedChat;
func (m *EncryptedChat) To_EncryptedChat() *TLEncryptedChat {
	return &TLEncryptedChat{
		Data2: m.Data2,
	}
}

// encryptedChatDiscarded#13d6dd27 id:int = EncryptedChat;
func (m *EncryptedChat) To_EncryptedChatDiscarded() *TLEncryptedChatDiscarded {
	return &TLEncryptedChatDiscarded{
		Data2: m.Data2,
	}
}

// encryptedChatEmpty#ab7ec0a0 id:int = EncryptedChat;
func (m *TLEncryptedChatEmpty) To_EncryptedChat() *EncryptedChat {
	return &EncryptedChat{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLEncryptedChatEmpty) SetId(v int32) { m.Data2.Id = v }
func (m *TLEncryptedChatEmpty) GetId() int32  { return m.Data2.Id }

func (m *TLEncryptedChatEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_encryptedChatEmpty))

	x.Int(m.Data2.Id)

	return x.buf
}

func (m *TLEncryptedChatEmpty) Decode(dbuf *DecodeBuf) error {
	m.Data2.Id = x.Int()

	return dbuf.err
}

// encryptedChatWaiting#3bf703dc id:int access_hash:long date:int admin_id:int participant_id:int = EncryptedChat;
func (m *TLEncryptedChatWaiting) To_EncryptedChat() *EncryptedChat {
	return &EncryptedChat{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLEncryptedChatWaiting) SetId(v int32) { m.Data2.Id = v }
func (m *TLEncryptedChatWaiting) GetId() int32  { return m.Data2.Id }

func (m *TLEncryptedChatWaiting) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLEncryptedChatWaiting) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLEncryptedChatWaiting) SetDate(v int32) { m.Data2.Date = v }
func (m *TLEncryptedChatWaiting) GetDate() int32  { return m.Data2.Date }

func (m *TLEncryptedChatWaiting) SetAdminId(v int32) { m.Data2.AdminId = v }
func (m *TLEncryptedChatWaiting) GetAdminId() int32  { return m.Data2.AdminId }

func (m *TLEncryptedChatWaiting) SetParticipantId(v int32) { m.Data2.ParticipantId = v }
func (m *TLEncryptedChatWaiting) GetParticipantId() int32  { return m.Data2.ParticipantId }

func (m *TLEncryptedChatWaiting) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_encryptedChatWaiting))

	x.Int(m.Data2.Id)
	x.Long(m.Data2.AccessHash)
	x.Int(m.Data2.Date)
	x.Int(m.Data2.AdminId)
	x.Int(m.Data2.ParticipantId)

	return x.buf
}

func (m *TLEncryptedChatWaiting) Decode(dbuf *DecodeBuf) error {
	m.Data2.Id = x.Int()
	m.Data2.AccessHash = x.Long()
	m.Data2.Date = x.Int()
	m.Data2.AdminId = x.Int()
	m.Data2.ParticipantId = x.Int()

	return dbuf.err
}

// encryptedChatRequested#c878527e id:int access_hash:long date:int admin_id:int participant_id:int g_a:bytes = EncryptedChat;
func (m *TLEncryptedChatRequested) To_EncryptedChat() *EncryptedChat {
	return &EncryptedChat{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLEncryptedChatRequested) SetId(v int32) { m.Data2.Id = v }
func (m *TLEncryptedChatRequested) GetId() int32  { return m.Data2.Id }

func (m *TLEncryptedChatRequested) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLEncryptedChatRequested) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLEncryptedChatRequested) SetDate(v int32) { m.Data2.Date = v }
func (m *TLEncryptedChatRequested) GetDate() int32  { return m.Data2.Date }

func (m *TLEncryptedChatRequested) SetAdminId(v int32) { m.Data2.AdminId = v }
func (m *TLEncryptedChatRequested) GetAdminId() int32  { return m.Data2.AdminId }

func (m *TLEncryptedChatRequested) SetParticipantId(v int32) { m.Data2.ParticipantId = v }
func (m *TLEncryptedChatRequested) GetParticipantId() int32  { return m.Data2.ParticipantId }

func (m *TLEncryptedChatRequested) SetGA(v []byte) { m.Data2.GA = v }
func (m *TLEncryptedChatRequested) GetGA() []byte  { return m.Data2.GA }

func (m *TLEncryptedChatRequested) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_encryptedChatRequested))

	x.Int(m.Data2.Id)
	x.Long(m.Data2.AccessHash)
	x.Int(m.Data2.Date)
	x.Int(m.Data2.AdminId)
	x.Int(m.Data2.ParticipantId)
	x.StringBytes(m.Data2.GA)

	return x.buf
}

func (m *TLEncryptedChatRequested) Decode(dbuf *DecodeBuf) error {
	m.Data2.Id = x.Int()
	m.Data2.AccessHash = x.Long()
	m.Data2.Date = x.Int()
	m.Data2.AdminId = x.Int()
	m.Data2.ParticipantId = x.Int()
	m.Data2.GA = x.StringBytes()

	return dbuf.err
}

// encryptedChat#fa56ce36 id:int access_hash:long date:int admin_id:int participant_id:int g_a_or_b:bytes key_fingerprint:long = EncryptedChat;
func (m *TLEncryptedChat) To_EncryptedChat() *EncryptedChat {
	return &EncryptedChat{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLEncryptedChat) SetId(v int32) { m.Data2.Id = v }
func (m *TLEncryptedChat) GetId() int32  { return m.Data2.Id }

func (m *TLEncryptedChat) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLEncryptedChat) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLEncryptedChat) SetDate(v int32) { m.Data2.Date = v }
func (m *TLEncryptedChat) GetDate() int32  { return m.Data2.Date }

func (m *TLEncryptedChat) SetAdminId(v int32) { m.Data2.AdminId = v }
func (m *TLEncryptedChat) GetAdminId() int32  { return m.Data2.AdminId }

func (m *TLEncryptedChat) SetParticipantId(v int32) { m.Data2.ParticipantId = v }
func (m *TLEncryptedChat) GetParticipantId() int32  { return m.Data2.ParticipantId }

func (m *TLEncryptedChat) SetGAOrB(v []byte) { m.Data2.GAOrB = v }
func (m *TLEncryptedChat) GetGAOrB() []byte  { return m.Data2.GAOrB }

func (m *TLEncryptedChat) SetKeyFingerprint(v int64) { m.Data2.KeyFingerprint = v }
func (m *TLEncryptedChat) GetKeyFingerprint() int64  { return m.Data2.KeyFingerprint }

func (m *TLEncryptedChat) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_encryptedChat))

	x.Int(m.Data2.Id)
	x.Long(m.Data2.AccessHash)
	x.Int(m.Data2.Date)
	x.Int(m.Data2.AdminId)
	x.Int(m.Data2.ParticipantId)
	x.StringBytes(m.Data2.GAOrB)
	x.Long(m.Data2.KeyFingerprint)

	return x.buf
}

func (m *TLEncryptedChat) Decode(dbuf *DecodeBuf) error {
	m.Data2.Id = x.Int()
	m.Data2.AccessHash = x.Long()
	m.Data2.Date = x.Int()
	m.Data2.AdminId = x.Int()
	m.Data2.ParticipantId = x.Int()
	m.Data2.GAOrB = x.StringBytes()
	m.Data2.KeyFingerprint = x.Long()

	return dbuf.err
}

// encryptedChatDiscarded#13d6dd27 id:int = EncryptedChat;
func (m *TLEncryptedChatDiscarded) To_EncryptedChat() *EncryptedChat {
	return &EncryptedChat{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLEncryptedChatDiscarded) SetId(v int32) { m.Data2.Id = v }
func (m *TLEncryptedChatDiscarded) GetId() int32  { return m.Data2.Id }

func (m *TLEncryptedChatDiscarded) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_encryptedChatDiscarded))

	x.Int(m.Data2.Id)

	return x.buf
}

func (m *TLEncryptedChatDiscarded) Decode(dbuf *DecodeBuf) error {
	m.Data2.Id = x.Int()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Messages_DhConfig <--
//  + TL_MessagesDhConfigNotModified
//  + TL_MessagesDhConfig
//

func (m *Messages_DhConfig) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_messages_dhConfigNotModified:
		t := m.To_MessagesDhConfigNotModified()
		return t.Encode()
	case TLConstructor_CRC32_messages_dhConfig:
		t := m.To_MessagesDhConfig()
		return t.Encode()

	default:
		return []byte{}
	}
}

// messages.dhConfigNotModified#c0e24635 random:bytes = messages.DhConfig;
func (m *Messages_DhConfig) To_MessagesDhConfigNotModified() *TLMessagesDhConfigNotModified {
	return &TLMessagesDhConfigNotModified{
		Data2: m.Data2,
	}
}

// messages.dhConfig#2c221edd g:int p:bytes version:int random:bytes = messages.DhConfig;
func (m *Messages_DhConfig) To_MessagesDhConfig() *TLMessagesDhConfig {
	return &TLMessagesDhConfig{
		Data2: m.Data2,
	}
}

// messages.dhConfigNotModified#c0e24635 random:bytes = messages.DhConfig;
func (m *TLMessagesDhConfigNotModified) To_Messages_DhConfig() *Messages_DhConfig {
	return &Messages_DhConfig{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessagesDhConfigNotModified) SetRandom(v []byte) { m.Data2.Random = v }
func (m *TLMessagesDhConfigNotModified) GetRandom() []byte  { return m.Data2.Random }

func (m *TLMessagesDhConfigNotModified) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_dhConfigNotModified))

	x.StringBytes(m.Data2.Random)

	return x.buf
}

func (m *TLMessagesDhConfigNotModified) Decode(dbuf *DecodeBuf) error {
	m.Data2.Random = x.StringBytes()

	return dbuf.err
}

// messages.dhConfig#2c221edd g:int p:bytes version:int random:bytes = messages.DhConfig;
func (m *TLMessagesDhConfig) To_Messages_DhConfig() *Messages_DhConfig {
	return &Messages_DhConfig{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessagesDhConfig) SetG(v int32) { m.Data2.G = v }
func (m *TLMessagesDhConfig) GetG() int32  { return m.Data2.G }

func (m *TLMessagesDhConfig) SetP(v []byte) { m.Data2.P = v }
func (m *TLMessagesDhConfig) GetP() []byte  { return m.Data2.P }

func (m *TLMessagesDhConfig) SetVersion(v int32) { m.Data2.Version = v }
func (m *TLMessagesDhConfig) GetVersion() int32  { return m.Data2.Version }

func (m *TLMessagesDhConfig) SetRandom(v []byte) { m.Data2.Random = v }
func (m *TLMessagesDhConfig) GetRandom() []byte  { return m.Data2.Random }

func (m *TLMessagesDhConfig) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_dhConfig))

	x.Int(m.Data2.G)
	x.StringBytes(m.Data2.P)
	x.Int(m.Data2.Version)
	x.StringBytes(m.Data2.Random)

	return x.buf
}

func (m *TLMessagesDhConfig) Decode(dbuf *DecodeBuf) error {
	m.Data2.G = x.Int()
	m.Data2.P = x.StringBytes()
	m.Data2.Version = x.Int()
	m.Data2.Random = x.StringBytes()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// MessageFwdHeader <--
//  + TL_MessageFwdHeader
//

func (m *MessageFwdHeader) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_messageFwdHeader:
		t := m.To_MessageFwdHeader()
		return t.Encode()

	default:
		return []byte{}
	}
}

// messageFwdHeader#fadff4ac flags:# from_id:flags.0?int date:int channel_id:flags.1?int channel_post:flags.2?int post_author:flags.3?string = MessageFwdHeader;
func (m *MessageFwdHeader) To_MessageFwdHeader() *TLMessageFwdHeader {
	return &TLMessageFwdHeader{
		Data2: m.Data2,
	}
}

// messageFwdHeader#fadff4ac flags:# from_id:flags.0?int date:int channel_id:flags.1?int channel_post:flags.2?int post_author:flags.3?string = MessageFwdHeader;
func (m *TLMessageFwdHeader) To_MessageFwdHeader() *MessageFwdHeader {
	return &MessageFwdHeader{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessageFwdHeader) SetFromId(v int32) { m.Data2.FromId = v }
func (m *TLMessageFwdHeader) GetFromId() int32  { return m.Data2.FromId }

func (m *TLMessageFwdHeader) SetDate(v int32) { m.Data2.Date = v }
func (m *TLMessageFwdHeader) GetDate() int32  { return m.Data2.Date }

func (m *TLMessageFwdHeader) SetChannelId(v int32) { m.Data2.ChannelId = v }
func (m *TLMessageFwdHeader) GetChannelId() int32  { return m.Data2.ChannelId }

func (m *TLMessageFwdHeader) SetChannelPost(v int32) { m.Data2.ChannelPost = v }
func (m *TLMessageFwdHeader) GetChannelPost() int32  { return m.Data2.ChannelPost }

func (m *TLMessageFwdHeader) SetPostAuthor(v string) { m.Data2.PostAuthor = v }
func (m *TLMessageFwdHeader) GetPostAuthor() string  { return m.Data2.PostAuthor }

func (m *TLMessageFwdHeader) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageFwdHeader))

	// flags
	var flags uint32 = 0
	if m.GetFromId() != 0 {
		flags |= 1 << 1
	}
	if m.GetChannelId() != 0 {
		flags |= 1 << 3
	}
	if m.GetChannelPost() != 0 {
		flags |= 1 << 4
	}
	if m.GetPostAuthor() != "" {
		flags |= 1 << 5
	}
	x.UInt(flags)

	if m.GetFromId() != 0 {
		x.Int(m.Data2.FromId)
	}
	x.Int(m.Data2.Date)
	if m.GetChannelId() != 0 {
		x.Int(m.Data2.ChannelId)
	}
	if m.GetChannelPost() != 0 {
		x.Int(m.Data2.ChannelPost)
	}
	if m.GetPostAuthor() != 0 {
		x.StringBytes(m.Data2.PostAuthor)
	}

	return x.buf
}

func (m *TLMessageFwdHeader) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m.Data2.FromId = x.Int()
	}
	m.Data2.Date = x.Int()
	if (flags & (1 << 4)) != 0 {
		m.Data2.ChannelId = x.Int()
	}
	if (flags & (1 << 5)) != 0 {
		m.Data2.ChannelPost = x.Int()
	}
	if (flags & (1 << 6)) != 0 {
		m.Data2.PostAuthor = x.StringBytes()
	}

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Messages_FoundGifs <--
//  + TL_MessagesFoundGifs
//

func (m *Messages_FoundGifs) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_messages_foundGifs:
		t := m.To_MessagesFoundGifs()
		return t.Encode()

	default:
		return []byte{}
	}
}

// messages.foundGifs#450a1c0a next_offset:int results:Vector<FoundGif> = messages.FoundGifs;
func (m *Messages_FoundGifs) To_MessagesFoundGifs() *TLMessagesFoundGifs {
	return &TLMessagesFoundGifs{
		Data2: m.Data2,
	}
}

// messages.foundGifs#450a1c0a next_offset:int results:Vector<FoundGif> = messages.FoundGifs;
func (m *TLMessagesFoundGifs) To_Messages_FoundGifs() *Messages_FoundGifs {
	return &Messages_FoundGifs{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessagesFoundGifs) SetNextOffset(v int32) { m.Data2.NextOffset = v }
func (m *TLMessagesFoundGifs) GetNextOffset() int32  { return m.Data2.NextOffset }

func (m *TLMessagesFoundGifs) SetResults(v []*FoundGif) { m.Data2.Results = v }
func (m *TLMessagesFoundGifs) GetResults() []*FoundGif  { return m.Data2.Results }

func (m *TLMessagesFoundGifs) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_foundGifs))

	x.Int(m.Data2.NextOffset)

	return x.buf
}

func (m *TLMessagesFoundGifs) Decode(dbuf *DecodeBuf) error {
	m.Data2.NextOffset = x.Int()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Game <--
//  + TL_Game
//

func (m *Game) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_game:
		t := m.To_Game()
		return t.Encode()

	default:
		return []byte{}
	}
}

// game#bdf9653b flags:# id:long access_hash:long short_name:string title:string description:string photo:Photo document:flags.0?Document = Game;
func (m *Game) To_Game() *TLGame {
	return &TLGame{
		Data2: m.Data2,
	}
}

// game#bdf9653b flags:# id:long access_hash:long short_name:string title:string description:string photo:Photo document:flags.0?Document = Game;
func (m *TLGame) To_Game() *Game {
	return &Game{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLGame) SetId(v int64) { m.Data2.Id = v }
func (m *TLGame) GetId() int64  { return m.Data2.Id }

func (m *TLGame) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLGame) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLGame) SetShortName(v string) { m.Data2.ShortName = v }
func (m *TLGame) GetShortName() string  { return m.Data2.ShortName }

func (m *TLGame) SetTitle(v string) { m.Data2.Title = v }
func (m *TLGame) GetTitle() string  { return m.Data2.Title }

func (m *TLGame) SetDescription(v string) { m.Data2.Description = v }
func (m *TLGame) GetDescription() string  { return m.Data2.Description }

func (m *TLGame) SetPhoto(v *Photo) { m.Data2.Photo = v }
func (m *TLGame) GetPhoto() *Photo  { return m.Data2.Photo }

func (m *TLGame) SetDocument(v *Document) { m.Data2.Document = v }
func (m *TLGame) GetDocument() *Document  { return m.Data2.Document }

func (m *TLGame) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_game))

	// flags
	var flags uint32 = 0
	if m.GetDocument() != nil {
		flags |= 1 << 7
	}
	x.UInt(flags)

	x.Long(m.Data2.Id)
	x.Long(m.Data2.AccessHash)
	x.StringBytes(m.Data2.ShortName)
	x.StringBytes(m.Data2.Title)
	x.StringBytes(m.Data2.Description)
	x.Bytes(m.Data2Photo.Encode())
	if m.GetDocument() != 0 {
		x.Bytes(m.Data2Document.Encode())
	}

	return x.buf
}

func (m *TLGame) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	m.Data2.Id = x.Long()
	m.Data2.AccessHash = x.Long()
	m.Data2.ShortName = x.StringBytes()
	m.Data2.Title = x.StringBytes()
	m.Data2.Description = x.StringBytes()
	m7 := &Photo{}
	m7.Decode(dbuf)
	m.SetPhoto(m7)
	if (flags & (1 << 8)) != 0 {
		m8 := &Document{}
		m8.Decode(dbuf)
		m.SetDocument(m8)
	}

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// HighScore <--
//  + TL_HighScore
//

func (m *HighScore) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_highScore:
		t := m.To_HighScore()
		return t.Encode()

	default:
		return []byte{}
	}
}

// highScore#58fffcd0 pos:int user_id:int score:int = HighScore;
func (m *HighScore) To_HighScore() *TLHighScore {
	return &TLHighScore{
		Data2: m.Data2,
	}
}

// highScore#58fffcd0 pos:int user_id:int score:int = HighScore;
func (m *TLHighScore) To_HighScore() *HighScore {
	return &HighScore{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLHighScore) SetPos(v int32) { m.Data2.Pos = v }
func (m *TLHighScore) GetPos() int32  { return m.Data2.Pos }

func (m *TLHighScore) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLHighScore) GetUserId() int32  { return m.Data2.UserId }

func (m *TLHighScore) SetScore(v int32) { m.Data2.Score = v }
func (m *TLHighScore) GetScore() int32  { return m.Data2.Score }

func (m *TLHighScore) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_highScore))

	x.Int(m.Data2.Pos)
	x.Int(m.Data2.UserId)
	x.Int(m.Data2.Score)

	return x.buf
}

func (m *TLHighScore) Decode(dbuf *DecodeBuf) error {
	m.Data2.Pos = x.Int()
	m.Data2.UserId = x.Int()
	m.Data2.Score = x.Int()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// PhoneCallProtocol <--
//  + TL_PhoneCallProtocol
//

func (m *PhoneCallProtocol) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_phoneCallProtocol:
		t := m.To_PhoneCallProtocol()
		return t.Encode()

	default:
		return []byte{}
	}
}

// phoneCallProtocol#a2bb35cb flags:# udp_p2p:flags.0?true udp_reflector:flags.1?true min_layer:int max_layer:int = PhoneCallProtocol;
func (m *PhoneCallProtocol) To_PhoneCallProtocol() *TLPhoneCallProtocol {
	return &TLPhoneCallProtocol{
		Data2: m.Data2,
	}
}

// phoneCallProtocol#a2bb35cb flags:# udp_p2p:flags.0?true udp_reflector:flags.1?true min_layer:int max_layer:int = PhoneCallProtocol;
func (m *TLPhoneCallProtocol) To_PhoneCallProtocol() *PhoneCallProtocol {
	return &PhoneCallProtocol{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPhoneCallProtocol) SetUdpP2P(v bool) { m.Data2.UdpP2P = v }
func (m *TLPhoneCallProtocol) GetUdpP2P() bool  { return m.Data2.UdpP2P }

func (m *TLPhoneCallProtocol) SetUdpReflector(v bool) { m.Data2.UdpReflector = v }
func (m *TLPhoneCallProtocol) GetUdpReflector() bool  { return m.Data2.UdpReflector }

func (m *TLPhoneCallProtocol) SetMinLayer(v int32) { m.Data2.MinLayer = v }
func (m *TLPhoneCallProtocol) GetMinLayer() int32  { return m.Data2.MinLayer }

func (m *TLPhoneCallProtocol) SetMaxLayer(v int32) { m.Data2.MaxLayer = v }
func (m *TLPhoneCallProtocol) GetMaxLayer() int32  { return m.Data2.MaxLayer }

func (m *TLPhoneCallProtocol) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_phoneCallProtocol))

	// flags
	var flags uint32 = 0
	if m.GetUdpP2P() == true {
		flags |= 1 << 1
	}
	if m.GetUdpReflector() == true {
		flags |= 1 << 2
	}
	x.UInt(flags)

	x.Int(m.Data2.MinLayer)
	x.Int(m.Data2.MaxLayer)

	return x.buf
}

func (m *TLPhoneCallProtocol) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m.Data2.UdpP2P = true
	}
	if (flags & (1 << 3)) != 0 {
		m.Data2.UdpReflector = true
	}
	m.Data2.MinLayer = x.Int()
	m.Data2.MaxLayer = x.Int()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// UserStatus <--
//  + TL_UserStatusEmpty
//  + TL_UserStatusOnline
//  + TL_UserStatusOffline
//  + TL_UserStatusRecently
//  + TL_UserStatusLastWeek
//  + TL_UserStatusLastMonth
//

func (m *UserStatus) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_userStatusEmpty:
		t := m.To_UserStatusEmpty()
		return t.Encode()
	case TLConstructor_CRC32_userStatusOnline:
		t := m.To_UserStatusOnline()
		return t.Encode()
	case TLConstructor_CRC32_userStatusOffline:
		t := m.To_UserStatusOffline()
		return t.Encode()
	case TLConstructor_CRC32_userStatusRecently:
		t := m.To_UserStatusRecently()
		return t.Encode()
	case TLConstructor_CRC32_userStatusLastWeek:
		t := m.To_UserStatusLastWeek()
		return t.Encode()
	case TLConstructor_CRC32_userStatusLastMonth:
		t := m.To_UserStatusLastMonth()
		return t.Encode()

	default:
		return []byte{}
	}
}

// userStatusEmpty#9d05049 = UserStatus;
func (m *UserStatus) To_UserStatusEmpty() *TLUserStatusEmpty {
	return &TLUserStatusEmpty{
		Data2: m.Data2,
	}
}

// userStatusOnline#edb93949 expires:int = UserStatus;
func (m *UserStatus) To_UserStatusOnline() *TLUserStatusOnline {
	return &TLUserStatusOnline{
		Data2: m.Data2,
	}
}

// userStatusOffline#8c703f was_online:int = UserStatus;
func (m *UserStatus) To_UserStatusOffline() *TLUserStatusOffline {
	return &TLUserStatusOffline{
		Data2: m.Data2,
	}
}

// userStatusRecently#e26f42f1 = UserStatus;
func (m *UserStatus) To_UserStatusRecently() *TLUserStatusRecently {
	return &TLUserStatusRecently{
		Data2: m.Data2,
	}
}

// userStatusLastWeek#7bf09fc = UserStatus;
func (m *UserStatus) To_UserStatusLastWeek() *TLUserStatusLastWeek {
	return &TLUserStatusLastWeek{
		Data2: m.Data2,
	}
}

// userStatusLastMonth#77ebc742 = UserStatus;
func (m *UserStatus) To_UserStatusLastMonth() *TLUserStatusLastMonth {
	return &TLUserStatusLastMonth{
		Data2: m.Data2,
	}
}

// userStatusEmpty#9d05049 = UserStatus;
func (m *TLUserStatusEmpty) To_UserStatus() *UserStatus {
	return &UserStatus{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUserStatusEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_userStatusEmpty))

	return x.buf
}

func (m *TLUserStatusEmpty) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// userStatusOnline#edb93949 expires:int = UserStatus;
func (m *TLUserStatusOnline) To_UserStatus() *UserStatus {
	return &UserStatus{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUserStatusOnline) SetExpires(v int32) { m.Data2.Expires = v }
func (m *TLUserStatusOnline) GetExpires() int32  { return m.Data2.Expires }

func (m *TLUserStatusOnline) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_userStatusOnline))

	x.Int(m.Data2.Expires)

	return x.buf
}

func (m *TLUserStatusOnline) Decode(dbuf *DecodeBuf) error {
	m.Data2.Expires = x.Int()

	return dbuf.err
}

// userStatusOffline#8c703f was_online:int = UserStatus;
func (m *TLUserStatusOffline) To_UserStatus() *UserStatus {
	return &UserStatus{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUserStatusOffline) SetWasOnline(v int32) { m.Data2.WasOnline = v }
func (m *TLUserStatusOffline) GetWasOnline() int32  { return m.Data2.WasOnline }

func (m *TLUserStatusOffline) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_userStatusOffline))

	x.Int(m.Data2.WasOnline)

	return x.buf
}

func (m *TLUserStatusOffline) Decode(dbuf *DecodeBuf) error {
	m.Data2.WasOnline = x.Int()

	return dbuf.err
}

// userStatusRecently#e26f42f1 = UserStatus;
func (m *TLUserStatusRecently) To_UserStatus() *UserStatus {
	return &UserStatus{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUserStatusRecently) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_userStatusRecently))

	return x.buf
}

func (m *TLUserStatusRecently) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// userStatusLastWeek#7bf09fc = UserStatus;
func (m *TLUserStatusLastWeek) To_UserStatus() *UserStatus {
	return &UserStatus{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUserStatusLastWeek) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_userStatusLastWeek))

	return x.buf
}

func (m *TLUserStatusLastWeek) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// userStatusLastMonth#77ebc742 = UserStatus;
func (m *TLUserStatusLastMonth) To_UserStatus() *UserStatus {
	return &UserStatus{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUserStatusLastMonth) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_userStatusLastMonth))

	return x.buf
}

func (m *TLUserStatusLastMonth) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// BotInfo <--
//  + TL_BotInfo
//

func (m *BotInfo) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_botInfo:
		t := m.To_BotInfo()
		return t.Encode()

	default:
		return []byte{}
	}
}

// botInfo#98e81d3a user_id:int description:string commands:Vector<BotCommand> = BotInfo;
func (m *BotInfo) To_BotInfo() *TLBotInfo {
	return &TLBotInfo{
		Data2: m.Data2,
	}
}

// botInfo#98e81d3a user_id:int description:string commands:Vector<BotCommand> = BotInfo;
func (m *TLBotInfo) To_BotInfo() *BotInfo {
	return &BotInfo{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLBotInfo) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLBotInfo) GetUserId() int32  { return m.Data2.UserId }

func (m *TLBotInfo) SetDescription(v string) { m.Data2.Description = v }
func (m *TLBotInfo) GetDescription() string  { return m.Data2.Description }

func (m *TLBotInfo) SetCommands(v []*BotCommand) { m.Data2.Commands = v }
func (m *TLBotInfo) GetCommands() []*BotCommand  { return m.Data2.Commands }

func (m *TLBotInfo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_botInfo))

	x.Int(m.Data2.UserId)
	x.StringBytes(m.Data2.Description)

	return x.buf
}

func (m *TLBotInfo) Decode(dbuf *DecodeBuf) error {
	m.Data2.UserId = x.Int()
	m.Data2.Description = x.StringBytes()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Contacts_ResolvedPeer <--
//  + TL_ContactsResolvedPeer
//

func (m *Contacts_ResolvedPeer) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_contacts_resolvedPeer:
		t := m.To_ContactsResolvedPeer()
		return t.Encode()

	default:
		return []byte{}
	}
}

// contacts.resolvedPeer#7f077ad9 peer:Peer chats:Vector<Chat> users:Vector<User> = contacts.ResolvedPeer;
func (m *Contacts_ResolvedPeer) To_ContactsResolvedPeer() *TLContactsResolvedPeer {
	return &TLContactsResolvedPeer{
		Data2: m.Data2,
	}
}

// contacts.resolvedPeer#7f077ad9 peer:Peer chats:Vector<Chat> users:Vector<User> = contacts.ResolvedPeer;
func (m *TLContactsResolvedPeer) To_Contacts_ResolvedPeer() *Contacts_ResolvedPeer {
	return &Contacts_ResolvedPeer{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLContactsResolvedPeer) SetPeer(v *Peer) { m.Data2.Peer = v }
func (m *TLContactsResolvedPeer) GetPeer() *Peer  { return m.Data2.Peer }

func (m *TLContactsResolvedPeer) SetChats(v []*Chat) { m.Data2.Chats = v }
func (m *TLContactsResolvedPeer) GetChats() []*Chat  { return m.Data2.Chats }

func (m *TLContactsResolvedPeer) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLContactsResolvedPeer) GetUsers() []*User  { return m.Data2.Users }

func (m *TLContactsResolvedPeer) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_contacts_resolvedPeer))

	x.Bytes(m.Data2Peer.Encode())

	return x.buf
}

func (m *TLContactsResolvedPeer) Decode(dbuf *DecodeBuf) error {
	m1 := &Peer{}
	m1.Decode(dbuf)
	m.SetPeer(m1)

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Help_TermsOfService <--
//  + TL_HelpTermsOfService
//

func (m *Help_TermsOfService) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_help_termsOfService:
		t := m.To_HelpTermsOfService()
		return t.Encode()

	default:
		return []byte{}
	}
}

// help.termsOfService#f1ee3e90 text:string = help.TermsOfService;
func (m *Help_TermsOfService) To_HelpTermsOfService() *TLHelpTermsOfService {
	return &TLHelpTermsOfService{
		Data2: m.Data2,
	}
}

// help.termsOfService#f1ee3e90 text:string = help.TermsOfService;
func (m *TLHelpTermsOfService) To_Help_TermsOfService() *Help_TermsOfService {
	return &Help_TermsOfService{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLHelpTermsOfService) SetText(v string) { m.Data2.Text = v }
func (m *TLHelpTermsOfService) GetText() string  { return m.Data2.Text }

func (m *TLHelpTermsOfService) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_help_termsOfService))

	x.StringBytes(m.Data2.Text)

	return x.buf
}

func (m *TLHelpTermsOfService) Decode(dbuf *DecodeBuf) error {
	m.Data2.Text = x.StringBytes()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Messages_ArchivedStickers <--
//  + TL_MessagesArchivedStickers
//

func (m *Messages_ArchivedStickers) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_messages_archivedStickers:
		t := m.To_MessagesArchivedStickers()
		return t.Encode()

	default:
		return []byte{}
	}
}

// messages.archivedStickers#4fcba9c8 count:int sets:Vector<StickerSetCovered> = messages.ArchivedStickers;
func (m *Messages_ArchivedStickers) To_MessagesArchivedStickers() *TLMessagesArchivedStickers {
	return &TLMessagesArchivedStickers{
		Data2: m.Data2,
	}
}

// messages.archivedStickers#4fcba9c8 count:int sets:Vector<StickerSetCovered> = messages.ArchivedStickers;
func (m *TLMessagesArchivedStickers) To_Messages_ArchivedStickers() *Messages_ArchivedStickers {
	return &Messages_ArchivedStickers{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessagesArchivedStickers) SetCount(v int32) { m.Data2.Count = v }
func (m *TLMessagesArchivedStickers) GetCount() int32  { return m.Data2.Count }

func (m *TLMessagesArchivedStickers) SetSets(v []*StickerSetCovered) { m.Data2.Sets = v }
func (m *TLMessagesArchivedStickers) GetSets() []*StickerSetCovered  { return m.Data2.Sets }

func (m *TLMessagesArchivedStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_archivedStickers))

	x.Int(m.Data2.Count)

	return x.buf
}

func (m *TLMessagesArchivedStickers) Decode(dbuf *DecodeBuf) error {
	m.Data2.Count = x.Int()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// InputGame <--
//  + TL_InputGameID
//  + TL_InputGameShortName
//

func (m *InputGame) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_inputGameID:
		t := m.To_InputGameID()
		return t.Encode()
	case TLConstructor_CRC32_inputGameShortName:
		t := m.To_InputGameShortName()
		return t.Encode()

	default:
		return []byte{}
	}
}

// inputGameID#32c3e77 id:long access_hash:long = InputGame;
func (m *InputGame) To_InputGameID() *TLInputGameID {
	return &TLInputGameID{
		Data2: m.Data2,
	}
}

// inputGameShortName#c331e80a bot_id:InputUser short_name:string = InputGame;
func (m *InputGame) To_InputGameShortName() *TLInputGameShortName {
	return &TLInputGameShortName{
		Data2: m.Data2,
	}
}

// inputGameID#32c3e77 id:long access_hash:long = InputGame;
func (m *TLInputGameID) To_InputGame() *InputGame {
	return &InputGame{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputGameID) SetId(v int64) { m.Data2.Id = v }
func (m *TLInputGameID) GetId() int64  { return m.Data2.Id }

func (m *TLInputGameID) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLInputGameID) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLInputGameID) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputGameID))

	x.Long(m.Data2.Id)
	x.Long(m.Data2.AccessHash)

	return x.buf
}

func (m *TLInputGameID) Decode(dbuf *DecodeBuf) error {
	m.Data2.Id = x.Long()
	m.Data2.AccessHash = x.Long()

	return dbuf.err
}

// inputGameShortName#c331e80a bot_id:InputUser short_name:string = InputGame;
func (m *TLInputGameShortName) To_InputGame() *InputGame {
	return &InputGame{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputGameShortName) SetBotId(v *InputUser) { m.Data2.BotId = v }
func (m *TLInputGameShortName) GetBotId() *InputUser  { return m.Data2.BotId }

func (m *TLInputGameShortName) SetShortName(v string) { m.Data2.ShortName = v }
func (m *TLInputGameShortName) GetShortName() string  { return m.Data2.ShortName }

func (m *TLInputGameShortName) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputGameShortName))

	x.Bytes(m.Data2BotId.Encode())
	x.StringBytes(m.Data2.ShortName)

	return x.buf
}

func (m *TLInputGameShortName) Decode(dbuf *DecodeBuf) error {
	m1 := &BotId{}
	m1.Decode(dbuf)
	m.SetBotId(m1)
	m.Data2.ShortName = x.StringBytes()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// PaymentRequestedInfo <--
//  + TL_PaymentRequestedInfo
//

func (m *PaymentRequestedInfo) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_paymentRequestedInfo:
		t := m.To_PaymentRequestedInfo()
		return t.Encode()

	default:
		return []byte{}
	}
}

// paymentRequestedInfo#909c3f94 flags:# name:flags.0?string phone:flags.1?string email:flags.2?string shipping_address:flags.3?PostAddress = PaymentRequestedInfo;
func (m *PaymentRequestedInfo) To_PaymentRequestedInfo() *TLPaymentRequestedInfo {
	return &TLPaymentRequestedInfo{
		Data2: m.Data2,
	}
}

// paymentRequestedInfo#909c3f94 flags:# name:flags.0?string phone:flags.1?string email:flags.2?string shipping_address:flags.3?PostAddress = PaymentRequestedInfo;
func (m *TLPaymentRequestedInfo) To_PaymentRequestedInfo() *PaymentRequestedInfo {
	return &PaymentRequestedInfo{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPaymentRequestedInfo) SetName(v string) { m.Data2.Name = v }
func (m *TLPaymentRequestedInfo) GetName() string  { return m.Data2.Name }

func (m *TLPaymentRequestedInfo) SetPhone(v string) { m.Data2.Phone = v }
func (m *TLPaymentRequestedInfo) GetPhone() string  { return m.Data2.Phone }

func (m *TLPaymentRequestedInfo) SetEmail(v string) { m.Data2.Email = v }
func (m *TLPaymentRequestedInfo) GetEmail() string  { return m.Data2.Email }

func (m *TLPaymentRequestedInfo) SetShippingAddress(v *PostAddress) { m.Data2.ShippingAddress = v }
func (m *TLPaymentRequestedInfo) GetShippingAddress() *PostAddress  { return m.Data2.ShippingAddress }

func (m *TLPaymentRequestedInfo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_paymentRequestedInfo))

	// flags
	var flags uint32 = 0
	if m.GetName() != "" {
		flags |= 1 << 1
	}
	if m.GetPhone() != "" {
		flags |= 1 << 2
	}
	if m.GetEmail() != "" {
		flags |= 1 << 3
	}
	if m.GetShippingAddress() != nil {
		flags |= 1 << 4
	}
	x.UInt(flags)

	if m.GetName() != 0 {
		x.StringBytes(m.Data2.Name)
	}
	if m.GetPhone() != 0 {
		x.StringBytes(m.Data2.Phone)
	}
	if m.GetEmail() != 0 {
		x.StringBytes(m.Data2.Email)
	}
	if m.GetShippingAddress() != 0 {
		x.Bytes(m.Data2ShippingAddress.Encode())
	}

	return x.buf
}

func (m *TLPaymentRequestedInfo) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m.Data2.Name = x.StringBytes()
	}
	if (flags & (1 << 3)) != 0 {
		m.Data2.Phone = x.StringBytes()
	}
	if (flags & (1 << 4)) != 0 {
		m.Data2.Email = x.StringBytes()
	}
	if (flags & (1 << 5)) != 0 {
		m5 := &ShippingAddress{}
		m5.Decode(dbuf)
		m.SetShippingAddress(m5)
	}

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// ChatParticipants <--
//  + TL_ChatParticipantsForbidden
//  + TL_ChatParticipants
//

func (m *ChatParticipants) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_chatParticipantsForbidden:
		t := m.To_ChatParticipantsForbidden()
		return t.Encode()
	case TLConstructor_CRC32_chatParticipants:
		t := m.To_ChatParticipants()
		return t.Encode()

	default:
		return []byte{}
	}
}

// chatParticipantsForbidden#fc900c2b flags:# chat_id:int self_participant:flags.0?ChatParticipant = ChatParticipants;
func (m *ChatParticipants) To_ChatParticipantsForbidden() *TLChatParticipantsForbidden {
	return &TLChatParticipantsForbidden{
		Data2: m.Data2,
	}
}

// chatParticipants#3f460fed chat_id:int participants:Vector<ChatParticipant> version:int = ChatParticipants;
func (m *ChatParticipants) To_ChatParticipants() *TLChatParticipants {
	return &TLChatParticipants{
		Data2: m.Data2,
	}
}

// chatParticipantsForbidden#fc900c2b flags:# chat_id:int self_participant:flags.0?ChatParticipant = ChatParticipants;
func (m *TLChatParticipantsForbidden) To_ChatParticipants() *ChatParticipants {
	return &ChatParticipants{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLChatParticipantsForbidden) SetChatId(v int32) { m.Data2.ChatId = v }
func (m *TLChatParticipantsForbidden) GetChatId() int32  { return m.Data2.ChatId }

func (m *TLChatParticipantsForbidden) SetSelfParticipant(v *ChatParticipant) {
	m.Data2.SelfParticipant = v
}
func (m *TLChatParticipantsForbidden) GetSelfParticipant() *ChatParticipant {
	return m.Data2.SelfParticipant
}

func (m *TLChatParticipantsForbidden) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_chatParticipantsForbidden))

	// flags
	var flags uint32 = 0
	if m.GetSelfParticipant() != nil {
		flags |= 1 << 2
	}
	x.UInt(flags)

	x.Int(m.Data2.ChatId)
	if m.GetSelfParticipant() != 0 {
		x.Bytes(m.Data2SelfParticipant.Encode())
	}

	return x.buf
}

func (m *TLChatParticipantsForbidden) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	m.Data2.ChatId = x.Int()
	if (flags & (1 << 3)) != 0 {
		m3 := &SelfParticipant{}
		m3.Decode(dbuf)
		m.SetSelfParticipant(m3)
	}

	return dbuf.err
}

// chatParticipants#3f460fed chat_id:int participants:Vector<ChatParticipant> version:int = ChatParticipants;
func (m *TLChatParticipants) To_ChatParticipants() *ChatParticipants {
	return &ChatParticipants{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLChatParticipants) SetChatId(v int32) { m.Data2.ChatId = v }
func (m *TLChatParticipants) GetChatId() int32  { return m.Data2.ChatId }

func (m *TLChatParticipants) SetParticipants(v []*ChatParticipant) { m.Data2.Participants = v }
func (m *TLChatParticipants) GetParticipants() []*ChatParticipant  { return m.Data2.Participants }

func (m *TLChatParticipants) SetVersion(v int32) { m.Data2.Version = v }
func (m *TLChatParticipants) GetVersion() int32  { return m.Data2.Version }

func (m *TLChatParticipants) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_chatParticipants))

	x.Int(m.Data2.ChatId)

	x.Int(m.Data2.Version)

	return x.buf
}

func (m *TLChatParticipants) Decode(dbuf *DecodeBuf) error {
	m.Data2.ChatId = x.Int()

	m.Data2.Version = x.Int()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// ReportReason <--
//  + TL_InputReportReasonSpam
//  + TL_InputReportReasonViolence
//  + TL_InputReportReasonPornography
//  + TL_InputReportReasonOther
//

func (m *ReportReason) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_inputReportReasonSpam:
		t := m.To_InputReportReasonSpam()
		return t.Encode()
	case TLConstructor_CRC32_inputReportReasonViolence:
		t := m.To_InputReportReasonViolence()
		return t.Encode()
	case TLConstructor_CRC32_inputReportReasonPornography:
		t := m.To_InputReportReasonPornography()
		return t.Encode()
	case TLConstructor_CRC32_inputReportReasonOther:
		t := m.To_InputReportReasonOther()
		return t.Encode()

	default:
		return []byte{}
	}
}

// inputReportReasonSpam#58dbcab8 = ReportReason;
func (m *ReportReason) To_InputReportReasonSpam() *TLInputReportReasonSpam {
	return &TLInputReportReasonSpam{
		Data2: m.Data2,
	}
}

// inputReportReasonViolence#1e22c78d = ReportReason;
func (m *ReportReason) To_InputReportReasonViolence() *TLInputReportReasonViolence {
	return &TLInputReportReasonViolence{
		Data2: m.Data2,
	}
}

// inputReportReasonPornography#2e59d922 = ReportReason;
func (m *ReportReason) To_InputReportReasonPornography() *TLInputReportReasonPornography {
	return &TLInputReportReasonPornography{
		Data2: m.Data2,
	}
}

// inputReportReasonOther#e1746d0a text:string = ReportReason;
func (m *ReportReason) To_InputReportReasonOther() *TLInputReportReasonOther {
	return &TLInputReportReasonOther{
		Data2: m.Data2,
	}
}

// inputReportReasonSpam#58dbcab8 = ReportReason;
func (m *TLInputReportReasonSpam) To_ReportReason() *ReportReason {
	return &ReportReason{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputReportReasonSpam) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputReportReasonSpam))

	return x.buf
}

func (m *TLInputReportReasonSpam) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// inputReportReasonViolence#1e22c78d = ReportReason;
func (m *TLInputReportReasonViolence) To_ReportReason() *ReportReason {
	return &ReportReason{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputReportReasonViolence) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputReportReasonViolence))

	return x.buf
}

func (m *TLInputReportReasonViolence) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// inputReportReasonPornography#2e59d922 = ReportReason;
func (m *TLInputReportReasonPornography) To_ReportReason() *ReportReason {
	return &ReportReason{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputReportReasonPornography) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputReportReasonPornography))

	return x.buf
}

func (m *TLInputReportReasonPornography) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// inputReportReasonOther#e1746d0a text:string = ReportReason;
func (m *TLInputReportReasonOther) To_ReportReason() *ReportReason {
	return &ReportReason{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputReportReasonOther) SetText(v string) { m.Data2.Text = v }
func (m *TLInputReportReasonOther) GetText() string  { return m.Data2.Text }

func (m *TLInputReportReasonOther) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputReportReasonOther))

	x.StringBytes(m.Data2.Text)

	return x.buf
}

func (m *TLInputReportReasonOther) Decode(dbuf *DecodeBuf) error {
	m.Data2.Text = x.StringBytes()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Help_AppUpdate <--
//  + TL_HelpAppUpdate
//  + TL_HelpNoAppUpdate
//

func (m *Help_AppUpdate) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_help_appUpdate:
		t := m.To_HelpAppUpdate()
		return t.Encode()
	case TLConstructor_CRC32_help_noAppUpdate:
		t := m.To_HelpNoAppUpdate()
		return t.Encode()

	default:
		return []byte{}
	}
}

// help.appUpdate#8987f311 id:int critical:Bool url:string text:string = help.AppUpdate;
func (m *Help_AppUpdate) To_HelpAppUpdate() *TLHelpAppUpdate {
	return &TLHelpAppUpdate{
		Data2: m.Data2,
	}
}

// help.noAppUpdate#c45a6536 = help.AppUpdate;
func (m *Help_AppUpdate) To_HelpNoAppUpdate() *TLHelpNoAppUpdate {
	return &TLHelpNoAppUpdate{
		Data2: m.Data2,
	}
}

// help.appUpdate#8987f311 id:int critical:Bool url:string text:string = help.AppUpdate;
func (m *TLHelpAppUpdate) To_Help_AppUpdate() *Help_AppUpdate {
	return &Help_AppUpdate{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLHelpAppUpdate) SetId(v int32) { m.Data2.Id = v }
func (m *TLHelpAppUpdate) GetId() int32  { return m.Data2.Id }

func (m *TLHelpAppUpdate) SetCritical(v *Bool) { m.Data2.Critical = v }
func (m *TLHelpAppUpdate) GetCritical() *Bool  { return m.Data2.Critical }

func (m *TLHelpAppUpdate) SetUrl(v string) { m.Data2.Url = v }
func (m *TLHelpAppUpdate) GetUrl() string  { return m.Data2.Url }

func (m *TLHelpAppUpdate) SetText(v string) { m.Data2.Text = v }
func (m *TLHelpAppUpdate) GetText() string  { return m.Data2.Text }

func (m *TLHelpAppUpdate) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_help_appUpdate))

	x.Int(m.Data2.Id)
	x.Bytes(m.Data2Critical.Encode())
	x.StringBytes(m.Data2.Url)
	x.StringBytes(m.Data2.Text)

	return x.buf
}

func (m *TLHelpAppUpdate) Decode(dbuf *DecodeBuf) error {
	m.Data2.Id = x.Int()
	m2 := &Critical{}
	m2.Decode(dbuf)
	m.SetCritical(m2)
	m.Data2.Url = x.StringBytes()
	m.Data2.Text = x.StringBytes()

	return dbuf.err
}

// help.noAppUpdate#c45a6536 = help.AppUpdate;
func (m *TLHelpNoAppUpdate) To_Help_AppUpdate() *Help_AppUpdate {
	return &Help_AppUpdate{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLHelpNoAppUpdate) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_help_noAppUpdate))

	return x.buf
}

func (m *TLHelpNoAppUpdate) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// InputEncryptedFile <--
//  + TL_InputEncryptedFileEmpty
//  + TL_InputEncryptedFileUploaded
//  + TL_InputEncryptedFile
//  + TL_InputEncryptedFileBigUploaded
//

func (m *InputEncryptedFile) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_inputEncryptedFileEmpty:
		t := m.To_InputEncryptedFileEmpty()
		return t.Encode()
	case TLConstructor_CRC32_inputEncryptedFileUploaded:
		t := m.To_InputEncryptedFileUploaded()
		return t.Encode()
	case TLConstructor_CRC32_inputEncryptedFile:
		t := m.To_InputEncryptedFile()
		return t.Encode()
	case TLConstructor_CRC32_inputEncryptedFileBigUploaded:
		t := m.To_InputEncryptedFileBigUploaded()
		return t.Encode()

	default:
		return []byte{}
	}
}

// inputEncryptedFileEmpty#1837c364 = InputEncryptedFile;
func (m *InputEncryptedFile) To_InputEncryptedFileEmpty() *TLInputEncryptedFileEmpty {
	return &TLInputEncryptedFileEmpty{
		Data2: m.Data2,
	}
}

// inputEncryptedFileUploaded#64bd0306 id:long parts:int md5_checksum:string key_fingerprint:int = InputEncryptedFile;
func (m *InputEncryptedFile) To_InputEncryptedFileUploaded() *TLInputEncryptedFileUploaded {
	return &TLInputEncryptedFileUploaded{
		Data2: m.Data2,
	}
}

// inputEncryptedFile#5a17b5e5 id:long access_hash:long = InputEncryptedFile;
func (m *InputEncryptedFile) To_InputEncryptedFile() *TLInputEncryptedFile {
	return &TLInputEncryptedFile{
		Data2: m.Data2,
	}
}

// inputEncryptedFileBigUploaded#2dc173c8 id:long parts:int key_fingerprint:int = InputEncryptedFile;
func (m *InputEncryptedFile) To_InputEncryptedFileBigUploaded() *TLInputEncryptedFileBigUploaded {
	return &TLInputEncryptedFileBigUploaded{
		Data2: m.Data2,
	}
}

// inputEncryptedFileEmpty#1837c364 = InputEncryptedFile;
func (m *TLInputEncryptedFileEmpty) To_InputEncryptedFile() *InputEncryptedFile {
	return &InputEncryptedFile{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputEncryptedFileEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputEncryptedFileEmpty))

	return x.buf
}

func (m *TLInputEncryptedFileEmpty) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// inputEncryptedFileUploaded#64bd0306 id:long parts:int md5_checksum:string key_fingerprint:int = InputEncryptedFile;
func (m *TLInputEncryptedFileUploaded) To_InputEncryptedFile() *InputEncryptedFile {
	return &InputEncryptedFile{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputEncryptedFileUploaded) SetId(v int64) { m.Data2.Id = v }
func (m *TLInputEncryptedFileUploaded) GetId() int64  { return m.Data2.Id }

func (m *TLInputEncryptedFileUploaded) SetParts(v int32) { m.Data2.Parts = v }
func (m *TLInputEncryptedFileUploaded) GetParts() int32  { return m.Data2.Parts }

func (m *TLInputEncryptedFileUploaded) SetMd5Checksum(v string) { m.Data2.Md5Checksum = v }
func (m *TLInputEncryptedFileUploaded) GetMd5Checksum() string  { return m.Data2.Md5Checksum }

func (m *TLInputEncryptedFileUploaded) SetKeyFingerprint(v int32) { m.Data2.KeyFingerprint = v }
func (m *TLInputEncryptedFileUploaded) GetKeyFingerprint() int32  { return m.Data2.KeyFingerprint }

func (m *TLInputEncryptedFileUploaded) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputEncryptedFileUploaded))

	x.Long(m.Data2.Id)
	x.Int(m.Data2.Parts)
	x.StringBytes(m.Data2.Md5Checksum)
	x.Int(m.Data2.KeyFingerprint)

	return x.buf
}

func (m *TLInputEncryptedFileUploaded) Decode(dbuf *DecodeBuf) error {
	m.Data2.Id = x.Long()
	m.Data2.Parts = x.Int()
	m.Data2.Md5Checksum = x.StringBytes()
	m.Data2.KeyFingerprint = x.Int()

	return dbuf.err
}

// inputEncryptedFile#5a17b5e5 id:long access_hash:long = InputEncryptedFile;
func (m *TLInputEncryptedFile) To_InputEncryptedFile() *InputEncryptedFile {
	return &InputEncryptedFile{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputEncryptedFile) SetId(v int64) { m.Data2.Id = v }
func (m *TLInputEncryptedFile) GetId() int64  { return m.Data2.Id }

func (m *TLInputEncryptedFile) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLInputEncryptedFile) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLInputEncryptedFile) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputEncryptedFile))

	x.Long(m.Data2.Id)
	x.Long(m.Data2.AccessHash)

	return x.buf
}

func (m *TLInputEncryptedFile) Decode(dbuf *DecodeBuf) error {
	m.Data2.Id = x.Long()
	m.Data2.AccessHash = x.Long()

	return dbuf.err
}

// inputEncryptedFileBigUploaded#2dc173c8 id:long parts:int key_fingerprint:int = InputEncryptedFile;
func (m *TLInputEncryptedFileBigUploaded) To_InputEncryptedFile() *InputEncryptedFile {
	return &InputEncryptedFile{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputEncryptedFileBigUploaded) SetId(v int64) { m.Data2.Id = v }
func (m *TLInputEncryptedFileBigUploaded) GetId() int64  { return m.Data2.Id }

func (m *TLInputEncryptedFileBigUploaded) SetParts(v int32) { m.Data2.Parts = v }
func (m *TLInputEncryptedFileBigUploaded) GetParts() int32  { return m.Data2.Parts }

func (m *TLInputEncryptedFileBigUploaded) SetKeyFingerprint(v int32) { m.Data2.KeyFingerprint = v }
func (m *TLInputEncryptedFileBigUploaded) GetKeyFingerprint() int32  { return m.Data2.KeyFingerprint }

func (m *TLInputEncryptedFileBigUploaded) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputEncryptedFileBigUploaded))

	x.Long(m.Data2.Id)
	x.Int(m.Data2.Parts)
	x.Int(m.Data2.KeyFingerprint)

	return x.buf
}

func (m *TLInputEncryptedFileBigUploaded) Decode(dbuf *DecodeBuf) error {
	m.Data2.Id = x.Long()
	m.Data2.Parts = x.Int()
	m.Data2.KeyFingerprint = x.Int()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Messages_RecentStickers <--
//  + TL_MessagesRecentStickersNotModified
//  + TL_MessagesRecentStickers
//

func (m *Messages_RecentStickers) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_messages_recentStickersNotModified:
		t := m.To_MessagesRecentStickersNotModified()
		return t.Encode()
	case TLConstructor_CRC32_messages_recentStickers:
		t := m.To_MessagesRecentStickers()
		return t.Encode()

	default:
		return []byte{}
	}
}

// messages.recentStickersNotModified#b17f890 = messages.RecentStickers;
func (m *Messages_RecentStickers) To_MessagesRecentStickersNotModified() *TLMessagesRecentStickersNotModified {
	return &TLMessagesRecentStickersNotModified{
		Data2: m.Data2,
	}
}

// messages.recentStickers#5ce20970 hash:int stickers:Vector<Document> = messages.RecentStickers;
func (m *Messages_RecentStickers) To_MessagesRecentStickers() *TLMessagesRecentStickers {
	return &TLMessagesRecentStickers{
		Data2: m.Data2,
	}
}

// messages.recentStickersNotModified#b17f890 = messages.RecentStickers;
func (m *TLMessagesRecentStickersNotModified) To_Messages_RecentStickers() *Messages_RecentStickers {
	return &Messages_RecentStickers{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessagesRecentStickersNotModified) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_recentStickersNotModified))

	return x.buf
}

func (m *TLMessagesRecentStickersNotModified) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// messages.recentStickers#5ce20970 hash:int stickers:Vector<Document> = messages.RecentStickers;
func (m *TLMessagesRecentStickers) To_Messages_RecentStickers() *Messages_RecentStickers {
	return &Messages_RecentStickers{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessagesRecentStickers) SetHash(v int32) { m.Data2.Hash = v }
func (m *TLMessagesRecentStickers) GetHash() int32  { return m.Data2.Hash }

func (m *TLMessagesRecentStickers) SetStickers(v []*Document) { m.Data2.Stickers = v }
func (m *TLMessagesRecentStickers) GetStickers() []*Document  { return m.Data2.Stickers }

func (m *TLMessagesRecentStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_recentStickers))

	x.Int(m.Data2.Hash)

	return x.buf
}

func (m *TLMessagesRecentStickers) Decode(dbuf *DecodeBuf) error {
	m.Data2.Hash = x.Int()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// InputStickeredMedia <--
//  + TL_InputStickeredMediaPhoto
//  + TL_InputStickeredMediaDocument
//

func (m *InputStickeredMedia) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_inputStickeredMediaPhoto:
		t := m.To_InputStickeredMediaPhoto()
		return t.Encode()
	case TLConstructor_CRC32_inputStickeredMediaDocument:
		t := m.To_InputStickeredMediaDocument()
		return t.Encode()

	default:
		return []byte{}
	}
}

// inputStickeredMediaPhoto#4a992157 id:InputPhoto = InputStickeredMedia;
func (m *InputStickeredMedia) To_InputStickeredMediaPhoto() *TLInputStickeredMediaPhoto {
	return &TLInputStickeredMediaPhoto{
		Data2: m.Data2,
	}
}

// inputStickeredMediaDocument#438865b id:InputDocument = InputStickeredMedia;
func (m *InputStickeredMedia) To_InputStickeredMediaDocument() *TLInputStickeredMediaDocument {
	return &TLInputStickeredMediaDocument{
		Data2: m.Data2,
	}
}

// inputStickeredMediaPhoto#4a992157 id:InputPhoto = InputStickeredMedia;
func (m *TLInputStickeredMediaPhoto) To_InputStickeredMedia() *InputStickeredMedia {
	return &InputStickeredMedia{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputStickeredMediaPhoto) SetId(v *InputPhoto) { m.Data2.Id_1 = v }
func (m *TLInputStickeredMediaPhoto) GetId() *InputPhoto  { return m.Data2.Id_1 }

func (m *TLInputStickeredMediaPhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputStickeredMediaPhoto))

	x.Bytes(m.Data2Id.Encode())

	return x.buf
}

func (m *TLInputStickeredMediaPhoto) Decode(dbuf *DecodeBuf) error {
	m1 := &Id{}
	m1.Decode(dbuf)
	m.SetId(m1)

	return dbuf.err
}

// inputStickeredMediaDocument#438865b id:InputDocument = InputStickeredMedia;
func (m *TLInputStickeredMediaDocument) To_InputStickeredMedia() *InputStickeredMedia {
	return &InputStickeredMedia{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputStickeredMediaDocument) SetId(v *InputDocument) { m.Data2.Id_2 = v }
func (m *TLInputStickeredMediaDocument) GetId() *InputDocument  { return m.Data2.Id_2 }

func (m *TLInputStickeredMediaDocument) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputStickeredMediaDocument))

	x.Bytes(m.Data2Id.Encode())

	return x.buf
}

func (m *TLInputStickeredMediaDocument) Decode(dbuf *DecodeBuf) error {
	m1 := &Id{}
	m1.Decode(dbuf)
	m.SetId(m1)

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// InputWebFileLocation <--
//  + TL_InputWebFileLocation
//

func (m *InputWebFileLocation) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_inputWebFileLocation:
		t := m.To_InputWebFileLocation()
		return t.Encode()

	default:
		return []byte{}
	}
}

// inputWebFileLocation#c239d686 url:string access_hash:long = InputWebFileLocation;
func (m *InputWebFileLocation) To_InputWebFileLocation() *TLInputWebFileLocation {
	return &TLInputWebFileLocation{
		Data2: m.Data2,
	}
}

// inputWebFileLocation#c239d686 url:string access_hash:long = InputWebFileLocation;
func (m *TLInputWebFileLocation) To_InputWebFileLocation() *InputWebFileLocation {
	return &InputWebFileLocation{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputWebFileLocation) SetUrl(v string) { m.Data2.Url = v }
func (m *TLInputWebFileLocation) GetUrl() string  { return m.Data2.Url }

func (m *TLInputWebFileLocation) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLInputWebFileLocation) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLInputWebFileLocation) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputWebFileLocation))

	x.StringBytes(m.Data2.Url)
	x.Long(m.Data2.AccessHash)

	return x.buf
}

func (m *TLInputWebFileLocation) Decode(dbuf *DecodeBuf) error {
	m.Data2.Url = x.StringBytes()
	m.Data2.AccessHash = x.Long()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// InputPeerNotifySettings <--
//  + TL_InputPeerNotifySettings
//

func (m *InputPeerNotifySettings) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_inputPeerNotifySettings:
		t := m.To_InputPeerNotifySettings()
		return t.Encode()

	default:
		return []byte{}
	}
}

// inputPeerNotifySettings#38935eb2 flags:# show_previews:flags.0?true silent:flags.1?true mute_until:int sound:string = InputPeerNotifySettings;
func (m *InputPeerNotifySettings) To_InputPeerNotifySettings() *TLInputPeerNotifySettings {
	return &TLInputPeerNotifySettings{
		Data2: m.Data2,
	}
}

// inputPeerNotifySettings#38935eb2 flags:# show_previews:flags.0?true silent:flags.1?true mute_until:int sound:string = InputPeerNotifySettings;
func (m *TLInputPeerNotifySettings) To_InputPeerNotifySettings() *InputPeerNotifySettings {
	return &InputPeerNotifySettings{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputPeerNotifySettings) SetShowPreviews(v bool) { m.Data2.ShowPreviews = v }
func (m *TLInputPeerNotifySettings) GetShowPreviews() bool  { return m.Data2.ShowPreviews }

func (m *TLInputPeerNotifySettings) SetSilent(v bool) { m.Data2.Silent = v }
func (m *TLInputPeerNotifySettings) GetSilent() bool  { return m.Data2.Silent }

func (m *TLInputPeerNotifySettings) SetMuteUntil(v int32) { m.Data2.MuteUntil = v }
func (m *TLInputPeerNotifySettings) GetMuteUntil() int32  { return m.Data2.MuteUntil }

func (m *TLInputPeerNotifySettings) SetSound(v string) { m.Data2.Sound = v }
func (m *TLInputPeerNotifySettings) GetSound() string  { return m.Data2.Sound }

func (m *TLInputPeerNotifySettings) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputPeerNotifySettings))

	// flags
	var flags uint32 = 0
	if m.GetShowPreviews() == true {
		flags |= 1 << 1
	}
	if m.GetSilent() == true {
		flags |= 1 << 2
	}
	x.UInt(flags)

	x.Int(m.Data2.MuteUntil)
	x.StringBytes(m.Data2.Sound)

	return x.buf
}

func (m *TLInputPeerNotifySettings) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m.Data2.ShowPreviews = true
	}
	if (flags & (1 << 3)) != 0 {
		m.Data2.Silent = true
	}
	m.Data2.MuteUntil = x.Int()
	m.Data2.Sound = x.StringBytes()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Photos_Photos <--
//  + TL_PhotosPhotos
//  + TL_PhotosPhotosSlice
//

func (m *Photos_Photos) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_photos_photos:
		t := m.To_PhotosPhotos()
		return t.Encode()
	case TLConstructor_CRC32_photos_photosSlice:
		t := m.To_PhotosPhotosSlice()
		return t.Encode()

	default:
		return []byte{}
	}
}

// photos.photos#8dca6aa5 photos:Vector<Photo> users:Vector<User> = photos.Photos;
func (m *Photos_Photos) To_PhotosPhotos() *TLPhotosPhotos {
	return &TLPhotosPhotos{
		Data2: m.Data2,
	}
}

// photos.photosSlice#15051f54 count:int photos:Vector<Photo> users:Vector<User> = photos.Photos;
func (m *Photos_Photos) To_PhotosPhotosSlice() *TLPhotosPhotosSlice {
	return &TLPhotosPhotosSlice{
		Data2: m.Data2,
	}
}

// photos.photos#8dca6aa5 photos:Vector<Photo> users:Vector<User> = photos.Photos;
func (m *TLPhotosPhotos) To_Photos_Photos() *Photos_Photos {
	return &Photos_Photos{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPhotosPhotos) SetPhotos(v []*Photo) { m.Data2.Photos = v }
func (m *TLPhotosPhotos) GetPhotos() []*Photo  { return m.Data2.Photos }

func (m *TLPhotosPhotos) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLPhotosPhotos) GetUsers() []*User  { return m.Data2.Users }

func (m *TLPhotosPhotos) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_photos_photos))

	return x.buf
}

func (m *TLPhotosPhotos) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// photos.photosSlice#15051f54 count:int photos:Vector<Photo> users:Vector<User> = photos.Photos;
func (m *TLPhotosPhotosSlice) To_Photos_Photos() *Photos_Photos {
	return &Photos_Photos{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPhotosPhotosSlice) SetCount(v int32) { m.Data2.Count = v }
func (m *TLPhotosPhotosSlice) GetCount() int32  { return m.Data2.Count }

func (m *TLPhotosPhotosSlice) SetPhotos(v []*Photo) { m.Data2.Photos = v }
func (m *TLPhotosPhotosSlice) GetPhotos() []*Photo  { return m.Data2.Photos }

func (m *TLPhotosPhotosSlice) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLPhotosPhotosSlice) GetUsers() []*User  { return m.Data2.Users }

func (m *TLPhotosPhotosSlice) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_photos_photosSlice))

	x.Int(m.Data2.Count)

	return x.buf
}

func (m *TLPhotosPhotosSlice) Decode(dbuf *DecodeBuf) error {
	m.Data2.Count = x.Int()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// ReplyMarkup <--
//  + TL_ReplyKeyboardHide
//  + TL_ReplyKeyboardForceReply
//  + TL_ReplyKeyboardMarkup
//  + TL_ReplyInlineMarkup
//

func (m *ReplyMarkup) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_replyKeyboardHide:
		t := m.To_ReplyKeyboardHide()
		return t.Encode()
	case TLConstructor_CRC32_replyKeyboardForceReply:
		t := m.To_ReplyKeyboardForceReply()
		return t.Encode()
	case TLConstructor_CRC32_replyKeyboardMarkup:
		t := m.To_ReplyKeyboardMarkup()
		return t.Encode()
	case TLConstructor_CRC32_replyInlineMarkup:
		t := m.To_ReplyInlineMarkup()
		return t.Encode()

	default:
		return []byte{}
	}
}

// replyKeyboardHide#a03e5b85 flags:# selective:flags.2?true = ReplyMarkup;
func (m *ReplyMarkup) To_ReplyKeyboardHide() *TLReplyKeyboardHide {
	return &TLReplyKeyboardHide{
		Data2: m.Data2,
	}
}

// replyKeyboardForceReply#f4108aa0 flags:# single_use:flags.1?true selective:flags.2?true = ReplyMarkup;
func (m *ReplyMarkup) To_ReplyKeyboardForceReply() *TLReplyKeyboardForceReply {
	return &TLReplyKeyboardForceReply{
		Data2: m.Data2,
	}
}

// replyKeyboardMarkup#3502758c flags:# resize:flags.0?true single_use:flags.1?true selective:flags.2?true rows:Vector<KeyboardButtonRow> = ReplyMarkup;
func (m *ReplyMarkup) To_ReplyKeyboardMarkup() *TLReplyKeyboardMarkup {
	return &TLReplyKeyboardMarkup{
		Data2: m.Data2,
	}
}

// replyInlineMarkup#48a30254 rows:Vector<KeyboardButtonRow> = ReplyMarkup;
func (m *ReplyMarkup) To_ReplyInlineMarkup() *TLReplyInlineMarkup {
	return &TLReplyInlineMarkup{
		Data2: m.Data2,
	}
}

// replyKeyboardHide#a03e5b85 flags:# selective:flags.2?true = ReplyMarkup;
func (m *TLReplyKeyboardHide) To_ReplyMarkup() *ReplyMarkup {
	return &ReplyMarkup{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLReplyKeyboardHide) SetSelective(v bool) { m.Data2.Selective = v }
func (m *TLReplyKeyboardHide) GetSelective() bool  { return m.Data2.Selective }

func (m *TLReplyKeyboardHide) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_replyKeyboardHide))

	// flags
	var flags uint32 = 0
	if m.GetSelective() == true {
		flags |= 1 << 1
	}
	x.UInt(flags)

	return x.buf
}

func (m *TLReplyKeyboardHide) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m.Data2.Selective = true
	}

	return dbuf.err
}

// replyKeyboardForceReply#f4108aa0 flags:# single_use:flags.1?true selective:flags.2?true = ReplyMarkup;
func (m *TLReplyKeyboardForceReply) To_ReplyMarkup() *ReplyMarkup {
	return &ReplyMarkup{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLReplyKeyboardForceReply) SetSingleUse(v bool) { m.Data2.SingleUse = v }
func (m *TLReplyKeyboardForceReply) GetSingleUse() bool  { return m.Data2.SingleUse }

func (m *TLReplyKeyboardForceReply) SetSelective(v bool) { m.Data2.Selective = v }
func (m *TLReplyKeyboardForceReply) GetSelective() bool  { return m.Data2.Selective }

func (m *TLReplyKeyboardForceReply) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_replyKeyboardForceReply))

	// flags
	var flags uint32 = 0
	if m.GetSingleUse() == true {
		flags |= 1 << 1
	}
	if m.GetSelective() == true {
		flags |= 1 << 2
	}
	x.UInt(flags)

	return x.buf
}

func (m *TLReplyKeyboardForceReply) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m.Data2.SingleUse = true
	}
	if (flags & (1 << 3)) != 0 {
		m.Data2.Selective = true
	}

	return dbuf.err
}

// replyKeyboardMarkup#3502758c flags:# resize:flags.0?true single_use:flags.1?true selective:flags.2?true rows:Vector<KeyboardButtonRow> = ReplyMarkup;
func (m *TLReplyKeyboardMarkup) To_ReplyMarkup() *ReplyMarkup {
	return &ReplyMarkup{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLReplyKeyboardMarkup) SetResize(v bool) { m.Data2.Resize = v }
func (m *TLReplyKeyboardMarkup) GetResize() bool  { return m.Data2.Resize }

func (m *TLReplyKeyboardMarkup) SetSingleUse(v bool) { m.Data2.SingleUse = v }
func (m *TLReplyKeyboardMarkup) GetSingleUse() bool  { return m.Data2.SingleUse }

func (m *TLReplyKeyboardMarkup) SetSelective(v bool) { m.Data2.Selective = v }
func (m *TLReplyKeyboardMarkup) GetSelective() bool  { return m.Data2.Selective }

func (m *TLReplyKeyboardMarkup) SetRows(v []*KeyboardButtonRow) { m.Data2.Rows = v }
func (m *TLReplyKeyboardMarkup) GetRows() []*KeyboardButtonRow  { return m.Data2.Rows }

func (m *TLReplyKeyboardMarkup) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_replyKeyboardMarkup))

	// flags
	var flags uint32 = 0
	if m.GetResize() == true {
		flags |= 1 << 1
	}
	if m.GetSingleUse() == true {
		flags |= 1 << 2
	}
	if m.GetSelective() == true {
		flags |= 1 << 3
	}
	x.UInt(flags)

	return x.buf
}

func (m *TLReplyKeyboardMarkup) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m.Data2.Resize = true
	}
	if (flags & (1 << 3)) != 0 {
		m.Data2.SingleUse = true
	}
	if (flags & (1 << 4)) != 0 {
		m.Data2.Selective = true
	}

	return dbuf.err
}

// replyInlineMarkup#48a30254 rows:Vector<KeyboardButtonRow> = ReplyMarkup;
func (m *TLReplyInlineMarkup) To_ReplyMarkup() *ReplyMarkup {
	return &ReplyMarkup{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLReplyInlineMarkup) SetRows(v []*KeyboardButtonRow) { m.Data2.Rows = v }
func (m *TLReplyInlineMarkup) GetRows() []*KeyboardButtonRow  { return m.Data2.Rows }

func (m *TLReplyInlineMarkup) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_replyInlineMarkup))

	return x.buf
}

func (m *TLReplyInlineMarkup) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Contacts_TopPeers <--
//  + TL_ContactsTopPeersNotModified
//  + TL_ContactsTopPeers
//

func (m *Contacts_TopPeers) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_contacts_topPeersNotModified:
		t := m.To_ContactsTopPeersNotModified()
		return t.Encode()
	case TLConstructor_CRC32_contacts_topPeers:
		t := m.To_ContactsTopPeers()
		return t.Encode()

	default:
		return []byte{}
	}
}

// contacts.topPeersNotModified#de266ef5 = contacts.TopPeers;
func (m *Contacts_TopPeers) To_ContactsTopPeersNotModified() *TLContactsTopPeersNotModified {
	return &TLContactsTopPeersNotModified{
		Data2: m.Data2,
	}
}

// contacts.topPeers#70b772a8 categories:Vector<TopPeerCategoryPeers> chats:Vector<Chat> users:Vector<User> = contacts.TopPeers;
func (m *Contacts_TopPeers) To_ContactsTopPeers() *TLContactsTopPeers {
	return &TLContactsTopPeers{
		Data2: m.Data2,
	}
}

// contacts.topPeersNotModified#de266ef5 = contacts.TopPeers;
func (m *TLContactsTopPeersNotModified) To_Contacts_TopPeers() *Contacts_TopPeers {
	return &Contacts_TopPeers{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLContactsTopPeersNotModified) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_contacts_topPeersNotModified))

	return x.buf
}

func (m *TLContactsTopPeersNotModified) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// contacts.topPeers#70b772a8 categories:Vector<TopPeerCategoryPeers> chats:Vector<Chat> users:Vector<User> = contacts.TopPeers;
func (m *TLContactsTopPeers) To_Contacts_TopPeers() *Contacts_TopPeers {
	return &Contacts_TopPeers{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLContactsTopPeers) SetCategories(v []*TopPeerCategoryPeers) { m.Data2.Categories = v }
func (m *TLContactsTopPeers) GetCategories() []*TopPeerCategoryPeers  { return m.Data2.Categories }

func (m *TLContactsTopPeers) SetChats(v []*Chat) { m.Data2.Chats = v }
func (m *TLContactsTopPeers) GetChats() []*Chat  { return m.Data2.Chats }

func (m *TLContactsTopPeers) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLContactsTopPeers) GetUsers() []*User  { return m.Data2.Users }

func (m *TLContactsTopPeers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_contacts_topPeers))

	return x.buf
}

func (m *TLContactsTopPeers) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Messages_FavedStickers <--
//  + TL_MessagesFavedStickersNotModified
//  + TL_MessagesFavedStickers
//

func (m *Messages_FavedStickers) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_messages_favedStickersNotModified:
		t := m.To_MessagesFavedStickersNotModified()
		return t.Encode()
	case TLConstructor_CRC32_messages_favedStickers:
		t := m.To_MessagesFavedStickers()
		return t.Encode()

	default:
		return []byte{}
	}
}

// messages.favedStickersNotModified#9e8fa6d3 = messages.FavedStickers;
func (m *Messages_FavedStickers) To_MessagesFavedStickersNotModified() *TLMessagesFavedStickersNotModified {
	return &TLMessagesFavedStickersNotModified{
		Data2: m.Data2,
	}
}

// messages.favedStickers#f37f2f16 hash:int packs:Vector<StickerPack> stickers:Vector<Document> = messages.FavedStickers;
func (m *Messages_FavedStickers) To_MessagesFavedStickers() *TLMessagesFavedStickers {
	return &TLMessagesFavedStickers{
		Data2: m.Data2,
	}
}

// messages.favedStickersNotModified#9e8fa6d3 = messages.FavedStickers;
func (m *TLMessagesFavedStickersNotModified) To_Messages_FavedStickers() *Messages_FavedStickers {
	return &Messages_FavedStickers{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessagesFavedStickersNotModified) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_favedStickersNotModified))

	return x.buf
}

func (m *TLMessagesFavedStickersNotModified) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// messages.favedStickers#f37f2f16 hash:int packs:Vector<StickerPack> stickers:Vector<Document> = messages.FavedStickers;
func (m *TLMessagesFavedStickers) To_Messages_FavedStickers() *Messages_FavedStickers {
	return &Messages_FavedStickers{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessagesFavedStickers) SetHash(v int32) { m.Data2.Hash = v }
func (m *TLMessagesFavedStickers) GetHash() int32  { return m.Data2.Hash }

func (m *TLMessagesFavedStickers) SetPacks(v []*StickerPack) { m.Data2.Packs = v }
func (m *TLMessagesFavedStickers) GetPacks() []*StickerPack  { return m.Data2.Packs }

func (m *TLMessagesFavedStickers) SetStickers(v []*Document) { m.Data2.Stickers = v }
func (m *TLMessagesFavedStickers) GetStickers() []*Document  { return m.Data2.Stickers }

func (m *TLMessagesFavedStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_favedStickers))

	x.Int(m.Data2.Hash)

	return x.buf
}

func (m *TLMessagesFavedStickers) Decode(dbuf *DecodeBuf) error {
	m.Data2.Hash = x.Int()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// True <--
//  + TL_True
//

func (m *True) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_true:
		t := m.To_True()
		return t.Encode()

	default:
		return []byte{}
	}
}

// true#3fedd339 = True;
func (m *True) To_True() *TLTrue {
	return &TLTrue{
		Data2: m.Data2,
	}
}

// true#3fedd339 = True;
func (m *TLTrue) To_True() *True {
	return &True{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLTrue) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_true))

	return x.buf
}

func (m *TLTrue) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// ImportedContact <--
//  + TL_ImportedContact
//

func (m *ImportedContact) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_importedContact:
		t := m.To_ImportedContact()
		return t.Encode()

	default:
		return []byte{}
	}
}

// importedContact#d0028438 user_id:int client_id:long = ImportedContact;
func (m *ImportedContact) To_ImportedContact() *TLImportedContact {
	return &TLImportedContact{
		Data2: m.Data2,
	}
}

// importedContact#d0028438 user_id:int client_id:long = ImportedContact;
func (m *TLImportedContact) To_ImportedContact() *ImportedContact {
	return &ImportedContact{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLImportedContact) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLImportedContact) GetUserId() int32  { return m.Data2.UserId }

func (m *TLImportedContact) SetClientId(v int64) { m.Data2.ClientId = v }
func (m *TLImportedContact) GetClientId() int64  { return m.Data2.ClientId }

func (m *TLImportedContact) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_importedContact))

	x.Int(m.Data2.UserId)
	x.Long(m.Data2.ClientId)

	return x.buf
}

func (m *TLImportedContact) Decode(dbuf *DecodeBuf) error {
	m.Data2.UserId = x.Int()
	m.Data2.ClientId = x.Long()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// NearestDc <--
//  + TL_NearestDc
//

func (m *NearestDc) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_nearestDc:
		t := m.To_NearestDc()
		return t.Encode()

	default:
		return []byte{}
	}
}

// nearestDc#8e1a1775 country:string this_dc:int nearest_dc:int = NearestDc;
func (m *NearestDc) To_NearestDc() *TLNearestDc {
	return &TLNearestDc{
		Data2: m.Data2,
	}
}

// nearestDc#8e1a1775 country:string this_dc:int nearest_dc:int = NearestDc;
func (m *TLNearestDc) To_NearestDc() *NearestDc {
	return &NearestDc{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLNearestDc) SetCountry(v string) { m.Data2.Country = v }
func (m *TLNearestDc) GetCountry() string  { return m.Data2.Country }

func (m *TLNearestDc) SetThisDc(v int32) { m.Data2.ThisDc = v }
func (m *TLNearestDc) GetThisDc() int32  { return m.Data2.ThisDc }

func (m *TLNearestDc) SetNearestDc(v int32) { m.Data2.NearestDc = v }
func (m *TLNearestDc) GetNearestDc() int32  { return m.Data2.NearestDc }

func (m *TLNearestDc) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_nearestDc))

	x.StringBytes(m.Data2.Country)
	x.Int(m.Data2.ThisDc)
	x.Int(m.Data2.NearestDc)

	return x.buf
}

func (m *TLNearestDc) Decode(dbuf *DecodeBuf) error {
	m.Data2.Country = x.StringBytes()
	m.Data2.ThisDc = x.Int()
	m.Data2.NearestDc = x.Int()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// ChannelBannedRights <--
//  + TL_ChannelBannedRights
//

func (m *ChannelBannedRights) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_channelBannedRights:
		t := m.To_ChannelBannedRights()
		return t.Encode()

	default:
		return []byte{}
	}
}

// channelBannedRights#58cf4249 flags:# view_messages:flags.0?true send_messages:flags.1?true send_media:flags.2?true send_stickers:flags.3?true send_gifs:flags.4?true send_games:flags.5?true send_inline:flags.6?true embed_links:flags.7?true until_date:int = ChannelBannedRights;
func (m *ChannelBannedRights) To_ChannelBannedRights() *TLChannelBannedRights {
	return &TLChannelBannedRights{
		Data2: m.Data2,
	}
}

// channelBannedRights#58cf4249 flags:# view_messages:flags.0?true send_messages:flags.1?true send_media:flags.2?true send_stickers:flags.3?true send_gifs:flags.4?true send_games:flags.5?true send_inline:flags.6?true embed_links:flags.7?true until_date:int = ChannelBannedRights;
func (m *TLChannelBannedRights) To_ChannelBannedRights() *ChannelBannedRights {
	return &ChannelBannedRights{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLChannelBannedRights) SetViewMessages(v bool) { m.Data2.ViewMessages = v }
func (m *TLChannelBannedRights) GetViewMessages() bool  { return m.Data2.ViewMessages }

func (m *TLChannelBannedRights) SetSendMessages(v bool) { m.Data2.SendMessages = v }
func (m *TLChannelBannedRights) GetSendMessages() bool  { return m.Data2.SendMessages }

func (m *TLChannelBannedRights) SetSendMedia(v bool) { m.Data2.SendMedia = v }
func (m *TLChannelBannedRights) GetSendMedia() bool  { return m.Data2.SendMedia }

func (m *TLChannelBannedRights) SetSendStickers(v bool) { m.Data2.SendStickers = v }
func (m *TLChannelBannedRights) GetSendStickers() bool  { return m.Data2.SendStickers }

func (m *TLChannelBannedRights) SetSendGifs(v bool) { m.Data2.SendGifs = v }
func (m *TLChannelBannedRights) GetSendGifs() bool  { return m.Data2.SendGifs }

func (m *TLChannelBannedRights) SetSendGames(v bool) { m.Data2.SendGames = v }
func (m *TLChannelBannedRights) GetSendGames() bool  { return m.Data2.SendGames }

func (m *TLChannelBannedRights) SetSendInline(v bool) { m.Data2.SendInline = v }
func (m *TLChannelBannedRights) GetSendInline() bool  { return m.Data2.SendInline }

func (m *TLChannelBannedRights) SetEmbedLinks(v bool) { m.Data2.EmbedLinks = v }
func (m *TLChannelBannedRights) GetEmbedLinks() bool  { return m.Data2.EmbedLinks }

func (m *TLChannelBannedRights) SetUntilDate(v int32) { m.Data2.UntilDate = v }
func (m *TLChannelBannedRights) GetUntilDate() int32  { return m.Data2.UntilDate }

func (m *TLChannelBannedRights) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelBannedRights))

	// flags
	var flags uint32 = 0
	if m.GetViewMessages() == true {
		flags |= 1 << 1
	}
	if m.GetSendMessages() == true {
		flags |= 1 << 2
	}
	if m.GetSendMedia() == true {
		flags |= 1 << 3
	}
	if m.GetSendStickers() == true {
		flags |= 1 << 4
	}
	if m.GetSendGifs() == true {
		flags |= 1 << 5
	}
	if m.GetSendGames() == true {
		flags |= 1 << 6
	}
	if m.GetSendInline() == true {
		flags |= 1 << 7
	}
	if m.GetEmbedLinks() == true {
		flags |= 1 << 8
	}
	x.UInt(flags)

	x.Int(m.Data2.UntilDate)

	return x.buf
}

func (m *TLChannelBannedRights) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m.Data2.ViewMessages = true
	}
	if (flags & (1 << 3)) != 0 {
		m.Data2.SendMessages = true
	}
	if (flags & (1 << 4)) != 0 {
		m.Data2.SendMedia = true
	}
	if (flags & (1 << 5)) != 0 {
		m.Data2.SendStickers = true
	}
	if (flags & (1 << 6)) != 0 {
		m.Data2.SendGifs = true
	}
	if (flags & (1 << 7)) != 0 {
		m.Data2.SendGames = true
	}
	if (flags & (1 << 8)) != 0 {
		m.Data2.SendInline = true
	}
	if (flags & (1 << 9)) != 0 {
		m.Data2.EmbedLinks = true
	}
	m.Data2.UntilDate = x.Int()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Account_PasswordSettings <--
//  + TL_AccountPasswordSettings
//

func (m *Account_PasswordSettings) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_account_passwordSettings:
		t := m.To_AccountPasswordSettings()
		return t.Encode()

	default:
		return []byte{}
	}
}

// account.passwordSettings#b7b72ab3 email:string = account.PasswordSettings;
func (m *Account_PasswordSettings) To_AccountPasswordSettings() *TLAccountPasswordSettings {
	return &TLAccountPasswordSettings{
		Data2: m.Data2,
	}
}

// account.passwordSettings#b7b72ab3 email:string = account.PasswordSettings;
func (m *TLAccountPasswordSettings) To_Account_PasswordSettings() *Account_PasswordSettings {
	return &Account_PasswordSettings{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLAccountPasswordSettings) SetEmail(v string) { m.Data2.Email = v }
func (m *TLAccountPasswordSettings) GetEmail() string  { return m.Data2.Email }

func (m *TLAccountPasswordSettings) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_account_passwordSettings))

	x.StringBytes(m.Data2.Email)

	return x.buf
}

func (m *TLAccountPasswordSettings) Decode(dbuf *DecodeBuf) error {
	m.Data2.Email = x.StringBytes()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Auth_PasswordRecovery <--
//  + TL_AuthPasswordRecovery
//

func (m *Auth_PasswordRecovery) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_auth_passwordRecovery:
		t := m.To_AuthPasswordRecovery()
		return t.Encode()

	default:
		return []byte{}
	}
}

// auth.passwordRecovery#137948a5 email_pattern:string = auth.PasswordRecovery;
func (m *Auth_PasswordRecovery) To_AuthPasswordRecovery() *TLAuthPasswordRecovery {
	return &TLAuthPasswordRecovery{
		Data2: m.Data2,
	}
}

// auth.passwordRecovery#137948a5 email_pattern:string = auth.PasswordRecovery;
func (m *TLAuthPasswordRecovery) To_Auth_PasswordRecovery() *Auth_PasswordRecovery {
	return &Auth_PasswordRecovery{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLAuthPasswordRecovery) SetEmailPattern(v string) { m.Data2.EmailPattern = v }
func (m *TLAuthPasswordRecovery) GetEmailPattern() string  { return m.Data2.EmailPattern }

func (m *TLAuthPasswordRecovery) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_auth_passwordRecovery))

	x.StringBytes(m.Data2.EmailPattern)

	return x.buf
}

func (m *TLAuthPasswordRecovery) Decode(dbuf *DecodeBuf) error {
	m.Data2.EmailPattern = x.StringBytes()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Payments_PaymentForm <--
//  + TL_PaymentsPaymentForm
//

func (m *Payments_PaymentForm) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_payments_paymentForm:
		t := m.To_PaymentsPaymentForm()
		return t.Encode()

	default:
		return []byte{}
	}
}

// payments.paymentForm#3f56aea3 flags:# can_save_credentials:flags.2?true password_missing:flags.3?true bot_id:int invoice:Invoice provider_id:int url:string native_provider:flags.4?string native_params:flags.4?DataJSON saved_info:flags.0?PaymentRequestedInfo saved_credentials:flags.1?PaymentSavedCredentials users:Vector<User> = payments.PaymentForm;
func (m *Payments_PaymentForm) To_PaymentsPaymentForm() *TLPaymentsPaymentForm {
	return &TLPaymentsPaymentForm{
		Data2: m.Data2,
	}
}

// payments.paymentForm#3f56aea3 flags:# can_save_credentials:flags.2?true password_missing:flags.3?true bot_id:int invoice:Invoice provider_id:int url:string native_provider:flags.4?string native_params:flags.4?DataJSON saved_info:flags.0?PaymentRequestedInfo saved_credentials:flags.1?PaymentSavedCredentials users:Vector<User> = payments.PaymentForm;
func (m *TLPaymentsPaymentForm) To_Payments_PaymentForm() *Payments_PaymentForm {
	return &Payments_PaymentForm{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPaymentsPaymentForm) SetCanSaveCredentials(v bool) { m.Data2.CanSaveCredentials = v }
func (m *TLPaymentsPaymentForm) GetCanSaveCredentials() bool  { return m.Data2.CanSaveCredentials }

func (m *TLPaymentsPaymentForm) SetPasswordMissing(v bool) { m.Data2.PasswordMissing = v }
func (m *TLPaymentsPaymentForm) GetPasswordMissing() bool  { return m.Data2.PasswordMissing }

func (m *TLPaymentsPaymentForm) SetBotId(v int32) { m.Data2.BotId = v }
func (m *TLPaymentsPaymentForm) GetBotId() int32  { return m.Data2.BotId }

func (m *TLPaymentsPaymentForm) SetInvoice(v *Invoice) { m.Data2.Invoice = v }
func (m *TLPaymentsPaymentForm) GetInvoice() *Invoice  { return m.Data2.Invoice }

func (m *TLPaymentsPaymentForm) SetProviderId(v int32) { m.Data2.ProviderId = v }
func (m *TLPaymentsPaymentForm) GetProviderId() int32  { return m.Data2.ProviderId }

func (m *TLPaymentsPaymentForm) SetUrl(v string) { m.Data2.Url = v }
func (m *TLPaymentsPaymentForm) GetUrl() string  { return m.Data2.Url }

func (m *TLPaymentsPaymentForm) SetNativeProvider(v string) { m.Data2.NativeProvider = v }
func (m *TLPaymentsPaymentForm) GetNativeProvider() string  { return m.Data2.NativeProvider }

func (m *TLPaymentsPaymentForm) SetNativeParams(v *DataJSON) { m.Data2.NativeParams = v }
func (m *TLPaymentsPaymentForm) GetNativeParams() *DataJSON  { return m.Data2.NativeParams }

func (m *TLPaymentsPaymentForm) SetSavedInfo(v *PaymentRequestedInfo) { m.Data2.SavedInfo = v }
func (m *TLPaymentsPaymentForm) GetSavedInfo() *PaymentRequestedInfo  { return m.Data2.SavedInfo }

func (m *TLPaymentsPaymentForm) SetSavedCredentials(v *PaymentSavedCredentials) {
	m.Data2.SavedCredentials = v
}
func (m *TLPaymentsPaymentForm) GetSavedCredentials() *PaymentSavedCredentials {
	return m.Data2.SavedCredentials
}

func (m *TLPaymentsPaymentForm) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLPaymentsPaymentForm) GetUsers() []*User  { return m.Data2.Users }

func (m *TLPaymentsPaymentForm) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_payments_paymentForm))

	// flags
	var flags uint32 = 0
	if m.GetCanSaveCredentials() == true {
		flags |= 1 << 1
	}
	if m.GetPasswordMissing() == true {
		flags |= 1 << 2
	}
	if m.GetNativeProvider() != "" {
		flags |= 1 << 7
	}
	if m.GetNativeParams() != nil {
		flags |= 1 << 8
	}
	if m.GetSavedInfo() != nil {
		flags |= 1 << 9
	}
	if m.GetSavedCredentials() != nil {
		flags |= 1 << 10
	}
	x.UInt(flags)

	x.Int(m.Data2.BotId)
	x.Bytes(m.Data2Invoice.Encode())
	x.Int(m.Data2.ProviderId)
	x.StringBytes(m.Data2.Url)
	if m.GetNativeProvider() != 0 {
		x.StringBytes(m.Data2.NativeProvider)
	}
	if m.GetNativeParams() != 0 {
		x.Bytes(m.Data2NativeParams.Encode())
	}
	if m.GetSavedInfo() != 0 {
		x.Bytes(m.Data2SavedInfo.Encode())
	}
	if m.GetSavedCredentials() != 0 {
		x.Bytes(m.Data2SavedCredentials.Encode())
	}

	return x.buf
}

func (m *TLPaymentsPaymentForm) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m.Data2.CanSaveCredentials = true
	}
	if (flags & (1 << 3)) != 0 {
		m.Data2.PasswordMissing = true
	}
	m.Data2.BotId = x.Int()
	m5 := &Invoice{}
	m5.Decode(dbuf)
	m.SetInvoice(m5)
	m.Data2.ProviderId = x.Int()
	m.Data2.Url = x.StringBytes()
	if (flags & (1 << 8)) != 0 {
		m.Data2.NativeProvider = x.StringBytes()
	}
	if (flags & (1 << 9)) != 0 {
		m9 := &NativeParams{}
		m9.Decode(dbuf)
		m.SetNativeParams(m9)
	}
	if (flags & (1 << 10)) != 0 {
		m10 := &SavedInfo{}
		m10.Decode(dbuf)
		m.SetSavedInfo(m10)
	}
	if (flags & (1 << 11)) != 0 {
		m11 := &SavedCredentials{}
		m11.Decode(dbuf)
		m.SetSavedCredentials(m11)
	}

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// RichText <--
//  + TL_TextEmpty
//  + TL_TextPlain
//  + TL_TextBold
//  + TL_TextItalic
//  + TL_TextUnderline
//  + TL_TextStrike
//  + TL_TextFixed
//  + TL_TextUrl
//  + TL_TextEmail
//  + TL_TextConcat
//

func (m *RichText) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_textEmpty:
		t := m.To_TextEmpty()
		return t.Encode()
	case TLConstructor_CRC32_textPlain:
		t := m.To_TextPlain()
		return t.Encode()
	case TLConstructor_CRC32_textBold:
		t := m.To_TextBold()
		return t.Encode()
	case TLConstructor_CRC32_textItalic:
		t := m.To_TextItalic()
		return t.Encode()
	case TLConstructor_CRC32_textUnderline:
		t := m.To_TextUnderline()
		return t.Encode()
	case TLConstructor_CRC32_textStrike:
		t := m.To_TextStrike()
		return t.Encode()
	case TLConstructor_CRC32_textFixed:
		t := m.To_TextFixed()
		return t.Encode()
	case TLConstructor_CRC32_textUrl:
		t := m.To_TextUrl()
		return t.Encode()
	case TLConstructor_CRC32_textEmail:
		t := m.To_TextEmail()
		return t.Encode()
	case TLConstructor_CRC32_textConcat:
		t := m.To_TextConcat()
		return t.Encode()

	default:
		return []byte{}
	}
}

// textEmpty#dc3d824f = RichText;
func (m *RichText) To_TextEmpty() *TLTextEmpty {
	return &TLTextEmpty{
		Data2: m.Data2,
	}
}

// textPlain#744694e0 text:string = RichText;
func (m *RichText) To_TextPlain() *TLTextPlain {
	return &TLTextPlain{
		Data2: m.Data2,
	}
}

// textBold#6724abc4 text:RichText = RichText;
func (m *RichText) To_TextBold() *TLTextBold {
	return &TLTextBold{
		Data2: m.Data2,
	}
}

// textItalic#d912a59c text:RichText = RichText;
func (m *RichText) To_TextItalic() *TLTextItalic {
	return &TLTextItalic{
		Data2: m.Data2,
	}
}

// textUnderline#c12622c4 text:RichText = RichText;
func (m *RichText) To_TextUnderline() *TLTextUnderline {
	return &TLTextUnderline{
		Data2: m.Data2,
	}
}

// textStrike#9bf8bb95 text:RichText = RichText;
func (m *RichText) To_TextStrike() *TLTextStrike {
	return &TLTextStrike{
		Data2: m.Data2,
	}
}

// textFixed#6c3f19b9 text:RichText = RichText;
func (m *RichText) To_TextFixed() *TLTextFixed {
	return &TLTextFixed{
		Data2: m.Data2,
	}
}

// textUrl#3c2884c1 text:RichText url:string webpage_id:long = RichText;
func (m *RichText) To_TextUrl() *TLTextUrl {
	return &TLTextUrl{
		Data2: m.Data2,
	}
}

// textEmail#de5a0dd6 text:RichText email:string = RichText;
func (m *RichText) To_TextEmail() *TLTextEmail {
	return &TLTextEmail{
		Data2: m.Data2,
	}
}

// textConcat#7e6260d7 texts:Vector<RichText> = RichText;
func (m *RichText) To_TextConcat() *TLTextConcat {
	return &TLTextConcat{
		Data2: m.Data2,
	}
}

// textEmpty#dc3d824f = RichText;
func (m *TLTextEmpty) To_RichText() *RichText {
	return &RichText{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLTextEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_textEmpty))

	return x.buf
}

func (m *TLTextEmpty) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// textPlain#744694e0 text:string = RichText;
func (m *TLTextPlain) To_RichText() *RichText {
	return &RichText{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLTextPlain) SetText(v string) { m.Data2.Text_1 = v }
func (m *TLTextPlain) GetText() string  { return m.Data2.Text_1 }

func (m *TLTextPlain) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_textPlain))

	x.StringBytes(m.Data2.Text)

	return x.buf
}

func (m *TLTextPlain) Decode(dbuf *DecodeBuf) error {
	m.Data2.Text = x.StringBytes()

	return dbuf.err
}

// textBold#6724abc4 text:RichText = RichText;
func (m *TLTextBold) To_RichText() *RichText {
	return &RichText{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLTextBold) SetText(v *RichText) { m.Data2.Text_2 = v }
func (m *TLTextBold) GetText() *RichText  { return m.Data2.Text_2 }

func (m *TLTextBold) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_textBold))

	x.Bytes(m.Data2Text.Encode())

	return x.buf
}

func (m *TLTextBold) Decode(dbuf *DecodeBuf) error {
	m1 := &Text{}
	m1.Decode(dbuf)
	m.SetText(m1)

	return dbuf.err
}

// textItalic#d912a59c text:RichText = RichText;
func (m *TLTextItalic) To_RichText() *RichText {
	return &RichText{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLTextItalic) SetText(v *RichText) { m.Data2.Text_2 = v }
func (m *TLTextItalic) GetText() *RichText  { return m.Data2.Text_2 }

func (m *TLTextItalic) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_textItalic))

	x.Bytes(m.Data2Text.Encode())

	return x.buf
}

func (m *TLTextItalic) Decode(dbuf *DecodeBuf) error {
	m1 := &Text{}
	m1.Decode(dbuf)
	m.SetText(m1)

	return dbuf.err
}

// textUnderline#c12622c4 text:RichText = RichText;
func (m *TLTextUnderline) To_RichText() *RichText {
	return &RichText{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLTextUnderline) SetText(v *RichText) { m.Data2.Text_2 = v }
func (m *TLTextUnderline) GetText() *RichText  { return m.Data2.Text_2 }

func (m *TLTextUnderline) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_textUnderline))

	x.Bytes(m.Data2Text.Encode())

	return x.buf
}

func (m *TLTextUnderline) Decode(dbuf *DecodeBuf) error {
	m1 := &Text{}
	m1.Decode(dbuf)
	m.SetText(m1)

	return dbuf.err
}

// textStrike#9bf8bb95 text:RichText = RichText;
func (m *TLTextStrike) To_RichText() *RichText {
	return &RichText{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLTextStrike) SetText(v *RichText) { m.Data2.Text_2 = v }
func (m *TLTextStrike) GetText() *RichText  { return m.Data2.Text_2 }

func (m *TLTextStrike) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_textStrike))

	x.Bytes(m.Data2Text.Encode())

	return x.buf
}

func (m *TLTextStrike) Decode(dbuf *DecodeBuf) error {
	m1 := &Text{}
	m1.Decode(dbuf)
	m.SetText(m1)

	return dbuf.err
}

// textFixed#6c3f19b9 text:RichText = RichText;
func (m *TLTextFixed) To_RichText() *RichText {
	return &RichText{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLTextFixed) SetText(v *RichText) { m.Data2.Text_2 = v }
func (m *TLTextFixed) GetText() *RichText  { return m.Data2.Text_2 }

func (m *TLTextFixed) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_textFixed))

	x.Bytes(m.Data2Text.Encode())

	return x.buf
}

func (m *TLTextFixed) Decode(dbuf *DecodeBuf) error {
	m1 := &Text{}
	m1.Decode(dbuf)
	m.SetText(m1)

	return dbuf.err
}

// textUrl#3c2884c1 text:RichText url:string webpage_id:long = RichText;
func (m *TLTextUrl) To_RichText() *RichText {
	return &RichText{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLTextUrl) SetText(v *RichText) { m.Data2.Text_2 = v }
func (m *TLTextUrl) GetText() *RichText  { return m.Data2.Text_2 }

func (m *TLTextUrl) SetUrl(v string) { m.Data2.Url = v }
func (m *TLTextUrl) GetUrl() string  { return m.Data2.Url }

func (m *TLTextUrl) SetWebpageId(v int64) { m.Data2.WebpageId = v }
func (m *TLTextUrl) GetWebpageId() int64  { return m.Data2.WebpageId }

func (m *TLTextUrl) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_textUrl))

	x.Bytes(m.Data2Text.Encode())
	x.StringBytes(m.Data2.Url)
	x.Long(m.Data2.WebpageId)

	return x.buf
}

func (m *TLTextUrl) Decode(dbuf *DecodeBuf) error {
	m1 := &Text{}
	m1.Decode(dbuf)
	m.SetText(m1)
	m.Data2.Url = x.StringBytes()
	m.Data2.WebpageId = x.Long()

	return dbuf.err
}

// textEmail#de5a0dd6 text:RichText email:string = RichText;
func (m *TLTextEmail) To_RichText() *RichText {
	return &RichText{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLTextEmail) SetText(v *RichText) { m.Data2.Text_2 = v }
func (m *TLTextEmail) GetText() *RichText  { return m.Data2.Text_2 }

func (m *TLTextEmail) SetEmail(v string) { m.Data2.Email = v }
func (m *TLTextEmail) GetEmail() string  { return m.Data2.Email }

func (m *TLTextEmail) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_textEmail))

	x.Bytes(m.Data2Text.Encode())
	x.StringBytes(m.Data2.Email)

	return x.buf
}

func (m *TLTextEmail) Decode(dbuf *DecodeBuf) error {
	m1 := &Text{}
	m1.Decode(dbuf)
	m.SetText(m1)
	m.Data2.Email = x.StringBytes()

	return dbuf.err
}

// textConcat#7e6260d7 texts:Vector<RichText> = RichText;
func (m *TLTextConcat) To_RichText() *RichText {
	return &RichText{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLTextConcat) SetTexts(v []*RichText) { m.Data2.Texts = v }
func (m *TLTextConcat) GetTexts() []*RichText  { return m.Data2.Texts }

func (m *TLTextConcat) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_textConcat))

	return x.buf
}

func (m *TLTextConcat) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Dialog <--
//  + TL_Dialog
//

func (m *Dialog) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_dialog:
		t := m.To_Dialog()
		return t.Encode()

	default:
		return []byte{}
	}
}

// dialog#e4def5db flags:# pinned:flags.2?true peer:Peer top_message:int read_inbox_max_id:int read_outbox_max_id:int unread_count:int unread_mentions_count:int notify_settings:PeerNotifySettings pts:flags.0?int draft:flags.1?DraftMessage = Dialog;
func (m *Dialog) To_Dialog() *TLDialog {
	return &TLDialog{
		Data2: m.Data2,
	}
}

// dialog#e4def5db flags:# pinned:flags.2?true peer:Peer top_message:int read_inbox_max_id:int read_outbox_max_id:int unread_count:int unread_mentions_count:int notify_settings:PeerNotifySettings pts:flags.0?int draft:flags.1?DraftMessage = Dialog;
func (m *TLDialog) To_Dialog() *Dialog {
	return &Dialog{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLDialog) SetPinned(v bool) { m.Data2.Pinned = v }
func (m *TLDialog) GetPinned() bool  { return m.Data2.Pinned }

func (m *TLDialog) SetPeer(v *Peer) { m.Data2.Peer = v }
func (m *TLDialog) GetPeer() *Peer  { return m.Data2.Peer }

func (m *TLDialog) SetTopMessage(v int32) { m.Data2.TopMessage = v }
func (m *TLDialog) GetTopMessage() int32  { return m.Data2.TopMessage }

func (m *TLDialog) SetReadInboxMaxId(v int32) { m.Data2.ReadInboxMaxId = v }
func (m *TLDialog) GetReadInboxMaxId() int32  { return m.Data2.ReadInboxMaxId }

func (m *TLDialog) SetReadOutboxMaxId(v int32) { m.Data2.ReadOutboxMaxId = v }
func (m *TLDialog) GetReadOutboxMaxId() int32  { return m.Data2.ReadOutboxMaxId }

func (m *TLDialog) SetUnreadCount(v int32) { m.Data2.UnreadCount = v }
func (m *TLDialog) GetUnreadCount() int32  { return m.Data2.UnreadCount }

func (m *TLDialog) SetUnreadMentionsCount(v int32) { m.Data2.UnreadMentionsCount = v }
func (m *TLDialog) GetUnreadMentionsCount() int32  { return m.Data2.UnreadMentionsCount }

func (m *TLDialog) SetNotifySettings(v *PeerNotifySettings) { m.Data2.NotifySettings = v }
func (m *TLDialog) GetNotifySettings() *PeerNotifySettings  { return m.Data2.NotifySettings }

func (m *TLDialog) SetPts(v int32) { m.Data2.Pts = v }
func (m *TLDialog) GetPts() int32  { return m.Data2.Pts }

func (m *TLDialog) SetDraft(v *DraftMessage) { m.Data2.Draft = v }
func (m *TLDialog) GetDraft() *DraftMessage  { return m.Data2.Draft }

func (m *TLDialog) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_dialog))

	// flags
	var flags uint32 = 0
	if m.GetPinned() == true {
		flags |= 1 << 1
	}
	if m.GetPts() != 0 {
		flags |= 1 << 9
	}
	if m.GetDraft() != nil {
		flags |= 1 << 10
	}
	x.UInt(flags)

	x.Bytes(m.Data2Peer.Encode())
	x.Int(m.Data2.TopMessage)
	x.Int(m.Data2.ReadInboxMaxId)
	x.Int(m.Data2.ReadOutboxMaxId)
	x.Int(m.Data2.UnreadCount)
	x.Int(m.Data2.UnreadMentionsCount)
	x.Bytes(m.Data2NotifySettings.Encode())
	if m.GetPts() != 0 {
		x.Int(m.Data2.Pts)
	}
	if m.GetDraft() != 0 {
		x.Bytes(m.Data2Draft.Encode())
	}

	return x.buf
}

func (m *TLDialog) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m.Data2.Pinned = true
	}
	m3 := &Peer{}
	m3.Decode(dbuf)
	m.SetPeer(m3)
	m.Data2.TopMessage = x.Int()
	m.Data2.ReadInboxMaxId = x.Int()
	m.Data2.ReadOutboxMaxId = x.Int()
	m.Data2.UnreadCount = x.Int()
	m.Data2.UnreadMentionsCount = x.Int()
	m9 := &NotifySettings{}
	m9.Decode(dbuf)
	m.SetNotifySettings(m9)
	if (flags & (1 << 10)) != 0 {
		m.Data2.Pts = x.Int()
	}
	if (flags & (1 << 11)) != 0 {
		m11 := &Draft{}
		m11.Decode(dbuf)
		m.SetDraft(m11)
	}

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// InputPeerNotifyEvents <--
//  + TL_InputPeerNotifyEventsEmpty
//  + TL_InputPeerNotifyEventsAll
//

func (m *InputPeerNotifyEvents) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_inputPeerNotifyEventsEmpty:
		t := m.To_InputPeerNotifyEventsEmpty()
		return t.Encode()
	case TLConstructor_CRC32_inputPeerNotifyEventsAll:
		t := m.To_InputPeerNotifyEventsAll()
		return t.Encode()

	default:
		return []byte{}
	}
}

// inputPeerNotifyEventsEmpty#f03064d8 = InputPeerNotifyEvents;
func (m *InputPeerNotifyEvents) To_InputPeerNotifyEventsEmpty() *TLInputPeerNotifyEventsEmpty {
	return &TLInputPeerNotifyEventsEmpty{
		Data2: m.Data2,
	}
}

// inputPeerNotifyEventsAll#e86a2c74 = InputPeerNotifyEvents;
func (m *InputPeerNotifyEvents) To_InputPeerNotifyEventsAll() *TLInputPeerNotifyEventsAll {
	return &TLInputPeerNotifyEventsAll{
		Data2: m.Data2,
	}
}

// inputPeerNotifyEventsEmpty#f03064d8 = InputPeerNotifyEvents;
func (m *TLInputPeerNotifyEventsEmpty) To_InputPeerNotifyEvents() *InputPeerNotifyEvents {
	return &InputPeerNotifyEvents{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputPeerNotifyEventsEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputPeerNotifyEventsEmpty))

	return x.buf
}

func (m *TLInputPeerNotifyEventsEmpty) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// inputPeerNotifyEventsAll#e86a2c74 = InputPeerNotifyEvents;
func (m *TLInputPeerNotifyEventsAll) To_InputPeerNotifyEvents() *InputPeerNotifyEvents {
	return &InputPeerNotifyEvents{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputPeerNotifyEventsAll) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputPeerNotifyEventsAll))

	return x.buf
}

func (m *TLInputPeerNotifyEventsAll) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Contact <--
//  + TL_Contact
//

func (m *Contact) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_contact:
		t := m.To_Contact()
		return t.Encode()

	default:
		return []byte{}
	}
}

// contact#f911c994 user_id:int mutual:Bool = Contact;
func (m *Contact) To_Contact() *TLContact {
	return &TLContact{
		Data2: m.Data2,
	}
}

// contact#f911c994 user_id:int mutual:Bool = Contact;
func (m *TLContact) To_Contact() *Contact {
	return &Contact{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLContact) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLContact) GetUserId() int32  { return m.Data2.UserId }

func (m *TLContact) SetMutual(v *Bool) { m.Data2.Mutual = v }
func (m *TLContact) GetMutual() *Bool  { return m.Data2.Mutual }

func (m *TLContact) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_contact))

	x.Int(m.Data2.UserId)
	x.Bytes(m.Data2Mutual.Encode())

	return x.buf
}

func (m *TLContact) Decode(dbuf *DecodeBuf) error {
	m.Data2.UserId = x.Int()
	m2 := &Mutual{}
	m2.Decode(dbuf)
	m.SetMutual(m2)

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Messages_BotResults <--
//  + TL_MessagesBotResults
//

func (m *Messages_BotResults) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_messages_botResults:
		t := m.To_MessagesBotResults()
		return t.Encode()

	default:
		return []byte{}
	}
}

// messages.botResults#ccd3563d flags:# gallery:flags.0?true query_id:long next_offset:flags.1?string switch_pm:flags.2?InlineBotSwitchPM results:Vector<BotInlineResult> cache_time:int = messages.BotResults;
func (m *Messages_BotResults) To_MessagesBotResults() *TLMessagesBotResults {
	return &TLMessagesBotResults{
		Data2: m.Data2,
	}
}

// messages.botResults#ccd3563d flags:# gallery:flags.0?true query_id:long next_offset:flags.1?string switch_pm:flags.2?InlineBotSwitchPM results:Vector<BotInlineResult> cache_time:int = messages.BotResults;
func (m *TLMessagesBotResults) To_Messages_BotResults() *Messages_BotResults {
	return &Messages_BotResults{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessagesBotResults) SetGallery(v bool) { m.Data2.Gallery = v }
func (m *TLMessagesBotResults) GetGallery() bool  { return m.Data2.Gallery }

func (m *TLMessagesBotResults) SetQueryId(v int64) { m.Data2.QueryId = v }
func (m *TLMessagesBotResults) GetQueryId() int64  { return m.Data2.QueryId }

func (m *TLMessagesBotResults) SetNextOffset(v string) { m.Data2.NextOffset = v }
func (m *TLMessagesBotResults) GetNextOffset() string  { return m.Data2.NextOffset }

func (m *TLMessagesBotResults) SetSwitchPm(v *InlineBotSwitchPM) { m.Data2.SwitchPm = v }
func (m *TLMessagesBotResults) GetSwitchPm() *InlineBotSwitchPM  { return m.Data2.SwitchPm }

func (m *TLMessagesBotResults) SetResults(v []*BotInlineResult) { m.Data2.Results = v }
func (m *TLMessagesBotResults) GetResults() []*BotInlineResult  { return m.Data2.Results }

func (m *TLMessagesBotResults) SetCacheTime(v int32) { m.Data2.CacheTime = v }
func (m *TLMessagesBotResults) GetCacheTime() int32  { return m.Data2.CacheTime }

func (m *TLMessagesBotResults) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_botResults))

	// flags
	var flags uint32 = 0
	if m.GetGallery() == true {
		flags |= 1 << 1
	}
	if m.GetNextOffset() != "" {
		flags |= 1 << 3
	}
	if m.GetSwitchPm() != nil {
		flags |= 1 << 4
	}
	x.UInt(flags)

	x.Long(m.Data2.QueryId)
	if m.GetNextOffset() != 0 {
		x.StringBytes(m.Data2.NextOffset)
	}
	if m.GetSwitchPm() != 0 {
		x.Bytes(m.Data2SwitchPm.Encode())
	}

	x.Int(m.Data2.CacheTime)

	return x.buf
}

func (m *TLMessagesBotResults) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m.Data2.Gallery = true
	}
	m.Data2.QueryId = x.Long()
	if (flags & (1 << 4)) != 0 {
		m.Data2.NextOffset = x.StringBytes()
	}
	if (flags & (1 << 5)) != 0 {
		m5 := &SwitchPm{}
		m5.Decode(dbuf)
		m.SetSwitchPm(m5)
	}

	m.Data2.CacheTime = x.Int()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// MessageEntity <--
//  + TL_MessageEntityUnknown
//  + TL_MessageEntityMention
//  + TL_MessageEntityHashtag
//  + TL_MessageEntityBotCommand
//  + TL_MessageEntityUrl
//  + TL_MessageEntityEmail
//  + TL_MessageEntityBold
//  + TL_MessageEntityItalic
//  + TL_MessageEntityCode
//  + TL_MessageEntityPre
//  + TL_MessageEntityTextUrl
//  + TL_MessageEntityMentionName
//  + TL_InputMessageEntityMentionName
//

func (m *MessageEntity) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_messageEntityUnknown:
		t := m.To_MessageEntityUnknown()
		return t.Encode()
	case TLConstructor_CRC32_messageEntityMention:
		t := m.To_MessageEntityMention()
		return t.Encode()
	case TLConstructor_CRC32_messageEntityHashtag:
		t := m.To_MessageEntityHashtag()
		return t.Encode()
	case TLConstructor_CRC32_messageEntityBotCommand:
		t := m.To_MessageEntityBotCommand()
		return t.Encode()
	case TLConstructor_CRC32_messageEntityUrl:
		t := m.To_MessageEntityUrl()
		return t.Encode()
	case TLConstructor_CRC32_messageEntityEmail:
		t := m.To_MessageEntityEmail()
		return t.Encode()
	case TLConstructor_CRC32_messageEntityBold:
		t := m.To_MessageEntityBold()
		return t.Encode()
	case TLConstructor_CRC32_messageEntityItalic:
		t := m.To_MessageEntityItalic()
		return t.Encode()
	case TLConstructor_CRC32_messageEntityCode:
		t := m.To_MessageEntityCode()
		return t.Encode()
	case TLConstructor_CRC32_messageEntityPre:
		t := m.To_MessageEntityPre()
		return t.Encode()
	case TLConstructor_CRC32_messageEntityTextUrl:
		t := m.To_MessageEntityTextUrl()
		return t.Encode()
	case TLConstructor_CRC32_messageEntityMentionName:
		t := m.To_MessageEntityMentionName()
		return t.Encode()
	case TLConstructor_CRC32_inputMessageEntityMentionName:
		t := m.To_InputMessageEntityMentionName()
		return t.Encode()

	default:
		return []byte{}
	}
}

// messageEntityUnknown#bb92ba95 offset:int length:int = MessageEntity;
func (m *MessageEntity) To_MessageEntityUnknown() *TLMessageEntityUnknown {
	return &TLMessageEntityUnknown{
		Data2: m.Data2,
	}
}

// messageEntityMention#fa04579d offset:int length:int = MessageEntity;
func (m *MessageEntity) To_MessageEntityMention() *TLMessageEntityMention {
	return &TLMessageEntityMention{
		Data2: m.Data2,
	}
}

// messageEntityHashtag#6f635b0d offset:int length:int = MessageEntity;
func (m *MessageEntity) To_MessageEntityHashtag() *TLMessageEntityHashtag {
	return &TLMessageEntityHashtag{
		Data2: m.Data2,
	}
}

// messageEntityBotCommand#6cef8ac7 offset:int length:int = MessageEntity;
func (m *MessageEntity) To_MessageEntityBotCommand() *TLMessageEntityBotCommand {
	return &TLMessageEntityBotCommand{
		Data2: m.Data2,
	}
}

// messageEntityUrl#6ed02538 offset:int length:int = MessageEntity;
func (m *MessageEntity) To_MessageEntityUrl() *TLMessageEntityUrl {
	return &TLMessageEntityUrl{
		Data2: m.Data2,
	}
}

// messageEntityEmail#64e475c2 offset:int length:int = MessageEntity;
func (m *MessageEntity) To_MessageEntityEmail() *TLMessageEntityEmail {
	return &TLMessageEntityEmail{
		Data2: m.Data2,
	}
}

// messageEntityBold#bd610bc9 offset:int length:int = MessageEntity;
func (m *MessageEntity) To_MessageEntityBold() *TLMessageEntityBold {
	return &TLMessageEntityBold{
		Data2: m.Data2,
	}
}

// messageEntityItalic#826f8b60 offset:int length:int = MessageEntity;
func (m *MessageEntity) To_MessageEntityItalic() *TLMessageEntityItalic {
	return &TLMessageEntityItalic{
		Data2: m.Data2,
	}
}

// messageEntityCode#28a20571 offset:int length:int = MessageEntity;
func (m *MessageEntity) To_MessageEntityCode() *TLMessageEntityCode {
	return &TLMessageEntityCode{
		Data2: m.Data2,
	}
}

// messageEntityPre#73924be0 offset:int length:int language:string = MessageEntity;
func (m *MessageEntity) To_MessageEntityPre() *TLMessageEntityPre {
	return &TLMessageEntityPre{
		Data2: m.Data2,
	}
}

// messageEntityTextUrl#76a6d327 offset:int length:int url:string = MessageEntity;
func (m *MessageEntity) To_MessageEntityTextUrl() *TLMessageEntityTextUrl {
	return &TLMessageEntityTextUrl{
		Data2: m.Data2,
	}
}

// messageEntityMentionName#352dca58 offset:int length:int user_id:int = MessageEntity;
func (m *MessageEntity) To_MessageEntityMentionName() *TLMessageEntityMentionName {
	return &TLMessageEntityMentionName{
		Data2: m.Data2,
	}
}

// inputMessageEntityMentionName#208e68c9 offset:int length:int user_id:InputUser = MessageEntity;
func (m *MessageEntity) To_InputMessageEntityMentionName() *TLInputMessageEntityMentionName {
	return &TLInputMessageEntityMentionName{
		Data2: m.Data2,
	}
}

// messageEntityUnknown#bb92ba95 offset:int length:int = MessageEntity;
func (m *TLMessageEntityUnknown) To_MessageEntity() *MessageEntity {
	return &MessageEntity{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessageEntityUnknown) SetOffset(v int32) { m.Data2.Offset = v }
func (m *TLMessageEntityUnknown) GetOffset() int32  { return m.Data2.Offset }

func (m *TLMessageEntityUnknown) SetLength(v int32) { m.Data2.Length = v }
func (m *TLMessageEntityUnknown) GetLength() int32  { return m.Data2.Length }

func (m *TLMessageEntityUnknown) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageEntityUnknown))

	x.Int(m.Data2.Offset)
	x.Int(m.Data2.Length)

	return x.buf
}

func (m *TLMessageEntityUnknown) Decode(dbuf *DecodeBuf) error {
	m.Data2.Offset = x.Int()
	m.Data2.Length = x.Int()

	return dbuf.err
}

// messageEntityMention#fa04579d offset:int length:int = MessageEntity;
func (m *TLMessageEntityMention) To_MessageEntity() *MessageEntity {
	return &MessageEntity{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessageEntityMention) SetOffset(v int32) { m.Data2.Offset = v }
func (m *TLMessageEntityMention) GetOffset() int32  { return m.Data2.Offset }

func (m *TLMessageEntityMention) SetLength(v int32) { m.Data2.Length = v }
func (m *TLMessageEntityMention) GetLength() int32  { return m.Data2.Length }

func (m *TLMessageEntityMention) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageEntityMention))

	x.Int(m.Data2.Offset)
	x.Int(m.Data2.Length)

	return x.buf
}

func (m *TLMessageEntityMention) Decode(dbuf *DecodeBuf) error {
	m.Data2.Offset = x.Int()
	m.Data2.Length = x.Int()

	return dbuf.err
}

// messageEntityHashtag#6f635b0d offset:int length:int = MessageEntity;
func (m *TLMessageEntityHashtag) To_MessageEntity() *MessageEntity {
	return &MessageEntity{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessageEntityHashtag) SetOffset(v int32) { m.Data2.Offset = v }
func (m *TLMessageEntityHashtag) GetOffset() int32  { return m.Data2.Offset }

func (m *TLMessageEntityHashtag) SetLength(v int32) { m.Data2.Length = v }
func (m *TLMessageEntityHashtag) GetLength() int32  { return m.Data2.Length }

func (m *TLMessageEntityHashtag) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageEntityHashtag))

	x.Int(m.Data2.Offset)
	x.Int(m.Data2.Length)

	return x.buf
}

func (m *TLMessageEntityHashtag) Decode(dbuf *DecodeBuf) error {
	m.Data2.Offset = x.Int()
	m.Data2.Length = x.Int()

	return dbuf.err
}

// messageEntityBotCommand#6cef8ac7 offset:int length:int = MessageEntity;
func (m *TLMessageEntityBotCommand) To_MessageEntity() *MessageEntity {
	return &MessageEntity{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessageEntityBotCommand) SetOffset(v int32) { m.Data2.Offset = v }
func (m *TLMessageEntityBotCommand) GetOffset() int32  { return m.Data2.Offset }

func (m *TLMessageEntityBotCommand) SetLength(v int32) { m.Data2.Length = v }
func (m *TLMessageEntityBotCommand) GetLength() int32  { return m.Data2.Length }

func (m *TLMessageEntityBotCommand) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageEntityBotCommand))

	x.Int(m.Data2.Offset)
	x.Int(m.Data2.Length)

	return x.buf
}

func (m *TLMessageEntityBotCommand) Decode(dbuf *DecodeBuf) error {
	m.Data2.Offset = x.Int()
	m.Data2.Length = x.Int()

	return dbuf.err
}

// messageEntityUrl#6ed02538 offset:int length:int = MessageEntity;
func (m *TLMessageEntityUrl) To_MessageEntity() *MessageEntity {
	return &MessageEntity{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessageEntityUrl) SetOffset(v int32) { m.Data2.Offset = v }
func (m *TLMessageEntityUrl) GetOffset() int32  { return m.Data2.Offset }

func (m *TLMessageEntityUrl) SetLength(v int32) { m.Data2.Length = v }
func (m *TLMessageEntityUrl) GetLength() int32  { return m.Data2.Length }

func (m *TLMessageEntityUrl) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageEntityUrl))

	x.Int(m.Data2.Offset)
	x.Int(m.Data2.Length)

	return x.buf
}

func (m *TLMessageEntityUrl) Decode(dbuf *DecodeBuf) error {
	m.Data2.Offset = x.Int()
	m.Data2.Length = x.Int()

	return dbuf.err
}

// messageEntityEmail#64e475c2 offset:int length:int = MessageEntity;
func (m *TLMessageEntityEmail) To_MessageEntity() *MessageEntity {
	return &MessageEntity{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessageEntityEmail) SetOffset(v int32) { m.Data2.Offset = v }
func (m *TLMessageEntityEmail) GetOffset() int32  { return m.Data2.Offset }

func (m *TLMessageEntityEmail) SetLength(v int32) { m.Data2.Length = v }
func (m *TLMessageEntityEmail) GetLength() int32  { return m.Data2.Length }

func (m *TLMessageEntityEmail) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageEntityEmail))

	x.Int(m.Data2.Offset)
	x.Int(m.Data2.Length)

	return x.buf
}

func (m *TLMessageEntityEmail) Decode(dbuf *DecodeBuf) error {
	m.Data2.Offset = x.Int()
	m.Data2.Length = x.Int()

	return dbuf.err
}

// messageEntityBold#bd610bc9 offset:int length:int = MessageEntity;
func (m *TLMessageEntityBold) To_MessageEntity() *MessageEntity {
	return &MessageEntity{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessageEntityBold) SetOffset(v int32) { m.Data2.Offset = v }
func (m *TLMessageEntityBold) GetOffset() int32  { return m.Data2.Offset }

func (m *TLMessageEntityBold) SetLength(v int32) { m.Data2.Length = v }
func (m *TLMessageEntityBold) GetLength() int32  { return m.Data2.Length }

func (m *TLMessageEntityBold) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageEntityBold))

	x.Int(m.Data2.Offset)
	x.Int(m.Data2.Length)

	return x.buf
}

func (m *TLMessageEntityBold) Decode(dbuf *DecodeBuf) error {
	m.Data2.Offset = x.Int()
	m.Data2.Length = x.Int()

	return dbuf.err
}

// messageEntityItalic#826f8b60 offset:int length:int = MessageEntity;
func (m *TLMessageEntityItalic) To_MessageEntity() *MessageEntity {
	return &MessageEntity{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessageEntityItalic) SetOffset(v int32) { m.Data2.Offset = v }
func (m *TLMessageEntityItalic) GetOffset() int32  { return m.Data2.Offset }

func (m *TLMessageEntityItalic) SetLength(v int32) { m.Data2.Length = v }
func (m *TLMessageEntityItalic) GetLength() int32  { return m.Data2.Length }

func (m *TLMessageEntityItalic) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageEntityItalic))

	x.Int(m.Data2.Offset)
	x.Int(m.Data2.Length)

	return x.buf
}

func (m *TLMessageEntityItalic) Decode(dbuf *DecodeBuf) error {
	m.Data2.Offset = x.Int()
	m.Data2.Length = x.Int()

	return dbuf.err
}

// messageEntityCode#28a20571 offset:int length:int = MessageEntity;
func (m *TLMessageEntityCode) To_MessageEntity() *MessageEntity {
	return &MessageEntity{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessageEntityCode) SetOffset(v int32) { m.Data2.Offset = v }
func (m *TLMessageEntityCode) GetOffset() int32  { return m.Data2.Offset }

func (m *TLMessageEntityCode) SetLength(v int32) { m.Data2.Length = v }
func (m *TLMessageEntityCode) GetLength() int32  { return m.Data2.Length }

func (m *TLMessageEntityCode) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageEntityCode))

	x.Int(m.Data2.Offset)
	x.Int(m.Data2.Length)

	return x.buf
}

func (m *TLMessageEntityCode) Decode(dbuf *DecodeBuf) error {
	m.Data2.Offset = x.Int()
	m.Data2.Length = x.Int()

	return dbuf.err
}

// messageEntityPre#73924be0 offset:int length:int language:string = MessageEntity;
func (m *TLMessageEntityPre) To_MessageEntity() *MessageEntity {
	return &MessageEntity{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessageEntityPre) SetOffset(v int32) { m.Data2.Offset = v }
func (m *TLMessageEntityPre) GetOffset() int32  { return m.Data2.Offset }

func (m *TLMessageEntityPre) SetLength(v int32) { m.Data2.Length = v }
func (m *TLMessageEntityPre) GetLength() int32  { return m.Data2.Length }

func (m *TLMessageEntityPre) SetLanguage(v string) { m.Data2.Language = v }
func (m *TLMessageEntityPre) GetLanguage() string  { return m.Data2.Language }

func (m *TLMessageEntityPre) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageEntityPre))

	x.Int(m.Data2.Offset)
	x.Int(m.Data2.Length)
	x.StringBytes(m.Data2.Language)

	return x.buf
}

func (m *TLMessageEntityPre) Decode(dbuf *DecodeBuf) error {
	m.Data2.Offset = x.Int()
	m.Data2.Length = x.Int()
	m.Data2.Language = x.StringBytes()

	return dbuf.err
}

// messageEntityTextUrl#76a6d327 offset:int length:int url:string = MessageEntity;
func (m *TLMessageEntityTextUrl) To_MessageEntity() *MessageEntity {
	return &MessageEntity{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessageEntityTextUrl) SetOffset(v int32) { m.Data2.Offset = v }
func (m *TLMessageEntityTextUrl) GetOffset() int32  { return m.Data2.Offset }

func (m *TLMessageEntityTextUrl) SetLength(v int32) { m.Data2.Length = v }
func (m *TLMessageEntityTextUrl) GetLength() int32  { return m.Data2.Length }

func (m *TLMessageEntityTextUrl) SetUrl(v string) { m.Data2.Url = v }
func (m *TLMessageEntityTextUrl) GetUrl() string  { return m.Data2.Url }

func (m *TLMessageEntityTextUrl) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageEntityTextUrl))

	x.Int(m.Data2.Offset)
	x.Int(m.Data2.Length)
	x.StringBytes(m.Data2.Url)

	return x.buf
}

func (m *TLMessageEntityTextUrl) Decode(dbuf *DecodeBuf) error {
	m.Data2.Offset = x.Int()
	m.Data2.Length = x.Int()
	m.Data2.Url = x.StringBytes()

	return dbuf.err
}

// messageEntityMentionName#352dca58 offset:int length:int user_id:int = MessageEntity;
func (m *TLMessageEntityMentionName) To_MessageEntity() *MessageEntity {
	return &MessageEntity{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessageEntityMentionName) SetOffset(v int32) { m.Data2.Offset = v }
func (m *TLMessageEntityMentionName) GetOffset() int32  { return m.Data2.Offset }

func (m *TLMessageEntityMentionName) SetLength(v int32) { m.Data2.Length = v }
func (m *TLMessageEntityMentionName) GetLength() int32  { return m.Data2.Length }

func (m *TLMessageEntityMentionName) SetUserId(v int32) { m.Data2.UserId_5 = v }
func (m *TLMessageEntityMentionName) GetUserId() int32  { return m.Data2.UserId_5 }

func (m *TLMessageEntityMentionName) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageEntityMentionName))

	x.Int(m.Data2.Offset)
	x.Int(m.Data2.Length)
	x.Int(m.Data2.UserId)

	return x.buf
}

func (m *TLMessageEntityMentionName) Decode(dbuf *DecodeBuf) error {
	m.Data2.Offset = x.Int()
	m.Data2.Length = x.Int()
	m.Data2.UserId = x.Int()

	return dbuf.err
}

// inputMessageEntityMentionName#208e68c9 offset:int length:int user_id:InputUser = MessageEntity;
func (m *TLInputMessageEntityMentionName) To_MessageEntity() *MessageEntity {
	return &MessageEntity{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputMessageEntityMentionName) SetOffset(v int32) { m.Data2.Offset = v }
func (m *TLInputMessageEntityMentionName) GetOffset() int32  { return m.Data2.Offset }

func (m *TLInputMessageEntityMentionName) SetLength(v int32) { m.Data2.Length = v }
func (m *TLInputMessageEntityMentionName) GetLength() int32  { return m.Data2.Length }

func (m *TLInputMessageEntityMentionName) SetUserId(v *InputUser) { m.Data2.UserId_6 = v }
func (m *TLInputMessageEntityMentionName) GetUserId() *InputUser  { return m.Data2.UserId_6 }

func (m *TLInputMessageEntityMentionName) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputMessageEntityMentionName))

	x.Int(m.Data2.Offset)
	x.Int(m.Data2.Length)
	x.Bytes(m.Data2UserId.Encode())

	return x.buf
}

func (m *TLInputMessageEntityMentionName) Decode(dbuf *DecodeBuf) error {
	m.Data2.Offset = x.Int()
	m.Data2.Length = x.Int()
	m3 := &UserId{}
	m3.Decode(dbuf)
	m.SetUserId(m3)

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// FoundGif <--
//  + TL_FoundGif
//  + TL_FoundGifCached
//

func (m *FoundGif) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_foundGif:
		t := m.To_FoundGif()
		return t.Encode()
	case TLConstructor_CRC32_foundGifCached:
		t := m.To_FoundGifCached()
		return t.Encode()

	default:
		return []byte{}
	}
}

// foundGif#162ecc1f url:string thumb_url:string content_url:string content_type:string w:int h:int = FoundGif;
func (m *FoundGif) To_FoundGif() *TLFoundGif {
	return &TLFoundGif{
		Data2: m.Data2,
	}
}

// foundGifCached#9c750409 url:string photo:Photo document:Document = FoundGif;
func (m *FoundGif) To_FoundGifCached() *TLFoundGifCached {
	return &TLFoundGifCached{
		Data2: m.Data2,
	}
}

// foundGif#162ecc1f url:string thumb_url:string content_url:string content_type:string w:int h:int = FoundGif;
func (m *TLFoundGif) To_FoundGif() *FoundGif {
	return &FoundGif{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLFoundGif) SetUrl(v string) { m.Data2.Url = v }
func (m *TLFoundGif) GetUrl() string  { return m.Data2.Url }

func (m *TLFoundGif) SetThumbUrl(v string) { m.Data2.ThumbUrl = v }
func (m *TLFoundGif) GetThumbUrl() string  { return m.Data2.ThumbUrl }

func (m *TLFoundGif) SetContentUrl(v string) { m.Data2.ContentUrl = v }
func (m *TLFoundGif) GetContentUrl() string  { return m.Data2.ContentUrl }

func (m *TLFoundGif) SetContentType(v string) { m.Data2.ContentType = v }
func (m *TLFoundGif) GetContentType() string  { return m.Data2.ContentType }

func (m *TLFoundGif) SetW(v int32) { m.Data2.W = v }
func (m *TLFoundGif) GetW() int32  { return m.Data2.W }

func (m *TLFoundGif) SetH(v int32) { m.Data2.H = v }
func (m *TLFoundGif) GetH() int32  { return m.Data2.H }

func (m *TLFoundGif) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_foundGif))

	x.StringBytes(m.Data2.Url)
	x.StringBytes(m.Data2.ThumbUrl)
	x.StringBytes(m.Data2.ContentUrl)
	x.StringBytes(m.Data2.ContentType)
	x.Int(m.Data2.W)
	x.Int(m.Data2.H)

	return x.buf
}

func (m *TLFoundGif) Decode(dbuf *DecodeBuf) error {
	m.Data2.Url = x.StringBytes()
	m.Data2.ThumbUrl = x.StringBytes()
	m.Data2.ContentUrl = x.StringBytes()
	m.Data2.ContentType = x.StringBytes()
	m.Data2.W = x.Int()
	m.Data2.H = x.Int()

	return dbuf.err
}

// foundGifCached#9c750409 url:string photo:Photo document:Document = FoundGif;
func (m *TLFoundGifCached) To_FoundGif() *FoundGif {
	return &FoundGif{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLFoundGifCached) SetUrl(v string) { m.Data2.Url = v }
func (m *TLFoundGifCached) GetUrl() string  { return m.Data2.Url }

func (m *TLFoundGifCached) SetPhoto(v *Photo) { m.Data2.Photo = v }
func (m *TLFoundGifCached) GetPhoto() *Photo  { return m.Data2.Photo }

func (m *TLFoundGifCached) SetDocument(v *Document) { m.Data2.Document = v }
func (m *TLFoundGifCached) GetDocument() *Document  { return m.Data2.Document }

func (m *TLFoundGifCached) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_foundGifCached))

	x.StringBytes(m.Data2.Url)
	x.Bytes(m.Data2Photo.Encode())
	x.Bytes(m.Data2Document.Encode())

	return x.buf
}

func (m *TLFoundGifCached) Decode(dbuf *DecodeBuf) error {
	m.Data2.Url = x.StringBytes()
	m2 := &Photo{}
	m2.Decode(dbuf)
	m.SetPhoto(m2)
	m3 := &Document{}
	m3.Decode(dbuf)
	m.SetDocument(m3)

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// LabeledPrice <--
//  + TL_LabeledPrice
//

func (m *LabeledPrice) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_labeledPrice:
		t := m.To_LabeledPrice()
		return t.Encode()

	default:
		return []byte{}
	}
}

// labeledPrice#cb296bf8 label:string amount:long = LabeledPrice;
func (m *LabeledPrice) To_LabeledPrice() *TLLabeledPrice {
	return &TLLabeledPrice{
		Data2: m.Data2,
	}
}

// labeledPrice#cb296bf8 label:string amount:long = LabeledPrice;
func (m *TLLabeledPrice) To_LabeledPrice() *LabeledPrice {
	return &LabeledPrice{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLLabeledPrice) SetLabel(v string) { m.Data2.Label = v }
func (m *TLLabeledPrice) GetLabel() string  { return m.Data2.Label }

func (m *TLLabeledPrice) SetAmount(v int64) { m.Data2.Amount = v }
func (m *TLLabeledPrice) GetAmount() int64  { return m.Data2.Amount }

func (m *TLLabeledPrice) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_labeledPrice))

	x.StringBytes(m.Data2.Label)
	x.Long(m.Data2.Amount)

	return x.buf
}

func (m *TLLabeledPrice) Decode(dbuf *DecodeBuf) error {
	m.Data2.Label = x.StringBytes()
	m.Data2.Amount = x.Long()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Payments_ValidatedRequestedInfo <--
//  + TL_PaymentsValidatedRequestedInfo
//

func (m *Payments_ValidatedRequestedInfo) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_payments_validatedRequestedInfo:
		t := m.To_PaymentsValidatedRequestedInfo()
		return t.Encode()

	default:
		return []byte{}
	}
}

// payments.validatedRequestedInfo#d1451883 flags:# id:flags.0?string shipping_options:flags.1?Vector<ShippingOption> = payments.ValidatedRequestedInfo;
func (m *Payments_ValidatedRequestedInfo) To_PaymentsValidatedRequestedInfo() *TLPaymentsValidatedRequestedInfo {
	return &TLPaymentsValidatedRequestedInfo{
		Data2: m.Data2,
	}
}

// payments.validatedRequestedInfo#d1451883 flags:# id:flags.0?string shipping_options:flags.1?Vector<ShippingOption> = payments.ValidatedRequestedInfo;
func (m *TLPaymentsValidatedRequestedInfo) To_Payments_ValidatedRequestedInfo() *Payments_ValidatedRequestedInfo {
	return &Payments_ValidatedRequestedInfo{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPaymentsValidatedRequestedInfo) SetId(v string) { m.Data2.Id = v }
func (m *TLPaymentsValidatedRequestedInfo) GetId() string  { return m.Data2.Id }

func (m *TLPaymentsValidatedRequestedInfo) SetShippingOptions(v []*ShippingOption) {
	m.Data2.ShippingOptions = v
}
func (m *TLPaymentsValidatedRequestedInfo) GetShippingOptions() []*ShippingOption {
	return m.Data2.ShippingOptions
}

func (m *TLPaymentsValidatedRequestedInfo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_payments_validatedRequestedInfo))

	// flags
	var flags uint32 = 0
	if m.GetId() != "" {
		flags |= 1 << 1
	}
	if m.GetShippingOptions() != nil {
		flags |= 1 << 2
	}
	x.UInt(flags)

	if m.GetId() != 0 {
		x.StringBytes(m.Data2.Id)
	}
	if m.Data2.ShippingOptions != 0 {

	}

	return x.buf
}

func (m *TLPaymentsValidatedRequestedInfo) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m.Data2.Id = x.StringBytes()
	}
	if (flags & (1 << 3)) != 0 {

	}

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// MessageAction <--
//  + TL_MessageActionEmpty
//  + TL_MessageActionChatCreate
//  + TL_MessageActionChatEditTitle
//  + TL_MessageActionChatEditPhoto
//  + TL_MessageActionChatDeletePhoto
//  + TL_MessageActionChatAddUser
//  + TL_MessageActionChatDeleteUser
//  + TL_MessageActionChatJoinedByLink
//  + TL_MessageActionChannelCreate
//  + TL_MessageActionChatMigrateTo
//  + TL_MessageActionChannelMigrateFrom
//  + TL_MessageActionPinMessage
//  + TL_MessageActionHistoryClear
//  + TL_MessageActionGameScore
//  + TL_MessageActionPaymentSentMe
//  + TL_MessageActionPaymentSent
//  + TL_MessageActionPhoneCall
//  + TL_MessageActionScreenshotTaken
//

func (m *MessageAction) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_messageActionEmpty:
		t := m.To_MessageActionEmpty()
		return t.Encode()
	case TLConstructor_CRC32_messageActionChatCreate:
		t := m.To_MessageActionChatCreate()
		return t.Encode()
	case TLConstructor_CRC32_messageActionChatEditTitle:
		t := m.To_MessageActionChatEditTitle()
		return t.Encode()
	case TLConstructor_CRC32_messageActionChatEditPhoto:
		t := m.To_MessageActionChatEditPhoto()
		return t.Encode()
	case TLConstructor_CRC32_messageActionChatDeletePhoto:
		t := m.To_MessageActionChatDeletePhoto()
		return t.Encode()
	case TLConstructor_CRC32_messageActionChatAddUser:
		t := m.To_MessageActionChatAddUser()
		return t.Encode()
	case TLConstructor_CRC32_messageActionChatDeleteUser:
		t := m.To_MessageActionChatDeleteUser()
		return t.Encode()
	case TLConstructor_CRC32_messageActionChatJoinedByLink:
		t := m.To_MessageActionChatJoinedByLink()
		return t.Encode()
	case TLConstructor_CRC32_messageActionChannelCreate:
		t := m.To_MessageActionChannelCreate()
		return t.Encode()
	case TLConstructor_CRC32_messageActionChatMigrateTo:
		t := m.To_MessageActionChatMigrateTo()
		return t.Encode()
	case TLConstructor_CRC32_messageActionChannelMigrateFrom:
		t := m.To_MessageActionChannelMigrateFrom()
		return t.Encode()
	case TLConstructor_CRC32_messageActionPinMessage:
		t := m.To_MessageActionPinMessage()
		return t.Encode()
	case TLConstructor_CRC32_messageActionHistoryClear:
		t := m.To_MessageActionHistoryClear()
		return t.Encode()
	case TLConstructor_CRC32_messageActionGameScore:
		t := m.To_MessageActionGameScore()
		return t.Encode()
	case TLConstructor_CRC32_messageActionPaymentSentMe:
		t := m.To_MessageActionPaymentSentMe()
		return t.Encode()
	case TLConstructor_CRC32_messageActionPaymentSent:
		t := m.To_MessageActionPaymentSent()
		return t.Encode()
	case TLConstructor_CRC32_messageActionPhoneCall:
		t := m.To_MessageActionPhoneCall()
		return t.Encode()
	case TLConstructor_CRC32_messageActionScreenshotTaken:
		t := m.To_MessageActionScreenshotTaken()
		return t.Encode()

	default:
		return []byte{}
	}
}

// messageActionEmpty#b6aef7b0 = MessageAction;
func (m *MessageAction) To_MessageActionEmpty() *TLMessageActionEmpty {
	return &TLMessageActionEmpty{
		Data2: m.Data2,
	}
}

// messageActionChatCreate#a6638b9a title:string users:Vector<int> = MessageAction;
func (m *MessageAction) To_MessageActionChatCreate() *TLMessageActionChatCreate {
	return &TLMessageActionChatCreate{
		Data2: m.Data2,
	}
}

// messageActionChatEditTitle#b5a1ce5a title:string = MessageAction;
func (m *MessageAction) To_MessageActionChatEditTitle() *TLMessageActionChatEditTitle {
	return &TLMessageActionChatEditTitle{
		Data2: m.Data2,
	}
}

// messageActionChatEditPhoto#7fcb13a8 photo:Photo = MessageAction;
func (m *MessageAction) To_MessageActionChatEditPhoto() *TLMessageActionChatEditPhoto {
	return &TLMessageActionChatEditPhoto{
		Data2: m.Data2,
	}
}

// messageActionChatDeletePhoto#95e3fbef = MessageAction;
func (m *MessageAction) To_MessageActionChatDeletePhoto() *TLMessageActionChatDeletePhoto {
	return &TLMessageActionChatDeletePhoto{
		Data2: m.Data2,
	}
}

// messageActionChatAddUser#488a7337 users:Vector<int> = MessageAction;
func (m *MessageAction) To_MessageActionChatAddUser() *TLMessageActionChatAddUser {
	return &TLMessageActionChatAddUser{
		Data2: m.Data2,
	}
}

// messageActionChatDeleteUser#b2ae9b0c user_id:int = MessageAction;
func (m *MessageAction) To_MessageActionChatDeleteUser() *TLMessageActionChatDeleteUser {
	return &TLMessageActionChatDeleteUser{
		Data2: m.Data2,
	}
}

// messageActionChatJoinedByLink#f89cf5e8 inviter_id:int = MessageAction;
func (m *MessageAction) To_MessageActionChatJoinedByLink() *TLMessageActionChatJoinedByLink {
	return &TLMessageActionChatJoinedByLink{
		Data2: m.Data2,
	}
}

// messageActionChannelCreate#95d2ac92 title:string = MessageAction;
func (m *MessageAction) To_MessageActionChannelCreate() *TLMessageActionChannelCreate {
	return &TLMessageActionChannelCreate{
		Data2: m.Data2,
	}
}

// messageActionChatMigrateTo#51bdb021 channel_id:int = MessageAction;
func (m *MessageAction) To_MessageActionChatMigrateTo() *TLMessageActionChatMigrateTo {
	return &TLMessageActionChatMigrateTo{
		Data2: m.Data2,
	}
}

// messageActionChannelMigrateFrom#b055eaee title:string chat_id:int = MessageAction;
func (m *MessageAction) To_MessageActionChannelMigrateFrom() *TLMessageActionChannelMigrateFrom {
	return &TLMessageActionChannelMigrateFrom{
		Data2: m.Data2,
	}
}

// messageActionPinMessage#94bd38ed = MessageAction;
func (m *MessageAction) To_MessageActionPinMessage() *TLMessageActionPinMessage {
	return &TLMessageActionPinMessage{
		Data2: m.Data2,
	}
}

// messageActionHistoryClear#9fbab604 = MessageAction;
func (m *MessageAction) To_MessageActionHistoryClear() *TLMessageActionHistoryClear {
	return &TLMessageActionHistoryClear{
		Data2: m.Data2,
	}
}

// messageActionGameScore#92a72876 game_id:long score:int = MessageAction;
func (m *MessageAction) To_MessageActionGameScore() *TLMessageActionGameScore {
	return &TLMessageActionGameScore{
		Data2: m.Data2,
	}
}

// messageActionPaymentSentMe#8f31b327 flags:# currency:string total_amount:long payload:bytes info:flags.0?PaymentRequestedInfo shipping_option_id:flags.1?string charge:PaymentCharge = MessageAction;
func (m *MessageAction) To_MessageActionPaymentSentMe() *TLMessageActionPaymentSentMe {
	return &TLMessageActionPaymentSentMe{
		Data2: m.Data2,
	}
}

// messageActionPaymentSent#40699cd0 currency:string total_amount:long = MessageAction;
func (m *MessageAction) To_MessageActionPaymentSent() *TLMessageActionPaymentSent {
	return &TLMessageActionPaymentSent{
		Data2: m.Data2,
	}
}

// messageActionPhoneCall#80e11a7f flags:# call_id:long reason:flags.0?PhoneCallDiscardReason duration:flags.1?int = MessageAction;
func (m *MessageAction) To_MessageActionPhoneCall() *TLMessageActionPhoneCall {
	return &TLMessageActionPhoneCall{
		Data2: m.Data2,
	}
}

// messageActionScreenshotTaken#4792929b = MessageAction;
func (m *MessageAction) To_MessageActionScreenshotTaken() *TLMessageActionScreenshotTaken {
	return &TLMessageActionScreenshotTaken{
		Data2: m.Data2,
	}
}

// messageActionEmpty#b6aef7b0 = MessageAction;
func (m *TLMessageActionEmpty) To_MessageAction() *MessageAction {
	return &MessageAction{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessageActionEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageActionEmpty))

	return x.buf
}

func (m *TLMessageActionEmpty) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// messageActionChatCreate#a6638b9a title:string users:Vector<int> = MessageAction;
func (m *TLMessageActionChatCreate) To_MessageAction() *MessageAction {
	return &MessageAction{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessageActionChatCreate) SetTitle(v string) { m.Data2.Title = v }
func (m *TLMessageActionChatCreate) GetTitle() string  { return m.Data2.Title }

func (m *TLMessageActionChatCreate) SetUsers(v []int32) { m.Data2.Users = v }
func (m *TLMessageActionChatCreate) GetUsers() []int32  { return m.Data2.Users }

func (m *TLMessageActionChatCreate) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageActionChatCreate))

	x.StringBytes(m.Data2.Title)

	return x.buf
}

func (m *TLMessageActionChatCreate) Decode(dbuf *DecodeBuf) error {
	m.Data2.Title = x.StringBytes()

	return dbuf.err
}

// messageActionChatEditTitle#b5a1ce5a title:string = MessageAction;
func (m *TLMessageActionChatEditTitle) To_MessageAction() *MessageAction {
	return &MessageAction{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessageActionChatEditTitle) SetTitle(v string) { m.Data2.Title = v }
func (m *TLMessageActionChatEditTitle) GetTitle() string  { return m.Data2.Title }

func (m *TLMessageActionChatEditTitle) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageActionChatEditTitle))

	x.StringBytes(m.Data2.Title)

	return x.buf
}

func (m *TLMessageActionChatEditTitle) Decode(dbuf *DecodeBuf) error {
	m.Data2.Title = x.StringBytes()

	return dbuf.err
}

// messageActionChatEditPhoto#7fcb13a8 photo:Photo = MessageAction;
func (m *TLMessageActionChatEditPhoto) To_MessageAction() *MessageAction {
	return &MessageAction{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessageActionChatEditPhoto) SetPhoto(v *Photo) { m.Data2.Photo = v }
func (m *TLMessageActionChatEditPhoto) GetPhoto() *Photo  { return m.Data2.Photo }

func (m *TLMessageActionChatEditPhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageActionChatEditPhoto))

	x.Bytes(m.Data2Photo.Encode())

	return x.buf
}

func (m *TLMessageActionChatEditPhoto) Decode(dbuf *DecodeBuf) error {
	m1 := &Photo{}
	m1.Decode(dbuf)
	m.SetPhoto(m1)

	return dbuf.err
}

// messageActionChatDeletePhoto#95e3fbef = MessageAction;
func (m *TLMessageActionChatDeletePhoto) To_MessageAction() *MessageAction {
	return &MessageAction{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessageActionChatDeletePhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageActionChatDeletePhoto))

	return x.buf
}

func (m *TLMessageActionChatDeletePhoto) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// messageActionChatAddUser#488a7337 users:Vector<int> = MessageAction;
func (m *TLMessageActionChatAddUser) To_MessageAction() *MessageAction {
	return &MessageAction{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessageActionChatAddUser) SetUsers(v []int32) { m.Data2.Users = v }
func (m *TLMessageActionChatAddUser) GetUsers() []int32  { return m.Data2.Users }

func (m *TLMessageActionChatAddUser) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageActionChatAddUser))

	return x.buf
}

func (m *TLMessageActionChatAddUser) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// messageActionChatDeleteUser#b2ae9b0c user_id:int = MessageAction;
func (m *TLMessageActionChatDeleteUser) To_MessageAction() *MessageAction {
	return &MessageAction{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessageActionChatDeleteUser) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLMessageActionChatDeleteUser) GetUserId() int32  { return m.Data2.UserId }

func (m *TLMessageActionChatDeleteUser) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageActionChatDeleteUser))

	x.Int(m.Data2.UserId)

	return x.buf
}

func (m *TLMessageActionChatDeleteUser) Decode(dbuf *DecodeBuf) error {
	m.Data2.UserId = x.Int()

	return dbuf.err
}

// messageActionChatJoinedByLink#f89cf5e8 inviter_id:int = MessageAction;
func (m *TLMessageActionChatJoinedByLink) To_MessageAction() *MessageAction {
	return &MessageAction{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessageActionChatJoinedByLink) SetInviterId(v int32) { m.Data2.InviterId = v }
func (m *TLMessageActionChatJoinedByLink) GetInviterId() int32  { return m.Data2.InviterId }

func (m *TLMessageActionChatJoinedByLink) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageActionChatJoinedByLink))

	x.Int(m.Data2.InviterId)

	return x.buf
}

func (m *TLMessageActionChatJoinedByLink) Decode(dbuf *DecodeBuf) error {
	m.Data2.InviterId = x.Int()

	return dbuf.err
}

// messageActionChannelCreate#95d2ac92 title:string = MessageAction;
func (m *TLMessageActionChannelCreate) To_MessageAction() *MessageAction {
	return &MessageAction{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessageActionChannelCreate) SetTitle(v string) { m.Data2.Title = v }
func (m *TLMessageActionChannelCreate) GetTitle() string  { return m.Data2.Title }

func (m *TLMessageActionChannelCreate) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageActionChannelCreate))

	x.StringBytes(m.Data2.Title)

	return x.buf
}

func (m *TLMessageActionChannelCreate) Decode(dbuf *DecodeBuf) error {
	m.Data2.Title = x.StringBytes()

	return dbuf.err
}

// messageActionChatMigrateTo#51bdb021 channel_id:int = MessageAction;
func (m *TLMessageActionChatMigrateTo) To_MessageAction() *MessageAction {
	return &MessageAction{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessageActionChatMigrateTo) SetChannelId(v int32) { m.Data2.ChannelId = v }
func (m *TLMessageActionChatMigrateTo) GetChannelId() int32  { return m.Data2.ChannelId }

func (m *TLMessageActionChatMigrateTo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageActionChatMigrateTo))

	x.Int(m.Data2.ChannelId)

	return x.buf
}

func (m *TLMessageActionChatMigrateTo) Decode(dbuf *DecodeBuf) error {
	m.Data2.ChannelId = x.Int()

	return dbuf.err
}

// messageActionChannelMigrateFrom#b055eaee title:string chat_id:int = MessageAction;
func (m *TLMessageActionChannelMigrateFrom) To_MessageAction() *MessageAction {
	return &MessageAction{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessageActionChannelMigrateFrom) SetTitle(v string) { m.Data2.Title = v }
func (m *TLMessageActionChannelMigrateFrom) GetTitle() string  { return m.Data2.Title }

func (m *TLMessageActionChannelMigrateFrom) SetChatId(v int32) { m.Data2.ChatId = v }
func (m *TLMessageActionChannelMigrateFrom) GetChatId() int32  { return m.Data2.ChatId }

func (m *TLMessageActionChannelMigrateFrom) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageActionChannelMigrateFrom))

	x.StringBytes(m.Data2.Title)
	x.Int(m.Data2.ChatId)

	return x.buf
}

func (m *TLMessageActionChannelMigrateFrom) Decode(dbuf *DecodeBuf) error {
	m.Data2.Title = x.StringBytes()
	m.Data2.ChatId = x.Int()

	return dbuf.err
}

// messageActionPinMessage#94bd38ed = MessageAction;
func (m *TLMessageActionPinMessage) To_MessageAction() *MessageAction {
	return &MessageAction{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessageActionPinMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageActionPinMessage))

	return x.buf
}

func (m *TLMessageActionPinMessage) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// messageActionHistoryClear#9fbab604 = MessageAction;
func (m *TLMessageActionHistoryClear) To_MessageAction() *MessageAction {
	return &MessageAction{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessageActionHistoryClear) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageActionHistoryClear))

	return x.buf
}

func (m *TLMessageActionHistoryClear) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// messageActionGameScore#92a72876 game_id:long score:int = MessageAction;
func (m *TLMessageActionGameScore) To_MessageAction() *MessageAction {
	return &MessageAction{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessageActionGameScore) SetGameId(v int64) { m.Data2.GameId = v }
func (m *TLMessageActionGameScore) GetGameId() int64  { return m.Data2.GameId }

func (m *TLMessageActionGameScore) SetScore(v int32) { m.Data2.Score = v }
func (m *TLMessageActionGameScore) GetScore() int32  { return m.Data2.Score }

func (m *TLMessageActionGameScore) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageActionGameScore))

	x.Long(m.Data2.GameId)
	x.Int(m.Data2.Score)

	return x.buf
}

func (m *TLMessageActionGameScore) Decode(dbuf *DecodeBuf) error {
	m.Data2.GameId = x.Long()
	m.Data2.Score = x.Int()

	return dbuf.err
}

// messageActionPaymentSentMe#8f31b327 flags:# currency:string total_amount:long payload:bytes info:flags.0?PaymentRequestedInfo shipping_option_id:flags.1?string charge:PaymentCharge = MessageAction;
func (m *TLMessageActionPaymentSentMe) To_MessageAction() *MessageAction {
	return &MessageAction{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessageActionPaymentSentMe) SetCurrency(v string) { m.Data2.Currency = v }
func (m *TLMessageActionPaymentSentMe) GetCurrency() string  { return m.Data2.Currency }

func (m *TLMessageActionPaymentSentMe) SetTotalAmount(v int64) { m.Data2.TotalAmount = v }
func (m *TLMessageActionPaymentSentMe) GetTotalAmount() int64  { return m.Data2.TotalAmount }

func (m *TLMessageActionPaymentSentMe) SetPayload(v []byte) { m.Data2.Payload = v }
func (m *TLMessageActionPaymentSentMe) GetPayload() []byte  { return m.Data2.Payload }

func (m *TLMessageActionPaymentSentMe) SetInfo(v *PaymentRequestedInfo) { m.Data2.Info = v }
func (m *TLMessageActionPaymentSentMe) GetInfo() *PaymentRequestedInfo  { return m.Data2.Info }

func (m *TLMessageActionPaymentSentMe) SetShippingOptionId(v string) { m.Data2.ShippingOptionId = v }
func (m *TLMessageActionPaymentSentMe) GetShippingOptionId() string  { return m.Data2.ShippingOptionId }

func (m *TLMessageActionPaymentSentMe) SetCharge(v *PaymentCharge) { m.Data2.Charge = v }
func (m *TLMessageActionPaymentSentMe) GetCharge() *PaymentCharge  { return m.Data2.Charge }

func (m *TLMessageActionPaymentSentMe) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageActionPaymentSentMe))

	// flags
	var flags uint32 = 0
	if m.GetInfo() != nil {
		flags |= 1 << 4
	}
	if m.GetShippingOptionId() != "" {
		flags |= 1 << 5
	}
	x.UInt(flags)

	x.StringBytes(m.Data2.Currency)
	x.Long(m.Data2.TotalAmount)
	x.StringBytes(m.Data2.Payload)
	if m.GetInfo() != 0 {
		x.Bytes(m.Data2Info.Encode())
	}
	if m.GetShippingOptionId() != 0 {
		x.StringBytes(m.Data2.ShippingOptionId)
	}
	x.Bytes(m.Data2Charge.Encode())

	return x.buf
}

func (m *TLMessageActionPaymentSentMe) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	m.Data2.Currency = x.StringBytes()
	m.Data2.TotalAmount = x.Long()
	m.Data2.Payload = x.StringBytes()
	if (flags & (1 << 5)) != 0 {
		m5 := &Info{}
		m5.Decode(dbuf)
		m.SetInfo(m5)
	}
	if (flags & (1 << 6)) != 0 {
		m.Data2.ShippingOptionId = x.StringBytes()
	}
	m7 := &Charge{}
	m7.Decode(dbuf)
	m.SetCharge(m7)

	return dbuf.err
}

// messageActionPaymentSent#40699cd0 currency:string total_amount:long = MessageAction;
func (m *TLMessageActionPaymentSent) To_MessageAction() *MessageAction {
	return &MessageAction{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessageActionPaymentSent) SetCurrency(v string) { m.Data2.Currency = v }
func (m *TLMessageActionPaymentSent) GetCurrency() string  { return m.Data2.Currency }

func (m *TLMessageActionPaymentSent) SetTotalAmount(v int64) { m.Data2.TotalAmount = v }
func (m *TLMessageActionPaymentSent) GetTotalAmount() int64  { return m.Data2.TotalAmount }

func (m *TLMessageActionPaymentSent) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageActionPaymentSent))

	x.StringBytes(m.Data2.Currency)
	x.Long(m.Data2.TotalAmount)

	return x.buf
}

func (m *TLMessageActionPaymentSent) Decode(dbuf *DecodeBuf) error {
	m.Data2.Currency = x.StringBytes()
	m.Data2.TotalAmount = x.Long()

	return dbuf.err
}

// messageActionPhoneCall#80e11a7f flags:# call_id:long reason:flags.0?PhoneCallDiscardReason duration:flags.1?int = MessageAction;
func (m *TLMessageActionPhoneCall) To_MessageAction() *MessageAction {
	return &MessageAction{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessageActionPhoneCall) SetCallId(v int64) { m.Data2.CallId = v }
func (m *TLMessageActionPhoneCall) GetCallId() int64  { return m.Data2.CallId }

func (m *TLMessageActionPhoneCall) SetReason(v *PhoneCallDiscardReason) { m.Data2.Reason = v }
func (m *TLMessageActionPhoneCall) GetReason() *PhoneCallDiscardReason  { return m.Data2.Reason }

func (m *TLMessageActionPhoneCall) SetDuration(v int32) { m.Data2.Duration = v }
func (m *TLMessageActionPhoneCall) GetDuration() int32  { return m.Data2.Duration }

func (m *TLMessageActionPhoneCall) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageActionPhoneCall))

	// flags
	var flags uint32 = 0
	if m.GetReason() != nil {
		flags |= 1 << 2
	}
	if m.GetDuration() != 0 {
		flags |= 1 << 3
	}
	x.UInt(flags)

	x.Long(m.Data2.CallId)
	if m.GetReason() != 0 {
		x.Bytes(m.Data2Reason.Encode())
	}
	if m.GetDuration() != 0 {
		x.Int(m.Data2.Duration)
	}

	return x.buf
}

func (m *TLMessageActionPhoneCall) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	m.Data2.CallId = x.Long()
	if (flags & (1 << 3)) != 0 {
		m3 := &Reason{}
		m3.Decode(dbuf)
		m.SetReason(m3)
	}
	if (flags & (1 << 4)) != 0 {
		m.Data2.Duration = x.Int()
	}

	return dbuf.err
}

// messageActionScreenshotTaken#4792929b = MessageAction;
func (m *TLMessageActionScreenshotTaken) To_MessageAction() *MessageAction {
	return &MessageAction{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessageActionScreenshotTaken) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageActionScreenshotTaken))

	return x.buf
}

func (m *TLMessageActionScreenshotTaken) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// InputNotifyPeer <--
//  + TL_InputNotifyPeer
//  + TL_InputNotifyUsers
//  + TL_InputNotifyChats
//  + TL_InputNotifyAll
//

func (m *InputNotifyPeer) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_inputNotifyPeer:
		t := m.To_InputNotifyPeer()
		return t.Encode()
	case TLConstructor_CRC32_inputNotifyUsers:
		t := m.To_InputNotifyUsers()
		return t.Encode()
	case TLConstructor_CRC32_inputNotifyChats:
		t := m.To_InputNotifyChats()
		return t.Encode()
	case TLConstructor_CRC32_inputNotifyAll:
		t := m.To_InputNotifyAll()
		return t.Encode()

	default:
		return []byte{}
	}
}

// inputNotifyPeer#b8bc5b0c peer:InputPeer = InputNotifyPeer;
func (m *InputNotifyPeer) To_InputNotifyPeer() *TLInputNotifyPeer {
	return &TLInputNotifyPeer{
		Data2: m.Data2,
	}
}

// inputNotifyUsers#193b4417 = InputNotifyPeer;
func (m *InputNotifyPeer) To_InputNotifyUsers() *TLInputNotifyUsers {
	return &TLInputNotifyUsers{
		Data2: m.Data2,
	}
}

// inputNotifyChats#4a95e84e = InputNotifyPeer;
func (m *InputNotifyPeer) To_InputNotifyChats() *TLInputNotifyChats {
	return &TLInputNotifyChats{
		Data2: m.Data2,
	}
}

// inputNotifyAll#a429b886 = InputNotifyPeer;
func (m *InputNotifyPeer) To_InputNotifyAll() *TLInputNotifyAll {
	return &TLInputNotifyAll{
		Data2: m.Data2,
	}
}

// inputNotifyPeer#b8bc5b0c peer:InputPeer = InputNotifyPeer;
func (m *TLInputNotifyPeer) To_InputNotifyPeer() *InputNotifyPeer {
	return &InputNotifyPeer{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputNotifyPeer) SetPeer(v *InputPeer) { m.Data2.Peer = v }
func (m *TLInputNotifyPeer) GetPeer() *InputPeer  { return m.Data2.Peer }

func (m *TLInputNotifyPeer) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputNotifyPeer))

	x.Bytes(m.Data2Peer.Encode())

	return x.buf
}

func (m *TLInputNotifyPeer) Decode(dbuf *DecodeBuf) error {
	m1 := &Peer{}
	m1.Decode(dbuf)
	m.SetPeer(m1)

	return dbuf.err
}

// inputNotifyUsers#193b4417 = InputNotifyPeer;
func (m *TLInputNotifyUsers) To_InputNotifyPeer() *InputNotifyPeer {
	return &InputNotifyPeer{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputNotifyUsers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputNotifyUsers))

	return x.buf
}

func (m *TLInputNotifyUsers) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// inputNotifyChats#4a95e84e = InputNotifyPeer;
func (m *TLInputNotifyChats) To_InputNotifyPeer() *InputNotifyPeer {
	return &InputNotifyPeer{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputNotifyChats) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputNotifyChats))

	return x.buf
}

func (m *TLInputNotifyChats) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// inputNotifyAll#a429b886 = InputNotifyPeer;
func (m *TLInputNotifyAll) To_InputNotifyPeer() *InputNotifyPeer {
	return &InputNotifyPeer{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputNotifyAll) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputNotifyAll))

	return x.buf
}

func (m *TLInputNotifyAll) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// AccountDaysTTL <--
//  + TL_AccountDaysTTL
//

func (m *AccountDaysTTL) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_accountDaysTTL:
		t := m.To_AccountDaysTTL()
		return t.Encode()

	default:
		return []byte{}
	}
}

// accountDaysTTL#b8d0afdf days:int = AccountDaysTTL;
func (m *AccountDaysTTL) To_AccountDaysTTL() *TLAccountDaysTTL {
	return &TLAccountDaysTTL{
		Data2: m.Data2,
	}
}

// accountDaysTTL#b8d0afdf days:int = AccountDaysTTL;
func (m *TLAccountDaysTTL) To_AccountDaysTTL() *AccountDaysTTL {
	return &AccountDaysTTL{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLAccountDaysTTL) SetDays(v int32) { m.Data2.Days = v }
func (m *TLAccountDaysTTL) GetDays() int32  { return m.Data2.Days }

func (m *TLAccountDaysTTL) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_accountDaysTTL))

	x.Int(m.Data2.Days)

	return x.buf
}

func (m *TLAccountDaysTTL) Decode(dbuf *DecodeBuf) error {
	m.Data2.Days = x.Int()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Messages_AllStickers <--
//  + TL_MessagesAllStickersNotModified
//  + TL_MessagesAllStickers
//

func (m *Messages_AllStickers) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_messages_allStickersNotModified:
		t := m.To_MessagesAllStickersNotModified()
		return t.Encode()
	case TLConstructor_CRC32_messages_allStickers:
		t := m.To_MessagesAllStickers()
		return t.Encode()

	default:
		return []byte{}
	}
}

// messages.allStickersNotModified#e86602c3 = messages.AllStickers;
func (m *Messages_AllStickers) To_MessagesAllStickersNotModified() *TLMessagesAllStickersNotModified {
	return &TLMessagesAllStickersNotModified{
		Data2: m.Data2,
	}
}

// messages.allStickers#edfd405f hash:int sets:Vector<StickerSet> = messages.AllStickers;
func (m *Messages_AllStickers) To_MessagesAllStickers() *TLMessagesAllStickers {
	return &TLMessagesAllStickers{
		Data2: m.Data2,
	}
}

// messages.allStickersNotModified#e86602c3 = messages.AllStickers;
func (m *TLMessagesAllStickersNotModified) To_Messages_AllStickers() *Messages_AllStickers {
	return &Messages_AllStickers{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessagesAllStickersNotModified) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_allStickersNotModified))

	return x.buf
}

func (m *TLMessagesAllStickersNotModified) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// messages.allStickers#edfd405f hash:int sets:Vector<StickerSet> = messages.AllStickers;
func (m *TLMessagesAllStickers) To_Messages_AllStickers() *Messages_AllStickers {
	return &Messages_AllStickers{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessagesAllStickers) SetHash(v int32) { m.Data2.Hash = v }
func (m *TLMessagesAllStickers) GetHash() int32  { return m.Data2.Hash }

func (m *TLMessagesAllStickers) SetSets(v []*StickerSet) { m.Data2.Sets = v }
func (m *TLMessagesAllStickers) GetSets() []*StickerSet  { return m.Data2.Sets }

func (m *TLMessagesAllStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_allStickers))

	x.Int(m.Data2.Hash)

	return x.buf
}

func (m *TLMessagesAllStickers) Decode(dbuf *DecodeBuf) error {
	m.Data2.Hash = x.Int()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// ChannelAdminLogEventsFilter <--
//  + TL_ChannelAdminLogEventsFilter
//

func (m *ChannelAdminLogEventsFilter) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_channelAdminLogEventsFilter:
		t := m.To_ChannelAdminLogEventsFilter()
		return t.Encode()

	default:
		return []byte{}
	}
}

// channelAdminLogEventsFilter#ea107ae4 flags:# join:flags.0?true leave:flags.1?true invite:flags.2?true ban:flags.3?true unban:flags.4?true kick:flags.5?true unkick:flags.6?true promote:flags.7?true demote:flags.8?true info:flags.9?true settings:flags.10?true pinned:flags.11?true edit:flags.12?true delete:flags.13?true = ChannelAdminLogEventsFilter;
func (m *ChannelAdminLogEventsFilter) To_ChannelAdminLogEventsFilter() *TLChannelAdminLogEventsFilter {
	return &TLChannelAdminLogEventsFilter{
		Data2: m.Data2,
	}
}

// channelAdminLogEventsFilter#ea107ae4 flags:# join:flags.0?true leave:flags.1?true invite:flags.2?true ban:flags.3?true unban:flags.4?true kick:flags.5?true unkick:flags.6?true promote:flags.7?true demote:flags.8?true info:flags.9?true settings:flags.10?true pinned:flags.11?true edit:flags.12?true delete:flags.13?true = ChannelAdminLogEventsFilter;
func (m *TLChannelAdminLogEventsFilter) To_ChannelAdminLogEventsFilter() *ChannelAdminLogEventsFilter {
	return &ChannelAdminLogEventsFilter{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLChannelAdminLogEventsFilter) SetJoin(v bool) { m.Data2.Join = v }
func (m *TLChannelAdminLogEventsFilter) GetJoin() bool  { return m.Data2.Join }

func (m *TLChannelAdminLogEventsFilter) SetLeave(v bool) { m.Data2.Leave = v }
func (m *TLChannelAdminLogEventsFilter) GetLeave() bool  { return m.Data2.Leave }

func (m *TLChannelAdminLogEventsFilter) SetInvite(v bool) { m.Data2.Invite = v }
func (m *TLChannelAdminLogEventsFilter) GetInvite() bool  { return m.Data2.Invite }

func (m *TLChannelAdminLogEventsFilter) SetBan(v bool) { m.Data2.Ban = v }
func (m *TLChannelAdminLogEventsFilter) GetBan() bool  { return m.Data2.Ban }

func (m *TLChannelAdminLogEventsFilter) SetUnban(v bool) { m.Data2.Unban = v }
func (m *TLChannelAdminLogEventsFilter) GetUnban() bool  { return m.Data2.Unban }

func (m *TLChannelAdminLogEventsFilter) SetKick(v bool) { m.Data2.Kick = v }
func (m *TLChannelAdminLogEventsFilter) GetKick() bool  { return m.Data2.Kick }

func (m *TLChannelAdminLogEventsFilter) SetUnkick(v bool) { m.Data2.Unkick = v }
func (m *TLChannelAdminLogEventsFilter) GetUnkick() bool  { return m.Data2.Unkick }

func (m *TLChannelAdminLogEventsFilter) SetPromote(v bool) { m.Data2.Promote = v }
func (m *TLChannelAdminLogEventsFilter) GetPromote() bool  { return m.Data2.Promote }

func (m *TLChannelAdminLogEventsFilter) SetDemote(v bool) { m.Data2.Demote = v }
func (m *TLChannelAdminLogEventsFilter) GetDemote() bool  { return m.Data2.Demote }

func (m *TLChannelAdminLogEventsFilter) SetInfo(v bool) { m.Data2.Info = v }
func (m *TLChannelAdminLogEventsFilter) GetInfo() bool  { return m.Data2.Info }

func (m *TLChannelAdminLogEventsFilter) SetSettings(v bool) { m.Data2.Settings = v }
func (m *TLChannelAdminLogEventsFilter) GetSettings() bool  { return m.Data2.Settings }

func (m *TLChannelAdminLogEventsFilter) SetPinned(v bool) { m.Data2.Pinned = v }
func (m *TLChannelAdminLogEventsFilter) GetPinned() bool  { return m.Data2.Pinned }

func (m *TLChannelAdminLogEventsFilter) SetEdit(v bool) { m.Data2.Edit = v }
func (m *TLChannelAdminLogEventsFilter) GetEdit() bool  { return m.Data2.Edit }

func (m *TLChannelAdminLogEventsFilter) SetDelete(v bool) { m.Data2.Delete = v }
func (m *TLChannelAdminLogEventsFilter) GetDelete() bool  { return m.Data2.Delete }

func (m *TLChannelAdminLogEventsFilter) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelAdminLogEventsFilter))

	// flags
	var flags uint32 = 0
	if m.GetJoin() == true {
		flags |= 1 << 1
	}
	if m.GetLeave() == true {
		flags |= 1 << 2
	}
	if m.GetInvite() == true {
		flags |= 1 << 3
	}
	if m.GetBan() == true {
		flags |= 1 << 4
	}
	if m.GetUnban() == true {
		flags |= 1 << 5
	}
	if m.GetKick() == true {
		flags |= 1 << 6
	}
	if m.GetUnkick() == true {
		flags |= 1 << 7
	}
	if m.GetPromote() == true {
		flags |= 1 << 8
	}
	if m.GetDemote() == true {
		flags |= 1 << 9
	}
	if m.GetInfo() == true {
		flags |= 1 << 10
	}
	if m.GetSettings() == true {
		flags |= 1 << 11
	}
	if m.GetPinned() == true {
		flags |= 1 << 12
	}
	if m.GetEdit() == true {
		flags |= 1 << 13
	}
	if m.GetDelete() == true {
		flags |= 1 << 14
	}
	x.UInt(flags)

	return x.buf
}

func (m *TLChannelAdminLogEventsFilter) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m.Data2.Join = true
	}
	if (flags & (1 << 3)) != 0 {
		m.Data2.Leave = true
	}
	if (flags & (1 << 4)) != 0 {
		m.Data2.Invite = true
	}
	if (flags & (1 << 5)) != 0 {
		m.Data2.Ban = true
	}
	if (flags & (1 << 6)) != 0 {
		m.Data2.Unban = true
	}
	if (flags & (1 << 7)) != 0 {
		m.Data2.Kick = true
	}
	if (flags & (1 << 8)) != 0 {
		m.Data2.Unkick = true
	}
	if (flags & (1 << 9)) != 0 {
		m.Data2.Promote = true
	}
	if (flags & (1 << 10)) != 0 {
		m.Data2.Demote = true
	}
	if (flags & (1 << 11)) != 0 {
		m.Data2.Info = true
	}
	if (flags & (1 << 12)) != 0 {
		m.Data2.Settings = true
	}
	if (flags & (1 << 13)) != 0 {
		m.Data2.Pinned = true
	}
	if (flags & (1 << 14)) != 0 {
		m.Data2.Edit = true
	}
	if (flags & (1 << 15)) != 0 {
		m.Data2.Delete = true
	}

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Contacts_Found <--
//  + TL_ContactsFound
//

func (m *Contacts_Found) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_contacts_found:
		t := m.To_ContactsFound()
		return t.Encode()

	default:
		return []byte{}
	}
}

// contacts.found#1aa1f784 results:Vector<Peer> chats:Vector<Chat> users:Vector<User> = contacts.Found;
func (m *Contacts_Found) To_ContactsFound() *TLContactsFound {
	return &TLContactsFound{
		Data2: m.Data2,
	}
}

// contacts.found#1aa1f784 results:Vector<Peer> chats:Vector<Chat> users:Vector<User> = contacts.Found;
func (m *TLContactsFound) To_Contacts_Found() *Contacts_Found {
	return &Contacts_Found{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLContactsFound) SetResults(v []*Peer) { m.Data2.Results = v }
func (m *TLContactsFound) GetResults() []*Peer  { return m.Data2.Results }

func (m *TLContactsFound) SetChats(v []*Chat) { m.Data2.Chats = v }
func (m *TLContactsFound) GetChats() []*Chat  { return m.Data2.Chats }

func (m *TLContactsFound) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLContactsFound) GetUsers() []*User  { return m.Data2.Users }

func (m *TLContactsFound) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_contacts_found))

	return x.buf
}

func (m *TLContactsFound) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Payments_PaymentResult <--
//  + TL_PaymentsPaymentResult
//  + TL_PaymentsPaymentVerficationNeeded
//

func (m *Payments_PaymentResult) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_payments_paymentResult:
		t := m.To_PaymentsPaymentResult()
		return t.Encode()
	case TLConstructor_CRC32_payments_paymentVerficationNeeded:
		t := m.To_PaymentsPaymentVerficationNeeded()
		return t.Encode()

	default:
		return []byte{}
	}
}

// payments.paymentResult#4e5f810d updates:Updates = payments.PaymentResult;
func (m *Payments_PaymentResult) To_PaymentsPaymentResult() *TLPaymentsPaymentResult {
	return &TLPaymentsPaymentResult{
		Data2: m.Data2,
	}
}

// payments.paymentVerficationNeeded#6b56b921 url:string = payments.PaymentResult;
func (m *Payments_PaymentResult) To_PaymentsPaymentVerficationNeeded() *TLPaymentsPaymentVerficationNeeded {
	return &TLPaymentsPaymentVerficationNeeded{
		Data2: m.Data2,
	}
}

// payments.paymentResult#4e5f810d updates:Updates = payments.PaymentResult;
func (m *TLPaymentsPaymentResult) To_Payments_PaymentResult() *Payments_PaymentResult {
	return &Payments_PaymentResult{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPaymentsPaymentResult) SetUpdates(v *Updates) { m.Data2.Updates = v }
func (m *TLPaymentsPaymentResult) GetUpdates() *Updates  { return m.Data2.Updates }

func (m *TLPaymentsPaymentResult) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_payments_paymentResult))

	x.Bytes(m.Data2Updates.Encode())

	return x.buf
}

func (m *TLPaymentsPaymentResult) Decode(dbuf *DecodeBuf) error {
	m1 := &Updates{}
	m1.Decode(dbuf)
	m.SetUpdates(m1)

	return dbuf.err
}

// payments.paymentVerficationNeeded#6b56b921 url:string = payments.PaymentResult;
func (m *TLPaymentsPaymentVerficationNeeded) To_Payments_PaymentResult() *Payments_PaymentResult {
	return &Payments_PaymentResult{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPaymentsPaymentVerficationNeeded) SetUrl(v string) { m.Data2.Url = v }
func (m *TLPaymentsPaymentVerficationNeeded) GetUrl() string  { return m.Data2.Url }

func (m *TLPaymentsPaymentVerficationNeeded) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_payments_paymentVerficationNeeded))

	x.StringBytes(m.Data2.Url)

	return x.buf
}

func (m *TLPaymentsPaymentVerficationNeeded) Decode(dbuf *DecodeBuf) error {
	m.Data2.Url = x.StringBytes()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// CdnPublicKey <--
//  + TL_CdnPublicKey
//

func (m *CdnPublicKey) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_cdnPublicKey:
		t := m.To_CdnPublicKey()
		return t.Encode()

	default:
		return []byte{}
	}
}

// cdnPublicKey#c982eaba dc_id:int public_key:string = CdnPublicKey;
func (m *CdnPublicKey) To_CdnPublicKey() *TLCdnPublicKey {
	return &TLCdnPublicKey{
		Data2: m.Data2,
	}
}

// cdnPublicKey#c982eaba dc_id:int public_key:string = CdnPublicKey;
func (m *TLCdnPublicKey) To_CdnPublicKey() *CdnPublicKey {
	return &CdnPublicKey{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLCdnPublicKey) SetDcId(v int32) { m.Data2.DcId = v }
func (m *TLCdnPublicKey) GetDcId() int32  { return m.Data2.DcId }

func (m *TLCdnPublicKey) SetPublicKey(v string) { m.Data2.PublicKey = v }
func (m *TLCdnPublicKey) GetPublicKey() string  { return m.Data2.PublicKey }

func (m *TLCdnPublicKey) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_cdnPublicKey))

	x.Int(m.Data2.DcId)
	x.StringBytes(m.Data2.PublicKey)

	return x.buf
}

func (m *TLCdnPublicKey) Decode(dbuf *DecodeBuf) error {
	m.Data2.DcId = x.Int()
	m.Data2.PublicKey = x.StringBytes()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Bool <--
//  + TL_BoolFalse
//  + TL_BoolTrue
//

func (m *Bool) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_boolFalse:
		t := m.To_BoolFalse()
		return t.Encode()
	case TLConstructor_CRC32_boolTrue:
		t := m.To_BoolTrue()
		return t.Encode()

	default:
		return []byte{}
	}
}

// boolFalse#bc799737 = Bool;
func (m *Bool) To_BoolFalse() *TLBoolFalse {
	return &TLBoolFalse{
		Data2: m.Data2,
	}
}

// boolTrue#997275b5 = Bool;
func (m *Bool) To_BoolTrue() *TLBoolTrue {
	return &TLBoolTrue{
		Data2: m.Data2,
	}
}

// boolFalse#bc799737 = Bool;
func (m *TLBoolFalse) To_Bool() *Bool {
	return &Bool{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLBoolFalse) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_boolFalse))

	return x.buf
}

func (m *TLBoolFalse) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// boolTrue#997275b5 = Bool;
func (m *TLBoolTrue) To_Bool() *Bool {
	return &Bool{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLBoolTrue) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_boolTrue))

	return x.buf
}

func (m *TLBoolTrue) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Auth_SentCode <--
//  + TL_AuthSentCode
//

func (m *Auth_SentCode) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_auth_sentCode:
		t := m.To_AuthSentCode()
		return t.Encode()

	default:
		return []byte{}
	}
}

// auth.sentCode#5e002502 flags:# phone_registered:flags.0?true type:auth.SentCodeType phone_code_hash:string next_type:flags.1?auth.CodeType timeout:flags.2?int = auth.SentCode;
func (m *Auth_SentCode) To_AuthSentCode() *TLAuthSentCode {
	return &TLAuthSentCode{
		Data2: m.Data2,
	}
}

// auth.sentCode#5e002502 flags:# phone_registered:flags.0?true type:auth.SentCodeType phone_code_hash:string next_type:flags.1?auth.CodeType timeout:flags.2?int = auth.SentCode;
func (m *TLAuthSentCode) To_Auth_SentCode() *Auth_SentCode {
	return &Auth_SentCode{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLAuthSentCode) SetPhoneRegistered(v bool) { m.Data2.PhoneRegistered = v }
func (m *TLAuthSentCode) GetPhoneRegistered() bool  { return m.Data2.PhoneRegistered }

func (m *TLAuthSentCode) SetType(v *Auth_SentCodeType) { m.Data2.Type = v }
func (m *TLAuthSentCode) GetType() *Auth_SentCodeType  { return m.Data2.Type }

func (m *TLAuthSentCode) SetPhoneCodeHash(v string) { m.Data2.PhoneCodeHash = v }
func (m *TLAuthSentCode) GetPhoneCodeHash() string  { return m.Data2.PhoneCodeHash }

func (m *TLAuthSentCode) SetNextType(v *Auth_CodeType) { m.Data2.NextType = v }
func (m *TLAuthSentCode) GetNextType() *Auth_CodeType  { return m.Data2.NextType }

func (m *TLAuthSentCode) SetTimeout(v int32) { m.Data2.Timeout = v }
func (m *TLAuthSentCode) GetTimeout() int32  { return m.Data2.Timeout }

func (m *TLAuthSentCode) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_auth_sentCode))

	// flags
	var flags uint32 = 0
	if m.GetPhoneRegistered() == true {
		flags |= 1 << 1
	}
	if m.GetNextType() != nil {
		flags |= 1 << 4
	}
	if m.GetTimeout() != 0 {
		flags |= 1 << 5
	}
	x.UInt(flags)

	x.Bytes(m.Data2Type.Encode())
	x.StringBytes(m.Data2.PhoneCodeHash)
	if m.GetNextType() != 0 {
		x.Bytes(m.Data2NextType.Encode())
	}
	if m.GetTimeout() != 0 {
		x.Int(m.Data2.Timeout)
	}

	return x.buf
}

func (m *TLAuthSentCode) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m.Data2.PhoneRegistered = true
	}
	m3 := &Type{}
	m3.Decode(dbuf)
	m.SetType(m3)
	m.Data2.PhoneCodeHash = x.StringBytes()
	if (flags & (1 << 5)) != 0 {
		m5 := &NextType{}
		m5.Decode(dbuf)
		m.SetNextType(m5)
	}
	if (flags & (1 << 6)) != 0 {
		m.Data2.Timeout = x.Int()
	}

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Auth_ExportedAuthorization <--
//  + TL_AuthExportedAuthorization
//

func (m *Auth_ExportedAuthorization) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_auth_exportedAuthorization:
		t := m.To_AuthExportedAuthorization()
		return t.Encode()

	default:
		return []byte{}
	}
}

// auth.exportedAuthorization#df969c2d id:int bytes:bytes = auth.ExportedAuthorization;
func (m *Auth_ExportedAuthorization) To_AuthExportedAuthorization() *TLAuthExportedAuthorization {
	return &TLAuthExportedAuthorization{
		Data2: m.Data2,
	}
}

// auth.exportedAuthorization#df969c2d id:int bytes:bytes = auth.ExportedAuthorization;
func (m *TLAuthExportedAuthorization) To_Auth_ExportedAuthorization() *Auth_ExportedAuthorization {
	return &Auth_ExportedAuthorization{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLAuthExportedAuthorization) SetId(v int32) { m.Data2.Id = v }
func (m *TLAuthExportedAuthorization) GetId() int32  { return m.Data2.Id }

func (m *TLAuthExportedAuthorization) SetBytes(v []byte) { m.Data2.Bytes = v }
func (m *TLAuthExportedAuthorization) GetBytes() []byte  { return m.Data2.Bytes }

func (m *TLAuthExportedAuthorization) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_auth_exportedAuthorization))

	x.Int(m.Data2.Id)
	x.StringBytes(m.Data2.Bytes)

	return x.buf
}

func (m *TLAuthExportedAuthorization) Decode(dbuf *DecodeBuf) error {
	m.Data2.Id = x.Int()
	m.Data2.Bytes = x.StringBytes()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// PeerNotifyEvents <--
//  + TL_PeerNotifyEventsEmpty
//  + TL_PeerNotifyEventsAll
//

func (m *PeerNotifyEvents) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_peerNotifyEventsEmpty:
		t := m.To_PeerNotifyEventsEmpty()
		return t.Encode()
	case TLConstructor_CRC32_peerNotifyEventsAll:
		t := m.To_PeerNotifyEventsAll()
		return t.Encode()

	default:
		return []byte{}
	}
}

// peerNotifyEventsEmpty#add53cb3 = PeerNotifyEvents;
func (m *PeerNotifyEvents) To_PeerNotifyEventsEmpty() *TLPeerNotifyEventsEmpty {
	return &TLPeerNotifyEventsEmpty{
		Data2: m.Data2,
	}
}

// peerNotifyEventsAll#6d1ded88 = PeerNotifyEvents;
func (m *PeerNotifyEvents) To_PeerNotifyEventsAll() *TLPeerNotifyEventsAll {
	return &TLPeerNotifyEventsAll{
		Data2: m.Data2,
	}
}

// peerNotifyEventsEmpty#add53cb3 = PeerNotifyEvents;
func (m *TLPeerNotifyEventsEmpty) To_PeerNotifyEvents() *PeerNotifyEvents {
	return &PeerNotifyEvents{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPeerNotifyEventsEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_peerNotifyEventsEmpty))

	return x.buf
}

func (m *TLPeerNotifyEventsEmpty) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// peerNotifyEventsAll#6d1ded88 = PeerNotifyEvents;
func (m *TLPeerNotifyEventsAll) To_PeerNotifyEvents() *PeerNotifyEvents {
	return &PeerNotifyEvents{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPeerNotifyEventsAll) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_peerNotifyEventsAll))

	return x.buf
}

func (m *TLPeerNotifyEventsAll) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// InputChatPhoto <--
//  + TL_InputChatPhotoEmpty
//  + TL_InputChatUploadedPhoto
//  + TL_InputChatPhoto
//

func (m *InputChatPhoto) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_inputChatPhotoEmpty:
		t := m.To_InputChatPhotoEmpty()
		return t.Encode()
	case TLConstructor_CRC32_inputChatUploadedPhoto:
		t := m.To_InputChatUploadedPhoto()
		return t.Encode()
	case TLConstructor_CRC32_inputChatPhoto:
		t := m.To_InputChatPhoto()
		return t.Encode()

	default:
		return []byte{}
	}
}

// inputChatPhotoEmpty#1ca48f57 = InputChatPhoto;
func (m *InputChatPhoto) To_InputChatPhotoEmpty() *TLInputChatPhotoEmpty {
	return &TLInputChatPhotoEmpty{
		Data2: m.Data2,
	}
}

// inputChatUploadedPhoto#927c55b4 file:InputFile = InputChatPhoto;
func (m *InputChatPhoto) To_InputChatUploadedPhoto() *TLInputChatUploadedPhoto {
	return &TLInputChatUploadedPhoto{
		Data2: m.Data2,
	}
}

// inputChatPhoto#8953ad37 id:InputPhoto = InputChatPhoto;
func (m *InputChatPhoto) To_InputChatPhoto() *TLInputChatPhoto {
	return &TLInputChatPhoto{
		Data2: m.Data2,
	}
}

// inputChatPhotoEmpty#1ca48f57 = InputChatPhoto;
func (m *TLInputChatPhotoEmpty) To_InputChatPhoto() *InputChatPhoto {
	return &InputChatPhoto{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputChatPhotoEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputChatPhotoEmpty))

	return x.buf
}

func (m *TLInputChatPhotoEmpty) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// inputChatUploadedPhoto#927c55b4 file:InputFile = InputChatPhoto;
func (m *TLInputChatUploadedPhoto) To_InputChatPhoto() *InputChatPhoto {
	return &InputChatPhoto{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputChatUploadedPhoto) SetFile(v *InputFile) { m.Data2.File = v }
func (m *TLInputChatUploadedPhoto) GetFile() *InputFile  { return m.Data2.File }

func (m *TLInputChatUploadedPhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputChatUploadedPhoto))

	x.Bytes(m.Data2File.Encode())

	return x.buf
}

func (m *TLInputChatUploadedPhoto) Decode(dbuf *DecodeBuf) error {
	m1 := &File{}
	m1.Decode(dbuf)
	m.SetFile(m1)

	return dbuf.err
}

// inputChatPhoto#8953ad37 id:InputPhoto = InputChatPhoto;
func (m *TLInputChatPhoto) To_InputChatPhoto() *InputChatPhoto {
	return &InputChatPhoto{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputChatPhoto) SetId(v *InputPhoto) { m.Data2.Id = v }
func (m *TLInputChatPhoto) GetId() *InputPhoto  { return m.Data2.Id }

func (m *TLInputChatPhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputChatPhoto))

	x.Bytes(m.Data2Id.Encode())

	return x.buf
}

func (m *TLInputChatPhoto) Decode(dbuf *DecodeBuf) error {
	m1 := &Id{}
	m1.Decode(dbuf)
	m.SetId(m1)

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Account_PrivacyRules <--
//  + TL_AccountPrivacyRules
//

func (m *Account_PrivacyRules) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_account_privacyRules:
		t := m.To_AccountPrivacyRules()
		return t.Encode()

	default:
		return []byte{}
	}
}

// account.privacyRules#554abb6f rules:Vector<PrivacyRule> users:Vector<User> = account.PrivacyRules;
func (m *Account_PrivacyRules) To_AccountPrivacyRules() *TLAccountPrivacyRules {
	return &TLAccountPrivacyRules{
		Data2: m.Data2,
	}
}

// account.privacyRules#554abb6f rules:Vector<PrivacyRule> users:Vector<User> = account.PrivacyRules;
func (m *TLAccountPrivacyRules) To_Account_PrivacyRules() *Account_PrivacyRules {
	return &Account_PrivacyRules{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLAccountPrivacyRules) SetRules(v []*PrivacyRule) { m.Data2.Rules = v }
func (m *TLAccountPrivacyRules) GetRules() []*PrivacyRule  { return m.Data2.Rules }

func (m *TLAccountPrivacyRules) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLAccountPrivacyRules) GetUsers() []*User  { return m.Data2.Users }

func (m *TLAccountPrivacyRules) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_account_privacyRules))

	return x.buf
}

func (m *TLAccountPrivacyRules) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Messages_HighScores <--
//  + TL_MessagesHighScores
//

func (m *Messages_HighScores) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_messages_highScores:
		t := m.To_MessagesHighScores()
		return t.Encode()

	default:
		return []byte{}
	}
}

// messages.highScores#9a3bfd99 scores:Vector<HighScore> users:Vector<User> = messages.HighScores;
func (m *Messages_HighScores) To_MessagesHighScores() *TLMessagesHighScores {
	return &TLMessagesHighScores{
		Data2: m.Data2,
	}
}

// messages.highScores#9a3bfd99 scores:Vector<HighScore> users:Vector<User> = messages.HighScores;
func (m *TLMessagesHighScores) To_Messages_HighScores() *Messages_HighScores {
	return &Messages_HighScores{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessagesHighScores) SetScores(v []*HighScore) { m.Data2.Scores = v }
func (m *TLMessagesHighScores) GetScores() []*HighScore  { return m.Data2.Scores }

func (m *TLMessagesHighScores) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLMessagesHighScores) GetUsers() []*User  { return m.Data2.Users }

func (m *TLMessagesHighScores) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_highScores))

	return x.buf
}

func (m *TLMessagesHighScores) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Channels_AdminLogResults <--
//  + TL_ChannelsAdminLogResults
//

func (m *Channels_AdminLogResults) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_channels_adminLogResults:
		t := m.To_ChannelsAdminLogResults()
		return t.Encode()

	default:
		return []byte{}
	}
}

// channels.adminLogResults#ed8af74d events:Vector<ChannelAdminLogEvent> chats:Vector<Chat> users:Vector<User> = channels.AdminLogResults;
func (m *Channels_AdminLogResults) To_ChannelsAdminLogResults() *TLChannelsAdminLogResults {
	return &TLChannelsAdminLogResults{
		Data2: m.Data2,
	}
}

// channels.adminLogResults#ed8af74d events:Vector<ChannelAdminLogEvent> chats:Vector<Chat> users:Vector<User> = channels.AdminLogResults;
func (m *TLChannelsAdminLogResults) To_Channels_AdminLogResults() *Channels_AdminLogResults {
	return &Channels_AdminLogResults{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLChannelsAdminLogResults) SetEvents(v []*ChannelAdminLogEvent) { m.Data2.Events = v }
func (m *TLChannelsAdminLogResults) GetEvents() []*ChannelAdminLogEvent  { return m.Data2.Events }

func (m *TLChannelsAdminLogResults) SetChats(v []*Chat) { m.Data2.Chats = v }
func (m *TLChannelsAdminLogResults) GetChats() []*Chat  { return m.Data2.Chats }

func (m *TLChannelsAdminLogResults) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLChannelsAdminLogResults) GetUsers() []*User  { return m.Data2.Users }

func (m *TLChannelsAdminLogResults) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channels_adminLogResults))

	return x.buf
}

func (m *TLChannelsAdminLogResults) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Payments_SavedInfo <--
//  + TL_PaymentsSavedInfo
//

func (m *Payments_SavedInfo) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_payments_savedInfo:
		t := m.To_PaymentsSavedInfo()
		return t.Encode()

	default:
		return []byte{}
	}
}

// payments.savedInfo#fb8fe43c flags:# has_saved_credentials:flags.1?true saved_info:flags.0?PaymentRequestedInfo = payments.SavedInfo;
func (m *Payments_SavedInfo) To_PaymentsSavedInfo() *TLPaymentsSavedInfo {
	return &TLPaymentsSavedInfo{
		Data2: m.Data2,
	}
}

// payments.savedInfo#fb8fe43c flags:# has_saved_credentials:flags.1?true saved_info:flags.0?PaymentRequestedInfo = payments.SavedInfo;
func (m *TLPaymentsSavedInfo) To_Payments_SavedInfo() *Payments_SavedInfo {
	return &Payments_SavedInfo{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPaymentsSavedInfo) SetHasSavedCredentials(v bool) { m.Data2.HasSavedCredentials = v }
func (m *TLPaymentsSavedInfo) GetHasSavedCredentials() bool  { return m.Data2.HasSavedCredentials }

func (m *TLPaymentsSavedInfo) SetSavedInfo(v *PaymentRequestedInfo) { m.Data2.SavedInfo = v }
func (m *TLPaymentsSavedInfo) GetSavedInfo() *PaymentRequestedInfo  { return m.Data2.SavedInfo }

func (m *TLPaymentsSavedInfo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_payments_savedInfo))

	// flags
	var flags uint32 = 0
	if m.GetHasSavedCredentials() == true {
		flags |= 1 << 1
	}
	if m.GetSavedInfo() != nil {
		flags |= 1 << 2
	}
	x.UInt(flags)

	if m.GetSavedInfo() != 0 {
		x.Bytes(m.Data2SavedInfo.Encode())
	}

	return x.buf
}

func (m *TLPaymentsSavedInfo) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m.Data2.HasSavedCredentials = true
	}
	if (flags & (1 << 3)) != 0 {
		m3 := &SavedInfo{}
		m3.Decode(dbuf)
		m.SetSavedInfo(m3)
	}

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// PhoneConnection <--
//  + TL_PhoneConnection
//

func (m *PhoneConnection) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_phoneConnection:
		t := m.To_PhoneConnection()
		return t.Encode()

	default:
		return []byte{}
	}
}

// phoneConnection#9d4c17c0 id:long ip:string ipv6:string port:int peer_tag:bytes = PhoneConnection;
func (m *PhoneConnection) To_PhoneConnection() *TLPhoneConnection {
	return &TLPhoneConnection{
		Data2: m.Data2,
	}
}

// phoneConnection#9d4c17c0 id:long ip:string ipv6:string port:int peer_tag:bytes = PhoneConnection;
func (m *TLPhoneConnection) To_PhoneConnection() *PhoneConnection {
	return &PhoneConnection{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPhoneConnection) SetId(v int64) { m.Data2.Id = v }
func (m *TLPhoneConnection) GetId() int64  { return m.Data2.Id }

func (m *TLPhoneConnection) SetIp(v string) { m.Data2.Ip = v }
func (m *TLPhoneConnection) GetIp() string  { return m.Data2.Ip }

func (m *TLPhoneConnection) SetIpv6(v string) { m.Data2.Ipv6 = v }
func (m *TLPhoneConnection) GetIpv6() string  { return m.Data2.Ipv6 }

func (m *TLPhoneConnection) SetPort(v int32) { m.Data2.Port = v }
func (m *TLPhoneConnection) GetPort() int32  { return m.Data2.Port }

func (m *TLPhoneConnection) SetPeerTag(v []byte) { m.Data2.PeerTag = v }
func (m *TLPhoneConnection) GetPeerTag() []byte  { return m.Data2.PeerTag }

func (m *TLPhoneConnection) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_phoneConnection))

	x.Long(m.Data2.Id)
	x.StringBytes(m.Data2.Ip)
	x.StringBytes(m.Data2.Ipv6)
	x.Int(m.Data2.Port)
	x.StringBytes(m.Data2.PeerTag)

	return x.buf
}

func (m *TLPhoneConnection) Decode(dbuf *DecodeBuf) error {
	m.Data2.Id = x.Long()
	m.Data2.Ip = x.StringBytes()
	m.Data2.Ipv6 = x.StringBytes()
	m.Data2.Port = x.Int()
	m.Data2.PeerTag = x.StringBytes()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Peer <--
//  + TL_PeerUser
//  + TL_PeerChat
//  + TL_PeerChannel
//

func (m *Peer) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_peerUser:
		t := m.To_PeerUser()
		return t.Encode()
	case TLConstructor_CRC32_peerChat:
		t := m.To_PeerChat()
		return t.Encode()
	case TLConstructor_CRC32_peerChannel:
		t := m.To_PeerChannel()
		return t.Encode()

	default:
		return []byte{}
	}
}

// peerUser#9db1bc6d user_id:int = Peer;
func (m *Peer) To_PeerUser() *TLPeerUser {
	return &TLPeerUser{
		Data2: m.Data2,
	}
}

// peerChat#bad0e5bb chat_id:int = Peer;
func (m *Peer) To_PeerChat() *TLPeerChat {
	return &TLPeerChat{
		Data2: m.Data2,
	}
}

// peerChannel#bddde532 channel_id:int = Peer;
func (m *Peer) To_PeerChannel() *TLPeerChannel {
	return &TLPeerChannel{
		Data2: m.Data2,
	}
}

// peerUser#9db1bc6d user_id:int = Peer;
func (m *TLPeerUser) To_Peer() *Peer {
	return &Peer{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPeerUser) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLPeerUser) GetUserId() int32  { return m.Data2.UserId }

func (m *TLPeerUser) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_peerUser))

	x.Int(m.Data2.UserId)

	return x.buf
}

func (m *TLPeerUser) Decode(dbuf *DecodeBuf) error {
	m.Data2.UserId = x.Int()

	return dbuf.err
}

// peerChat#bad0e5bb chat_id:int = Peer;
func (m *TLPeerChat) To_Peer() *Peer {
	return &Peer{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPeerChat) SetChatId(v int32) { m.Data2.ChatId = v }
func (m *TLPeerChat) GetChatId() int32  { return m.Data2.ChatId }

func (m *TLPeerChat) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_peerChat))

	x.Int(m.Data2.ChatId)

	return x.buf
}

func (m *TLPeerChat) Decode(dbuf *DecodeBuf) error {
	m.Data2.ChatId = x.Int()

	return dbuf.err
}

// peerChannel#bddde532 channel_id:int = Peer;
func (m *TLPeerChannel) To_Peer() *Peer {
	return &Peer{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPeerChannel) SetChannelId(v int32) { m.Data2.ChannelId = v }
func (m *TLPeerChannel) GetChannelId() int32  { return m.Data2.ChannelId }

func (m *TLPeerChannel) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_peerChannel))

	x.Int(m.Data2.ChannelId)

	return x.buf
}

func (m *TLPeerChannel) Decode(dbuf *DecodeBuf) error {
	m.Data2.ChannelId = x.Int()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Update <--
//  + TL_UpdateNewMessage
//  + TL_UpdateMessageID
//  + TL_UpdateDeleteMessages
//  + TL_UpdateUserTyping
//  + TL_UpdateChatUserTyping
//  + TL_UpdateChatParticipants
//  + TL_UpdateUserStatus
//  + TL_UpdateUserName
//  + TL_UpdateUserPhoto
//  + TL_UpdateContactRegistered
//  + TL_UpdateContactLink
//  + TL_UpdateNewEncryptedMessage
//  + TL_UpdateEncryptedChatTyping
//  + TL_UpdateEncryption
//  + TL_UpdateEncryptedMessagesRead
//  + TL_UpdateChatParticipantAdd
//  + TL_UpdateChatParticipantDelete
//  + TL_UpdateDcOptions
//  + TL_UpdateUserBlocked
//  + TL_UpdateNotifySettings
//  + TL_UpdateServiceNotification
//  + TL_UpdatePrivacy
//  + TL_UpdateUserPhone
//  + TL_UpdateReadHistoryInbox
//  + TL_UpdateReadHistoryOutbox
//  + TL_UpdateWebPage
//  + TL_UpdateReadMessagesContents
//  + TL_UpdateChannelTooLong
//  + TL_UpdateChannel
//  + TL_UpdateNewChannelMessage
//  + TL_UpdateReadChannelInbox
//  + TL_UpdateDeleteChannelMessages
//  + TL_UpdateChannelMessageViews
//  + TL_UpdateChatAdmins
//  + TL_UpdateChatParticipantAdmin
//  + TL_UpdateNewStickerSet
//  + TL_UpdateStickerSetsOrder
//  + TL_UpdateStickerSets
//  + TL_UpdateSavedGifs
//  + TL_UpdateBotInlineQuery
//  + TL_UpdateBotInlineSend
//  + TL_UpdateEditChannelMessage
//  + TL_UpdateChannelPinnedMessage
//  + TL_UpdateBotCallbackQuery
//  + TL_UpdateEditMessage
//  + TL_UpdateInlineBotCallbackQuery
//  + TL_UpdateReadChannelOutbox
//  + TL_UpdateDraftMessage
//  + TL_UpdateReadFeaturedStickers
//  + TL_UpdateRecentStickers
//  + TL_UpdateConfig
//  + TL_UpdatePtsChanged
//  + TL_UpdateChannelWebPage
//  + TL_UpdateDialogPinned
//  + TL_UpdatePinnedDialogs
//  + TL_UpdateBotWebhookJSON
//  + TL_UpdateBotWebhookJSONQuery
//  + TL_UpdateBotShippingQuery
//  + TL_UpdateBotPrecheckoutQuery
//  + TL_UpdatePhoneCall
//  + TL_UpdateLangPackTooLong
//  + TL_UpdateLangPack
//  + TL_UpdateFavedStickers
//  + TL_UpdateChannelReadMessagesContents
//  + TL_UpdateContactsReset
//

func (m *Update) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_updateNewMessage:
		t := m.To_UpdateNewMessage()
		return t.Encode()
	case TLConstructor_CRC32_updateMessageID:
		t := m.To_UpdateMessageID()
		return t.Encode()
	case TLConstructor_CRC32_updateDeleteMessages:
		t := m.To_UpdateDeleteMessages()
		return t.Encode()
	case TLConstructor_CRC32_updateUserTyping:
		t := m.To_UpdateUserTyping()
		return t.Encode()
	case TLConstructor_CRC32_updateChatUserTyping:
		t := m.To_UpdateChatUserTyping()
		return t.Encode()
	case TLConstructor_CRC32_updateChatParticipants:
		t := m.To_UpdateChatParticipants()
		return t.Encode()
	case TLConstructor_CRC32_updateUserStatus:
		t := m.To_UpdateUserStatus()
		return t.Encode()
	case TLConstructor_CRC32_updateUserName:
		t := m.To_UpdateUserName()
		return t.Encode()
	case TLConstructor_CRC32_updateUserPhoto:
		t := m.To_UpdateUserPhoto()
		return t.Encode()
	case TLConstructor_CRC32_updateContactRegistered:
		t := m.To_UpdateContactRegistered()
		return t.Encode()
	case TLConstructor_CRC32_updateContactLink:
		t := m.To_UpdateContactLink()
		return t.Encode()
	case TLConstructor_CRC32_updateNewEncryptedMessage:
		t := m.To_UpdateNewEncryptedMessage()
		return t.Encode()
	case TLConstructor_CRC32_updateEncryptedChatTyping:
		t := m.To_UpdateEncryptedChatTyping()
		return t.Encode()
	case TLConstructor_CRC32_updateEncryption:
		t := m.To_UpdateEncryption()
		return t.Encode()
	case TLConstructor_CRC32_updateEncryptedMessagesRead:
		t := m.To_UpdateEncryptedMessagesRead()
		return t.Encode()
	case TLConstructor_CRC32_updateChatParticipantAdd:
		t := m.To_UpdateChatParticipantAdd()
		return t.Encode()
	case TLConstructor_CRC32_updateChatParticipantDelete:
		t := m.To_UpdateChatParticipantDelete()
		return t.Encode()
	case TLConstructor_CRC32_updateDcOptions:
		t := m.To_UpdateDcOptions()
		return t.Encode()
	case TLConstructor_CRC32_updateUserBlocked:
		t := m.To_UpdateUserBlocked()
		return t.Encode()
	case TLConstructor_CRC32_updateNotifySettings:
		t := m.To_UpdateNotifySettings()
		return t.Encode()
	case TLConstructor_CRC32_updateServiceNotification:
		t := m.To_UpdateServiceNotification()
		return t.Encode()
	case TLConstructor_CRC32_updatePrivacy:
		t := m.To_UpdatePrivacy()
		return t.Encode()
	case TLConstructor_CRC32_updateUserPhone:
		t := m.To_UpdateUserPhone()
		return t.Encode()
	case TLConstructor_CRC32_updateReadHistoryInbox:
		t := m.To_UpdateReadHistoryInbox()
		return t.Encode()
	case TLConstructor_CRC32_updateReadHistoryOutbox:
		t := m.To_UpdateReadHistoryOutbox()
		return t.Encode()
	case TLConstructor_CRC32_updateWebPage:
		t := m.To_UpdateWebPage()
		return t.Encode()
	case TLConstructor_CRC32_updateReadMessagesContents:
		t := m.To_UpdateReadMessagesContents()
		return t.Encode()
	case TLConstructor_CRC32_updateChannelTooLong:
		t := m.To_UpdateChannelTooLong()
		return t.Encode()
	case TLConstructor_CRC32_updateChannel:
		t := m.To_UpdateChannel()
		return t.Encode()
	case TLConstructor_CRC32_updateNewChannelMessage:
		t := m.To_UpdateNewChannelMessage()
		return t.Encode()
	case TLConstructor_CRC32_updateReadChannelInbox:
		t := m.To_UpdateReadChannelInbox()
		return t.Encode()
	case TLConstructor_CRC32_updateDeleteChannelMessages:
		t := m.To_UpdateDeleteChannelMessages()
		return t.Encode()
	case TLConstructor_CRC32_updateChannelMessageViews:
		t := m.To_UpdateChannelMessageViews()
		return t.Encode()
	case TLConstructor_CRC32_updateChatAdmins:
		t := m.To_UpdateChatAdmins()
		return t.Encode()
	case TLConstructor_CRC32_updateChatParticipantAdmin:
		t := m.To_UpdateChatParticipantAdmin()
		return t.Encode()
	case TLConstructor_CRC32_updateNewStickerSet:
		t := m.To_UpdateNewStickerSet()
		return t.Encode()
	case TLConstructor_CRC32_updateStickerSetsOrder:
		t := m.To_UpdateStickerSetsOrder()
		return t.Encode()
	case TLConstructor_CRC32_updateStickerSets:
		t := m.To_UpdateStickerSets()
		return t.Encode()
	case TLConstructor_CRC32_updateSavedGifs:
		t := m.To_UpdateSavedGifs()
		return t.Encode()
	case TLConstructor_CRC32_updateBotInlineQuery:
		t := m.To_UpdateBotInlineQuery()
		return t.Encode()
	case TLConstructor_CRC32_updateBotInlineSend:
		t := m.To_UpdateBotInlineSend()
		return t.Encode()
	case TLConstructor_CRC32_updateEditChannelMessage:
		t := m.To_UpdateEditChannelMessage()
		return t.Encode()
	case TLConstructor_CRC32_updateChannelPinnedMessage:
		t := m.To_UpdateChannelPinnedMessage()
		return t.Encode()
	case TLConstructor_CRC32_updateBotCallbackQuery:
		t := m.To_UpdateBotCallbackQuery()
		return t.Encode()
	case TLConstructor_CRC32_updateEditMessage:
		t := m.To_UpdateEditMessage()
		return t.Encode()
	case TLConstructor_CRC32_updateInlineBotCallbackQuery:
		t := m.To_UpdateInlineBotCallbackQuery()
		return t.Encode()
	case TLConstructor_CRC32_updateReadChannelOutbox:
		t := m.To_UpdateReadChannelOutbox()
		return t.Encode()
	case TLConstructor_CRC32_updateDraftMessage:
		t := m.To_UpdateDraftMessage()
		return t.Encode()
	case TLConstructor_CRC32_updateReadFeaturedStickers:
		t := m.To_UpdateReadFeaturedStickers()
		return t.Encode()
	case TLConstructor_CRC32_updateRecentStickers:
		t := m.To_UpdateRecentStickers()
		return t.Encode()
	case TLConstructor_CRC32_updateConfig:
		t := m.To_UpdateConfig()
		return t.Encode()
	case TLConstructor_CRC32_updatePtsChanged:
		t := m.To_UpdatePtsChanged()
		return t.Encode()
	case TLConstructor_CRC32_updateChannelWebPage:
		t := m.To_UpdateChannelWebPage()
		return t.Encode()
	case TLConstructor_CRC32_updateDialogPinned:
		t := m.To_UpdateDialogPinned()
		return t.Encode()
	case TLConstructor_CRC32_updatePinnedDialogs:
		t := m.To_UpdatePinnedDialogs()
		return t.Encode()
	case TLConstructor_CRC32_updateBotWebhookJSON:
		t := m.To_UpdateBotWebhookJSON()
		return t.Encode()
	case TLConstructor_CRC32_updateBotWebhookJSONQuery:
		t := m.To_UpdateBotWebhookJSONQuery()
		return t.Encode()
	case TLConstructor_CRC32_updateBotShippingQuery:
		t := m.To_UpdateBotShippingQuery()
		return t.Encode()
	case TLConstructor_CRC32_updateBotPrecheckoutQuery:
		t := m.To_UpdateBotPrecheckoutQuery()
		return t.Encode()
	case TLConstructor_CRC32_updatePhoneCall:
		t := m.To_UpdatePhoneCall()
		return t.Encode()
	case TLConstructor_CRC32_updateLangPackTooLong:
		t := m.To_UpdateLangPackTooLong()
		return t.Encode()
	case TLConstructor_CRC32_updateLangPack:
		t := m.To_UpdateLangPack()
		return t.Encode()
	case TLConstructor_CRC32_updateFavedStickers:
		t := m.To_UpdateFavedStickers()
		return t.Encode()
	case TLConstructor_CRC32_updateChannelReadMessagesContents:
		t := m.To_UpdateChannelReadMessagesContents()
		return t.Encode()
	case TLConstructor_CRC32_updateContactsReset:
		t := m.To_UpdateContactsReset()
		return t.Encode()

	default:
		return []byte{}
	}
}

// updateNewMessage#1f2b0afd message:Message pts:int pts_count:int = Update;
func (m *Update) To_UpdateNewMessage() *TLUpdateNewMessage {
	return &TLUpdateNewMessage{
		Data2: m.Data2,
	}
}

// updateMessageID#4e90bfd6 id:int random_id:long = Update;
func (m *Update) To_UpdateMessageID() *TLUpdateMessageID {
	return &TLUpdateMessageID{
		Data2: m.Data2,
	}
}

// updateDeleteMessages#a20db0e5 messages:Vector<int> pts:int pts_count:int = Update;
func (m *Update) To_UpdateDeleteMessages() *TLUpdateDeleteMessages {
	return &TLUpdateDeleteMessages{
		Data2: m.Data2,
	}
}

// updateUserTyping#5c486927 user_id:int action:SendMessageAction = Update;
func (m *Update) To_UpdateUserTyping() *TLUpdateUserTyping {
	return &TLUpdateUserTyping{
		Data2: m.Data2,
	}
}

// updateChatUserTyping#9a65ea1f chat_id:int user_id:int action:SendMessageAction = Update;
func (m *Update) To_UpdateChatUserTyping() *TLUpdateChatUserTyping {
	return &TLUpdateChatUserTyping{
		Data2: m.Data2,
	}
}

// updateChatParticipants#7761198 participants:ChatParticipants = Update;
func (m *Update) To_UpdateChatParticipants() *TLUpdateChatParticipants {
	return &TLUpdateChatParticipants{
		Data2: m.Data2,
	}
}

// updateUserStatus#1bfbd823 user_id:int status:UserStatus = Update;
func (m *Update) To_UpdateUserStatus() *TLUpdateUserStatus {
	return &TLUpdateUserStatus{
		Data2: m.Data2,
	}
}

// updateUserName#a7332b73 user_id:int first_name:string last_name:string username:string = Update;
func (m *Update) To_UpdateUserName() *TLUpdateUserName {
	return &TLUpdateUserName{
		Data2: m.Data2,
	}
}

// updateUserPhoto#95313b0c user_id:int date:int photo:UserProfilePhoto previous:Bool = Update;
func (m *Update) To_UpdateUserPhoto() *TLUpdateUserPhoto {
	return &TLUpdateUserPhoto{
		Data2: m.Data2,
	}
}

// updateContactRegistered#2575bbb9 user_id:int date:int = Update;
func (m *Update) To_UpdateContactRegistered() *TLUpdateContactRegistered {
	return &TLUpdateContactRegistered{
		Data2: m.Data2,
	}
}

// updateContactLink#9d2e67c5 user_id:int my_link:ContactLink foreign_link:ContactLink = Update;
func (m *Update) To_UpdateContactLink() *TLUpdateContactLink {
	return &TLUpdateContactLink{
		Data2: m.Data2,
	}
}

// updateNewEncryptedMessage#12bcbd9a message:EncryptedMessage qts:int = Update;
func (m *Update) To_UpdateNewEncryptedMessage() *TLUpdateNewEncryptedMessage {
	return &TLUpdateNewEncryptedMessage{
		Data2: m.Data2,
	}
}

// updateEncryptedChatTyping#1710f156 chat_id:int = Update;
func (m *Update) To_UpdateEncryptedChatTyping() *TLUpdateEncryptedChatTyping {
	return &TLUpdateEncryptedChatTyping{
		Data2: m.Data2,
	}
}

// updateEncryption#b4a2e88d chat:EncryptedChat date:int = Update;
func (m *Update) To_UpdateEncryption() *TLUpdateEncryption {
	return &TLUpdateEncryption{
		Data2: m.Data2,
	}
}

// updateEncryptedMessagesRead#38fe25b7 chat_id:int max_date:int date:int = Update;
func (m *Update) To_UpdateEncryptedMessagesRead() *TLUpdateEncryptedMessagesRead {
	return &TLUpdateEncryptedMessagesRead{
		Data2: m.Data2,
	}
}

// updateChatParticipantAdd#ea4b0e5c chat_id:int user_id:int inviter_id:int date:int version:int = Update;
func (m *Update) To_UpdateChatParticipantAdd() *TLUpdateChatParticipantAdd {
	return &TLUpdateChatParticipantAdd{
		Data2: m.Data2,
	}
}

// updateChatParticipantDelete#6e5f8c22 chat_id:int user_id:int version:int = Update;
func (m *Update) To_UpdateChatParticipantDelete() *TLUpdateChatParticipantDelete {
	return &TLUpdateChatParticipantDelete{
		Data2: m.Data2,
	}
}

// updateDcOptions#8e5e9873 dc_options:Vector<DcOption> = Update;
func (m *Update) To_UpdateDcOptions() *TLUpdateDcOptions {
	return &TLUpdateDcOptions{
		Data2: m.Data2,
	}
}

// updateUserBlocked#80ece81a user_id:int blocked:Bool = Update;
func (m *Update) To_UpdateUserBlocked() *TLUpdateUserBlocked {
	return &TLUpdateUserBlocked{
		Data2: m.Data2,
	}
}

// updateNotifySettings#bec268ef peer:NotifyPeer notify_settings:PeerNotifySettings = Update;
func (m *Update) To_UpdateNotifySettings() *TLUpdateNotifySettings {
	return &TLUpdateNotifySettings{
		Data2: m.Data2,
	}
}

// updateServiceNotification#ebe46819 flags:# popup:flags.0?true inbox_date:flags.1?int type:string message:string media:MessageMedia entities:Vector<MessageEntity> = Update;
func (m *Update) To_UpdateServiceNotification() *TLUpdateServiceNotification {
	return &TLUpdateServiceNotification{
		Data2: m.Data2,
	}
}

// updatePrivacy#ee3b272a key:PrivacyKey rules:Vector<PrivacyRule> = Update;
func (m *Update) To_UpdatePrivacy() *TLUpdatePrivacy {
	return &TLUpdatePrivacy{
		Data2: m.Data2,
	}
}

// updateUserPhone#12b9417b user_id:int phone:string = Update;
func (m *Update) To_UpdateUserPhone() *TLUpdateUserPhone {
	return &TLUpdateUserPhone{
		Data2: m.Data2,
	}
}

// updateReadHistoryInbox#9961fd5c peer:Peer max_id:int pts:int pts_count:int = Update;
func (m *Update) To_UpdateReadHistoryInbox() *TLUpdateReadHistoryInbox {
	return &TLUpdateReadHistoryInbox{
		Data2: m.Data2,
	}
}

// updateReadHistoryOutbox#2f2f21bf peer:Peer max_id:int pts:int pts_count:int = Update;
func (m *Update) To_UpdateReadHistoryOutbox() *TLUpdateReadHistoryOutbox {
	return &TLUpdateReadHistoryOutbox{
		Data2: m.Data2,
	}
}

// updateWebPage#7f891213 webpage:WebPage pts:int pts_count:int = Update;
func (m *Update) To_UpdateWebPage() *TLUpdateWebPage {
	return &TLUpdateWebPage{
		Data2: m.Data2,
	}
}

// updateReadMessagesContents#68c13933 messages:Vector<int> pts:int pts_count:int = Update;
func (m *Update) To_UpdateReadMessagesContents() *TLUpdateReadMessagesContents {
	return &TLUpdateReadMessagesContents{
		Data2: m.Data2,
	}
}

// updateChannelTooLong#eb0467fb flags:# channel_id:int pts:flags.0?int = Update;
func (m *Update) To_UpdateChannelTooLong() *TLUpdateChannelTooLong {
	return &TLUpdateChannelTooLong{
		Data2: m.Data2,
	}
}

// updateChannel#b6d45656 channel_id:int = Update;
func (m *Update) To_UpdateChannel() *TLUpdateChannel {
	return &TLUpdateChannel{
		Data2: m.Data2,
	}
}

// updateNewChannelMessage#62ba04d9 message:Message pts:int pts_count:int = Update;
func (m *Update) To_UpdateNewChannelMessage() *TLUpdateNewChannelMessage {
	return &TLUpdateNewChannelMessage{
		Data2: m.Data2,
	}
}

// updateReadChannelInbox#4214f37f channel_id:int max_id:int = Update;
func (m *Update) To_UpdateReadChannelInbox() *TLUpdateReadChannelInbox {
	return &TLUpdateReadChannelInbox{
		Data2: m.Data2,
	}
}

// updateDeleteChannelMessages#c37521c9 channel_id:int messages:Vector<int> pts:int pts_count:int = Update;
func (m *Update) To_UpdateDeleteChannelMessages() *TLUpdateDeleteChannelMessages {
	return &TLUpdateDeleteChannelMessages{
		Data2: m.Data2,
	}
}

// updateChannelMessageViews#98a12b4b channel_id:int id:int views:int = Update;
func (m *Update) To_UpdateChannelMessageViews() *TLUpdateChannelMessageViews {
	return &TLUpdateChannelMessageViews{
		Data2: m.Data2,
	}
}

// updateChatAdmins#6e947941 chat_id:int enabled:Bool version:int = Update;
func (m *Update) To_UpdateChatAdmins() *TLUpdateChatAdmins {
	return &TLUpdateChatAdmins{
		Data2: m.Data2,
	}
}

// updateChatParticipantAdmin#b6901959 chat_id:int user_id:int is_admin:Bool version:int = Update;
func (m *Update) To_UpdateChatParticipantAdmin() *TLUpdateChatParticipantAdmin {
	return &TLUpdateChatParticipantAdmin{
		Data2: m.Data2,
	}
}

// updateNewStickerSet#688a30aa stickerset:messages.StickerSet = Update;
func (m *Update) To_UpdateNewStickerSet() *TLUpdateNewStickerSet {
	return &TLUpdateNewStickerSet{
		Data2: m.Data2,
	}
}

// updateStickerSetsOrder#bb2d201 flags:# masks:flags.0?true order:Vector<long> = Update;
func (m *Update) To_UpdateStickerSetsOrder() *TLUpdateStickerSetsOrder {
	return &TLUpdateStickerSetsOrder{
		Data2: m.Data2,
	}
}

// updateStickerSets#43ae3dec = Update;
func (m *Update) To_UpdateStickerSets() *TLUpdateStickerSets {
	return &TLUpdateStickerSets{
		Data2: m.Data2,
	}
}

// updateSavedGifs#9375341e = Update;
func (m *Update) To_UpdateSavedGifs() *TLUpdateSavedGifs {
	return &TLUpdateSavedGifs{
		Data2: m.Data2,
	}
}

// updateBotInlineQuery#54826690 flags:# query_id:long user_id:int query:string geo:flags.0?GeoPoint offset:string = Update;
func (m *Update) To_UpdateBotInlineQuery() *TLUpdateBotInlineQuery {
	return &TLUpdateBotInlineQuery{
		Data2: m.Data2,
	}
}

// updateBotInlineSend#e48f964 flags:# user_id:int query:string geo:flags.0?GeoPoint id:string msg_id:flags.1?InputBotInlineMessageID = Update;
func (m *Update) To_UpdateBotInlineSend() *TLUpdateBotInlineSend {
	return &TLUpdateBotInlineSend{
		Data2: m.Data2,
	}
}

// updateEditChannelMessage#1b3f4df7 message:Message pts:int pts_count:int = Update;
func (m *Update) To_UpdateEditChannelMessage() *TLUpdateEditChannelMessage {
	return &TLUpdateEditChannelMessage{
		Data2: m.Data2,
	}
}

// updateChannelPinnedMessage#98592475 channel_id:int id:int = Update;
func (m *Update) To_UpdateChannelPinnedMessage() *TLUpdateChannelPinnedMessage {
	return &TLUpdateChannelPinnedMessage{
		Data2: m.Data2,
	}
}

// updateBotCallbackQuery#e73547e1 flags:# query_id:long user_id:int peer:Peer msg_id:int chat_instance:long data:flags.0?bytes game_short_name:flags.1?string = Update;
func (m *Update) To_UpdateBotCallbackQuery() *TLUpdateBotCallbackQuery {
	return &TLUpdateBotCallbackQuery{
		Data2: m.Data2,
	}
}

// updateEditMessage#e40370a3 message:Message pts:int pts_count:int = Update;
func (m *Update) To_UpdateEditMessage() *TLUpdateEditMessage {
	return &TLUpdateEditMessage{
		Data2: m.Data2,
	}
}

// updateInlineBotCallbackQuery#f9d27a5a flags:# query_id:long user_id:int msg_id:InputBotInlineMessageID chat_instance:long data:flags.0?bytes game_short_name:flags.1?string = Update;
func (m *Update) To_UpdateInlineBotCallbackQuery() *TLUpdateInlineBotCallbackQuery {
	return &TLUpdateInlineBotCallbackQuery{
		Data2: m.Data2,
	}
}

// updateReadChannelOutbox#25d6c9c7 channel_id:int max_id:int = Update;
func (m *Update) To_UpdateReadChannelOutbox() *TLUpdateReadChannelOutbox {
	return &TLUpdateReadChannelOutbox{
		Data2: m.Data2,
	}
}

// updateDraftMessage#ee2bb969 peer:Peer draft:DraftMessage = Update;
func (m *Update) To_UpdateDraftMessage() *TLUpdateDraftMessage {
	return &TLUpdateDraftMessage{
		Data2: m.Data2,
	}
}

// updateReadFeaturedStickers#571d2742 = Update;
func (m *Update) To_UpdateReadFeaturedStickers() *TLUpdateReadFeaturedStickers {
	return &TLUpdateReadFeaturedStickers{
		Data2: m.Data2,
	}
}

// updateRecentStickers#9a422c20 = Update;
func (m *Update) To_UpdateRecentStickers() *TLUpdateRecentStickers {
	return &TLUpdateRecentStickers{
		Data2: m.Data2,
	}
}

// updateConfig#a229dd06 = Update;
func (m *Update) To_UpdateConfig() *TLUpdateConfig {
	return &TLUpdateConfig{
		Data2: m.Data2,
	}
}

// updatePtsChanged#3354678f = Update;
func (m *Update) To_UpdatePtsChanged() *TLUpdatePtsChanged {
	return &TLUpdatePtsChanged{
		Data2: m.Data2,
	}
}

// updateChannelWebPage#40771900 channel_id:int webpage:WebPage pts:int pts_count:int = Update;
func (m *Update) To_UpdateChannelWebPage() *TLUpdateChannelWebPage {
	return &TLUpdateChannelWebPage{
		Data2: m.Data2,
	}
}

// updateDialogPinned#d711a2cc flags:# pinned:flags.0?true peer:Peer = Update;
func (m *Update) To_UpdateDialogPinned() *TLUpdateDialogPinned {
	return &TLUpdateDialogPinned{
		Data2: m.Data2,
	}
}

// updatePinnedDialogs#d8caf68d flags:# order:flags.0?Vector<Peer> = Update;
func (m *Update) To_UpdatePinnedDialogs() *TLUpdatePinnedDialogs {
	return &TLUpdatePinnedDialogs{
		Data2: m.Data2,
	}
}

// updateBotWebhookJSON#8317c0c3 data:DataJSON = Update;
func (m *Update) To_UpdateBotWebhookJSON() *TLUpdateBotWebhookJSON {
	return &TLUpdateBotWebhookJSON{
		Data2: m.Data2,
	}
}

// updateBotWebhookJSONQuery#9b9240a6 query_id:long data:DataJSON timeout:int = Update;
func (m *Update) To_UpdateBotWebhookJSONQuery() *TLUpdateBotWebhookJSONQuery {
	return &TLUpdateBotWebhookJSONQuery{
		Data2: m.Data2,
	}
}

// updateBotShippingQuery#e0cdc940 query_id:long user_id:int payload:bytes shipping_address:PostAddress = Update;
func (m *Update) To_UpdateBotShippingQuery() *TLUpdateBotShippingQuery {
	return &TLUpdateBotShippingQuery{
		Data2: m.Data2,
	}
}

// updateBotPrecheckoutQuery#5d2f3aa9 flags:# query_id:long user_id:int payload:bytes info:flags.0?PaymentRequestedInfo shipping_option_id:flags.1?string currency:string total_amount:long = Update;
func (m *Update) To_UpdateBotPrecheckoutQuery() *TLUpdateBotPrecheckoutQuery {
	return &TLUpdateBotPrecheckoutQuery{
		Data2: m.Data2,
	}
}

// updatePhoneCall#ab0f6b1e phone_call:PhoneCall = Update;
func (m *Update) To_UpdatePhoneCall() *TLUpdatePhoneCall {
	return &TLUpdatePhoneCall{
		Data2: m.Data2,
	}
}

// updateLangPackTooLong#10c2404b = Update;
func (m *Update) To_UpdateLangPackTooLong() *TLUpdateLangPackTooLong {
	return &TLUpdateLangPackTooLong{
		Data2: m.Data2,
	}
}

// updateLangPack#56022f4d difference:LangPackDifference = Update;
func (m *Update) To_UpdateLangPack() *TLUpdateLangPack {
	return &TLUpdateLangPack{
		Data2: m.Data2,
	}
}

// updateFavedStickers#e511996d = Update;
func (m *Update) To_UpdateFavedStickers() *TLUpdateFavedStickers {
	return &TLUpdateFavedStickers{
		Data2: m.Data2,
	}
}

// updateChannelReadMessagesContents#89893b45 channel_id:int messages:Vector<int> = Update;
func (m *Update) To_UpdateChannelReadMessagesContents() *TLUpdateChannelReadMessagesContents {
	return &TLUpdateChannelReadMessagesContents{
		Data2: m.Data2,
	}
}

// updateContactsReset#7084a7be = Update;
func (m *Update) To_UpdateContactsReset() *TLUpdateContactsReset {
	return &TLUpdateContactsReset{
		Data2: m.Data2,
	}
}

// updateNewMessage#1f2b0afd message:Message pts:int pts_count:int = Update;
func (m *TLUpdateNewMessage) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateNewMessage) SetMessage(v *Message) { m.Data2.Message_1 = v }
func (m *TLUpdateNewMessage) GetMessage() *Message  { return m.Data2.Message_1 }

func (m *TLUpdateNewMessage) SetPts(v int32) { m.Data2.Pts = v }
func (m *TLUpdateNewMessage) GetPts() int32  { return m.Data2.Pts }

func (m *TLUpdateNewMessage) SetPtsCount(v int32) { m.Data2.PtsCount = v }
func (m *TLUpdateNewMessage) GetPtsCount() int32  { return m.Data2.PtsCount }

func (m *TLUpdateNewMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateNewMessage))

	x.Bytes(m.Data2Message.Encode())
	x.Int(m.Data2.Pts)
	x.Int(m.Data2.PtsCount)

	return x.buf
}

func (m *TLUpdateNewMessage) Decode(dbuf *DecodeBuf) error {
	m1 := &Message{}
	m1.Decode(dbuf)
	m.SetMessage(m1)
	m.Data2.Pts = x.Int()
	m.Data2.PtsCount = x.Int()

	return dbuf.err
}

// updateMessageID#4e90bfd6 id:int random_id:long = Update;
func (m *TLUpdateMessageID) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateMessageID) SetId(v int32) { m.Data2.Id_4 = v }
func (m *TLUpdateMessageID) GetId() int32  { return m.Data2.Id_4 }

func (m *TLUpdateMessageID) SetRandomId(v int64) { m.Data2.RandomId = v }
func (m *TLUpdateMessageID) GetRandomId() int64  { return m.Data2.RandomId }

func (m *TLUpdateMessageID) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateMessageID))

	x.Int(m.Data2.Id)
	x.Long(m.Data2.RandomId)

	return x.buf
}

func (m *TLUpdateMessageID) Decode(dbuf *DecodeBuf) error {
	m.Data2.Id = x.Int()
	m.Data2.RandomId = x.Long()

	return dbuf.err
}

// updateDeleteMessages#a20db0e5 messages:Vector<int> pts:int pts_count:int = Update;
func (m *TLUpdateDeleteMessages) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateDeleteMessages) SetMessages(v []int32) { m.Data2.Messages = v }
func (m *TLUpdateDeleteMessages) GetMessages() []int32  { return m.Data2.Messages }

func (m *TLUpdateDeleteMessages) SetPts(v int32) { m.Data2.Pts = v }
func (m *TLUpdateDeleteMessages) GetPts() int32  { return m.Data2.Pts }

func (m *TLUpdateDeleteMessages) SetPtsCount(v int32) { m.Data2.PtsCount = v }
func (m *TLUpdateDeleteMessages) GetPtsCount() int32  { return m.Data2.PtsCount }

func (m *TLUpdateDeleteMessages) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateDeleteMessages))

	x.Int(m.Data2.Pts)
	x.Int(m.Data2.PtsCount)

	return x.buf
}

func (m *TLUpdateDeleteMessages) Decode(dbuf *DecodeBuf) error {

	m.Data2.Pts = x.Int()
	m.Data2.PtsCount = x.Int()

	return dbuf.err
}

// updateUserTyping#5c486927 user_id:int action:SendMessageAction = Update;
func (m *TLUpdateUserTyping) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateUserTyping) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLUpdateUserTyping) GetUserId() int32  { return m.Data2.UserId }

func (m *TLUpdateUserTyping) SetAction(v *SendMessageAction) { m.Data2.Action = v }
func (m *TLUpdateUserTyping) GetAction() *SendMessageAction  { return m.Data2.Action }

func (m *TLUpdateUserTyping) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateUserTyping))

	x.Int(m.Data2.UserId)
	x.Bytes(m.Data2Action.Encode())

	return x.buf
}

func (m *TLUpdateUserTyping) Decode(dbuf *DecodeBuf) error {
	m.Data2.UserId = x.Int()
	m2 := &Action{}
	m2.Decode(dbuf)
	m.SetAction(m2)

	return dbuf.err
}

// updateChatUserTyping#9a65ea1f chat_id:int user_id:int action:SendMessageAction = Update;
func (m *TLUpdateChatUserTyping) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateChatUserTyping) SetChatId(v int32) { m.Data2.ChatId = v }
func (m *TLUpdateChatUserTyping) GetChatId() int32  { return m.Data2.ChatId }

func (m *TLUpdateChatUserTyping) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLUpdateChatUserTyping) GetUserId() int32  { return m.Data2.UserId }

func (m *TLUpdateChatUserTyping) SetAction(v *SendMessageAction) { m.Data2.Action = v }
func (m *TLUpdateChatUserTyping) GetAction() *SendMessageAction  { return m.Data2.Action }

func (m *TLUpdateChatUserTyping) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateChatUserTyping))

	x.Int(m.Data2.ChatId)
	x.Int(m.Data2.UserId)
	x.Bytes(m.Data2Action.Encode())

	return x.buf
}

func (m *TLUpdateChatUserTyping) Decode(dbuf *DecodeBuf) error {
	m.Data2.ChatId = x.Int()
	m.Data2.UserId = x.Int()
	m3 := &Action{}
	m3.Decode(dbuf)
	m.SetAction(m3)

	return dbuf.err
}

// updateChatParticipants#7761198 participants:ChatParticipants = Update;
func (m *TLUpdateChatParticipants) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateChatParticipants) SetParticipants(v *ChatParticipants) { m.Data2.Participants = v }
func (m *TLUpdateChatParticipants) GetParticipants() *ChatParticipants  { return m.Data2.Participants }

func (m *TLUpdateChatParticipants) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateChatParticipants))

	x.Bytes(m.Data2Participants.Encode())

	return x.buf
}

func (m *TLUpdateChatParticipants) Decode(dbuf *DecodeBuf) error {
	m1 := &Participants{}
	m1.Decode(dbuf)
	m.SetParticipants(m1)

	return dbuf.err
}

// updateUserStatus#1bfbd823 user_id:int status:UserStatus = Update;
func (m *TLUpdateUserStatus) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateUserStatus) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLUpdateUserStatus) GetUserId() int32  { return m.Data2.UserId }

func (m *TLUpdateUserStatus) SetStatus(v *UserStatus) { m.Data2.Status = v }
func (m *TLUpdateUserStatus) GetStatus() *UserStatus  { return m.Data2.Status }

func (m *TLUpdateUserStatus) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateUserStatus))

	x.Int(m.Data2.UserId)
	x.Bytes(m.Data2Status.Encode())

	return x.buf
}

func (m *TLUpdateUserStatus) Decode(dbuf *DecodeBuf) error {
	m.Data2.UserId = x.Int()
	m2 := &Status{}
	m2.Decode(dbuf)
	m.SetStatus(m2)

	return dbuf.err
}

// updateUserName#a7332b73 user_id:int first_name:string last_name:string username:string = Update;
func (m *TLUpdateUserName) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateUserName) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLUpdateUserName) GetUserId() int32  { return m.Data2.UserId }

func (m *TLUpdateUserName) SetFirstName(v string) { m.Data2.FirstName = v }
func (m *TLUpdateUserName) GetFirstName() string  { return m.Data2.FirstName }

func (m *TLUpdateUserName) SetLastName(v string) { m.Data2.LastName = v }
func (m *TLUpdateUserName) GetLastName() string  { return m.Data2.LastName }

func (m *TLUpdateUserName) SetUsername(v string) { m.Data2.Username = v }
func (m *TLUpdateUserName) GetUsername() string  { return m.Data2.Username }

func (m *TLUpdateUserName) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateUserName))

	x.Int(m.Data2.UserId)
	x.StringBytes(m.Data2.FirstName)
	x.StringBytes(m.Data2.LastName)
	x.StringBytes(m.Data2.Username)

	return x.buf
}

func (m *TLUpdateUserName) Decode(dbuf *DecodeBuf) error {
	m.Data2.UserId = x.Int()
	m.Data2.FirstName = x.StringBytes()
	m.Data2.LastName = x.StringBytes()
	m.Data2.Username = x.StringBytes()

	return dbuf.err
}

// updateUserPhoto#95313b0c user_id:int date:int photo:UserProfilePhoto previous:Bool = Update;
func (m *TLUpdateUserPhoto) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateUserPhoto) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLUpdateUserPhoto) GetUserId() int32  { return m.Data2.UserId }

func (m *TLUpdateUserPhoto) SetDate(v int32) { m.Data2.Date = v }
func (m *TLUpdateUserPhoto) GetDate() int32  { return m.Data2.Date }

func (m *TLUpdateUserPhoto) SetPhoto(v *UserProfilePhoto) { m.Data2.Photo = v }
func (m *TLUpdateUserPhoto) GetPhoto() *UserProfilePhoto  { return m.Data2.Photo }

func (m *TLUpdateUserPhoto) SetPrevious(v *Bool) { m.Data2.Previous = v }
func (m *TLUpdateUserPhoto) GetPrevious() *Bool  { return m.Data2.Previous }

func (m *TLUpdateUserPhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateUserPhoto))

	x.Int(m.Data2.UserId)
	x.Int(m.Data2.Date)
	x.Bytes(m.Data2Photo.Encode())
	x.Bytes(m.Data2Previous.Encode())

	return x.buf
}

func (m *TLUpdateUserPhoto) Decode(dbuf *DecodeBuf) error {
	m.Data2.UserId = x.Int()
	m.Data2.Date = x.Int()
	m3 := &Photo{}
	m3.Decode(dbuf)
	m.SetPhoto(m3)
	m4 := &Previous{}
	m4.Decode(dbuf)
	m.SetPrevious(m4)

	return dbuf.err
}

// updateContactRegistered#2575bbb9 user_id:int date:int = Update;
func (m *TLUpdateContactRegistered) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateContactRegistered) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLUpdateContactRegistered) GetUserId() int32  { return m.Data2.UserId }

func (m *TLUpdateContactRegistered) SetDate(v int32) { m.Data2.Date = v }
func (m *TLUpdateContactRegistered) GetDate() int32  { return m.Data2.Date }

func (m *TLUpdateContactRegistered) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateContactRegistered))

	x.Int(m.Data2.UserId)
	x.Int(m.Data2.Date)

	return x.buf
}

func (m *TLUpdateContactRegistered) Decode(dbuf *DecodeBuf) error {
	m.Data2.UserId = x.Int()
	m.Data2.Date = x.Int()

	return dbuf.err
}

// updateContactLink#9d2e67c5 user_id:int my_link:ContactLink foreign_link:ContactLink = Update;
func (m *TLUpdateContactLink) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateContactLink) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLUpdateContactLink) GetUserId() int32  { return m.Data2.UserId }

func (m *TLUpdateContactLink) SetMyLink(v *ContactLink) { m.Data2.MyLink = v }
func (m *TLUpdateContactLink) GetMyLink() *ContactLink  { return m.Data2.MyLink }

func (m *TLUpdateContactLink) SetForeignLink(v *ContactLink) { m.Data2.ForeignLink = v }
func (m *TLUpdateContactLink) GetForeignLink() *ContactLink  { return m.Data2.ForeignLink }

func (m *TLUpdateContactLink) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateContactLink))

	x.Int(m.Data2.UserId)
	x.Bytes(m.Data2MyLink.Encode())
	x.Bytes(m.Data2ForeignLink.Encode())

	return x.buf
}

func (m *TLUpdateContactLink) Decode(dbuf *DecodeBuf) error {
	m.Data2.UserId = x.Int()
	m2 := &MyLink{}
	m2.Decode(dbuf)
	m.SetMyLink(m2)
	m3 := &ForeignLink{}
	m3.Decode(dbuf)
	m.SetForeignLink(m3)

	return dbuf.err
}

// updateNewEncryptedMessage#12bcbd9a message:EncryptedMessage qts:int = Update;
func (m *TLUpdateNewEncryptedMessage) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateNewEncryptedMessage) SetMessage(v *EncryptedMessage) { m.Data2.Message_20 = v }
func (m *TLUpdateNewEncryptedMessage) GetMessage() *EncryptedMessage  { return m.Data2.Message_20 }

func (m *TLUpdateNewEncryptedMessage) SetQts(v int32) { m.Data2.Qts = v }
func (m *TLUpdateNewEncryptedMessage) GetQts() int32  { return m.Data2.Qts }

func (m *TLUpdateNewEncryptedMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateNewEncryptedMessage))

	x.Bytes(m.Data2Message.Encode())
	x.Int(m.Data2.Qts)

	return x.buf
}

func (m *TLUpdateNewEncryptedMessage) Decode(dbuf *DecodeBuf) error {
	m1 := &Message{}
	m1.Decode(dbuf)
	m.SetMessage(m1)
	m.Data2.Qts = x.Int()

	return dbuf.err
}

// updateEncryptedChatTyping#1710f156 chat_id:int = Update;
func (m *TLUpdateEncryptedChatTyping) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateEncryptedChatTyping) SetChatId(v int32) { m.Data2.ChatId = v }
func (m *TLUpdateEncryptedChatTyping) GetChatId() int32  { return m.Data2.ChatId }

func (m *TLUpdateEncryptedChatTyping) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateEncryptedChatTyping))

	x.Int(m.Data2.ChatId)

	return x.buf
}

func (m *TLUpdateEncryptedChatTyping) Decode(dbuf *DecodeBuf) error {
	m.Data2.ChatId = x.Int()

	return dbuf.err
}

// updateEncryption#b4a2e88d chat:EncryptedChat date:int = Update;
func (m *TLUpdateEncryption) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateEncryption) SetChat(v *EncryptedChat) { m.Data2.Chat = v }
func (m *TLUpdateEncryption) GetChat() *EncryptedChat  { return m.Data2.Chat }

func (m *TLUpdateEncryption) SetDate(v int32) { m.Data2.Date = v }
func (m *TLUpdateEncryption) GetDate() int32  { return m.Data2.Date }

func (m *TLUpdateEncryption) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateEncryption))

	x.Bytes(m.Data2Chat.Encode())
	x.Int(m.Data2.Date)

	return x.buf
}

func (m *TLUpdateEncryption) Decode(dbuf *DecodeBuf) error {
	m1 := &Chat{}
	m1.Decode(dbuf)
	m.SetChat(m1)
	m.Data2.Date = x.Int()

	return dbuf.err
}

// updateEncryptedMessagesRead#38fe25b7 chat_id:int max_date:int date:int = Update;
func (m *TLUpdateEncryptedMessagesRead) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateEncryptedMessagesRead) SetChatId(v int32) { m.Data2.ChatId = v }
func (m *TLUpdateEncryptedMessagesRead) GetChatId() int32  { return m.Data2.ChatId }

func (m *TLUpdateEncryptedMessagesRead) SetMaxDate(v int32) { m.Data2.MaxDate = v }
func (m *TLUpdateEncryptedMessagesRead) GetMaxDate() int32  { return m.Data2.MaxDate }

func (m *TLUpdateEncryptedMessagesRead) SetDate(v int32) { m.Data2.Date = v }
func (m *TLUpdateEncryptedMessagesRead) GetDate() int32  { return m.Data2.Date }

func (m *TLUpdateEncryptedMessagesRead) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateEncryptedMessagesRead))

	x.Int(m.Data2.ChatId)
	x.Int(m.Data2.MaxDate)
	x.Int(m.Data2.Date)

	return x.buf
}

func (m *TLUpdateEncryptedMessagesRead) Decode(dbuf *DecodeBuf) error {
	m.Data2.ChatId = x.Int()
	m.Data2.MaxDate = x.Int()
	m.Data2.Date = x.Int()

	return dbuf.err
}

// updateChatParticipantAdd#ea4b0e5c chat_id:int user_id:int inviter_id:int date:int version:int = Update;
func (m *TLUpdateChatParticipantAdd) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateChatParticipantAdd) SetChatId(v int32) { m.Data2.ChatId = v }
func (m *TLUpdateChatParticipantAdd) GetChatId() int32  { return m.Data2.ChatId }

func (m *TLUpdateChatParticipantAdd) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLUpdateChatParticipantAdd) GetUserId() int32  { return m.Data2.UserId }

func (m *TLUpdateChatParticipantAdd) SetInviterId(v int32) { m.Data2.InviterId = v }
func (m *TLUpdateChatParticipantAdd) GetInviterId() int32  { return m.Data2.InviterId }

func (m *TLUpdateChatParticipantAdd) SetDate(v int32) { m.Data2.Date = v }
func (m *TLUpdateChatParticipantAdd) GetDate() int32  { return m.Data2.Date }

func (m *TLUpdateChatParticipantAdd) SetVersion(v int32) { m.Data2.Version = v }
func (m *TLUpdateChatParticipantAdd) GetVersion() int32  { return m.Data2.Version }

func (m *TLUpdateChatParticipantAdd) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateChatParticipantAdd))

	x.Int(m.Data2.ChatId)
	x.Int(m.Data2.UserId)
	x.Int(m.Data2.InviterId)
	x.Int(m.Data2.Date)
	x.Int(m.Data2.Version)

	return x.buf
}

func (m *TLUpdateChatParticipantAdd) Decode(dbuf *DecodeBuf) error {
	m.Data2.ChatId = x.Int()
	m.Data2.UserId = x.Int()
	m.Data2.InviterId = x.Int()
	m.Data2.Date = x.Int()
	m.Data2.Version = x.Int()

	return dbuf.err
}

// updateChatParticipantDelete#6e5f8c22 chat_id:int user_id:int version:int = Update;
func (m *TLUpdateChatParticipantDelete) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateChatParticipantDelete) SetChatId(v int32) { m.Data2.ChatId = v }
func (m *TLUpdateChatParticipantDelete) GetChatId() int32  { return m.Data2.ChatId }

func (m *TLUpdateChatParticipantDelete) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLUpdateChatParticipantDelete) GetUserId() int32  { return m.Data2.UserId }

func (m *TLUpdateChatParticipantDelete) SetVersion(v int32) { m.Data2.Version = v }
func (m *TLUpdateChatParticipantDelete) GetVersion() int32  { return m.Data2.Version }

func (m *TLUpdateChatParticipantDelete) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateChatParticipantDelete))

	x.Int(m.Data2.ChatId)
	x.Int(m.Data2.UserId)
	x.Int(m.Data2.Version)

	return x.buf
}

func (m *TLUpdateChatParticipantDelete) Decode(dbuf *DecodeBuf) error {
	m.Data2.ChatId = x.Int()
	m.Data2.UserId = x.Int()
	m.Data2.Version = x.Int()

	return dbuf.err
}

// updateDcOptions#8e5e9873 dc_options:Vector<DcOption> = Update;
func (m *TLUpdateDcOptions) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateDcOptions) SetDcOptions(v []*DcOption) { m.Data2.DcOptions = v }
func (m *TLUpdateDcOptions) GetDcOptions() []*DcOption  { return m.Data2.DcOptions }

func (m *TLUpdateDcOptions) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateDcOptions))

	return x.buf
}

func (m *TLUpdateDcOptions) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// updateUserBlocked#80ece81a user_id:int blocked:Bool = Update;
func (m *TLUpdateUserBlocked) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateUserBlocked) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLUpdateUserBlocked) GetUserId() int32  { return m.Data2.UserId }

func (m *TLUpdateUserBlocked) SetBlocked(v *Bool) { m.Data2.Blocked = v }
func (m *TLUpdateUserBlocked) GetBlocked() *Bool  { return m.Data2.Blocked }

func (m *TLUpdateUserBlocked) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateUserBlocked))

	x.Int(m.Data2.UserId)
	x.Bytes(m.Data2Blocked.Encode())

	return x.buf
}

func (m *TLUpdateUserBlocked) Decode(dbuf *DecodeBuf) error {
	m.Data2.UserId = x.Int()
	m2 := &Blocked{}
	m2.Decode(dbuf)
	m.SetBlocked(m2)

	return dbuf.err
}

// updateNotifySettings#bec268ef peer:NotifyPeer notify_settings:PeerNotifySettings = Update;
func (m *TLUpdateNotifySettings) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateNotifySettings) SetPeer(v *NotifyPeer) { m.Data2.Peer_28 = v }
func (m *TLUpdateNotifySettings) GetPeer() *NotifyPeer  { return m.Data2.Peer_28 }

func (m *TLUpdateNotifySettings) SetNotifySettings(v *PeerNotifySettings) { m.Data2.NotifySettings = v }
func (m *TLUpdateNotifySettings) GetNotifySettings() *PeerNotifySettings {
	return m.Data2.NotifySettings
}

func (m *TLUpdateNotifySettings) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateNotifySettings))

	x.Bytes(m.Data2Peer.Encode())
	x.Bytes(m.Data2NotifySettings.Encode())

	return x.buf
}

func (m *TLUpdateNotifySettings) Decode(dbuf *DecodeBuf) error {
	m1 := &Peer{}
	m1.Decode(dbuf)
	m.SetPeer(m1)
	m2 := &NotifySettings{}
	m2.Decode(dbuf)
	m.SetNotifySettings(m2)

	return dbuf.err
}

// updateServiceNotification#ebe46819 flags:# popup:flags.0?true inbox_date:flags.1?int type:string message:string media:MessageMedia entities:Vector<MessageEntity> = Update;
func (m *TLUpdateServiceNotification) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateServiceNotification) SetPopup(v bool) { m.Data2.Popup = v }
func (m *TLUpdateServiceNotification) GetPopup() bool  { return m.Data2.Popup }

func (m *TLUpdateServiceNotification) SetInboxDate(v int32) { m.Data2.InboxDate = v }
func (m *TLUpdateServiceNotification) GetInboxDate() int32  { return m.Data2.InboxDate }

func (m *TLUpdateServiceNotification) SetType(v string) { m.Data2.Type = v }
func (m *TLUpdateServiceNotification) GetType() string  { return m.Data2.Type }

func (m *TLUpdateServiceNotification) SetMessage(v string) { m.Data2.Message_33 = v }
func (m *TLUpdateServiceNotification) GetMessage() string  { return m.Data2.Message_33 }

func (m *TLUpdateServiceNotification) SetMedia(v *MessageMedia) { m.Data2.Media = v }
func (m *TLUpdateServiceNotification) GetMedia() *MessageMedia  { return m.Data2.Media }

func (m *TLUpdateServiceNotification) SetEntities(v []*MessageEntity) { m.Data2.Entities = v }
func (m *TLUpdateServiceNotification) GetEntities() []*MessageEntity  { return m.Data2.Entities }

func (m *TLUpdateServiceNotification) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateServiceNotification))

	// flags
	var flags uint32 = 0
	if m.GetPopup() == true {
		flags |= 1 << 1
	}
	if m.GetInboxDate() != 0 {
		flags |= 1 << 2
	}
	x.UInt(flags)

	if m.GetInboxDate() != 0 {
		x.Int(m.Data2.InboxDate)
	}
	x.StringBytes(m.Data2.Type)
	x.StringBytes(m.Data2.Message)
	x.Bytes(m.Data2Media.Encode())

	return x.buf
}

func (m *TLUpdateServiceNotification) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m.Data2.Popup = true
	}
	if (flags & (1 << 3)) != 0 {
		m.Data2.InboxDate = x.Int()
	}
	m.Data2.Type = x.StringBytes()
	m.Data2.Message = x.StringBytes()
	m6 := &Media{}
	m6.Decode(dbuf)
	m.SetMedia(m6)

	return dbuf.err
}

// updatePrivacy#ee3b272a key:PrivacyKey rules:Vector<PrivacyRule> = Update;
func (m *TLUpdatePrivacy) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdatePrivacy) SetKey(v *PrivacyKey) { m.Data2.Key = v }
func (m *TLUpdatePrivacy) GetKey() *PrivacyKey  { return m.Data2.Key }

func (m *TLUpdatePrivacy) SetRules(v []*PrivacyRule) { m.Data2.Rules = v }
func (m *TLUpdatePrivacy) GetRules() []*PrivacyRule  { return m.Data2.Rules }

func (m *TLUpdatePrivacy) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updatePrivacy))

	x.Bytes(m.Data2Key.Encode())

	return x.buf
}

func (m *TLUpdatePrivacy) Decode(dbuf *DecodeBuf) error {
	m1 := &Key{}
	m1.Decode(dbuf)
	m.SetKey(m1)

	return dbuf.err
}

// updateUserPhone#12b9417b user_id:int phone:string = Update;
func (m *TLUpdateUserPhone) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateUserPhone) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLUpdateUserPhone) GetUserId() int32  { return m.Data2.UserId }

func (m *TLUpdateUserPhone) SetPhone(v string) { m.Data2.Phone = v }
func (m *TLUpdateUserPhone) GetPhone() string  { return m.Data2.Phone }

func (m *TLUpdateUserPhone) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateUserPhone))

	x.Int(m.Data2.UserId)
	x.StringBytes(m.Data2.Phone)

	return x.buf
}

func (m *TLUpdateUserPhone) Decode(dbuf *DecodeBuf) error {
	m.Data2.UserId = x.Int()
	m.Data2.Phone = x.StringBytes()

	return dbuf.err
}

// updateReadHistoryInbox#9961fd5c peer:Peer max_id:int pts:int pts_count:int = Update;
func (m *TLUpdateReadHistoryInbox) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateReadHistoryInbox) SetPeer(v *Peer) { m.Data2.Peer_39 = v }
func (m *TLUpdateReadHistoryInbox) GetPeer() *Peer  { return m.Data2.Peer_39 }

func (m *TLUpdateReadHistoryInbox) SetMaxId(v int32) { m.Data2.MaxId = v }
func (m *TLUpdateReadHistoryInbox) GetMaxId() int32  { return m.Data2.MaxId }

func (m *TLUpdateReadHistoryInbox) SetPts(v int32) { m.Data2.Pts = v }
func (m *TLUpdateReadHistoryInbox) GetPts() int32  { return m.Data2.Pts }

func (m *TLUpdateReadHistoryInbox) SetPtsCount(v int32) { m.Data2.PtsCount = v }
func (m *TLUpdateReadHistoryInbox) GetPtsCount() int32  { return m.Data2.PtsCount }

func (m *TLUpdateReadHistoryInbox) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateReadHistoryInbox))

	x.Bytes(m.Data2Peer.Encode())
	x.Int(m.Data2.MaxId)
	x.Int(m.Data2.Pts)
	x.Int(m.Data2.PtsCount)

	return x.buf
}

func (m *TLUpdateReadHistoryInbox) Decode(dbuf *DecodeBuf) error {
	m1 := &Peer{}
	m1.Decode(dbuf)
	m.SetPeer(m1)
	m.Data2.MaxId = x.Int()
	m.Data2.Pts = x.Int()
	m.Data2.PtsCount = x.Int()

	return dbuf.err
}

// updateReadHistoryOutbox#2f2f21bf peer:Peer max_id:int pts:int pts_count:int = Update;
func (m *TLUpdateReadHistoryOutbox) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateReadHistoryOutbox) SetPeer(v *Peer) { m.Data2.Peer_39 = v }
func (m *TLUpdateReadHistoryOutbox) GetPeer() *Peer  { return m.Data2.Peer_39 }

func (m *TLUpdateReadHistoryOutbox) SetMaxId(v int32) { m.Data2.MaxId = v }
func (m *TLUpdateReadHistoryOutbox) GetMaxId() int32  { return m.Data2.MaxId }

func (m *TLUpdateReadHistoryOutbox) SetPts(v int32) { m.Data2.Pts = v }
func (m *TLUpdateReadHistoryOutbox) GetPts() int32  { return m.Data2.Pts }

func (m *TLUpdateReadHistoryOutbox) SetPtsCount(v int32) { m.Data2.PtsCount = v }
func (m *TLUpdateReadHistoryOutbox) GetPtsCount() int32  { return m.Data2.PtsCount }

func (m *TLUpdateReadHistoryOutbox) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateReadHistoryOutbox))

	x.Bytes(m.Data2Peer.Encode())
	x.Int(m.Data2.MaxId)
	x.Int(m.Data2.Pts)
	x.Int(m.Data2.PtsCount)

	return x.buf
}

func (m *TLUpdateReadHistoryOutbox) Decode(dbuf *DecodeBuf) error {
	m1 := &Peer{}
	m1.Decode(dbuf)
	m.SetPeer(m1)
	m.Data2.MaxId = x.Int()
	m.Data2.Pts = x.Int()
	m.Data2.PtsCount = x.Int()

	return dbuf.err
}

// updateWebPage#7f891213 webpage:WebPage pts:int pts_count:int = Update;
func (m *TLUpdateWebPage) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateWebPage) SetWebpage(v *WebPage) { m.Data2.Webpage = v }
func (m *TLUpdateWebPage) GetWebpage() *WebPage  { return m.Data2.Webpage }

func (m *TLUpdateWebPage) SetPts(v int32) { m.Data2.Pts = v }
func (m *TLUpdateWebPage) GetPts() int32  { return m.Data2.Pts }

func (m *TLUpdateWebPage) SetPtsCount(v int32) { m.Data2.PtsCount = v }
func (m *TLUpdateWebPage) GetPtsCount() int32  { return m.Data2.PtsCount }

func (m *TLUpdateWebPage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateWebPage))

	x.Bytes(m.Data2Webpage.Encode())
	x.Int(m.Data2.Pts)
	x.Int(m.Data2.PtsCount)

	return x.buf
}

func (m *TLUpdateWebPage) Decode(dbuf *DecodeBuf) error {
	m1 := &Webpage{}
	m1.Decode(dbuf)
	m.SetWebpage(m1)
	m.Data2.Pts = x.Int()
	m.Data2.PtsCount = x.Int()

	return dbuf.err
}

// updateReadMessagesContents#68c13933 messages:Vector<int> pts:int pts_count:int = Update;
func (m *TLUpdateReadMessagesContents) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateReadMessagesContents) SetMessages(v []int32) { m.Data2.Messages = v }
func (m *TLUpdateReadMessagesContents) GetMessages() []int32  { return m.Data2.Messages }

func (m *TLUpdateReadMessagesContents) SetPts(v int32) { m.Data2.Pts = v }
func (m *TLUpdateReadMessagesContents) GetPts() int32  { return m.Data2.Pts }

func (m *TLUpdateReadMessagesContents) SetPtsCount(v int32) { m.Data2.PtsCount = v }
func (m *TLUpdateReadMessagesContents) GetPtsCount() int32  { return m.Data2.PtsCount }

func (m *TLUpdateReadMessagesContents) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateReadMessagesContents))

	x.Int(m.Data2.Pts)
	x.Int(m.Data2.PtsCount)

	return x.buf
}

func (m *TLUpdateReadMessagesContents) Decode(dbuf *DecodeBuf) error {

	m.Data2.Pts = x.Int()
	m.Data2.PtsCount = x.Int()

	return dbuf.err
}

// updateChannelTooLong#eb0467fb flags:# channel_id:int pts:flags.0?int = Update;
func (m *TLUpdateChannelTooLong) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateChannelTooLong) SetChannelId(v int32) { m.Data2.ChannelId = v }
func (m *TLUpdateChannelTooLong) GetChannelId() int32  { return m.Data2.ChannelId }

func (m *TLUpdateChannelTooLong) SetPts(v int32) { m.Data2.Pts = v }
func (m *TLUpdateChannelTooLong) GetPts() int32  { return m.Data2.Pts }

func (m *TLUpdateChannelTooLong) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateChannelTooLong))

	// flags
	var flags uint32 = 0
	if m.GetPts() != 0 {
		flags |= 1 << 2
	}
	x.UInt(flags)

	x.Int(m.Data2.ChannelId)
	if m.GetPts() != 0 {
		x.Int(m.Data2.Pts)
	}

	return x.buf
}

func (m *TLUpdateChannelTooLong) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	m.Data2.ChannelId = x.Int()
	if (flags & (1 << 3)) != 0 {
		m.Data2.Pts = x.Int()
	}

	return dbuf.err
}

// updateChannel#b6d45656 channel_id:int = Update;
func (m *TLUpdateChannel) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateChannel) SetChannelId(v int32) { m.Data2.ChannelId = v }
func (m *TLUpdateChannel) GetChannelId() int32  { return m.Data2.ChannelId }

func (m *TLUpdateChannel) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateChannel))

	x.Int(m.Data2.ChannelId)

	return x.buf
}

func (m *TLUpdateChannel) Decode(dbuf *DecodeBuf) error {
	m.Data2.ChannelId = x.Int()

	return dbuf.err
}

// updateNewChannelMessage#62ba04d9 message:Message pts:int pts_count:int = Update;
func (m *TLUpdateNewChannelMessage) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateNewChannelMessage) SetMessage(v *Message) { m.Data2.Message_1 = v }
func (m *TLUpdateNewChannelMessage) GetMessage() *Message  { return m.Data2.Message_1 }

func (m *TLUpdateNewChannelMessage) SetPts(v int32) { m.Data2.Pts = v }
func (m *TLUpdateNewChannelMessage) GetPts() int32  { return m.Data2.Pts }

func (m *TLUpdateNewChannelMessage) SetPtsCount(v int32) { m.Data2.PtsCount = v }
func (m *TLUpdateNewChannelMessage) GetPtsCount() int32  { return m.Data2.PtsCount }

func (m *TLUpdateNewChannelMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateNewChannelMessage))

	x.Bytes(m.Data2Message.Encode())
	x.Int(m.Data2.Pts)
	x.Int(m.Data2.PtsCount)

	return x.buf
}

func (m *TLUpdateNewChannelMessage) Decode(dbuf *DecodeBuf) error {
	m1 := &Message{}
	m1.Decode(dbuf)
	m.SetMessage(m1)
	m.Data2.Pts = x.Int()
	m.Data2.PtsCount = x.Int()

	return dbuf.err
}

// updateReadChannelInbox#4214f37f channel_id:int max_id:int = Update;
func (m *TLUpdateReadChannelInbox) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateReadChannelInbox) SetChannelId(v int32) { m.Data2.ChannelId = v }
func (m *TLUpdateReadChannelInbox) GetChannelId() int32  { return m.Data2.ChannelId }

func (m *TLUpdateReadChannelInbox) SetMaxId(v int32) { m.Data2.MaxId = v }
func (m *TLUpdateReadChannelInbox) GetMaxId() int32  { return m.Data2.MaxId }

func (m *TLUpdateReadChannelInbox) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateReadChannelInbox))

	x.Int(m.Data2.ChannelId)
	x.Int(m.Data2.MaxId)

	return x.buf
}

func (m *TLUpdateReadChannelInbox) Decode(dbuf *DecodeBuf) error {
	m.Data2.ChannelId = x.Int()
	m.Data2.MaxId = x.Int()

	return dbuf.err
}

// updateDeleteChannelMessages#c37521c9 channel_id:int messages:Vector<int> pts:int pts_count:int = Update;
func (m *TLUpdateDeleteChannelMessages) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateDeleteChannelMessages) SetChannelId(v int32) { m.Data2.ChannelId = v }
func (m *TLUpdateDeleteChannelMessages) GetChannelId() int32  { return m.Data2.ChannelId }

func (m *TLUpdateDeleteChannelMessages) SetMessages(v []int32) { m.Data2.Messages = v }
func (m *TLUpdateDeleteChannelMessages) GetMessages() []int32  { return m.Data2.Messages }

func (m *TLUpdateDeleteChannelMessages) SetPts(v int32) { m.Data2.Pts = v }
func (m *TLUpdateDeleteChannelMessages) GetPts() int32  { return m.Data2.Pts }

func (m *TLUpdateDeleteChannelMessages) SetPtsCount(v int32) { m.Data2.PtsCount = v }
func (m *TLUpdateDeleteChannelMessages) GetPtsCount() int32  { return m.Data2.PtsCount }

func (m *TLUpdateDeleteChannelMessages) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateDeleteChannelMessages))

	x.Int(m.Data2.ChannelId)

	x.Int(m.Data2.Pts)
	x.Int(m.Data2.PtsCount)

	return x.buf
}

func (m *TLUpdateDeleteChannelMessages) Decode(dbuf *DecodeBuf) error {
	m.Data2.ChannelId = x.Int()

	m.Data2.Pts = x.Int()
	m.Data2.PtsCount = x.Int()

	return dbuf.err
}

// updateChannelMessageViews#98a12b4b channel_id:int id:int views:int = Update;
func (m *TLUpdateChannelMessageViews) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateChannelMessageViews) SetChannelId(v int32) { m.Data2.ChannelId = v }
func (m *TLUpdateChannelMessageViews) GetChannelId() int32  { return m.Data2.ChannelId }

func (m *TLUpdateChannelMessageViews) SetId(v int32) { m.Data2.Id_4 = v }
func (m *TLUpdateChannelMessageViews) GetId() int32  { return m.Data2.Id_4 }

func (m *TLUpdateChannelMessageViews) SetViews(v int32) { m.Data2.Views = v }
func (m *TLUpdateChannelMessageViews) GetViews() int32  { return m.Data2.Views }

func (m *TLUpdateChannelMessageViews) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateChannelMessageViews))

	x.Int(m.Data2.ChannelId)
	x.Int(m.Data2.Id)
	x.Int(m.Data2.Views)

	return x.buf
}

func (m *TLUpdateChannelMessageViews) Decode(dbuf *DecodeBuf) error {
	m.Data2.ChannelId = x.Int()
	m.Data2.Id = x.Int()
	m.Data2.Views = x.Int()

	return dbuf.err
}

// updateChatAdmins#6e947941 chat_id:int enabled:Bool version:int = Update;
func (m *TLUpdateChatAdmins) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateChatAdmins) SetChatId(v int32) { m.Data2.ChatId = v }
func (m *TLUpdateChatAdmins) GetChatId() int32  { return m.Data2.ChatId }

func (m *TLUpdateChatAdmins) SetEnabled(v *Bool) { m.Data2.Enabled = v }
func (m *TLUpdateChatAdmins) GetEnabled() *Bool  { return m.Data2.Enabled }

func (m *TLUpdateChatAdmins) SetVersion(v int32) { m.Data2.Version = v }
func (m *TLUpdateChatAdmins) GetVersion() int32  { return m.Data2.Version }

func (m *TLUpdateChatAdmins) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateChatAdmins))

	x.Int(m.Data2.ChatId)
	x.Bytes(m.Data2Enabled.Encode())
	x.Int(m.Data2.Version)

	return x.buf
}

func (m *TLUpdateChatAdmins) Decode(dbuf *DecodeBuf) error {
	m.Data2.ChatId = x.Int()
	m2 := &Enabled{}
	m2.Decode(dbuf)
	m.SetEnabled(m2)
	m.Data2.Version = x.Int()

	return dbuf.err
}

// updateChatParticipantAdmin#b6901959 chat_id:int user_id:int is_admin:Bool version:int = Update;
func (m *TLUpdateChatParticipantAdmin) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateChatParticipantAdmin) SetChatId(v int32) { m.Data2.ChatId = v }
func (m *TLUpdateChatParticipantAdmin) GetChatId() int32  { return m.Data2.ChatId }

func (m *TLUpdateChatParticipantAdmin) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLUpdateChatParticipantAdmin) GetUserId() int32  { return m.Data2.UserId }

func (m *TLUpdateChatParticipantAdmin) SetIsAdmin(v *Bool) { m.Data2.IsAdmin = v }
func (m *TLUpdateChatParticipantAdmin) GetIsAdmin() *Bool  { return m.Data2.IsAdmin }

func (m *TLUpdateChatParticipantAdmin) SetVersion(v int32) { m.Data2.Version = v }
func (m *TLUpdateChatParticipantAdmin) GetVersion() int32  { return m.Data2.Version }

func (m *TLUpdateChatParticipantAdmin) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateChatParticipantAdmin))

	x.Int(m.Data2.ChatId)
	x.Int(m.Data2.UserId)
	x.Bytes(m.Data2IsAdmin.Encode())
	x.Int(m.Data2.Version)

	return x.buf
}

func (m *TLUpdateChatParticipantAdmin) Decode(dbuf *DecodeBuf) error {
	m.Data2.ChatId = x.Int()
	m.Data2.UserId = x.Int()
	m3 := &IsAdmin{}
	m3.Decode(dbuf)
	m.SetIsAdmin(m3)
	m.Data2.Version = x.Int()

	return dbuf.err
}

// updateNewStickerSet#688a30aa stickerset:messages.StickerSet = Update;
func (m *TLUpdateNewStickerSet) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateNewStickerSet) SetStickerset(v *Messages_StickerSet) { m.Data2.Stickerset = v }
func (m *TLUpdateNewStickerSet) GetStickerset() *Messages_StickerSet  { return m.Data2.Stickerset }

func (m *TLUpdateNewStickerSet) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateNewStickerSet))

	x.Bytes(m.Data2Stickerset.Encode())

	return x.buf
}

func (m *TLUpdateNewStickerSet) Decode(dbuf *DecodeBuf) error {
	m1 := &Stickerset{}
	m1.Decode(dbuf)
	m.SetStickerset(m1)

	return dbuf.err
}

// updateStickerSetsOrder#bb2d201 flags:# masks:flags.0?true order:Vector<long> = Update;
func (m *TLUpdateStickerSetsOrder) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateStickerSetsOrder) SetMasks(v bool) { m.Data2.Masks = v }
func (m *TLUpdateStickerSetsOrder) GetMasks() bool  { return m.Data2.Masks }

func (m *TLUpdateStickerSetsOrder) SetOrder(v []int64) { m.Data2.Order_48 = v }
func (m *TLUpdateStickerSetsOrder) GetOrder() []int64  { return m.Data2.Order_48 }

func (m *TLUpdateStickerSetsOrder) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateStickerSetsOrder))

	// flags
	var flags uint32 = 0
	if m.GetMasks() == true {
		flags |= 1 << 1
	}
	x.UInt(flags)

	return x.buf
}

func (m *TLUpdateStickerSetsOrder) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m.Data2.Masks = true
	}

	return dbuf.err
}

// updateStickerSets#43ae3dec = Update;
func (m *TLUpdateStickerSets) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateStickerSets) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateStickerSets))

	return x.buf
}

func (m *TLUpdateStickerSets) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// updateSavedGifs#9375341e = Update;
func (m *TLUpdateSavedGifs) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateSavedGifs) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateSavedGifs))

	return x.buf
}

func (m *TLUpdateSavedGifs) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// updateBotInlineQuery#54826690 flags:# query_id:long user_id:int query:string geo:flags.0?GeoPoint offset:string = Update;
func (m *TLUpdateBotInlineQuery) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateBotInlineQuery) SetQueryId(v int64) { m.Data2.QueryId = v }
func (m *TLUpdateBotInlineQuery) GetQueryId() int64  { return m.Data2.QueryId }

func (m *TLUpdateBotInlineQuery) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLUpdateBotInlineQuery) GetUserId() int32  { return m.Data2.UserId }

func (m *TLUpdateBotInlineQuery) SetQuery(v string) { m.Data2.Query = v }
func (m *TLUpdateBotInlineQuery) GetQuery() string  { return m.Data2.Query }

func (m *TLUpdateBotInlineQuery) SetGeo(v *GeoPoint) { m.Data2.Geo = v }
func (m *TLUpdateBotInlineQuery) GetGeo() *GeoPoint  { return m.Data2.Geo }

func (m *TLUpdateBotInlineQuery) SetOffset(v string) { m.Data2.Offset = v }
func (m *TLUpdateBotInlineQuery) GetOffset() string  { return m.Data2.Offset }

func (m *TLUpdateBotInlineQuery) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateBotInlineQuery))

	// flags
	var flags uint32 = 0
	if m.GetGeo() != nil {
		flags |= 1 << 4
	}
	x.UInt(flags)

	x.Long(m.Data2.QueryId)
	x.Int(m.Data2.UserId)
	x.StringBytes(m.Data2.Query)
	if m.GetGeo() != 0 {
		x.Bytes(m.Data2Geo.Encode())
	}
	x.StringBytes(m.Data2.Offset)

	return x.buf
}

func (m *TLUpdateBotInlineQuery) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	m.Data2.QueryId = x.Long()
	m.Data2.UserId = x.Int()
	m.Data2.Query = x.StringBytes()
	if (flags & (1 << 5)) != 0 {
		m5 := &Geo{}
		m5.Decode(dbuf)
		m.SetGeo(m5)
	}
	m.Data2.Offset = x.StringBytes()

	return dbuf.err
}

// updateBotInlineSend#e48f964 flags:# user_id:int query:string geo:flags.0?GeoPoint id:string msg_id:flags.1?InputBotInlineMessageID = Update;
func (m *TLUpdateBotInlineSend) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateBotInlineSend) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLUpdateBotInlineSend) GetUserId() int32  { return m.Data2.UserId }

func (m *TLUpdateBotInlineSend) SetQuery(v string) { m.Data2.Query = v }
func (m *TLUpdateBotInlineSend) GetQuery() string  { return m.Data2.Query }

func (m *TLUpdateBotInlineSend) SetGeo(v *GeoPoint) { m.Data2.Geo = v }
func (m *TLUpdateBotInlineSend) GetGeo() *GeoPoint  { return m.Data2.Geo }

func (m *TLUpdateBotInlineSend) SetId(v string) { m.Data2.Id_53 = v }
func (m *TLUpdateBotInlineSend) GetId() string  { return m.Data2.Id_53 }

func (m *TLUpdateBotInlineSend) SetMsgId(v *InputBotInlineMessageID) { m.Data2.MsgId_54 = v }
func (m *TLUpdateBotInlineSend) GetMsgId() *InputBotInlineMessageID  { return m.Data2.MsgId_54 }

func (m *TLUpdateBotInlineSend) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateBotInlineSend))

	// flags
	var flags uint32 = 0
	if m.GetGeo() != nil {
		flags |= 1 << 3
	}
	if m.GetMsgId() != nil {
		flags |= 1 << 5
	}
	x.UInt(flags)

	x.Int(m.Data2.UserId)
	x.StringBytes(m.Data2.Query)
	if m.GetGeo() != 0 {
		x.Bytes(m.Data2Geo.Encode())
	}
	x.StringBytes(m.Data2.Id)
	if m.GetMsgId() != 0 {
		x.Bytes(m.Data2MsgId.Encode())
	}

	return x.buf
}

func (m *TLUpdateBotInlineSend) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	m.Data2.UserId = x.Int()
	m.Data2.Query = x.StringBytes()
	if (flags & (1 << 4)) != 0 {
		m4 := &Geo{}
		m4.Decode(dbuf)
		m.SetGeo(m4)
	}
	m.Data2.Id = x.StringBytes()
	if (flags & (1 << 6)) != 0 {
		m6 := &MsgId{}
		m6.Decode(dbuf)
		m.SetMsgId(m6)
	}

	return dbuf.err
}

// updateEditChannelMessage#1b3f4df7 message:Message pts:int pts_count:int = Update;
func (m *TLUpdateEditChannelMessage) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateEditChannelMessage) SetMessage(v *Message) { m.Data2.Message_1 = v }
func (m *TLUpdateEditChannelMessage) GetMessage() *Message  { return m.Data2.Message_1 }

func (m *TLUpdateEditChannelMessage) SetPts(v int32) { m.Data2.Pts = v }
func (m *TLUpdateEditChannelMessage) GetPts() int32  { return m.Data2.Pts }

func (m *TLUpdateEditChannelMessage) SetPtsCount(v int32) { m.Data2.PtsCount = v }
func (m *TLUpdateEditChannelMessage) GetPtsCount() int32  { return m.Data2.PtsCount }

func (m *TLUpdateEditChannelMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateEditChannelMessage))

	x.Bytes(m.Data2Message.Encode())
	x.Int(m.Data2.Pts)
	x.Int(m.Data2.PtsCount)

	return x.buf
}

func (m *TLUpdateEditChannelMessage) Decode(dbuf *DecodeBuf) error {
	m1 := &Message{}
	m1.Decode(dbuf)
	m.SetMessage(m1)
	m.Data2.Pts = x.Int()
	m.Data2.PtsCount = x.Int()

	return dbuf.err
}

// updateChannelPinnedMessage#98592475 channel_id:int id:int = Update;
func (m *TLUpdateChannelPinnedMessage) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateChannelPinnedMessage) SetChannelId(v int32) { m.Data2.ChannelId = v }
func (m *TLUpdateChannelPinnedMessage) GetChannelId() int32  { return m.Data2.ChannelId }

func (m *TLUpdateChannelPinnedMessage) SetId(v int32) { m.Data2.Id_4 = v }
func (m *TLUpdateChannelPinnedMessage) GetId() int32  { return m.Data2.Id_4 }

func (m *TLUpdateChannelPinnedMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateChannelPinnedMessage))

	x.Int(m.Data2.ChannelId)
	x.Int(m.Data2.Id)

	return x.buf
}

func (m *TLUpdateChannelPinnedMessage) Decode(dbuf *DecodeBuf) error {
	m.Data2.ChannelId = x.Int()
	m.Data2.Id = x.Int()

	return dbuf.err
}

// updateBotCallbackQuery#e73547e1 flags:# query_id:long user_id:int peer:Peer msg_id:int chat_instance:long data:flags.0?bytes game_short_name:flags.1?string = Update;
func (m *TLUpdateBotCallbackQuery) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateBotCallbackQuery) SetQueryId(v int64) { m.Data2.QueryId = v }
func (m *TLUpdateBotCallbackQuery) GetQueryId() int64  { return m.Data2.QueryId }

func (m *TLUpdateBotCallbackQuery) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLUpdateBotCallbackQuery) GetUserId() int32  { return m.Data2.UserId }

func (m *TLUpdateBotCallbackQuery) SetPeer(v *Peer) { m.Data2.Peer_39 = v }
func (m *TLUpdateBotCallbackQuery) GetPeer() *Peer  { return m.Data2.Peer_39 }

func (m *TLUpdateBotCallbackQuery) SetMsgId(v int32) { m.Data2.MsgId_55 = v }
func (m *TLUpdateBotCallbackQuery) GetMsgId() int32  { return m.Data2.MsgId_55 }

func (m *TLUpdateBotCallbackQuery) SetChatInstance(v int64) { m.Data2.ChatInstance = v }
func (m *TLUpdateBotCallbackQuery) GetChatInstance() int64  { return m.Data2.ChatInstance }

func (m *TLUpdateBotCallbackQuery) SetData(v []byte) { m.Data2.Data_57 = v }
func (m *TLUpdateBotCallbackQuery) GetData() []byte  { return m.Data2.Data_57 }

func (m *TLUpdateBotCallbackQuery) SetGameShortName(v string) { m.Data2.GameShortName = v }
func (m *TLUpdateBotCallbackQuery) GetGameShortName() string  { return m.Data2.GameShortName }

func (m *TLUpdateBotCallbackQuery) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateBotCallbackQuery))

	// flags
	var flags uint32 = 0
	if m.GetData() != nil {
		flags |= 1 << 6
	}
	if m.GetGameShortName() != "" {
		flags |= 1 << 7
	}
	x.UInt(flags)

	x.Long(m.Data2.QueryId)
	x.Int(m.Data2.UserId)
	x.Bytes(m.Data2Peer.Encode())
	x.Int(m.Data2.MsgId)
	x.Long(m.Data2.ChatInstance)
	if m.GetData() != 0 {
		x.StringBytes(m.Data2.Data)
	}
	if m.GetGameShortName() != 0 {
		x.StringBytes(m.Data2.GameShortName)
	}

	return x.buf
}

func (m *TLUpdateBotCallbackQuery) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	m.Data2.QueryId = x.Long()
	m.Data2.UserId = x.Int()
	m4 := &Peer{}
	m4.Decode(dbuf)
	m.SetPeer(m4)
	m.Data2.MsgId = x.Int()
	m.Data2.ChatInstance = x.Long()
	if (flags & (1 << 7)) != 0 {
		m.Data2.Data = x.StringBytes()
	}
	if (flags & (1 << 8)) != 0 {
		m.Data2.GameShortName = x.StringBytes()
	}

	return dbuf.err
}

// updateEditMessage#e40370a3 message:Message pts:int pts_count:int = Update;
func (m *TLUpdateEditMessage) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateEditMessage) SetMessage(v *Message) { m.Data2.Message_1 = v }
func (m *TLUpdateEditMessage) GetMessage() *Message  { return m.Data2.Message_1 }

func (m *TLUpdateEditMessage) SetPts(v int32) { m.Data2.Pts = v }
func (m *TLUpdateEditMessage) GetPts() int32  { return m.Data2.Pts }

func (m *TLUpdateEditMessage) SetPtsCount(v int32) { m.Data2.PtsCount = v }
func (m *TLUpdateEditMessage) GetPtsCount() int32  { return m.Data2.PtsCount }

func (m *TLUpdateEditMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateEditMessage))

	x.Bytes(m.Data2Message.Encode())
	x.Int(m.Data2.Pts)
	x.Int(m.Data2.PtsCount)

	return x.buf
}

func (m *TLUpdateEditMessage) Decode(dbuf *DecodeBuf) error {
	m1 := &Message{}
	m1.Decode(dbuf)
	m.SetMessage(m1)
	m.Data2.Pts = x.Int()
	m.Data2.PtsCount = x.Int()

	return dbuf.err
}

// updateInlineBotCallbackQuery#f9d27a5a flags:# query_id:long user_id:int msg_id:InputBotInlineMessageID chat_instance:long data:flags.0?bytes game_short_name:flags.1?string = Update;
func (m *TLUpdateInlineBotCallbackQuery) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateInlineBotCallbackQuery) SetQueryId(v int64) { m.Data2.QueryId = v }
func (m *TLUpdateInlineBotCallbackQuery) GetQueryId() int64  { return m.Data2.QueryId }

func (m *TLUpdateInlineBotCallbackQuery) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLUpdateInlineBotCallbackQuery) GetUserId() int32  { return m.Data2.UserId }

func (m *TLUpdateInlineBotCallbackQuery) SetMsgId(v *InputBotInlineMessageID) { m.Data2.MsgId_54 = v }
func (m *TLUpdateInlineBotCallbackQuery) GetMsgId() *InputBotInlineMessageID  { return m.Data2.MsgId_54 }

func (m *TLUpdateInlineBotCallbackQuery) SetChatInstance(v int64) { m.Data2.ChatInstance = v }
func (m *TLUpdateInlineBotCallbackQuery) GetChatInstance() int64  { return m.Data2.ChatInstance }

func (m *TLUpdateInlineBotCallbackQuery) SetData(v []byte) { m.Data2.Data_57 = v }
func (m *TLUpdateInlineBotCallbackQuery) GetData() []byte  { return m.Data2.Data_57 }

func (m *TLUpdateInlineBotCallbackQuery) SetGameShortName(v string) { m.Data2.GameShortName = v }
func (m *TLUpdateInlineBotCallbackQuery) GetGameShortName() string  { return m.Data2.GameShortName }

func (m *TLUpdateInlineBotCallbackQuery) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateInlineBotCallbackQuery))

	// flags
	var flags uint32 = 0
	if m.GetData() != nil {
		flags |= 1 << 5
	}
	if m.GetGameShortName() != "" {
		flags |= 1 << 6
	}
	x.UInt(flags)

	x.Long(m.Data2.QueryId)
	x.Int(m.Data2.UserId)
	x.Bytes(m.Data2MsgId.Encode())
	x.Long(m.Data2.ChatInstance)
	if m.GetData() != 0 {
		x.StringBytes(m.Data2.Data)
	}
	if m.GetGameShortName() != 0 {
		x.StringBytes(m.Data2.GameShortName)
	}

	return x.buf
}

func (m *TLUpdateInlineBotCallbackQuery) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	m.Data2.QueryId = x.Long()
	m.Data2.UserId = x.Int()
	m4 := &MsgId{}
	m4.Decode(dbuf)
	m.SetMsgId(m4)
	m.Data2.ChatInstance = x.Long()
	if (flags & (1 << 6)) != 0 {
		m.Data2.Data = x.StringBytes()
	}
	if (flags & (1 << 7)) != 0 {
		m.Data2.GameShortName = x.StringBytes()
	}

	return dbuf.err
}

// updateReadChannelOutbox#25d6c9c7 channel_id:int max_id:int = Update;
func (m *TLUpdateReadChannelOutbox) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateReadChannelOutbox) SetChannelId(v int32) { m.Data2.ChannelId = v }
func (m *TLUpdateReadChannelOutbox) GetChannelId() int32  { return m.Data2.ChannelId }

func (m *TLUpdateReadChannelOutbox) SetMaxId(v int32) { m.Data2.MaxId = v }
func (m *TLUpdateReadChannelOutbox) GetMaxId() int32  { return m.Data2.MaxId }

func (m *TLUpdateReadChannelOutbox) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateReadChannelOutbox))

	x.Int(m.Data2.ChannelId)
	x.Int(m.Data2.MaxId)

	return x.buf
}

func (m *TLUpdateReadChannelOutbox) Decode(dbuf *DecodeBuf) error {
	m.Data2.ChannelId = x.Int()
	m.Data2.MaxId = x.Int()

	return dbuf.err
}

// updateDraftMessage#ee2bb969 peer:Peer draft:DraftMessage = Update;
func (m *TLUpdateDraftMessage) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateDraftMessage) SetPeer(v *Peer) { m.Data2.Peer_39 = v }
func (m *TLUpdateDraftMessage) GetPeer() *Peer  { return m.Data2.Peer_39 }

func (m *TLUpdateDraftMessage) SetDraft(v *DraftMessage) { m.Data2.Draft = v }
func (m *TLUpdateDraftMessage) GetDraft() *DraftMessage  { return m.Data2.Draft }

func (m *TLUpdateDraftMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateDraftMessage))

	x.Bytes(m.Data2Peer.Encode())
	x.Bytes(m.Data2Draft.Encode())

	return x.buf
}

func (m *TLUpdateDraftMessage) Decode(dbuf *DecodeBuf) error {
	m1 := &Peer{}
	m1.Decode(dbuf)
	m.SetPeer(m1)
	m2 := &Draft{}
	m2.Decode(dbuf)
	m.SetDraft(m2)

	return dbuf.err
}

// updateReadFeaturedStickers#571d2742 = Update;
func (m *TLUpdateReadFeaturedStickers) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateReadFeaturedStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateReadFeaturedStickers))

	return x.buf
}

func (m *TLUpdateReadFeaturedStickers) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// updateRecentStickers#9a422c20 = Update;
func (m *TLUpdateRecentStickers) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateRecentStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateRecentStickers))

	return x.buf
}

func (m *TLUpdateRecentStickers) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// updateConfig#a229dd06 = Update;
func (m *TLUpdateConfig) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateConfig) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateConfig))

	return x.buf
}

func (m *TLUpdateConfig) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// updatePtsChanged#3354678f = Update;
func (m *TLUpdatePtsChanged) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdatePtsChanged) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updatePtsChanged))

	return x.buf
}

func (m *TLUpdatePtsChanged) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// updateChannelWebPage#40771900 channel_id:int webpage:WebPage pts:int pts_count:int = Update;
func (m *TLUpdateChannelWebPage) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateChannelWebPage) SetChannelId(v int32) { m.Data2.ChannelId = v }
func (m *TLUpdateChannelWebPage) GetChannelId() int32  { return m.Data2.ChannelId }

func (m *TLUpdateChannelWebPage) SetWebpage(v *WebPage) { m.Data2.Webpage = v }
func (m *TLUpdateChannelWebPage) GetWebpage() *WebPage  { return m.Data2.Webpage }

func (m *TLUpdateChannelWebPage) SetPts(v int32) { m.Data2.Pts = v }
func (m *TLUpdateChannelWebPage) GetPts() int32  { return m.Data2.Pts }

func (m *TLUpdateChannelWebPage) SetPtsCount(v int32) { m.Data2.PtsCount = v }
func (m *TLUpdateChannelWebPage) GetPtsCount() int32  { return m.Data2.PtsCount }

func (m *TLUpdateChannelWebPage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateChannelWebPage))

	x.Int(m.Data2.ChannelId)
	x.Bytes(m.Data2Webpage.Encode())
	x.Int(m.Data2.Pts)
	x.Int(m.Data2.PtsCount)

	return x.buf
}

func (m *TLUpdateChannelWebPage) Decode(dbuf *DecodeBuf) error {
	m.Data2.ChannelId = x.Int()
	m2 := &Webpage{}
	m2.Decode(dbuf)
	m.SetWebpage(m2)
	m.Data2.Pts = x.Int()
	m.Data2.PtsCount = x.Int()

	return dbuf.err
}

// updateDialogPinned#d711a2cc flags:# pinned:flags.0?true peer:Peer = Update;
func (m *TLUpdateDialogPinned) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateDialogPinned) SetPinned(v bool) { m.Data2.Pinned = v }
func (m *TLUpdateDialogPinned) GetPinned() bool  { return m.Data2.Pinned }

func (m *TLUpdateDialogPinned) SetPeer(v *Peer) { m.Data2.Peer_39 = v }
func (m *TLUpdateDialogPinned) GetPeer() *Peer  { return m.Data2.Peer_39 }

func (m *TLUpdateDialogPinned) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateDialogPinned))

	// flags
	var flags uint32 = 0
	if m.GetPinned() == true {
		flags |= 1 << 1
	}
	x.UInt(flags)

	x.Bytes(m.Data2Peer.Encode())

	return x.buf
}

func (m *TLUpdateDialogPinned) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m.Data2.Pinned = true
	}
	m3 := &Peer{}
	m3.Decode(dbuf)
	m.SetPeer(m3)

	return dbuf.err
}

// updatePinnedDialogs#d8caf68d flags:# order:flags.0?Vector<Peer> = Update;
func (m *TLUpdatePinnedDialogs) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdatePinnedDialogs) SetOrder(v []*Peer) { m.Data2.Order_61 = v }
func (m *TLUpdatePinnedDialogs) GetOrder() []*Peer  { return m.Data2.Order_61 }

func (m *TLUpdatePinnedDialogs) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updatePinnedDialogs))

	// flags
	var flags uint32 = 0
	if m.GetOrder() != nil {
		flags |= 1 << 1
	}
	x.UInt(flags)

	if m.Data2.Order != 0 {

	}

	return x.buf
}

func (m *TLUpdatePinnedDialogs) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {

	}

	return dbuf.err
}

// updateBotWebhookJSON#8317c0c3 data:DataJSON = Update;
func (m *TLUpdateBotWebhookJSON) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateBotWebhookJSON) SetData(v *DataJSON) { m.Data2.Data_62 = v }
func (m *TLUpdateBotWebhookJSON) GetData() *DataJSON  { return m.Data2.Data_62 }

func (m *TLUpdateBotWebhookJSON) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateBotWebhookJSON))

	x.Bytes(m.Data2Data.Encode())

	return x.buf
}

func (m *TLUpdateBotWebhookJSON) Decode(dbuf *DecodeBuf) error {
	m1 := &Data{}
	m1.Decode(dbuf)
	m.SetData(m1)

	return dbuf.err
}

// updateBotWebhookJSONQuery#9b9240a6 query_id:long data:DataJSON timeout:int = Update;
func (m *TLUpdateBotWebhookJSONQuery) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateBotWebhookJSONQuery) SetQueryId(v int64) { m.Data2.QueryId = v }
func (m *TLUpdateBotWebhookJSONQuery) GetQueryId() int64  { return m.Data2.QueryId }

func (m *TLUpdateBotWebhookJSONQuery) SetData(v *DataJSON) { m.Data2.Data_62 = v }
func (m *TLUpdateBotWebhookJSONQuery) GetData() *DataJSON  { return m.Data2.Data_62 }

func (m *TLUpdateBotWebhookJSONQuery) SetTimeout(v int32) { m.Data2.Timeout = v }
func (m *TLUpdateBotWebhookJSONQuery) GetTimeout() int32  { return m.Data2.Timeout }

func (m *TLUpdateBotWebhookJSONQuery) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateBotWebhookJSONQuery))

	x.Long(m.Data2.QueryId)
	x.Bytes(m.Data2Data.Encode())
	x.Int(m.Data2.Timeout)

	return x.buf
}

func (m *TLUpdateBotWebhookJSONQuery) Decode(dbuf *DecodeBuf) error {
	m.Data2.QueryId = x.Long()
	m2 := &Data{}
	m2.Decode(dbuf)
	m.SetData(m2)
	m.Data2.Timeout = x.Int()

	return dbuf.err
}

// updateBotShippingQuery#e0cdc940 query_id:long user_id:int payload:bytes shipping_address:PostAddress = Update;
func (m *TLUpdateBotShippingQuery) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateBotShippingQuery) SetQueryId(v int64) { m.Data2.QueryId = v }
func (m *TLUpdateBotShippingQuery) GetQueryId() int64  { return m.Data2.QueryId }

func (m *TLUpdateBotShippingQuery) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLUpdateBotShippingQuery) GetUserId() int32  { return m.Data2.UserId }

func (m *TLUpdateBotShippingQuery) SetPayload(v []byte) { m.Data2.Payload = v }
func (m *TLUpdateBotShippingQuery) GetPayload() []byte  { return m.Data2.Payload }

func (m *TLUpdateBotShippingQuery) SetShippingAddress(v *PostAddress) { m.Data2.ShippingAddress = v }
func (m *TLUpdateBotShippingQuery) GetShippingAddress() *PostAddress  { return m.Data2.ShippingAddress }

func (m *TLUpdateBotShippingQuery) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateBotShippingQuery))

	x.Long(m.Data2.QueryId)
	x.Int(m.Data2.UserId)
	x.StringBytes(m.Data2.Payload)
	x.Bytes(m.Data2ShippingAddress.Encode())

	return x.buf
}

func (m *TLUpdateBotShippingQuery) Decode(dbuf *DecodeBuf) error {
	m.Data2.QueryId = x.Long()
	m.Data2.UserId = x.Int()
	m.Data2.Payload = x.StringBytes()
	m4 := &ShippingAddress{}
	m4.Decode(dbuf)
	m.SetShippingAddress(m4)

	return dbuf.err
}

// updateBotPrecheckoutQuery#5d2f3aa9 flags:# query_id:long user_id:int payload:bytes info:flags.0?PaymentRequestedInfo shipping_option_id:flags.1?string currency:string total_amount:long = Update;
func (m *TLUpdateBotPrecheckoutQuery) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateBotPrecheckoutQuery) SetQueryId(v int64) { m.Data2.QueryId = v }
func (m *TLUpdateBotPrecheckoutQuery) GetQueryId() int64  { return m.Data2.QueryId }

func (m *TLUpdateBotPrecheckoutQuery) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLUpdateBotPrecheckoutQuery) GetUserId() int32  { return m.Data2.UserId }

func (m *TLUpdateBotPrecheckoutQuery) SetPayload(v []byte) { m.Data2.Payload = v }
func (m *TLUpdateBotPrecheckoutQuery) GetPayload() []byte  { return m.Data2.Payload }

func (m *TLUpdateBotPrecheckoutQuery) SetInfo(v *PaymentRequestedInfo) { m.Data2.Info = v }
func (m *TLUpdateBotPrecheckoutQuery) GetInfo() *PaymentRequestedInfo  { return m.Data2.Info }

func (m *TLUpdateBotPrecheckoutQuery) SetShippingOptionId(v string) { m.Data2.ShippingOptionId = v }
func (m *TLUpdateBotPrecheckoutQuery) GetShippingOptionId() string  { return m.Data2.ShippingOptionId }

func (m *TLUpdateBotPrecheckoutQuery) SetCurrency(v string) { m.Data2.Currency = v }
func (m *TLUpdateBotPrecheckoutQuery) GetCurrency() string  { return m.Data2.Currency }

func (m *TLUpdateBotPrecheckoutQuery) SetTotalAmount(v int64) { m.Data2.TotalAmount = v }
func (m *TLUpdateBotPrecheckoutQuery) GetTotalAmount() int64  { return m.Data2.TotalAmount }

func (m *TLUpdateBotPrecheckoutQuery) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateBotPrecheckoutQuery))

	// flags
	var flags uint32 = 0
	if m.GetInfo() != nil {
		flags |= 1 << 4
	}
	if m.GetShippingOptionId() != "" {
		flags |= 1 << 5
	}
	x.UInt(flags)

	x.Long(m.Data2.QueryId)
	x.Int(m.Data2.UserId)
	x.StringBytes(m.Data2.Payload)
	if m.GetInfo() != 0 {
		x.Bytes(m.Data2Info.Encode())
	}
	if m.GetShippingOptionId() != 0 {
		x.StringBytes(m.Data2.ShippingOptionId)
	}
	x.StringBytes(m.Data2.Currency)
	x.Long(m.Data2.TotalAmount)

	return x.buf
}

func (m *TLUpdateBotPrecheckoutQuery) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	m.Data2.QueryId = x.Long()
	m.Data2.UserId = x.Int()
	m.Data2.Payload = x.StringBytes()
	if (flags & (1 << 5)) != 0 {
		m5 := &Info{}
		m5.Decode(dbuf)
		m.SetInfo(m5)
	}
	if (flags & (1 << 6)) != 0 {
		m.Data2.ShippingOptionId = x.StringBytes()
	}
	m.Data2.Currency = x.StringBytes()
	m.Data2.TotalAmount = x.Long()

	return dbuf.err
}

// updatePhoneCall#ab0f6b1e phone_call:PhoneCall = Update;
func (m *TLUpdatePhoneCall) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdatePhoneCall) SetPhoneCall(v *PhoneCall) { m.Data2.PhoneCall = v }
func (m *TLUpdatePhoneCall) GetPhoneCall() *PhoneCall  { return m.Data2.PhoneCall }

func (m *TLUpdatePhoneCall) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updatePhoneCall))

	x.Bytes(m.Data2PhoneCall.Encode())

	return x.buf
}

func (m *TLUpdatePhoneCall) Decode(dbuf *DecodeBuf) error {
	m1 := &PhoneCall{}
	m1.Decode(dbuf)
	m.SetPhoneCall(m1)

	return dbuf.err
}

// updateLangPackTooLong#10c2404b = Update;
func (m *TLUpdateLangPackTooLong) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateLangPackTooLong) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateLangPackTooLong))

	return x.buf
}

func (m *TLUpdateLangPackTooLong) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// updateLangPack#56022f4d difference:LangPackDifference = Update;
func (m *TLUpdateLangPack) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateLangPack) SetDifference(v *LangPackDifference) { m.Data2.Difference = v }
func (m *TLUpdateLangPack) GetDifference() *LangPackDifference  { return m.Data2.Difference }

func (m *TLUpdateLangPack) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateLangPack))

	x.Bytes(m.Data2Difference.Encode())

	return x.buf
}

func (m *TLUpdateLangPack) Decode(dbuf *DecodeBuf) error {
	m1 := &Difference{}
	m1.Decode(dbuf)
	m.SetDifference(m1)

	return dbuf.err
}

// updateFavedStickers#e511996d = Update;
func (m *TLUpdateFavedStickers) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateFavedStickers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateFavedStickers))

	return x.buf
}

func (m *TLUpdateFavedStickers) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// updateChannelReadMessagesContents#89893b45 channel_id:int messages:Vector<int> = Update;
func (m *TLUpdateChannelReadMessagesContents) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateChannelReadMessagesContents) SetChannelId(v int32) { m.Data2.ChannelId = v }
func (m *TLUpdateChannelReadMessagesContents) GetChannelId() int32  { return m.Data2.ChannelId }

func (m *TLUpdateChannelReadMessagesContents) SetMessages(v []int32) { m.Data2.Messages = v }
func (m *TLUpdateChannelReadMessagesContents) GetMessages() []int32  { return m.Data2.Messages }

func (m *TLUpdateChannelReadMessagesContents) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateChannelReadMessagesContents))

	x.Int(m.Data2.ChannelId)

	return x.buf
}

func (m *TLUpdateChannelReadMessagesContents) Decode(dbuf *DecodeBuf) error {
	m.Data2.ChannelId = x.Int()

	return dbuf.err
}

// updateContactsReset#7084a7be = Update;
func (m *TLUpdateContactsReset) To_Update() *Update {
	return &Update{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUpdateContactsReset) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_updateContactsReset))

	return x.buf
}

func (m *TLUpdateContactsReset) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// BotInlineResult <--
//  + TL_BotInlineResult
//  + TL_BotInlineMediaResult
//

func (m *BotInlineResult) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_botInlineResult:
		t := m.To_BotInlineResult()
		return t.Encode()
	case TLConstructor_CRC32_botInlineMediaResult:
		t := m.To_BotInlineMediaResult()
		return t.Encode()

	default:
		return []byte{}
	}
}

// botInlineResult#9bebaeb9 flags:# id:string type:string title:flags.1?string description:flags.2?string url:flags.3?string thumb_url:flags.4?string content_url:flags.5?string content_type:flags.5?string w:flags.6?int h:flags.6?int duration:flags.7?int send_message:BotInlineMessage = BotInlineResult;
func (m *BotInlineResult) To_BotInlineResult() *TLBotInlineResult {
	return &TLBotInlineResult{
		Data2: m.Data2,
	}
}

// botInlineMediaResult#17db940b flags:# id:string type:string photo:flags.0?Photo document:flags.1?Document title:flags.2?string description:flags.3?string send_message:BotInlineMessage = BotInlineResult;
func (m *BotInlineResult) To_BotInlineMediaResult() *TLBotInlineMediaResult {
	return &TLBotInlineMediaResult{
		Data2: m.Data2,
	}
}

// botInlineResult#9bebaeb9 flags:# id:string type:string title:flags.1?string description:flags.2?string url:flags.3?string thumb_url:flags.4?string content_url:flags.5?string content_type:flags.5?string w:flags.6?int h:flags.6?int duration:flags.7?int send_message:BotInlineMessage = BotInlineResult;
func (m *TLBotInlineResult) To_BotInlineResult() *BotInlineResult {
	return &BotInlineResult{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLBotInlineResult) SetId(v string) { m.Data2.Id = v }
func (m *TLBotInlineResult) GetId() string  { return m.Data2.Id }

func (m *TLBotInlineResult) SetType(v string) { m.Data2.Type = v }
func (m *TLBotInlineResult) GetType() string  { return m.Data2.Type }

func (m *TLBotInlineResult) SetTitle(v string) { m.Data2.Title = v }
func (m *TLBotInlineResult) GetTitle() string  { return m.Data2.Title }

func (m *TLBotInlineResult) SetDescription(v string) { m.Data2.Description = v }
func (m *TLBotInlineResult) GetDescription() string  { return m.Data2.Description }

func (m *TLBotInlineResult) SetUrl(v string) { m.Data2.Url = v }
func (m *TLBotInlineResult) GetUrl() string  { return m.Data2.Url }

func (m *TLBotInlineResult) SetThumbUrl(v string) { m.Data2.ThumbUrl = v }
func (m *TLBotInlineResult) GetThumbUrl() string  { return m.Data2.ThumbUrl }

func (m *TLBotInlineResult) SetContentUrl(v string) { m.Data2.ContentUrl = v }
func (m *TLBotInlineResult) GetContentUrl() string  { return m.Data2.ContentUrl }

func (m *TLBotInlineResult) SetContentType(v string) { m.Data2.ContentType = v }
func (m *TLBotInlineResult) GetContentType() string  { return m.Data2.ContentType }

func (m *TLBotInlineResult) SetW(v int32) { m.Data2.W = v }
func (m *TLBotInlineResult) GetW() int32  { return m.Data2.W }

func (m *TLBotInlineResult) SetH(v int32) { m.Data2.H = v }
func (m *TLBotInlineResult) GetH() int32  { return m.Data2.H }

func (m *TLBotInlineResult) SetDuration(v int32) { m.Data2.Duration = v }
func (m *TLBotInlineResult) GetDuration() int32  { return m.Data2.Duration }

func (m *TLBotInlineResult) SetSendMessage(v *BotInlineMessage) { m.Data2.SendMessage = v }
func (m *TLBotInlineResult) GetSendMessage() *BotInlineMessage  { return m.Data2.SendMessage }

func (m *TLBotInlineResult) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_botInlineResult))

	// flags
	var flags uint32 = 0
	if m.GetTitle() != "" {
		flags |= 1 << 3
	}
	if m.GetDescription() != "" {
		flags |= 1 << 4
	}
	if m.GetUrl() != "" {
		flags |= 1 << 5
	}
	if m.GetThumbUrl() != "" {
		flags |= 1 << 6
	}
	if m.GetContentUrl() != "" {
		flags |= 1 << 7
	}
	if m.GetContentType() != "" {
		flags |= 1 << 8
	}
	if m.GetW() != 0 {
		flags |= 1 << 9
	}
	if m.GetH() != 0 {
		flags |= 1 << 10
	}
	if m.GetDuration() != 0 {
		flags |= 1 << 11
	}
	x.UInt(flags)

	x.StringBytes(m.Data2.Id)
	x.StringBytes(m.Data2.Type)
	if m.GetTitle() != 0 {
		x.StringBytes(m.Data2.Title)
	}
	if m.GetDescription() != 0 {
		x.StringBytes(m.Data2.Description)
	}
	if m.GetUrl() != 0 {
		x.StringBytes(m.Data2.Url)
	}
	if m.GetThumbUrl() != 0 {
		x.StringBytes(m.Data2.ThumbUrl)
	}
	if m.GetContentUrl() != 0 {
		x.StringBytes(m.Data2.ContentUrl)
	}
	if m.GetContentType() != 0 {
		x.StringBytes(m.Data2.ContentType)
	}
	if m.GetW() != 0 {
		x.Int(m.Data2.W)
	}
	if m.GetH() != 0 {
		x.Int(m.Data2.H)
	}
	if m.GetDuration() != 0 {
		x.Int(m.Data2.Duration)
	}
	x.Bytes(m.Data2SendMessage.Encode())

	return x.buf
}

func (m *TLBotInlineResult) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	m.Data2.Id = x.StringBytes()
	m.Data2.Type = x.StringBytes()
	if (flags & (1 << 4)) != 0 {
		m.Data2.Title = x.StringBytes()
	}
	if (flags & (1 << 5)) != 0 {
		m.Data2.Description = x.StringBytes()
	}
	if (flags & (1 << 6)) != 0 {
		m.Data2.Url = x.StringBytes()
	}
	if (flags & (1 << 7)) != 0 {
		m.Data2.ThumbUrl = x.StringBytes()
	}
	if (flags & (1 << 8)) != 0 {
		m.Data2.ContentUrl = x.StringBytes()
	}
	if (flags & (1 << 9)) != 0 {
		m.Data2.ContentType = x.StringBytes()
	}
	if (flags & (1 << 10)) != 0 {
		m.Data2.W = x.Int()
	}
	if (flags & (1 << 11)) != 0 {
		m.Data2.H = x.Int()
	}
	if (flags & (1 << 12)) != 0 {
		m.Data2.Duration = x.Int()
	}
	m13 := &SendMessage{}
	m13.Decode(dbuf)
	m.SetSendMessage(m13)

	return dbuf.err
}

// botInlineMediaResult#17db940b flags:# id:string type:string photo:flags.0?Photo document:flags.1?Document title:flags.2?string description:flags.3?string send_message:BotInlineMessage = BotInlineResult;
func (m *TLBotInlineMediaResult) To_BotInlineResult() *BotInlineResult {
	return &BotInlineResult{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLBotInlineMediaResult) SetId(v string) { m.Data2.Id = v }
func (m *TLBotInlineMediaResult) GetId() string  { return m.Data2.Id }

func (m *TLBotInlineMediaResult) SetType(v string) { m.Data2.Type = v }
func (m *TLBotInlineMediaResult) GetType() string  { return m.Data2.Type }

func (m *TLBotInlineMediaResult) SetPhoto(v *Photo) { m.Data2.Photo = v }
func (m *TLBotInlineMediaResult) GetPhoto() *Photo  { return m.Data2.Photo }

func (m *TLBotInlineMediaResult) SetDocument(v *Document) { m.Data2.Document = v }
func (m *TLBotInlineMediaResult) GetDocument() *Document  { return m.Data2.Document }

func (m *TLBotInlineMediaResult) SetTitle(v string) { m.Data2.Title = v }
func (m *TLBotInlineMediaResult) GetTitle() string  { return m.Data2.Title }

func (m *TLBotInlineMediaResult) SetDescription(v string) { m.Data2.Description = v }
func (m *TLBotInlineMediaResult) GetDescription() string  { return m.Data2.Description }

func (m *TLBotInlineMediaResult) SetSendMessage(v *BotInlineMessage) { m.Data2.SendMessage = v }
func (m *TLBotInlineMediaResult) GetSendMessage() *BotInlineMessage  { return m.Data2.SendMessage }

func (m *TLBotInlineMediaResult) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_botInlineMediaResult))

	// flags
	var flags uint32 = 0
	if m.GetPhoto() != nil {
		flags |= 1 << 3
	}
	if m.GetDocument() != nil {
		flags |= 1 << 4
	}
	if m.GetTitle() != "" {
		flags |= 1 << 5
	}
	if m.GetDescription() != "" {
		flags |= 1 << 6
	}
	x.UInt(flags)

	x.StringBytes(m.Data2.Id)
	x.StringBytes(m.Data2.Type)
	if m.GetPhoto() != 0 {
		x.Bytes(m.Data2Photo.Encode())
	}
	if m.GetDocument() != 0 {
		x.Bytes(m.Data2Document.Encode())
	}
	if m.GetTitle() != 0 {
		x.StringBytes(m.Data2.Title)
	}
	if m.GetDescription() != 0 {
		x.StringBytes(m.Data2.Description)
	}
	x.Bytes(m.Data2SendMessage.Encode())

	return x.buf
}

func (m *TLBotInlineMediaResult) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	m.Data2.Id = x.StringBytes()
	m.Data2.Type = x.StringBytes()
	if (flags & (1 << 4)) != 0 {
		m4 := &Photo{}
		m4.Decode(dbuf)
		m.SetPhoto(m4)
	}
	if (flags & (1 << 5)) != 0 {
		m5 := &Document{}
		m5.Decode(dbuf)
		m.SetDocument(m5)
	}
	if (flags & (1 << 6)) != 0 {
		m.Data2.Title = x.StringBytes()
	}
	if (flags & (1 << 7)) != 0 {
		m.Data2.Description = x.StringBytes()
	}
	m8 := &SendMessage{}
	m8.Decode(dbuf)
	m.SetSendMessage(m8)

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Auth_CodeType <--
//  + TL_AuthCodeTypeSms
//  + TL_AuthCodeTypeCall
//  + TL_AuthCodeTypeFlashCall
//

func (m *Auth_CodeType) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_auth_codeTypeSms:
		t := m.To_AuthCodeTypeSms()
		return t.Encode()
	case TLConstructor_CRC32_auth_codeTypeCall:
		t := m.To_AuthCodeTypeCall()
		return t.Encode()
	case TLConstructor_CRC32_auth_codeTypeFlashCall:
		t := m.To_AuthCodeTypeFlashCall()
		return t.Encode()

	default:
		return []byte{}
	}
}

// auth.codeTypeSms#72a3158c = auth.CodeType;
func (m *Auth_CodeType) To_AuthCodeTypeSms() *TLAuthCodeTypeSms {
	return &TLAuthCodeTypeSms{
		Data2: m.Data2,
	}
}

// auth.codeTypeCall#741cd3e3 = auth.CodeType;
func (m *Auth_CodeType) To_AuthCodeTypeCall() *TLAuthCodeTypeCall {
	return &TLAuthCodeTypeCall{
		Data2: m.Data2,
	}
}

// auth.codeTypeFlashCall#226ccefb = auth.CodeType;
func (m *Auth_CodeType) To_AuthCodeTypeFlashCall() *TLAuthCodeTypeFlashCall {
	return &TLAuthCodeTypeFlashCall{
		Data2: m.Data2,
	}
}

// auth.codeTypeSms#72a3158c = auth.CodeType;
func (m *TLAuthCodeTypeSms) To_Auth_CodeType() *Auth_CodeType {
	return &Auth_CodeType{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLAuthCodeTypeSms) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_auth_codeTypeSms))

	return x.buf
}

func (m *TLAuthCodeTypeSms) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// auth.codeTypeCall#741cd3e3 = auth.CodeType;
func (m *TLAuthCodeTypeCall) To_Auth_CodeType() *Auth_CodeType {
	return &Auth_CodeType{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLAuthCodeTypeCall) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_auth_codeTypeCall))

	return x.buf
}

func (m *TLAuthCodeTypeCall) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// auth.codeTypeFlashCall#226ccefb = auth.CodeType;
func (m *TLAuthCodeTypeFlashCall) To_Auth_CodeType() *Auth_CodeType {
	return &Auth_CodeType{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLAuthCodeTypeFlashCall) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_auth_codeTypeFlashCall))

	return x.buf
}

func (m *TLAuthCodeTypeFlashCall) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// PeerSettings <--
//  + TL_PeerSettings
//

func (m *PeerSettings) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_peerSettings:
		t := m.To_PeerSettings()
		return t.Encode()

	default:
		return []byte{}
	}
}

// peerSettings#818426cd flags:# report_spam:flags.0?true = PeerSettings;
func (m *PeerSettings) To_PeerSettings() *TLPeerSettings {
	return &TLPeerSettings{
		Data2: m.Data2,
	}
}

// peerSettings#818426cd flags:# report_spam:flags.0?true = PeerSettings;
func (m *TLPeerSettings) To_PeerSettings() *PeerSettings {
	return &PeerSettings{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPeerSettings) SetReportSpam(v bool) { m.Data2.ReportSpam = v }
func (m *TLPeerSettings) GetReportSpam() bool  { return m.Data2.ReportSpam }

func (m *TLPeerSettings) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_peerSettings))

	// flags
	var flags uint32 = 0
	if m.GetReportSpam() == true {
		flags |= 1 << 1
	}
	x.UInt(flags)

	return x.buf
}

func (m *TLPeerSettings) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m.Data2.ReportSpam = true
	}

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// StickerPack <--
//  + TL_StickerPack
//

func (m *StickerPack) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_stickerPack:
		t := m.To_StickerPack()
		return t.Encode()

	default:
		return []byte{}
	}
}

// stickerPack#12b299d4 emoticon:string documents:Vector<long> = StickerPack;
func (m *StickerPack) To_StickerPack() *TLStickerPack {
	return &TLStickerPack{
		Data2: m.Data2,
	}
}

// stickerPack#12b299d4 emoticon:string documents:Vector<long> = StickerPack;
func (m *TLStickerPack) To_StickerPack() *StickerPack {
	return &StickerPack{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLStickerPack) SetEmoticon(v string) { m.Data2.Emoticon = v }
func (m *TLStickerPack) GetEmoticon() string  { return m.Data2.Emoticon }

func (m *TLStickerPack) SetDocuments(v []int64) { m.Data2.Documents = v }
func (m *TLStickerPack) GetDocuments() []int64  { return m.Data2.Documents }

func (m *TLStickerPack) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_stickerPack))

	x.StringBytes(m.Data2.Emoticon)

	return x.buf
}

func (m *TLStickerPack) Decode(dbuf *DecodeBuf) error {
	m.Data2.Emoticon = x.StringBytes()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// MessageRange <--
//  + TL_MessageRange
//

func (m *MessageRange) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_messageRange:
		t := m.To_MessageRange()
		return t.Encode()

	default:
		return []byte{}
	}
}

// messageRange#ae30253 min_id:int max_id:int = MessageRange;
func (m *MessageRange) To_MessageRange() *TLMessageRange {
	return &TLMessageRange{
		Data2: m.Data2,
	}
}

// messageRange#ae30253 min_id:int max_id:int = MessageRange;
func (m *TLMessageRange) To_MessageRange() *MessageRange {
	return &MessageRange{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessageRange) SetMinId(v int32) { m.Data2.MinId = v }
func (m *TLMessageRange) GetMinId() int32  { return m.Data2.MinId }

func (m *TLMessageRange) SetMaxId(v int32) { m.Data2.MaxId = v }
func (m *TLMessageRange) GetMaxId() int32  { return m.Data2.MaxId }

func (m *TLMessageRange) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageRange))

	x.Int(m.Data2.MinId)
	x.Int(m.Data2.MaxId)

	return x.buf
}

func (m *TLMessageRange) Decode(dbuf *DecodeBuf) error {
	m.Data2.MinId = x.Int()
	m.Data2.MaxId = x.Int()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Invoice <--
//  + TL_Invoice
//

func (m *Invoice) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_invoice:
		t := m.To_Invoice()
		return t.Encode()

	default:
		return []byte{}
	}
}

// invoice#c30aa358 flags:# test:flags.0?true name_requested:flags.1?true phone_requested:flags.2?true email_requested:flags.3?true shipping_address_requested:flags.4?true flexible:flags.5?true currency:string prices:Vector<LabeledPrice> = Invoice;
func (m *Invoice) To_Invoice() *TLInvoice {
	return &TLInvoice{
		Data2: m.Data2,
	}
}

// invoice#c30aa358 flags:# test:flags.0?true name_requested:flags.1?true phone_requested:flags.2?true email_requested:flags.3?true shipping_address_requested:flags.4?true flexible:flags.5?true currency:string prices:Vector<LabeledPrice> = Invoice;
func (m *TLInvoice) To_Invoice() *Invoice {
	return &Invoice{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInvoice) SetTest(v bool) { m.Data2.Test = v }
func (m *TLInvoice) GetTest() bool  { return m.Data2.Test }

func (m *TLInvoice) SetNameRequested(v bool) { m.Data2.NameRequested = v }
func (m *TLInvoice) GetNameRequested() bool  { return m.Data2.NameRequested }

func (m *TLInvoice) SetPhoneRequested(v bool) { m.Data2.PhoneRequested = v }
func (m *TLInvoice) GetPhoneRequested() bool  { return m.Data2.PhoneRequested }

func (m *TLInvoice) SetEmailRequested(v bool) { m.Data2.EmailRequested = v }
func (m *TLInvoice) GetEmailRequested() bool  { return m.Data2.EmailRequested }

func (m *TLInvoice) SetShippingAddressRequested(v bool) { m.Data2.ShippingAddressRequested = v }
func (m *TLInvoice) GetShippingAddressRequested() bool  { return m.Data2.ShippingAddressRequested }

func (m *TLInvoice) SetFlexible(v bool) { m.Data2.Flexible = v }
func (m *TLInvoice) GetFlexible() bool  { return m.Data2.Flexible }

func (m *TLInvoice) SetCurrency(v string) { m.Data2.Currency = v }
func (m *TLInvoice) GetCurrency() string  { return m.Data2.Currency }

func (m *TLInvoice) SetPrices(v []*LabeledPrice) { m.Data2.Prices = v }
func (m *TLInvoice) GetPrices() []*LabeledPrice  { return m.Data2.Prices }

func (m *TLInvoice) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_invoice))

	// flags
	var flags uint32 = 0
	if m.GetTest() == true {
		flags |= 1 << 1
	}
	if m.GetNameRequested() == true {
		flags |= 1 << 2
	}
	if m.GetPhoneRequested() == true {
		flags |= 1 << 3
	}
	if m.GetEmailRequested() == true {
		flags |= 1 << 4
	}
	if m.GetShippingAddressRequested() == true {
		flags |= 1 << 5
	}
	if m.GetFlexible() == true {
		flags |= 1 << 6
	}
	x.UInt(flags)

	x.StringBytes(m.Data2.Currency)

	return x.buf
}

func (m *TLInvoice) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m.Data2.Test = true
	}
	if (flags & (1 << 3)) != 0 {
		m.Data2.NameRequested = true
	}
	if (flags & (1 << 4)) != 0 {
		m.Data2.PhoneRequested = true
	}
	if (flags & (1 << 5)) != 0 {
		m.Data2.EmailRequested = true
	}
	if (flags & (1 << 6)) != 0 {
		m.Data2.ShippingAddressRequested = true
	}
	if (flags & (1 << 7)) != 0 {
		m.Data2.Flexible = true
	}
	m.Data2.Currency = x.StringBytes()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// StickerSetCovered <--
//  + TL_StickerSetCovered
//  + TL_StickerSetMultiCovered
//

func (m *StickerSetCovered) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_stickerSetCovered:
		t := m.To_StickerSetCovered()
		return t.Encode()
	case TLConstructor_CRC32_stickerSetMultiCovered:
		t := m.To_StickerSetMultiCovered()
		return t.Encode()

	default:
		return []byte{}
	}
}

// stickerSetCovered#6410a5d2 set:StickerSet cover:Document = StickerSetCovered;
func (m *StickerSetCovered) To_StickerSetCovered() *TLStickerSetCovered {
	return &TLStickerSetCovered{
		Data2: m.Data2,
	}
}

// stickerSetMultiCovered#3407e51b set:StickerSet covers:Vector<Document> = StickerSetCovered;
func (m *StickerSetCovered) To_StickerSetMultiCovered() *TLStickerSetMultiCovered {
	return &TLStickerSetMultiCovered{
		Data2: m.Data2,
	}
}

// stickerSetCovered#6410a5d2 set:StickerSet cover:Document = StickerSetCovered;
func (m *TLStickerSetCovered) To_StickerSetCovered() *StickerSetCovered {
	return &StickerSetCovered{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLStickerSetCovered) SetSet(v *StickerSet) { m.Data2.Set = v }
func (m *TLStickerSetCovered) GetSet() *StickerSet  { return m.Data2.Set }

func (m *TLStickerSetCovered) SetCover(v *Document) { m.Data2.Cover = v }
func (m *TLStickerSetCovered) GetCover() *Document  { return m.Data2.Cover }

func (m *TLStickerSetCovered) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_stickerSetCovered))

	x.Bytes(m.Data2Set.Encode())
	x.Bytes(m.Data2Cover.Encode())

	return x.buf
}

func (m *TLStickerSetCovered) Decode(dbuf *DecodeBuf) error {
	m1 := &Set{}
	m1.Decode(dbuf)
	m.SetSet(m1)
	m2 := &Cover{}
	m2.Decode(dbuf)
	m.SetCover(m2)

	return dbuf.err
}

// stickerSetMultiCovered#3407e51b set:StickerSet covers:Vector<Document> = StickerSetCovered;
func (m *TLStickerSetMultiCovered) To_StickerSetCovered() *StickerSetCovered {
	return &StickerSetCovered{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLStickerSetMultiCovered) SetSet(v *StickerSet) { m.Data2.Set = v }
func (m *TLStickerSetMultiCovered) GetSet() *StickerSet  { return m.Data2.Set }

func (m *TLStickerSetMultiCovered) SetCovers(v []*Document) { m.Data2.Covers = v }
func (m *TLStickerSetMultiCovered) GetCovers() []*Document  { return m.Data2.Covers }

func (m *TLStickerSetMultiCovered) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_stickerSetMultiCovered))

	x.Bytes(m.Data2Set.Encode())

	return x.buf
}

func (m *TLStickerSetMultiCovered) Decode(dbuf *DecodeBuf) error {
	m1 := &Set{}
	m1.Decode(dbuf)
	m.SetSet(m1)

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// PhoneCallDiscardReason <--
//  + TL_PhoneCallDiscardReasonMissed
//  + TL_PhoneCallDiscardReasonDisconnect
//  + TL_PhoneCallDiscardReasonHangup
//  + TL_PhoneCallDiscardReasonBusy
//

func (m *PhoneCallDiscardReason) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_phoneCallDiscardReasonMissed:
		t := m.To_PhoneCallDiscardReasonMissed()
		return t.Encode()
	case TLConstructor_CRC32_phoneCallDiscardReasonDisconnect:
		t := m.To_PhoneCallDiscardReasonDisconnect()
		return t.Encode()
	case TLConstructor_CRC32_phoneCallDiscardReasonHangup:
		t := m.To_PhoneCallDiscardReasonHangup()
		return t.Encode()
	case TLConstructor_CRC32_phoneCallDiscardReasonBusy:
		t := m.To_PhoneCallDiscardReasonBusy()
		return t.Encode()

	default:
		return []byte{}
	}
}

// phoneCallDiscardReasonMissed#85e42301 = PhoneCallDiscardReason;
func (m *PhoneCallDiscardReason) To_PhoneCallDiscardReasonMissed() *TLPhoneCallDiscardReasonMissed {
	return &TLPhoneCallDiscardReasonMissed{
		Data2: m.Data2,
	}
}

// phoneCallDiscardReasonDisconnect#e095c1a0 = PhoneCallDiscardReason;
func (m *PhoneCallDiscardReason) To_PhoneCallDiscardReasonDisconnect() *TLPhoneCallDiscardReasonDisconnect {
	return &TLPhoneCallDiscardReasonDisconnect{
		Data2: m.Data2,
	}
}

// phoneCallDiscardReasonHangup#57adc690 = PhoneCallDiscardReason;
func (m *PhoneCallDiscardReason) To_PhoneCallDiscardReasonHangup() *TLPhoneCallDiscardReasonHangup {
	return &TLPhoneCallDiscardReasonHangup{
		Data2: m.Data2,
	}
}

// phoneCallDiscardReasonBusy#faf7e8c9 = PhoneCallDiscardReason;
func (m *PhoneCallDiscardReason) To_PhoneCallDiscardReasonBusy() *TLPhoneCallDiscardReasonBusy {
	return &TLPhoneCallDiscardReasonBusy{
		Data2: m.Data2,
	}
}

// phoneCallDiscardReasonMissed#85e42301 = PhoneCallDiscardReason;
func (m *TLPhoneCallDiscardReasonMissed) To_PhoneCallDiscardReason() *PhoneCallDiscardReason {
	return &PhoneCallDiscardReason{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPhoneCallDiscardReasonMissed) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_phoneCallDiscardReasonMissed))

	return x.buf
}

func (m *TLPhoneCallDiscardReasonMissed) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// phoneCallDiscardReasonDisconnect#e095c1a0 = PhoneCallDiscardReason;
func (m *TLPhoneCallDiscardReasonDisconnect) To_PhoneCallDiscardReason() *PhoneCallDiscardReason {
	return &PhoneCallDiscardReason{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPhoneCallDiscardReasonDisconnect) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_phoneCallDiscardReasonDisconnect))

	return x.buf
}

func (m *TLPhoneCallDiscardReasonDisconnect) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// phoneCallDiscardReasonHangup#57adc690 = PhoneCallDiscardReason;
func (m *TLPhoneCallDiscardReasonHangup) To_PhoneCallDiscardReason() *PhoneCallDiscardReason {
	return &PhoneCallDiscardReason{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPhoneCallDiscardReasonHangup) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_phoneCallDiscardReasonHangup))

	return x.buf
}

func (m *TLPhoneCallDiscardReasonHangup) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// phoneCallDiscardReasonBusy#faf7e8c9 = PhoneCallDiscardReason;
func (m *TLPhoneCallDiscardReasonBusy) To_PhoneCallDiscardReason() *PhoneCallDiscardReason {
	return &PhoneCallDiscardReason{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPhoneCallDiscardReasonBusy) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_phoneCallDiscardReasonBusy))

	return x.buf
}

func (m *TLPhoneCallDiscardReasonBusy) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// DataJSON <--
//  + TL_DataJSON
//

func (m *DataJSON) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_dataJSON:
		t := m.To_DataJSON()
		return t.Encode()

	default:
		return []byte{}
	}
}

// dataJSON#7d748d04 data:string = DataJSON;
func (m *DataJSON) To_DataJSON() *TLDataJSON {
	return &TLDataJSON{
		Data2: m.Data2,
	}
}

// dataJSON#7d748d04 data:string = DataJSON;
func (m *TLDataJSON) To_DataJSON() *DataJSON {
	return &DataJSON{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLDataJSON) SetData(v string) { m.Data2.Data = v }
func (m *TLDataJSON) GetData() string  { return m.Data2.Data }

func (m *TLDataJSON) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_dataJSON))

	x.StringBytes(m.Data2.Data)

	return x.buf
}

func (m *TLDataJSON) Decode(dbuf *DecodeBuf) error {
	m.Data2.Data = x.StringBytes()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// InputWebDocument <--
//  + TL_InputWebDocument
//

func (m *InputWebDocument) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_inputWebDocument:
		t := m.To_InputWebDocument()
		return t.Encode()

	default:
		return []byte{}
	}
}

// inputWebDocument#9bed434d url:string size:int mime_type:string attributes:Vector<DocumentAttribute> = InputWebDocument;
func (m *InputWebDocument) To_InputWebDocument() *TLInputWebDocument {
	return &TLInputWebDocument{
		Data2: m.Data2,
	}
}

// inputWebDocument#9bed434d url:string size:int mime_type:string attributes:Vector<DocumentAttribute> = InputWebDocument;
func (m *TLInputWebDocument) To_InputWebDocument() *InputWebDocument {
	return &InputWebDocument{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputWebDocument) SetUrl(v string) { m.Data2.Url = v }
func (m *TLInputWebDocument) GetUrl() string  { return m.Data2.Url }

func (m *TLInputWebDocument) SetSize(v int32) { m.Data2.Size = v }
func (m *TLInputWebDocument) GetSize() int32  { return m.Data2.Size }

func (m *TLInputWebDocument) SetMimeType(v string) { m.Data2.MimeType = v }
func (m *TLInputWebDocument) GetMimeType() string  { return m.Data2.MimeType }

func (m *TLInputWebDocument) SetAttributes(v []*DocumentAttribute) { m.Data2.Attributes = v }
func (m *TLInputWebDocument) GetAttributes() []*DocumentAttribute  { return m.Data2.Attributes }

func (m *TLInputWebDocument) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputWebDocument))

	x.StringBytes(m.Data2.Url)
	x.Int(m.Data2.Size)
	x.StringBytes(m.Data2.MimeType)

	return x.buf
}

func (m *TLInputWebDocument) Decode(dbuf *DecodeBuf) error {
	m.Data2.Url = x.StringBytes()
	m.Data2.Size = x.Int()
	m.Data2.MimeType = x.StringBytes()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// EncryptedFile <--
//  + TL_EncryptedFileEmpty
//  + TL_EncryptedFile
//

func (m *EncryptedFile) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_encryptedFileEmpty:
		t := m.To_EncryptedFileEmpty()
		return t.Encode()
	case TLConstructor_CRC32_encryptedFile:
		t := m.To_EncryptedFile()
		return t.Encode()

	default:
		return []byte{}
	}
}

// encryptedFileEmpty#c21f497e = EncryptedFile;
func (m *EncryptedFile) To_EncryptedFileEmpty() *TLEncryptedFileEmpty {
	return &TLEncryptedFileEmpty{
		Data2: m.Data2,
	}
}

// encryptedFile#4a70994c id:long access_hash:long size:int dc_id:int key_fingerprint:int = EncryptedFile;
func (m *EncryptedFile) To_EncryptedFile() *TLEncryptedFile {
	return &TLEncryptedFile{
		Data2: m.Data2,
	}
}

// encryptedFileEmpty#c21f497e = EncryptedFile;
func (m *TLEncryptedFileEmpty) To_EncryptedFile() *EncryptedFile {
	return &EncryptedFile{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLEncryptedFileEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_encryptedFileEmpty))

	return x.buf
}

func (m *TLEncryptedFileEmpty) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// encryptedFile#4a70994c id:long access_hash:long size:int dc_id:int key_fingerprint:int = EncryptedFile;
func (m *TLEncryptedFile) To_EncryptedFile() *EncryptedFile {
	return &EncryptedFile{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLEncryptedFile) SetId(v int64) { m.Data2.Id = v }
func (m *TLEncryptedFile) GetId() int64  { return m.Data2.Id }

func (m *TLEncryptedFile) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLEncryptedFile) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLEncryptedFile) SetSize(v int32) { m.Data2.Size = v }
func (m *TLEncryptedFile) GetSize() int32  { return m.Data2.Size }

func (m *TLEncryptedFile) SetDcId(v int32) { m.Data2.DcId = v }
func (m *TLEncryptedFile) GetDcId() int32  { return m.Data2.DcId }

func (m *TLEncryptedFile) SetKeyFingerprint(v int32) { m.Data2.KeyFingerprint = v }
func (m *TLEncryptedFile) GetKeyFingerprint() int32  { return m.Data2.KeyFingerprint }

func (m *TLEncryptedFile) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_encryptedFile))

	x.Long(m.Data2.Id)
	x.Long(m.Data2.AccessHash)
	x.Int(m.Data2.Size)
	x.Int(m.Data2.DcId)
	x.Int(m.Data2.KeyFingerprint)

	return x.buf
}

func (m *TLEncryptedFile) Decode(dbuf *DecodeBuf) error {
	m.Data2.Id = x.Long()
	m.Data2.AccessHash = x.Long()
	m.Data2.Size = x.Int()
	m.Data2.DcId = x.Int()
	m.Data2.KeyFingerprint = x.Int()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// InputStickerSet <--
//  + TL_InputStickerSetEmpty
//  + TL_InputStickerSetID
//  + TL_InputStickerSetShortName
//

func (m *InputStickerSet) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_inputStickerSetEmpty:
		t := m.To_InputStickerSetEmpty()
		return t.Encode()
	case TLConstructor_CRC32_inputStickerSetID:
		t := m.To_InputStickerSetID()
		return t.Encode()
	case TLConstructor_CRC32_inputStickerSetShortName:
		t := m.To_InputStickerSetShortName()
		return t.Encode()

	default:
		return []byte{}
	}
}

// inputStickerSetEmpty#ffb62b95 = InputStickerSet;
func (m *InputStickerSet) To_InputStickerSetEmpty() *TLInputStickerSetEmpty {
	return &TLInputStickerSetEmpty{
		Data2: m.Data2,
	}
}

// inputStickerSetID#9de7a269 id:long access_hash:long = InputStickerSet;
func (m *InputStickerSet) To_InputStickerSetID() *TLInputStickerSetID {
	return &TLInputStickerSetID{
		Data2: m.Data2,
	}
}

// inputStickerSetShortName#861cc8a0 short_name:string = InputStickerSet;
func (m *InputStickerSet) To_InputStickerSetShortName() *TLInputStickerSetShortName {
	return &TLInputStickerSetShortName{
		Data2: m.Data2,
	}
}

// inputStickerSetEmpty#ffb62b95 = InputStickerSet;
func (m *TLInputStickerSetEmpty) To_InputStickerSet() *InputStickerSet {
	return &InputStickerSet{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputStickerSetEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputStickerSetEmpty))

	return x.buf
}

func (m *TLInputStickerSetEmpty) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// inputStickerSetID#9de7a269 id:long access_hash:long = InputStickerSet;
func (m *TLInputStickerSetID) To_InputStickerSet() *InputStickerSet {
	return &InputStickerSet{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputStickerSetID) SetId(v int64) { m.Data2.Id = v }
func (m *TLInputStickerSetID) GetId() int64  { return m.Data2.Id }

func (m *TLInputStickerSetID) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLInputStickerSetID) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLInputStickerSetID) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputStickerSetID))

	x.Long(m.Data2.Id)
	x.Long(m.Data2.AccessHash)

	return x.buf
}

func (m *TLInputStickerSetID) Decode(dbuf *DecodeBuf) error {
	m.Data2.Id = x.Long()
	m.Data2.AccessHash = x.Long()

	return dbuf.err
}

// inputStickerSetShortName#861cc8a0 short_name:string = InputStickerSet;
func (m *TLInputStickerSetShortName) To_InputStickerSet() *InputStickerSet {
	return &InputStickerSet{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputStickerSetShortName) SetShortName(v string) { m.Data2.ShortName = v }
func (m *TLInputStickerSetShortName) GetShortName() string  { return m.Data2.ShortName }

func (m *TLInputStickerSetShortName) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputStickerSetShortName))

	x.StringBytes(m.Data2.ShortName)

	return x.buf
}

func (m *TLInputStickerSetShortName) Decode(dbuf *DecodeBuf) error {
	m.Data2.ShortName = x.StringBytes()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// KeyboardButtonRow <--
//  + TL_KeyboardButtonRow
//

func (m *KeyboardButtonRow) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_keyboardButtonRow:
		t := m.To_KeyboardButtonRow()
		return t.Encode()

	default:
		return []byte{}
	}
}

// keyboardButtonRow#77608b83 buttons:Vector<KeyboardButton> = KeyboardButtonRow;
func (m *KeyboardButtonRow) To_KeyboardButtonRow() *TLKeyboardButtonRow {
	return &TLKeyboardButtonRow{
		Data2: m.Data2,
	}
}

// keyboardButtonRow#77608b83 buttons:Vector<KeyboardButton> = KeyboardButtonRow;
func (m *TLKeyboardButtonRow) To_KeyboardButtonRow() *KeyboardButtonRow {
	return &KeyboardButtonRow{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLKeyboardButtonRow) SetButtons(v []*KeyboardButton) { m.Data2.Buttons = v }
func (m *TLKeyboardButtonRow) GetButtons() []*KeyboardButton  { return m.Data2.Buttons }

func (m *TLKeyboardButtonRow) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_keyboardButtonRow))

	return x.buf
}

func (m *TLKeyboardButtonRow) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Channels_ChannelParticipant <--
//  + TL_ChannelsChannelParticipant
//

func (m *Channels_ChannelParticipant) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_channels_channelParticipant:
		t := m.To_ChannelsChannelParticipant()
		return t.Encode()

	default:
		return []byte{}
	}
}

// channels.channelParticipant#d0d9b163 participant:ChannelParticipant users:Vector<User> = channels.ChannelParticipant;
func (m *Channels_ChannelParticipant) To_ChannelsChannelParticipant() *TLChannelsChannelParticipant {
	return &TLChannelsChannelParticipant{
		Data2: m.Data2,
	}
}

// channels.channelParticipant#d0d9b163 participant:ChannelParticipant users:Vector<User> = channels.ChannelParticipant;
func (m *TLChannelsChannelParticipant) To_Channels_ChannelParticipant() *Channels_ChannelParticipant {
	return &Channels_ChannelParticipant{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLChannelsChannelParticipant) SetParticipant(v *ChannelParticipant) { m.Data2.Participant = v }
func (m *TLChannelsChannelParticipant) GetParticipant() *ChannelParticipant {
	return m.Data2.Participant
}

func (m *TLChannelsChannelParticipant) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLChannelsChannelParticipant) GetUsers() []*User  { return m.Data2.Users }

func (m *TLChannelsChannelParticipant) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channels_channelParticipant))

	x.Bytes(m.Data2Participant.Encode())

	return x.buf
}

func (m *TLChannelsChannelParticipant) Decode(dbuf *DecodeBuf) error {
	m1 := &Participant{}
	m1.Decode(dbuf)
	m.SetParticipant(m1)

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Account_TmpPassword <--
//  + TL_AccountTmpPassword
//

func (m *Account_TmpPassword) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_account_tmpPassword:
		t := m.To_AccountTmpPassword()
		return t.Encode()

	default:
		return []byte{}
	}
}

// account.tmpPassword#db64fd34 tmp_password:bytes valid_until:int = account.TmpPassword;
func (m *Account_TmpPassword) To_AccountTmpPassword() *TLAccountTmpPassword {
	return &TLAccountTmpPassword{
		Data2: m.Data2,
	}
}

// account.tmpPassword#db64fd34 tmp_password:bytes valid_until:int = account.TmpPassword;
func (m *TLAccountTmpPassword) To_Account_TmpPassword() *Account_TmpPassword {
	return &Account_TmpPassword{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLAccountTmpPassword) SetTmpPassword(v []byte) { m.Data2.TmpPassword = v }
func (m *TLAccountTmpPassword) GetTmpPassword() []byte  { return m.Data2.TmpPassword }

func (m *TLAccountTmpPassword) SetValidUntil(v int32) { m.Data2.ValidUntil = v }
func (m *TLAccountTmpPassword) GetValidUntil() int32  { return m.Data2.ValidUntil }

func (m *TLAccountTmpPassword) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_account_tmpPassword))

	x.StringBytes(m.Data2.TmpPassword)
	x.Int(m.Data2.ValidUntil)

	return x.buf
}

func (m *TLAccountTmpPassword) Decode(dbuf *DecodeBuf) error {
	m.Data2.TmpPassword = x.StringBytes()
	m.Data2.ValidUntil = x.Int()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// InputStickerSetItem <--
//  + TL_InputStickerSetItem
//

func (m *InputStickerSetItem) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_inputStickerSetItem:
		t := m.To_InputStickerSetItem()
		return t.Encode()

	default:
		return []byte{}
	}
}

// inputStickerSetItem#ffa0a496 flags:# document:InputDocument emoji:string mask_coords:flags.0?MaskCoords = InputStickerSetItem;
func (m *InputStickerSetItem) To_InputStickerSetItem() *TLInputStickerSetItem {
	return &TLInputStickerSetItem{
		Data2: m.Data2,
	}
}

// inputStickerSetItem#ffa0a496 flags:# document:InputDocument emoji:string mask_coords:flags.0?MaskCoords = InputStickerSetItem;
func (m *TLInputStickerSetItem) To_InputStickerSetItem() *InputStickerSetItem {
	return &InputStickerSetItem{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputStickerSetItem) SetDocument(v *InputDocument) { m.Data2.Document = v }
func (m *TLInputStickerSetItem) GetDocument() *InputDocument  { return m.Data2.Document }

func (m *TLInputStickerSetItem) SetEmoji(v string) { m.Data2.Emoji = v }
func (m *TLInputStickerSetItem) GetEmoji() string  { return m.Data2.Emoji }

func (m *TLInputStickerSetItem) SetMaskCoords(v *MaskCoords) { m.Data2.MaskCoords = v }
func (m *TLInputStickerSetItem) GetMaskCoords() *MaskCoords  { return m.Data2.MaskCoords }

func (m *TLInputStickerSetItem) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputStickerSetItem))

	// flags
	var flags uint32 = 0
	if m.GetMaskCoords() != nil {
		flags |= 1 << 3
	}
	x.UInt(flags)

	x.Bytes(m.Data2Document.Encode())
	x.StringBytes(m.Data2.Emoji)
	if m.GetMaskCoords() != 0 {
		x.Bytes(m.Data2MaskCoords.Encode())
	}

	return x.buf
}

func (m *TLInputStickerSetItem) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	m2 := &Document{}
	m2.Decode(dbuf)
	m.SetDocument(m2)
	m.Data2.Emoji = x.StringBytes()
	if (flags & (1 << 4)) != 0 {
		m4 := &MaskCoords{}
		m4.Decode(dbuf)
		m.SetMaskCoords(m4)
	}

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// LangPackString <--
//  + TL_LangPackString
//  + TL_LangPackStringPluralized
//  + TL_LangPackStringDeleted
//

func (m *LangPackString) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_langPackString:
		t := m.To_LangPackString()
		return t.Encode()
	case TLConstructor_CRC32_langPackStringPluralized:
		t := m.To_LangPackStringPluralized()
		return t.Encode()
	case TLConstructor_CRC32_langPackStringDeleted:
		t := m.To_LangPackStringDeleted()
		return t.Encode()

	default:
		return []byte{}
	}
}

// langPackString#cad181f6 key:string value:string = LangPackString;
func (m *LangPackString) To_LangPackString() *TLLangPackString {
	return &TLLangPackString{
		Data2: m.Data2,
	}
}

// langPackStringPluralized#6c47ac9f flags:# key:string zero_value:flags.0?string one_value:flags.1?string two_value:flags.2?string few_value:flags.3?string many_value:flags.4?string other_value:string = LangPackString;
func (m *LangPackString) To_LangPackStringPluralized() *TLLangPackStringPluralized {
	return &TLLangPackStringPluralized{
		Data2: m.Data2,
	}
}

// langPackStringDeleted#2979eeb2 key:string = LangPackString;
func (m *LangPackString) To_LangPackStringDeleted() *TLLangPackStringDeleted {
	return &TLLangPackStringDeleted{
		Data2: m.Data2,
	}
}

// langPackString#cad181f6 key:string value:string = LangPackString;
func (m *TLLangPackString) To_LangPackString() *LangPackString {
	return &LangPackString{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLLangPackString) SetKey(v string) { m.Data2.Key = v }
func (m *TLLangPackString) GetKey() string  { return m.Data2.Key }

func (m *TLLangPackString) SetValue(v string) { m.Data2.Value = v }
func (m *TLLangPackString) GetValue() string  { return m.Data2.Value }

func (m *TLLangPackString) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_langPackString))

	x.StringBytes(m.Data2.Key)
	x.StringBytes(m.Data2.Value)

	return x.buf
}

func (m *TLLangPackString) Decode(dbuf *DecodeBuf) error {
	m.Data2.Key = x.StringBytes()
	m.Data2.Value = x.StringBytes()

	return dbuf.err
}

// langPackStringPluralized#6c47ac9f flags:# key:string zero_value:flags.0?string one_value:flags.1?string two_value:flags.2?string few_value:flags.3?string many_value:flags.4?string other_value:string = LangPackString;
func (m *TLLangPackStringPluralized) To_LangPackString() *LangPackString {
	return &LangPackString{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLLangPackStringPluralized) SetKey(v string) { m.Data2.Key = v }
func (m *TLLangPackStringPluralized) GetKey() string  { return m.Data2.Key }

func (m *TLLangPackStringPluralized) SetZeroValue(v string) { m.Data2.ZeroValue = v }
func (m *TLLangPackStringPluralized) GetZeroValue() string  { return m.Data2.ZeroValue }

func (m *TLLangPackStringPluralized) SetOneValue(v string) { m.Data2.OneValue = v }
func (m *TLLangPackStringPluralized) GetOneValue() string  { return m.Data2.OneValue }

func (m *TLLangPackStringPluralized) SetTwoValue(v string) { m.Data2.TwoValue = v }
func (m *TLLangPackStringPluralized) GetTwoValue() string  { return m.Data2.TwoValue }

func (m *TLLangPackStringPluralized) SetFewValue(v string) { m.Data2.FewValue = v }
func (m *TLLangPackStringPluralized) GetFewValue() string  { return m.Data2.FewValue }

func (m *TLLangPackStringPluralized) SetManyValue(v string) { m.Data2.ManyValue = v }
func (m *TLLangPackStringPluralized) GetManyValue() string  { return m.Data2.ManyValue }

func (m *TLLangPackStringPluralized) SetOtherValue(v string) { m.Data2.OtherValue = v }
func (m *TLLangPackStringPluralized) GetOtherValue() string  { return m.Data2.OtherValue }

func (m *TLLangPackStringPluralized) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_langPackStringPluralized))

	// flags
	var flags uint32 = 0
	if m.GetZeroValue() != "" {
		flags |= 1 << 2
	}
	if m.GetOneValue() != "" {
		flags |= 1 << 3
	}
	if m.GetTwoValue() != "" {
		flags |= 1 << 4
	}
	if m.GetFewValue() != "" {
		flags |= 1 << 5
	}
	if m.GetManyValue() != "" {
		flags |= 1 << 6
	}
	x.UInt(flags)

	x.StringBytes(m.Data2.Key)
	if m.GetZeroValue() != 0 {
		x.StringBytes(m.Data2.ZeroValue)
	}
	if m.GetOneValue() != 0 {
		x.StringBytes(m.Data2.OneValue)
	}
	if m.GetTwoValue() != 0 {
		x.StringBytes(m.Data2.TwoValue)
	}
	if m.GetFewValue() != 0 {
		x.StringBytes(m.Data2.FewValue)
	}
	if m.GetManyValue() != 0 {
		x.StringBytes(m.Data2.ManyValue)
	}
	x.StringBytes(m.Data2.OtherValue)

	return x.buf
}

func (m *TLLangPackStringPluralized) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	m.Data2.Key = x.StringBytes()
	if (flags & (1 << 3)) != 0 {
		m.Data2.ZeroValue = x.StringBytes()
	}
	if (flags & (1 << 4)) != 0 {
		m.Data2.OneValue = x.StringBytes()
	}
	if (flags & (1 << 5)) != 0 {
		m.Data2.TwoValue = x.StringBytes()
	}
	if (flags & (1 << 6)) != 0 {
		m.Data2.FewValue = x.StringBytes()
	}
	if (flags & (1 << 7)) != 0 {
		m.Data2.ManyValue = x.StringBytes()
	}
	m.Data2.OtherValue = x.StringBytes()

	return dbuf.err
}

// langPackStringDeleted#2979eeb2 key:string = LangPackString;
func (m *TLLangPackStringDeleted) To_LangPackString() *LangPackString {
	return &LangPackString{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLLangPackStringDeleted) SetKey(v string) { m.Data2.Key = v }
func (m *TLLangPackStringDeleted) GetKey() string  { return m.Data2.Key }

func (m *TLLangPackStringDeleted) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_langPackStringDeleted))

	x.StringBytes(m.Data2.Key)

	return x.buf
}

func (m *TLLangPackStringDeleted) Decode(dbuf *DecodeBuf) error {
	m.Data2.Key = x.StringBytes()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// PopularContact <--
//  + TL_PopularContact
//

func (m *PopularContact) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_popularContact:
		t := m.To_PopularContact()
		return t.Encode()

	default:
		return []byte{}
	}
}

// popularContact#5ce14175 client_id:long importers:int = PopularContact;
func (m *PopularContact) To_PopularContact() *TLPopularContact {
	return &TLPopularContact{
		Data2: m.Data2,
	}
}

// popularContact#5ce14175 client_id:long importers:int = PopularContact;
func (m *TLPopularContact) To_PopularContact() *PopularContact {
	return &PopularContact{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPopularContact) SetClientId(v int64) { m.Data2.ClientId = v }
func (m *TLPopularContact) GetClientId() int64  { return m.Data2.ClientId }

func (m *TLPopularContact) SetImporters(v int32) { m.Data2.Importers = v }
func (m *TLPopularContact) GetImporters() int32  { return m.Data2.Importers }

func (m *TLPopularContact) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_popularContact))

	x.Long(m.Data2.ClientId)
	x.Int(m.Data2.Importers)

	return x.buf
}

func (m *TLPopularContact) Decode(dbuf *DecodeBuf) error {
	m.Data2.ClientId = x.Long()
	m.Data2.Importers = x.Int()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Messages_Dialogs <--
//  + TL_MessagesDialogs
//  + TL_MessagesDialogsSlice
//

func (m *Messages_Dialogs) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_messages_dialogs:
		t := m.To_MessagesDialogs()
		return t.Encode()
	case TLConstructor_CRC32_messages_dialogsSlice:
		t := m.To_MessagesDialogsSlice()
		return t.Encode()

	default:
		return []byte{}
	}
}

// messages.dialogs#15ba6c40 dialogs:Vector<Dialog> messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = messages.Dialogs;
func (m *Messages_Dialogs) To_MessagesDialogs() *TLMessagesDialogs {
	return &TLMessagesDialogs{
		Data2: m.Data2,
	}
}

// messages.dialogsSlice#71e094f3 count:int dialogs:Vector<Dialog> messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = messages.Dialogs;
func (m *Messages_Dialogs) To_MessagesDialogsSlice() *TLMessagesDialogsSlice {
	return &TLMessagesDialogsSlice{
		Data2: m.Data2,
	}
}

// messages.dialogs#15ba6c40 dialogs:Vector<Dialog> messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = messages.Dialogs;
func (m *TLMessagesDialogs) To_Messages_Dialogs() *Messages_Dialogs {
	return &Messages_Dialogs{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessagesDialogs) SetDialogs(v []*Dialog) { m.Data2.Dialogs = v }
func (m *TLMessagesDialogs) GetDialogs() []*Dialog  { return m.Data2.Dialogs }

func (m *TLMessagesDialogs) SetMessages(v []*Message) { m.Data2.Messages = v }
func (m *TLMessagesDialogs) GetMessages() []*Message  { return m.Data2.Messages }

func (m *TLMessagesDialogs) SetChats(v []*Chat) { m.Data2.Chats = v }
func (m *TLMessagesDialogs) GetChats() []*Chat  { return m.Data2.Chats }

func (m *TLMessagesDialogs) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLMessagesDialogs) GetUsers() []*User  { return m.Data2.Users }

func (m *TLMessagesDialogs) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_dialogs))

	return x.buf
}

func (m *TLMessagesDialogs) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// messages.dialogsSlice#71e094f3 count:int dialogs:Vector<Dialog> messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = messages.Dialogs;
func (m *TLMessagesDialogsSlice) To_Messages_Dialogs() *Messages_Dialogs {
	return &Messages_Dialogs{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessagesDialogsSlice) SetCount(v int32) { m.Data2.Count = v }
func (m *TLMessagesDialogsSlice) GetCount() int32  { return m.Data2.Count }

func (m *TLMessagesDialogsSlice) SetDialogs(v []*Dialog) { m.Data2.Dialogs = v }
func (m *TLMessagesDialogsSlice) GetDialogs() []*Dialog  { return m.Data2.Dialogs }

func (m *TLMessagesDialogsSlice) SetMessages(v []*Message) { m.Data2.Messages = v }
func (m *TLMessagesDialogsSlice) GetMessages() []*Message  { return m.Data2.Messages }

func (m *TLMessagesDialogsSlice) SetChats(v []*Chat) { m.Data2.Chats = v }
func (m *TLMessagesDialogsSlice) GetChats() []*Chat  { return m.Data2.Chats }

func (m *TLMessagesDialogsSlice) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLMessagesDialogsSlice) GetUsers() []*User  { return m.Data2.Users }

func (m *TLMessagesDialogsSlice) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_dialogsSlice))

	x.Int(m.Data2.Count)

	return x.buf
}

func (m *TLMessagesDialogsSlice) Decode(dbuf *DecodeBuf) error {
	m.Data2.Count = x.Int()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Messages_Messages <--
//  + TL_MessagesMessages
//  + TL_MessagesMessagesSlice
//  + TL_MessagesChannelMessages
//

func (m *Messages_Messages) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_messages_messages:
		t := m.To_MessagesMessages()
		return t.Encode()
	case TLConstructor_CRC32_messages_messagesSlice:
		t := m.To_MessagesMessagesSlice()
		return t.Encode()
	case TLConstructor_CRC32_messages_channelMessages:
		t := m.To_MessagesChannelMessages()
		return t.Encode()

	default:
		return []byte{}
	}
}

// messages.messages#8c718e87 messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = messages.Messages;
func (m *Messages_Messages) To_MessagesMessages() *TLMessagesMessages {
	return &TLMessagesMessages{
		Data2: m.Data2,
	}
}

// messages.messagesSlice#b446ae3 count:int messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = messages.Messages;
func (m *Messages_Messages) To_MessagesMessagesSlice() *TLMessagesMessagesSlice {
	return &TLMessagesMessagesSlice{
		Data2: m.Data2,
	}
}

// messages.channelMessages#99262e37 flags:# pts:int count:int messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = messages.Messages;
func (m *Messages_Messages) To_MessagesChannelMessages() *TLMessagesChannelMessages {
	return &TLMessagesChannelMessages{
		Data2: m.Data2,
	}
}

// messages.messages#8c718e87 messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = messages.Messages;
func (m *TLMessagesMessages) To_Messages_Messages() *Messages_Messages {
	return &Messages_Messages{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessagesMessages) SetMessages(v []*Message) { m.Data2.Messages = v }
func (m *TLMessagesMessages) GetMessages() []*Message  { return m.Data2.Messages }

func (m *TLMessagesMessages) SetChats(v []*Chat) { m.Data2.Chats = v }
func (m *TLMessagesMessages) GetChats() []*Chat  { return m.Data2.Chats }

func (m *TLMessagesMessages) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLMessagesMessages) GetUsers() []*User  { return m.Data2.Users }

func (m *TLMessagesMessages) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_messages))

	return x.buf
}

func (m *TLMessagesMessages) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// messages.messagesSlice#b446ae3 count:int messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = messages.Messages;
func (m *TLMessagesMessagesSlice) To_Messages_Messages() *Messages_Messages {
	return &Messages_Messages{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessagesMessagesSlice) SetCount(v int32) { m.Data2.Count = v }
func (m *TLMessagesMessagesSlice) GetCount() int32  { return m.Data2.Count }

func (m *TLMessagesMessagesSlice) SetMessages(v []*Message) { m.Data2.Messages = v }
func (m *TLMessagesMessagesSlice) GetMessages() []*Message  { return m.Data2.Messages }

func (m *TLMessagesMessagesSlice) SetChats(v []*Chat) { m.Data2.Chats = v }
func (m *TLMessagesMessagesSlice) GetChats() []*Chat  { return m.Data2.Chats }

func (m *TLMessagesMessagesSlice) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLMessagesMessagesSlice) GetUsers() []*User  { return m.Data2.Users }

func (m *TLMessagesMessagesSlice) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_messagesSlice))

	x.Int(m.Data2.Count)

	return x.buf
}

func (m *TLMessagesMessagesSlice) Decode(dbuf *DecodeBuf) error {
	m.Data2.Count = x.Int()

	return dbuf.err
}

// messages.channelMessages#99262e37 flags:# pts:int count:int messages:Vector<Message> chats:Vector<Chat> users:Vector<User> = messages.Messages;
func (m *TLMessagesChannelMessages) To_Messages_Messages() *Messages_Messages {
	return &Messages_Messages{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessagesChannelMessages) SetPts(v int32) { m.Data2.Pts = v }
func (m *TLMessagesChannelMessages) GetPts() int32  { return m.Data2.Pts }

func (m *TLMessagesChannelMessages) SetCount(v int32) { m.Data2.Count = v }
func (m *TLMessagesChannelMessages) GetCount() int32  { return m.Data2.Count }

func (m *TLMessagesChannelMessages) SetMessages(v []*Message) { m.Data2.Messages = v }
func (m *TLMessagesChannelMessages) GetMessages() []*Message  { return m.Data2.Messages }

func (m *TLMessagesChannelMessages) SetChats(v []*Chat) { m.Data2.Chats = v }
func (m *TLMessagesChannelMessages) GetChats() []*Chat  { return m.Data2.Chats }

func (m *TLMessagesChannelMessages) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLMessagesChannelMessages) GetUsers() []*User  { return m.Data2.Users }

func (m *TLMessagesChannelMessages) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_channelMessages))

	// flags
	var flags uint32 = 0
	x.UInt(flags)

	x.Int(m.Data2.Pts)
	x.Int(m.Data2.Count)

	return x.buf
}

func (m *TLMessagesChannelMessages) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	m.Data2.Pts = x.Int()
	m.Data2.Count = x.Int()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// DcOption <--
//  + TL_DcOption
//

func (m *DcOption) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_dcOption:
		t := m.To_DcOption()
		return t.Encode()

	default:
		return []byte{}
	}
}

// dcOption#5d8c6cc flags:# ipv6:flags.0?true media_only:flags.1?true tcpo_only:flags.2?true cdn:flags.3?true static:flags.4?true id:int ip_address:string port:int = DcOption;
func (m *DcOption) To_DcOption() *TLDcOption {
	return &TLDcOption{
		Data2: m.Data2,
	}
}

// dcOption#5d8c6cc flags:# ipv6:flags.0?true media_only:flags.1?true tcpo_only:flags.2?true cdn:flags.3?true static:flags.4?true id:int ip_address:string port:int = DcOption;
func (m *TLDcOption) To_DcOption() *DcOption {
	return &DcOption{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLDcOption) SetIpv6(v bool) { m.Data2.Ipv6 = v }
func (m *TLDcOption) GetIpv6() bool  { return m.Data2.Ipv6 }

func (m *TLDcOption) SetMediaOnly(v bool) { m.Data2.MediaOnly = v }
func (m *TLDcOption) GetMediaOnly() bool  { return m.Data2.MediaOnly }

func (m *TLDcOption) SetTcpoOnly(v bool) { m.Data2.TcpoOnly = v }
func (m *TLDcOption) GetTcpoOnly() bool  { return m.Data2.TcpoOnly }

func (m *TLDcOption) SetCdn(v bool) { m.Data2.Cdn = v }
func (m *TLDcOption) GetCdn() bool  { return m.Data2.Cdn }

func (m *TLDcOption) SetStatic(v bool) { m.Data2.Static = v }
func (m *TLDcOption) GetStatic() bool  { return m.Data2.Static }

func (m *TLDcOption) SetId(v int32) { m.Data2.Id = v }
func (m *TLDcOption) GetId() int32  { return m.Data2.Id }

func (m *TLDcOption) SetIpAddress(v string) { m.Data2.IpAddress = v }
func (m *TLDcOption) GetIpAddress() string  { return m.Data2.IpAddress }

func (m *TLDcOption) SetPort(v int32) { m.Data2.Port = v }
func (m *TLDcOption) GetPort() int32  { return m.Data2.Port }

func (m *TLDcOption) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_dcOption))

	// flags
	var flags uint32 = 0
	if m.GetIpv6() == true {
		flags |= 1 << 1
	}
	if m.GetMediaOnly() == true {
		flags |= 1 << 2
	}
	if m.GetTcpoOnly() == true {
		flags |= 1 << 3
	}
	if m.GetCdn() == true {
		flags |= 1 << 4
	}
	if m.GetStatic() == true {
		flags |= 1 << 5
	}
	x.UInt(flags)

	x.Int(m.Data2.Id)
	x.StringBytes(m.Data2.IpAddress)
	x.Int(m.Data2.Port)

	return x.buf
}

func (m *TLDcOption) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m.Data2.Ipv6 = true
	}
	if (flags & (1 << 3)) != 0 {
		m.Data2.MediaOnly = true
	}
	if (flags & (1 << 4)) != 0 {
		m.Data2.TcpoOnly = true
	}
	if (flags & (1 << 5)) != 0 {
		m.Data2.Cdn = true
	}
	if (flags & (1 << 6)) != 0 {
		m.Data2.Static = true
	}
	m.Data2.Id = x.Int()
	m.Data2.IpAddress = x.StringBytes()
	m.Data2.Port = x.Int()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// ChannelAdminLogEventAction <--
//  + TL_ChannelAdminLogEventActionChangeTitle
//  + TL_ChannelAdminLogEventActionChangeAbout
//  + TL_ChannelAdminLogEventActionChangeUsername
//  + TL_ChannelAdminLogEventActionChangePhoto
//  + TL_ChannelAdminLogEventActionToggleInvites
//  + TL_ChannelAdminLogEventActionToggleSignatures
//  + TL_ChannelAdminLogEventActionUpdatePinned
//  + TL_ChannelAdminLogEventActionEditMessage
//  + TL_ChannelAdminLogEventActionDeleteMessage
//  + TL_ChannelAdminLogEventActionParticipantJoin
//  + TL_ChannelAdminLogEventActionParticipantLeave
//  + TL_ChannelAdminLogEventActionParticipantInvite
//  + TL_ChannelAdminLogEventActionParticipantToggleBan
//  + TL_ChannelAdminLogEventActionParticipantToggleAdmin
//  + TL_ChannelAdminLogEventActionChangeStickerSet
//

func (m *ChannelAdminLogEventAction) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_channelAdminLogEventActionChangeTitle:
		t := m.To_ChannelAdminLogEventActionChangeTitle()
		return t.Encode()
	case TLConstructor_CRC32_channelAdminLogEventActionChangeAbout:
		t := m.To_ChannelAdminLogEventActionChangeAbout()
		return t.Encode()
	case TLConstructor_CRC32_channelAdminLogEventActionChangeUsername:
		t := m.To_ChannelAdminLogEventActionChangeUsername()
		return t.Encode()
	case TLConstructor_CRC32_channelAdminLogEventActionChangePhoto:
		t := m.To_ChannelAdminLogEventActionChangePhoto()
		return t.Encode()
	case TLConstructor_CRC32_channelAdminLogEventActionToggleInvites:
		t := m.To_ChannelAdminLogEventActionToggleInvites()
		return t.Encode()
	case TLConstructor_CRC32_channelAdminLogEventActionToggleSignatures:
		t := m.To_ChannelAdminLogEventActionToggleSignatures()
		return t.Encode()
	case TLConstructor_CRC32_channelAdminLogEventActionUpdatePinned:
		t := m.To_ChannelAdminLogEventActionUpdatePinned()
		return t.Encode()
	case TLConstructor_CRC32_channelAdminLogEventActionEditMessage:
		t := m.To_ChannelAdminLogEventActionEditMessage()
		return t.Encode()
	case TLConstructor_CRC32_channelAdminLogEventActionDeleteMessage:
		t := m.To_ChannelAdminLogEventActionDeleteMessage()
		return t.Encode()
	case TLConstructor_CRC32_channelAdminLogEventActionParticipantJoin:
		t := m.To_ChannelAdminLogEventActionParticipantJoin()
		return t.Encode()
	case TLConstructor_CRC32_channelAdminLogEventActionParticipantLeave:
		t := m.To_ChannelAdminLogEventActionParticipantLeave()
		return t.Encode()
	case TLConstructor_CRC32_channelAdminLogEventActionParticipantInvite:
		t := m.To_ChannelAdminLogEventActionParticipantInvite()
		return t.Encode()
	case TLConstructor_CRC32_channelAdminLogEventActionParticipantToggleBan:
		t := m.To_ChannelAdminLogEventActionParticipantToggleBan()
		return t.Encode()
	case TLConstructor_CRC32_channelAdminLogEventActionParticipantToggleAdmin:
		t := m.To_ChannelAdminLogEventActionParticipantToggleAdmin()
		return t.Encode()
	case TLConstructor_CRC32_channelAdminLogEventActionChangeStickerSet:
		t := m.To_ChannelAdminLogEventActionChangeStickerSet()
		return t.Encode()

	default:
		return []byte{}
	}
}

// channelAdminLogEventActionChangeTitle#e6dfb825 prev_value:string new_value:string = ChannelAdminLogEventAction;
func (m *ChannelAdminLogEventAction) To_ChannelAdminLogEventActionChangeTitle() *TLChannelAdminLogEventActionChangeTitle {
	return &TLChannelAdminLogEventActionChangeTitle{
		Data2: m.Data2,
	}
}

// channelAdminLogEventActionChangeAbout#55188a2e prev_value:string new_value:string = ChannelAdminLogEventAction;
func (m *ChannelAdminLogEventAction) To_ChannelAdminLogEventActionChangeAbout() *TLChannelAdminLogEventActionChangeAbout {
	return &TLChannelAdminLogEventActionChangeAbout{
		Data2: m.Data2,
	}
}

// channelAdminLogEventActionChangeUsername#6a4afc38 prev_value:string new_value:string = ChannelAdminLogEventAction;
func (m *ChannelAdminLogEventAction) To_ChannelAdminLogEventActionChangeUsername() *TLChannelAdminLogEventActionChangeUsername {
	return &TLChannelAdminLogEventActionChangeUsername{
		Data2: m.Data2,
	}
}

// channelAdminLogEventActionChangePhoto#b82f55c3 prev_photo:ChatPhoto new_photo:ChatPhoto = ChannelAdminLogEventAction;
func (m *ChannelAdminLogEventAction) To_ChannelAdminLogEventActionChangePhoto() *TLChannelAdminLogEventActionChangePhoto {
	return &TLChannelAdminLogEventActionChangePhoto{
		Data2: m.Data2,
	}
}

// channelAdminLogEventActionToggleInvites#1b7907ae new_value:Bool = ChannelAdminLogEventAction;
func (m *ChannelAdminLogEventAction) To_ChannelAdminLogEventActionToggleInvites() *TLChannelAdminLogEventActionToggleInvites {
	return &TLChannelAdminLogEventActionToggleInvites{
		Data2: m.Data2,
	}
}

// channelAdminLogEventActionToggleSignatures#26ae0971 new_value:Bool = ChannelAdminLogEventAction;
func (m *ChannelAdminLogEventAction) To_ChannelAdminLogEventActionToggleSignatures() *TLChannelAdminLogEventActionToggleSignatures {
	return &TLChannelAdminLogEventActionToggleSignatures{
		Data2: m.Data2,
	}
}

// channelAdminLogEventActionUpdatePinned#e9e82c18 message:Message = ChannelAdminLogEventAction;
func (m *ChannelAdminLogEventAction) To_ChannelAdminLogEventActionUpdatePinned() *TLChannelAdminLogEventActionUpdatePinned {
	return &TLChannelAdminLogEventActionUpdatePinned{
		Data2: m.Data2,
	}
}

// channelAdminLogEventActionEditMessage#709b2405 prev_message:Message new_message:Message = ChannelAdminLogEventAction;
func (m *ChannelAdminLogEventAction) To_ChannelAdminLogEventActionEditMessage() *TLChannelAdminLogEventActionEditMessage {
	return &TLChannelAdminLogEventActionEditMessage{
		Data2: m.Data2,
	}
}

// channelAdminLogEventActionDeleteMessage#42e047bb message:Message = ChannelAdminLogEventAction;
func (m *ChannelAdminLogEventAction) To_ChannelAdminLogEventActionDeleteMessage() *TLChannelAdminLogEventActionDeleteMessage {
	return &TLChannelAdminLogEventActionDeleteMessage{
		Data2: m.Data2,
	}
}

// channelAdminLogEventActionParticipantJoin#183040d3 = ChannelAdminLogEventAction;
func (m *ChannelAdminLogEventAction) To_ChannelAdminLogEventActionParticipantJoin() *TLChannelAdminLogEventActionParticipantJoin {
	return &TLChannelAdminLogEventActionParticipantJoin{
		Data2: m.Data2,
	}
}

// channelAdminLogEventActionParticipantLeave#f89777f2 = ChannelAdminLogEventAction;
func (m *ChannelAdminLogEventAction) To_ChannelAdminLogEventActionParticipantLeave() *TLChannelAdminLogEventActionParticipantLeave {
	return &TLChannelAdminLogEventActionParticipantLeave{
		Data2: m.Data2,
	}
}

// channelAdminLogEventActionParticipantInvite#e31c34d8 participant:ChannelParticipant = ChannelAdminLogEventAction;
func (m *ChannelAdminLogEventAction) To_ChannelAdminLogEventActionParticipantInvite() *TLChannelAdminLogEventActionParticipantInvite {
	return &TLChannelAdminLogEventActionParticipantInvite{
		Data2: m.Data2,
	}
}

// channelAdminLogEventActionParticipantToggleBan#e6d83d7e prev_participant:ChannelParticipant new_participant:ChannelParticipant = ChannelAdminLogEventAction;
func (m *ChannelAdminLogEventAction) To_ChannelAdminLogEventActionParticipantToggleBan() *TLChannelAdminLogEventActionParticipantToggleBan {
	return &TLChannelAdminLogEventActionParticipantToggleBan{
		Data2: m.Data2,
	}
}

// channelAdminLogEventActionParticipantToggleAdmin#d5676710 prev_participant:ChannelParticipant new_participant:ChannelParticipant = ChannelAdminLogEventAction;
func (m *ChannelAdminLogEventAction) To_ChannelAdminLogEventActionParticipantToggleAdmin() *TLChannelAdminLogEventActionParticipantToggleAdmin {
	return &TLChannelAdminLogEventActionParticipantToggleAdmin{
		Data2: m.Data2,
	}
}

// channelAdminLogEventActionChangeStickerSet#b1c3caa7 prev_stickerset:InputStickerSet new_stickerset:InputStickerSet = ChannelAdminLogEventAction;
func (m *ChannelAdminLogEventAction) To_ChannelAdminLogEventActionChangeStickerSet() *TLChannelAdminLogEventActionChangeStickerSet {
	return &TLChannelAdminLogEventActionChangeStickerSet{
		Data2: m.Data2,
	}
}

// channelAdminLogEventActionChangeTitle#e6dfb825 prev_value:string new_value:string = ChannelAdminLogEventAction;
func (m *TLChannelAdminLogEventActionChangeTitle) To_ChannelAdminLogEventAction() *ChannelAdminLogEventAction {
	return &ChannelAdminLogEventAction{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLChannelAdminLogEventActionChangeTitle) SetPrevValue(v string) { m.Data2.PrevValue = v }
func (m *TLChannelAdminLogEventActionChangeTitle) GetPrevValue() string  { return m.Data2.PrevValue }

func (m *TLChannelAdminLogEventActionChangeTitle) SetNewValue(v string) { m.Data2.NewValue_2 = v }
func (m *TLChannelAdminLogEventActionChangeTitle) GetNewValue() string  { return m.Data2.NewValue_2 }

func (m *TLChannelAdminLogEventActionChangeTitle) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelAdminLogEventActionChangeTitle))

	x.StringBytes(m.Data2.PrevValue)
	x.StringBytes(m.Data2.NewValue)

	return x.buf
}

func (m *TLChannelAdminLogEventActionChangeTitle) Decode(dbuf *DecodeBuf) error {
	m.Data2.PrevValue = x.StringBytes()
	m.Data2.NewValue = x.StringBytes()

	return dbuf.err
}

// channelAdminLogEventActionChangeAbout#55188a2e prev_value:string new_value:string = ChannelAdminLogEventAction;
func (m *TLChannelAdminLogEventActionChangeAbout) To_ChannelAdminLogEventAction() *ChannelAdminLogEventAction {
	return &ChannelAdminLogEventAction{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLChannelAdminLogEventActionChangeAbout) SetPrevValue(v string) { m.Data2.PrevValue = v }
func (m *TLChannelAdminLogEventActionChangeAbout) GetPrevValue() string  { return m.Data2.PrevValue }

func (m *TLChannelAdminLogEventActionChangeAbout) SetNewValue(v string) { m.Data2.NewValue_2 = v }
func (m *TLChannelAdminLogEventActionChangeAbout) GetNewValue() string  { return m.Data2.NewValue_2 }

func (m *TLChannelAdminLogEventActionChangeAbout) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelAdminLogEventActionChangeAbout))

	x.StringBytes(m.Data2.PrevValue)
	x.StringBytes(m.Data2.NewValue)

	return x.buf
}

func (m *TLChannelAdminLogEventActionChangeAbout) Decode(dbuf *DecodeBuf) error {
	m.Data2.PrevValue = x.StringBytes()
	m.Data2.NewValue = x.StringBytes()

	return dbuf.err
}

// channelAdminLogEventActionChangeUsername#6a4afc38 prev_value:string new_value:string = ChannelAdminLogEventAction;
func (m *TLChannelAdminLogEventActionChangeUsername) To_ChannelAdminLogEventAction() *ChannelAdminLogEventAction {
	return &ChannelAdminLogEventAction{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLChannelAdminLogEventActionChangeUsername) SetPrevValue(v string) { m.Data2.PrevValue = v }
func (m *TLChannelAdminLogEventActionChangeUsername) GetPrevValue() string  { return m.Data2.PrevValue }

func (m *TLChannelAdminLogEventActionChangeUsername) SetNewValue(v string) { m.Data2.NewValue_2 = v }
func (m *TLChannelAdminLogEventActionChangeUsername) GetNewValue() string  { return m.Data2.NewValue_2 }

func (m *TLChannelAdminLogEventActionChangeUsername) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelAdminLogEventActionChangeUsername))

	x.StringBytes(m.Data2.PrevValue)
	x.StringBytes(m.Data2.NewValue)

	return x.buf
}

func (m *TLChannelAdminLogEventActionChangeUsername) Decode(dbuf *DecodeBuf) error {
	m.Data2.PrevValue = x.StringBytes()
	m.Data2.NewValue = x.StringBytes()

	return dbuf.err
}

// channelAdminLogEventActionChangePhoto#b82f55c3 prev_photo:ChatPhoto new_photo:ChatPhoto = ChannelAdminLogEventAction;
func (m *TLChannelAdminLogEventActionChangePhoto) To_ChannelAdminLogEventAction() *ChannelAdminLogEventAction {
	return &ChannelAdminLogEventAction{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLChannelAdminLogEventActionChangePhoto) SetPrevPhoto(v *ChatPhoto) { m.Data2.PrevPhoto = v }
func (m *TLChannelAdminLogEventActionChangePhoto) GetPrevPhoto() *ChatPhoto  { return m.Data2.PrevPhoto }

func (m *TLChannelAdminLogEventActionChangePhoto) SetNewPhoto(v *ChatPhoto) { m.Data2.NewPhoto = v }
func (m *TLChannelAdminLogEventActionChangePhoto) GetNewPhoto() *ChatPhoto  { return m.Data2.NewPhoto }

func (m *TLChannelAdminLogEventActionChangePhoto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelAdminLogEventActionChangePhoto))

	x.Bytes(m.Data2PrevPhoto.Encode())
	x.Bytes(m.Data2NewPhoto.Encode())

	return x.buf
}

func (m *TLChannelAdminLogEventActionChangePhoto) Decode(dbuf *DecodeBuf) error {
	m1 := &PrevPhoto{}
	m1.Decode(dbuf)
	m.SetPrevPhoto(m1)
	m2 := &NewPhoto{}
	m2.Decode(dbuf)
	m.SetNewPhoto(m2)

	return dbuf.err
}

// channelAdminLogEventActionToggleInvites#1b7907ae new_value:Bool = ChannelAdminLogEventAction;
func (m *TLChannelAdminLogEventActionToggleInvites) To_ChannelAdminLogEventAction() *ChannelAdminLogEventAction {
	return &ChannelAdminLogEventAction{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLChannelAdminLogEventActionToggleInvites) SetNewValue(v *Bool) { m.Data2.NewValue_5 = v }
func (m *TLChannelAdminLogEventActionToggleInvites) GetNewValue() *Bool  { return m.Data2.NewValue_5 }

func (m *TLChannelAdminLogEventActionToggleInvites) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelAdminLogEventActionToggleInvites))

	x.Bytes(m.Data2NewValue.Encode())

	return x.buf
}

func (m *TLChannelAdminLogEventActionToggleInvites) Decode(dbuf *DecodeBuf) error {
	m1 := &NewValue{}
	m1.Decode(dbuf)
	m.SetNewValue(m1)

	return dbuf.err
}

// channelAdminLogEventActionToggleSignatures#26ae0971 new_value:Bool = ChannelAdminLogEventAction;
func (m *TLChannelAdminLogEventActionToggleSignatures) To_ChannelAdminLogEventAction() *ChannelAdminLogEventAction {
	return &ChannelAdminLogEventAction{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLChannelAdminLogEventActionToggleSignatures) SetNewValue(v *Bool) { m.Data2.NewValue_5 = v }
func (m *TLChannelAdminLogEventActionToggleSignatures) GetNewValue() *Bool  { return m.Data2.NewValue_5 }

func (m *TLChannelAdminLogEventActionToggleSignatures) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelAdminLogEventActionToggleSignatures))

	x.Bytes(m.Data2NewValue.Encode())

	return x.buf
}

func (m *TLChannelAdminLogEventActionToggleSignatures) Decode(dbuf *DecodeBuf) error {
	m1 := &NewValue{}
	m1.Decode(dbuf)
	m.SetNewValue(m1)

	return dbuf.err
}

// channelAdminLogEventActionUpdatePinned#e9e82c18 message:Message = ChannelAdminLogEventAction;
func (m *TLChannelAdminLogEventActionUpdatePinned) To_ChannelAdminLogEventAction() *ChannelAdminLogEventAction {
	return &ChannelAdminLogEventAction{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLChannelAdminLogEventActionUpdatePinned) SetMessage(v *Message) { m.Data2.Message = v }
func (m *TLChannelAdminLogEventActionUpdatePinned) GetMessage() *Message  { return m.Data2.Message }

func (m *TLChannelAdminLogEventActionUpdatePinned) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelAdminLogEventActionUpdatePinned))

	x.Bytes(m.Data2Message.Encode())

	return x.buf
}

func (m *TLChannelAdminLogEventActionUpdatePinned) Decode(dbuf *DecodeBuf) error {
	m1 := &Message{}
	m1.Decode(dbuf)
	m.SetMessage(m1)

	return dbuf.err
}

// channelAdminLogEventActionEditMessage#709b2405 prev_message:Message new_message:Message = ChannelAdminLogEventAction;
func (m *TLChannelAdminLogEventActionEditMessage) To_ChannelAdminLogEventAction() *ChannelAdminLogEventAction {
	return &ChannelAdminLogEventAction{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLChannelAdminLogEventActionEditMessage) SetPrevMessage(v *Message) { m.Data2.PrevMessage = v }
func (m *TLChannelAdminLogEventActionEditMessage) GetPrevMessage() *Message {
	return m.Data2.PrevMessage
}

func (m *TLChannelAdminLogEventActionEditMessage) SetNewMessage(v *Message) { m.Data2.NewMessage = v }
func (m *TLChannelAdminLogEventActionEditMessage) GetNewMessage() *Message  { return m.Data2.NewMessage }

func (m *TLChannelAdminLogEventActionEditMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelAdminLogEventActionEditMessage))

	x.Bytes(m.Data2PrevMessage.Encode())
	x.Bytes(m.Data2NewMessage.Encode())

	return x.buf
}

func (m *TLChannelAdminLogEventActionEditMessage) Decode(dbuf *DecodeBuf) error {
	m1 := &PrevMessage{}
	m1.Decode(dbuf)
	m.SetPrevMessage(m1)
	m2 := &NewMessage{}
	m2.Decode(dbuf)
	m.SetNewMessage(m2)

	return dbuf.err
}

// channelAdminLogEventActionDeleteMessage#42e047bb message:Message = ChannelAdminLogEventAction;
func (m *TLChannelAdminLogEventActionDeleteMessage) To_ChannelAdminLogEventAction() *ChannelAdminLogEventAction {
	return &ChannelAdminLogEventAction{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLChannelAdminLogEventActionDeleteMessage) SetMessage(v *Message) { m.Data2.Message = v }
func (m *TLChannelAdminLogEventActionDeleteMessage) GetMessage() *Message  { return m.Data2.Message }

func (m *TLChannelAdminLogEventActionDeleteMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelAdminLogEventActionDeleteMessage))

	x.Bytes(m.Data2Message.Encode())

	return x.buf
}

func (m *TLChannelAdminLogEventActionDeleteMessage) Decode(dbuf *DecodeBuf) error {
	m1 := &Message{}
	m1.Decode(dbuf)
	m.SetMessage(m1)

	return dbuf.err
}

// channelAdminLogEventActionParticipantJoin#183040d3 = ChannelAdminLogEventAction;
func (m *TLChannelAdminLogEventActionParticipantJoin) To_ChannelAdminLogEventAction() *ChannelAdminLogEventAction {
	return &ChannelAdminLogEventAction{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLChannelAdminLogEventActionParticipantJoin) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelAdminLogEventActionParticipantJoin))

	return x.buf
}

func (m *TLChannelAdminLogEventActionParticipantJoin) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// channelAdminLogEventActionParticipantLeave#f89777f2 = ChannelAdminLogEventAction;
func (m *TLChannelAdminLogEventActionParticipantLeave) To_ChannelAdminLogEventAction() *ChannelAdminLogEventAction {
	return &ChannelAdminLogEventAction{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLChannelAdminLogEventActionParticipantLeave) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelAdminLogEventActionParticipantLeave))

	return x.buf
}

func (m *TLChannelAdminLogEventActionParticipantLeave) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// channelAdminLogEventActionParticipantInvite#e31c34d8 participant:ChannelParticipant = ChannelAdminLogEventAction;
func (m *TLChannelAdminLogEventActionParticipantInvite) To_ChannelAdminLogEventAction() *ChannelAdminLogEventAction {
	return &ChannelAdminLogEventAction{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLChannelAdminLogEventActionParticipantInvite) SetParticipant(v *ChannelParticipant) {
	m.Data2.Participant = v
}
func (m *TLChannelAdminLogEventActionParticipantInvite) GetParticipant() *ChannelParticipant {
	return m.Data2.Participant
}

func (m *TLChannelAdminLogEventActionParticipantInvite) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelAdminLogEventActionParticipantInvite))

	x.Bytes(m.Data2Participant.Encode())

	return x.buf
}

func (m *TLChannelAdminLogEventActionParticipantInvite) Decode(dbuf *DecodeBuf) error {
	m1 := &Participant{}
	m1.Decode(dbuf)
	m.SetParticipant(m1)

	return dbuf.err
}

// channelAdminLogEventActionParticipantToggleBan#e6d83d7e prev_participant:ChannelParticipant new_participant:ChannelParticipant = ChannelAdminLogEventAction;
func (m *TLChannelAdminLogEventActionParticipantToggleBan) To_ChannelAdminLogEventAction() *ChannelAdminLogEventAction {
	return &ChannelAdminLogEventAction{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLChannelAdminLogEventActionParticipantToggleBan) SetPrevParticipant(v *ChannelParticipant) {
	m.Data2.PrevParticipant = v
}
func (m *TLChannelAdminLogEventActionParticipantToggleBan) GetPrevParticipant() *ChannelParticipant {
	return m.Data2.PrevParticipant
}

func (m *TLChannelAdminLogEventActionParticipantToggleBan) SetNewParticipant(v *ChannelParticipant) {
	m.Data2.NewParticipant = v
}
func (m *TLChannelAdminLogEventActionParticipantToggleBan) GetNewParticipant() *ChannelParticipant {
	return m.Data2.NewParticipant
}

func (m *TLChannelAdminLogEventActionParticipantToggleBan) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelAdminLogEventActionParticipantToggleBan))

	x.Bytes(m.Data2PrevParticipant.Encode())
	x.Bytes(m.Data2NewParticipant.Encode())

	return x.buf
}

func (m *TLChannelAdminLogEventActionParticipantToggleBan) Decode(dbuf *DecodeBuf) error {
	m1 := &PrevParticipant{}
	m1.Decode(dbuf)
	m.SetPrevParticipant(m1)
	m2 := &NewParticipant{}
	m2.Decode(dbuf)
	m.SetNewParticipant(m2)

	return dbuf.err
}

// channelAdminLogEventActionParticipantToggleAdmin#d5676710 prev_participant:ChannelParticipant new_participant:ChannelParticipant = ChannelAdminLogEventAction;
func (m *TLChannelAdminLogEventActionParticipantToggleAdmin) To_ChannelAdminLogEventAction() *ChannelAdminLogEventAction {
	return &ChannelAdminLogEventAction{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLChannelAdminLogEventActionParticipantToggleAdmin) SetPrevParticipant(v *ChannelParticipant) {
	m.Data2.PrevParticipant = v
}
func (m *TLChannelAdminLogEventActionParticipantToggleAdmin) GetPrevParticipant() *ChannelParticipant {
	return m.Data2.PrevParticipant
}

func (m *TLChannelAdminLogEventActionParticipantToggleAdmin) SetNewParticipant(v *ChannelParticipant) {
	m.Data2.NewParticipant = v
}
func (m *TLChannelAdminLogEventActionParticipantToggleAdmin) GetNewParticipant() *ChannelParticipant {
	return m.Data2.NewParticipant
}

func (m *TLChannelAdminLogEventActionParticipantToggleAdmin) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelAdminLogEventActionParticipantToggleAdmin))

	x.Bytes(m.Data2PrevParticipant.Encode())
	x.Bytes(m.Data2NewParticipant.Encode())

	return x.buf
}

func (m *TLChannelAdminLogEventActionParticipantToggleAdmin) Decode(dbuf *DecodeBuf) error {
	m1 := &PrevParticipant{}
	m1.Decode(dbuf)
	m.SetPrevParticipant(m1)
	m2 := &NewParticipant{}
	m2.Decode(dbuf)
	m.SetNewParticipant(m2)

	return dbuf.err
}

// channelAdminLogEventActionChangeStickerSet#b1c3caa7 prev_stickerset:InputStickerSet new_stickerset:InputStickerSet = ChannelAdminLogEventAction;
func (m *TLChannelAdminLogEventActionChangeStickerSet) To_ChannelAdminLogEventAction() *ChannelAdminLogEventAction {
	return &ChannelAdminLogEventAction{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLChannelAdminLogEventActionChangeStickerSet) SetPrevStickerset(v *InputStickerSet) {
	m.Data2.PrevStickerset = v
}
func (m *TLChannelAdminLogEventActionChangeStickerSet) GetPrevStickerset() *InputStickerSet {
	return m.Data2.PrevStickerset
}

func (m *TLChannelAdminLogEventActionChangeStickerSet) SetNewStickerset(v *InputStickerSet) {
	m.Data2.NewStickerset = v
}
func (m *TLChannelAdminLogEventActionChangeStickerSet) GetNewStickerset() *InputStickerSet {
	return m.Data2.NewStickerset
}

func (m *TLChannelAdminLogEventActionChangeStickerSet) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelAdminLogEventActionChangeStickerSet))

	x.Bytes(m.Data2PrevStickerset.Encode())
	x.Bytes(m.Data2NewStickerset.Encode())

	return x.buf
}

func (m *TLChannelAdminLogEventActionChangeStickerSet) Decode(dbuf *DecodeBuf) error {
	m1 := &PrevStickerset{}
	m1.Decode(dbuf)
	m.SetPrevStickerset(m1)
	m2 := &NewStickerset{}
	m2.Decode(dbuf)
	m.SetNewStickerset(m2)

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Auth_Authorization <--
//  + TL_AuthAuthorization
//

func (m *Auth_Authorization) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_auth_authorization:
		t := m.To_AuthAuthorization()
		return t.Encode()

	default:
		return []byte{}
	}
}

// auth.authorization#cd050916 flags:# tmp_sessions:flags.0?int user:User = auth.Authorization;
func (m *Auth_Authorization) To_AuthAuthorization() *TLAuthAuthorization {
	return &TLAuthAuthorization{
		Data2: m.Data2,
	}
}

// auth.authorization#cd050916 flags:# tmp_sessions:flags.0?int user:User = auth.Authorization;
func (m *TLAuthAuthorization) To_Auth_Authorization() *Auth_Authorization {
	return &Auth_Authorization{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLAuthAuthorization) SetTmpSessions(v int32) { m.Data2.TmpSessions = v }
func (m *TLAuthAuthorization) GetTmpSessions() int32  { return m.Data2.TmpSessions }

func (m *TLAuthAuthorization) SetUser(v *User) { m.Data2.User = v }
func (m *TLAuthAuthorization) GetUser() *User  { return m.Data2.User }

func (m *TLAuthAuthorization) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_auth_authorization))

	// flags
	var flags uint32 = 0
	if m.GetTmpSessions() != 0 {
		flags |= 1 << 1
	}
	x.UInt(flags)

	if m.GetTmpSessions() != 0 {
		x.Int(m.Data2.TmpSessions)
	}
	x.Bytes(m.Data2User.Encode())

	return x.buf
}

func (m *TLAuthAuthorization) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m.Data2.TmpSessions = x.Int()
	}
	m3 := &User{}
	m3.Decode(dbuf)
	m.SetUser(m3)

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// PeerNotifySettings <--
//  + TL_PeerNotifySettingsEmpty
//  + TL_PeerNotifySettings
//

func (m *PeerNotifySettings) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_peerNotifySettingsEmpty:
		t := m.To_PeerNotifySettingsEmpty()
		return t.Encode()
	case TLConstructor_CRC32_peerNotifySettings:
		t := m.To_PeerNotifySettings()
		return t.Encode()

	default:
		return []byte{}
	}
}

// peerNotifySettingsEmpty#70a68512 = PeerNotifySettings;
func (m *PeerNotifySettings) To_PeerNotifySettingsEmpty() *TLPeerNotifySettingsEmpty {
	return &TLPeerNotifySettingsEmpty{
		Data2: m.Data2,
	}
}

// peerNotifySettings#9acda4c0 flags:# show_previews:flags.0?true silent:flags.1?true mute_until:int sound:string = PeerNotifySettings;
func (m *PeerNotifySettings) To_PeerNotifySettings() *TLPeerNotifySettings {
	return &TLPeerNotifySettings{
		Data2: m.Data2,
	}
}

// peerNotifySettingsEmpty#70a68512 = PeerNotifySettings;
func (m *TLPeerNotifySettingsEmpty) To_PeerNotifySettings() *PeerNotifySettings {
	return &PeerNotifySettings{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPeerNotifySettingsEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_peerNotifySettingsEmpty))

	return x.buf
}

func (m *TLPeerNotifySettingsEmpty) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// peerNotifySettings#9acda4c0 flags:# show_previews:flags.0?true silent:flags.1?true mute_until:int sound:string = PeerNotifySettings;
func (m *TLPeerNotifySettings) To_PeerNotifySettings() *PeerNotifySettings {
	return &PeerNotifySettings{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPeerNotifySettings) SetShowPreviews(v bool) { m.Data2.ShowPreviews = v }
func (m *TLPeerNotifySettings) GetShowPreviews() bool  { return m.Data2.ShowPreviews }

func (m *TLPeerNotifySettings) SetSilent(v bool) { m.Data2.Silent = v }
func (m *TLPeerNotifySettings) GetSilent() bool  { return m.Data2.Silent }

func (m *TLPeerNotifySettings) SetMuteUntil(v int32) { m.Data2.MuteUntil = v }
func (m *TLPeerNotifySettings) GetMuteUntil() int32  { return m.Data2.MuteUntil }

func (m *TLPeerNotifySettings) SetSound(v string) { m.Data2.Sound = v }
func (m *TLPeerNotifySettings) GetSound() string  { return m.Data2.Sound }

func (m *TLPeerNotifySettings) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_peerNotifySettings))

	// flags
	var flags uint32 = 0
	if m.GetShowPreviews() == true {
		flags |= 1 << 1
	}
	if m.GetSilent() == true {
		flags |= 1 << 2
	}
	x.UInt(flags)

	x.Int(m.Data2.MuteUntil)
	x.StringBytes(m.Data2.Sound)

	return x.buf
}

func (m *TLPeerNotifySettings) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m.Data2.ShowPreviews = true
	}
	if (flags & (1 << 3)) != 0 {
		m.Data2.Silent = true
	}
	m.Data2.MuteUntil = x.Int()
	m.Data2.Sound = x.StringBytes()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Help_InviteText <--
//  + TL_HelpInviteText
//

func (m *Help_InviteText) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_help_inviteText:
		t := m.To_HelpInviteText()
		return t.Encode()

	default:
		return []byte{}
	}
}

// help.inviteText#18cb9f78 message:string = help.InviteText;
func (m *Help_InviteText) To_HelpInviteText() *TLHelpInviteText {
	return &TLHelpInviteText{
		Data2: m.Data2,
	}
}

// help.inviteText#18cb9f78 message:string = help.InviteText;
func (m *TLHelpInviteText) To_Help_InviteText() *Help_InviteText {
	return &Help_InviteText{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLHelpInviteText) SetMessage(v string) { m.Data2.Message = v }
func (m *TLHelpInviteText) GetMessage() string  { return m.Data2.Message }

func (m *TLHelpInviteText) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_help_inviteText))

	x.StringBytes(m.Data2.Message)

	return x.buf
}

func (m *TLHelpInviteText) Decode(dbuf *DecodeBuf) error {
	m.Data2.Message = x.StringBytes()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// BotInlineMessage <--
//  + TL_BotInlineMessageMediaAuto
//  + TL_BotInlineMessageText
//  + TL_BotInlineMessageMediaGeo
//  + TL_BotInlineMessageMediaVenue
//  + TL_BotInlineMessageMediaContact
//

func (m *BotInlineMessage) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_botInlineMessageMediaAuto:
		t := m.To_BotInlineMessageMediaAuto()
		return t.Encode()
	case TLConstructor_CRC32_botInlineMessageText:
		t := m.To_BotInlineMessageText()
		return t.Encode()
	case TLConstructor_CRC32_botInlineMessageMediaGeo:
		t := m.To_BotInlineMessageMediaGeo()
		return t.Encode()
	case TLConstructor_CRC32_botInlineMessageMediaVenue:
		t := m.To_BotInlineMessageMediaVenue()
		return t.Encode()
	case TLConstructor_CRC32_botInlineMessageMediaContact:
		t := m.To_BotInlineMessageMediaContact()
		return t.Encode()

	default:
		return []byte{}
	}
}

// botInlineMessageMediaAuto#a74b15b flags:# caption:string reply_markup:flags.2?ReplyMarkup = BotInlineMessage;
func (m *BotInlineMessage) To_BotInlineMessageMediaAuto() *TLBotInlineMessageMediaAuto {
	return &TLBotInlineMessageMediaAuto{
		Data2: m.Data2,
	}
}

// botInlineMessageText#8c7f65e2 flags:# no_webpage:flags.0?true message:string entities:flags.1?Vector<MessageEntity> reply_markup:flags.2?ReplyMarkup = BotInlineMessage;
func (m *BotInlineMessage) To_BotInlineMessageText() *TLBotInlineMessageText {
	return &TLBotInlineMessageText{
		Data2: m.Data2,
	}
}

// botInlineMessageMediaGeo#3a8fd8b8 flags:# geo:GeoPoint reply_markup:flags.2?ReplyMarkup = BotInlineMessage;
func (m *BotInlineMessage) To_BotInlineMessageMediaGeo() *TLBotInlineMessageMediaGeo {
	return &TLBotInlineMessageMediaGeo{
		Data2: m.Data2,
	}
}

// botInlineMessageMediaVenue#4366232e flags:# geo:GeoPoint title:string address:string provider:string venue_id:string reply_markup:flags.2?ReplyMarkup = BotInlineMessage;
func (m *BotInlineMessage) To_BotInlineMessageMediaVenue() *TLBotInlineMessageMediaVenue {
	return &TLBotInlineMessageMediaVenue{
		Data2: m.Data2,
	}
}

// botInlineMessageMediaContact#35edb4d4 flags:# phone_number:string first_name:string last_name:string reply_markup:flags.2?ReplyMarkup = BotInlineMessage;
func (m *BotInlineMessage) To_BotInlineMessageMediaContact() *TLBotInlineMessageMediaContact {
	return &TLBotInlineMessageMediaContact{
		Data2: m.Data2,
	}
}

// botInlineMessageMediaAuto#a74b15b flags:# caption:string reply_markup:flags.2?ReplyMarkup = BotInlineMessage;
func (m *TLBotInlineMessageMediaAuto) To_BotInlineMessage() *BotInlineMessage {
	return &BotInlineMessage{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLBotInlineMessageMediaAuto) SetCaption(v string) { m.Data2.Caption = v }
func (m *TLBotInlineMessageMediaAuto) GetCaption() string  { return m.Data2.Caption }

func (m *TLBotInlineMessageMediaAuto) SetReplyMarkup(v *ReplyMarkup) { m.Data2.ReplyMarkup = v }
func (m *TLBotInlineMessageMediaAuto) GetReplyMarkup() *ReplyMarkup  { return m.Data2.ReplyMarkup }

func (m *TLBotInlineMessageMediaAuto) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_botInlineMessageMediaAuto))

	// flags
	var flags uint32 = 0
	if m.GetReplyMarkup() != nil {
		flags |= 1 << 2
	}
	x.UInt(flags)

	x.StringBytes(m.Data2.Caption)
	if m.GetReplyMarkup() != 0 {
		x.Bytes(m.Data2ReplyMarkup.Encode())
	}

	return x.buf
}

func (m *TLBotInlineMessageMediaAuto) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	m.Data2.Caption = x.StringBytes()
	if (flags & (1 << 3)) != 0 {
		m3 := &ReplyMarkup{}
		m3.Decode(dbuf)
		m.SetReplyMarkup(m3)
	}

	return dbuf.err
}

// botInlineMessageText#8c7f65e2 flags:# no_webpage:flags.0?true message:string entities:flags.1?Vector<MessageEntity> reply_markup:flags.2?ReplyMarkup = BotInlineMessage;
func (m *TLBotInlineMessageText) To_BotInlineMessage() *BotInlineMessage {
	return &BotInlineMessage{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLBotInlineMessageText) SetNoWebpage(v bool) { m.Data2.NoWebpage = v }
func (m *TLBotInlineMessageText) GetNoWebpage() bool  { return m.Data2.NoWebpage }

func (m *TLBotInlineMessageText) SetMessage(v string) { m.Data2.Message = v }
func (m *TLBotInlineMessageText) GetMessage() string  { return m.Data2.Message }

func (m *TLBotInlineMessageText) SetEntities(v []*MessageEntity) { m.Data2.Entities = v }
func (m *TLBotInlineMessageText) GetEntities() []*MessageEntity  { return m.Data2.Entities }

func (m *TLBotInlineMessageText) SetReplyMarkup(v *ReplyMarkup) { m.Data2.ReplyMarkup = v }
func (m *TLBotInlineMessageText) GetReplyMarkup() *ReplyMarkup  { return m.Data2.ReplyMarkup }

func (m *TLBotInlineMessageText) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_botInlineMessageText))

	// flags
	var flags uint32 = 0
	if m.GetNoWebpage() == true {
		flags |= 1 << 1
	}
	if m.GetEntities() != nil {
		flags |= 1 << 3
	}
	if m.GetReplyMarkup() != nil {
		flags |= 1 << 4
	}
	x.UInt(flags)

	x.StringBytes(m.Data2.Message)
	if m.Data2.Entities != 0 {

	}
	if m.GetReplyMarkup() != 0 {
		x.Bytes(m.Data2ReplyMarkup.Encode())
	}

	return x.buf
}

func (m *TLBotInlineMessageText) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m.Data2.NoWebpage = true
	}
	m.Data2.Message = x.StringBytes()
	if (flags & (1 << 4)) != 0 {

	}
	if (flags & (1 << 5)) != 0 {
		m5 := &ReplyMarkup{}
		m5.Decode(dbuf)
		m.SetReplyMarkup(m5)
	}

	return dbuf.err
}

// botInlineMessageMediaGeo#3a8fd8b8 flags:# geo:GeoPoint reply_markup:flags.2?ReplyMarkup = BotInlineMessage;
func (m *TLBotInlineMessageMediaGeo) To_BotInlineMessage() *BotInlineMessage {
	return &BotInlineMessage{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLBotInlineMessageMediaGeo) SetGeo(v *GeoPoint) { m.Data2.Geo = v }
func (m *TLBotInlineMessageMediaGeo) GetGeo() *GeoPoint  { return m.Data2.Geo }

func (m *TLBotInlineMessageMediaGeo) SetReplyMarkup(v *ReplyMarkup) { m.Data2.ReplyMarkup = v }
func (m *TLBotInlineMessageMediaGeo) GetReplyMarkup() *ReplyMarkup  { return m.Data2.ReplyMarkup }

func (m *TLBotInlineMessageMediaGeo) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_botInlineMessageMediaGeo))

	// flags
	var flags uint32 = 0
	if m.GetReplyMarkup() != nil {
		flags |= 1 << 2
	}
	x.UInt(flags)

	x.Bytes(m.Data2Geo.Encode())
	if m.GetReplyMarkup() != 0 {
		x.Bytes(m.Data2ReplyMarkup.Encode())
	}

	return x.buf
}

func (m *TLBotInlineMessageMediaGeo) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	m2 := &Geo{}
	m2.Decode(dbuf)
	m.SetGeo(m2)
	if (flags & (1 << 3)) != 0 {
		m3 := &ReplyMarkup{}
		m3.Decode(dbuf)
		m.SetReplyMarkup(m3)
	}

	return dbuf.err
}

// botInlineMessageMediaVenue#4366232e flags:# geo:GeoPoint title:string address:string provider:string venue_id:string reply_markup:flags.2?ReplyMarkup = BotInlineMessage;
func (m *TLBotInlineMessageMediaVenue) To_BotInlineMessage() *BotInlineMessage {
	return &BotInlineMessage{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLBotInlineMessageMediaVenue) SetGeo(v *GeoPoint) { m.Data2.Geo = v }
func (m *TLBotInlineMessageMediaVenue) GetGeo() *GeoPoint  { return m.Data2.Geo }

func (m *TLBotInlineMessageMediaVenue) SetTitle(v string) { m.Data2.Title = v }
func (m *TLBotInlineMessageMediaVenue) GetTitle() string  { return m.Data2.Title }

func (m *TLBotInlineMessageMediaVenue) SetAddress(v string) { m.Data2.Address = v }
func (m *TLBotInlineMessageMediaVenue) GetAddress() string  { return m.Data2.Address }

func (m *TLBotInlineMessageMediaVenue) SetProvider(v string) { m.Data2.Provider = v }
func (m *TLBotInlineMessageMediaVenue) GetProvider() string  { return m.Data2.Provider }

func (m *TLBotInlineMessageMediaVenue) SetVenueId(v string) { m.Data2.VenueId = v }
func (m *TLBotInlineMessageMediaVenue) GetVenueId() string  { return m.Data2.VenueId }

func (m *TLBotInlineMessageMediaVenue) SetReplyMarkup(v *ReplyMarkup) { m.Data2.ReplyMarkup = v }
func (m *TLBotInlineMessageMediaVenue) GetReplyMarkup() *ReplyMarkup  { return m.Data2.ReplyMarkup }

func (m *TLBotInlineMessageMediaVenue) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_botInlineMessageMediaVenue))

	// flags
	var flags uint32 = 0
	if m.GetReplyMarkup() != nil {
		flags |= 1 << 6
	}
	x.UInt(flags)

	x.Bytes(m.Data2Geo.Encode())
	x.StringBytes(m.Data2.Title)
	x.StringBytes(m.Data2.Address)
	x.StringBytes(m.Data2.Provider)
	x.StringBytes(m.Data2.VenueId)
	if m.GetReplyMarkup() != 0 {
		x.Bytes(m.Data2ReplyMarkup.Encode())
	}

	return x.buf
}

func (m *TLBotInlineMessageMediaVenue) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	m2 := &Geo{}
	m2.Decode(dbuf)
	m.SetGeo(m2)
	m.Data2.Title = x.StringBytes()
	m.Data2.Address = x.StringBytes()
	m.Data2.Provider = x.StringBytes()
	m.Data2.VenueId = x.StringBytes()
	if (flags & (1 << 7)) != 0 {
		m7 := &ReplyMarkup{}
		m7.Decode(dbuf)
		m.SetReplyMarkup(m7)
	}

	return dbuf.err
}

// botInlineMessageMediaContact#35edb4d4 flags:# phone_number:string first_name:string last_name:string reply_markup:flags.2?ReplyMarkup = BotInlineMessage;
func (m *TLBotInlineMessageMediaContact) To_BotInlineMessage() *BotInlineMessage {
	return &BotInlineMessage{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLBotInlineMessageMediaContact) SetPhoneNumber(v string) { m.Data2.PhoneNumber = v }
func (m *TLBotInlineMessageMediaContact) GetPhoneNumber() string  { return m.Data2.PhoneNumber }

func (m *TLBotInlineMessageMediaContact) SetFirstName(v string) { m.Data2.FirstName = v }
func (m *TLBotInlineMessageMediaContact) GetFirstName() string  { return m.Data2.FirstName }

func (m *TLBotInlineMessageMediaContact) SetLastName(v string) { m.Data2.LastName = v }
func (m *TLBotInlineMessageMediaContact) GetLastName() string  { return m.Data2.LastName }

func (m *TLBotInlineMessageMediaContact) SetReplyMarkup(v *ReplyMarkup) { m.Data2.ReplyMarkup = v }
func (m *TLBotInlineMessageMediaContact) GetReplyMarkup() *ReplyMarkup  { return m.Data2.ReplyMarkup }

func (m *TLBotInlineMessageMediaContact) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_botInlineMessageMediaContact))

	// flags
	var flags uint32 = 0
	if m.GetReplyMarkup() != nil {
		flags |= 1 << 4
	}
	x.UInt(flags)

	x.StringBytes(m.Data2.PhoneNumber)
	x.StringBytes(m.Data2.FirstName)
	x.StringBytes(m.Data2.LastName)
	if m.GetReplyMarkup() != 0 {
		x.Bytes(m.Data2ReplyMarkup.Encode())
	}

	return x.buf
}

func (m *TLBotInlineMessageMediaContact) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	m.Data2.PhoneNumber = x.StringBytes()
	m.Data2.FirstName = x.StringBytes()
	m.Data2.LastName = x.StringBytes()
	if (flags & (1 << 5)) != 0 {
		m5 := &ReplyMarkup{}
		m5.Decode(dbuf)
		m.SetReplyMarkup(m5)
	}

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Messages_StickerSet <--
//  + TL_MessagesStickerSet
//

func (m *Messages_StickerSet) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_messages_stickerSet:
		t := m.To_MessagesStickerSet()
		return t.Encode()

	default:
		return []byte{}
	}
}

// messages.stickerSet#b60a24a6 set:StickerSet packs:Vector<StickerPack> documents:Vector<Document> = messages.StickerSet;
func (m *Messages_StickerSet) To_MessagesStickerSet() *TLMessagesStickerSet {
	return &TLMessagesStickerSet{
		Data2: m.Data2,
	}
}

// messages.stickerSet#b60a24a6 set:StickerSet packs:Vector<StickerPack> documents:Vector<Document> = messages.StickerSet;
func (m *TLMessagesStickerSet) To_Messages_StickerSet() *Messages_StickerSet {
	return &Messages_StickerSet{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessagesStickerSet) SetSet(v *StickerSet) { m.Data2.Set = v }
func (m *TLMessagesStickerSet) GetSet() *StickerSet  { return m.Data2.Set }

func (m *TLMessagesStickerSet) SetPacks(v []*StickerPack) { m.Data2.Packs = v }
func (m *TLMessagesStickerSet) GetPacks() []*StickerPack  { return m.Data2.Packs }

func (m *TLMessagesStickerSet) SetDocuments(v []*Document) { m.Data2.Documents = v }
func (m *TLMessagesStickerSet) GetDocuments() []*Document  { return m.Data2.Documents }

func (m *TLMessagesStickerSet) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_stickerSet))

	x.Bytes(m.Data2Set.Encode())

	return x.buf
}

func (m *TLMessagesStickerSet) Decode(dbuf *DecodeBuf) error {
	m1 := &Set{}
	m1.Decode(dbuf)
	m.SetSet(m1)

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// PaymentCharge <--
//  + TL_PaymentCharge
//

func (m *PaymentCharge) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_paymentCharge:
		t := m.To_PaymentCharge()
		return t.Encode()

	default:
		return []byte{}
	}
}

// paymentCharge#ea02c27e id:string provider_charge_id:string = PaymentCharge;
func (m *PaymentCharge) To_PaymentCharge() *TLPaymentCharge {
	return &TLPaymentCharge{
		Data2: m.Data2,
	}
}

// paymentCharge#ea02c27e id:string provider_charge_id:string = PaymentCharge;
func (m *TLPaymentCharge) To_PaymentCharge() *PaymentCharge {
	return &PaymentCharge{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPaymentCharge) SetId(v string) { m.Data2.Id = v }
func (m *TLPaymentCharge) GetId() string  { return m.Data2.Id }

func (m *TLPaymentCharge) SetProviderChargeId(v string) { m.Data2.ProviderChargeId = v }
func (m *TLPaymentCharge) GetProviderChargeId() string  { return m.Data2.ProviderChargeId }

func (m *TLPaymentCharge) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_paymentCharge))

	x.StringBytes(m.Data2.Id)
	x.StringBytes(m.Data2.ProviderChargeId)

	return x.buf
}

func (m *TLPaymentCharge) Decode(dbuf *DecodeBuf) error {
	m.Data2.Id = x.StringBytes()
	m.Data2.ProviderChargeId = x.StringBytes()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// InputPhoneCall <--
//  + TL_InputPhoneCall
//

func (m *InputPhoneCall) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_inputPhoneCall:
		t := m.To_InputPhoneCall()
		return t.Encode()

	default:
		return []byte{}
	}
}

// inputPhoneCall#1e36fded id:long access_hash:long = InputPhoneCall;
func (m *InputPhoneCall) To_InputPhoneCall() *TLInputPhoneCall {
	return &TLInputPhoneCall{
		Data2: m.Data2,
	}
}

// inputPhoneCall#1e36fded id:long access_hash:long = InputPhoneCall;
func (m *TLInputPhoneCall) To_InputPhoneCall() *InputPhoneCall {
	return &InputPhoneCall{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputPhoneCall) SetId(v int64) { m.Data2.Id = v }
func (m *TLInputPhoneCall) GetId() int64  { return m.Data2.Id }

func (m *TLInputPhoneCall) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLInputPhoneCall) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLInputPhoneCall) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputPhoneCall))

	x.Long(m.Data2.Id)
	x.Long(m.Data2.AccessHash)

	return x.buf
}

func (m *TLInputPhoneCall) Decode(dbuf *DecodeBuf) error {
	m.Data2.Id = x.Long()
	m.Data2.AccessHash = x.Long()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// InputUser <--
//  + TL_InputUserEmpty
//  + TL_InputUserSelf
//  + TL_InputUser
//

func (m *InputUser) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_inputUserEmpty:
		t := m.To_InputUserEmpty()
		return t.Encode()
	case TLConstructor_CRC32_inputUserSelf:
		t := m.To_InputUserSelf()
		return t.Encode()
	case TLConstructor_CRC32_inputUser:
		t := m.To_InputUser()
		return t.Encode()

	default:
		return []byte{}
	}
}

// inputUserEmpty#b98886cf = InputUser;
func (m *InputUser) To_InputUserEmpty() *TLInputUserEmpty {
	return &TLInputUserEmpty{
		Data2: m.Data2,
	}
}

// inputUserSelf#f7c1b13f = InputUser;
func (m *InputUser) To_InputUserSelf() *TLInputUserSelf {
	return &TLInputUserSelf{
		Data2: m.Data2,
	}
}

// inputUser#d8292816 user_id:int access_hash:long = InputUser;
func (m *InputUser) To_InputUser() *TLInputUser {
	return &TLInputUser{
		Data2: m.Data2,
	}
}

// inputUserEmpty#b98886cf = InputUser;
func (m *TLInputUserEmpty) To_InputUser() *InputUser {
	return &InputUser{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputUserEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputUserEmpty))

	return x.buf
}

func (m *TLInputUserEmpty) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// inputUserSelf#f7c1b13f = InputUser;
func (m *TLInputUserSelf) To_InputUser() *InputUser {
	return &InputUser{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputUserSelf) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputUserSelf))

	return x.buf
}

func (m *TLInputUserSelf) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// inputUser#d8292816 user_id:int access_hash:long = InputUser;
func (m *TLInputUser) To_InputUser() *InputUser {
	return &InputUser{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputUser) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLInputUser) GetUserId() int32  { return m.Data2.UserId }

func (m *TLInputUser) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLInputUser) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLInputUser) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputUser))

	x.Int(m.Data2.UserId)
	x.Long(m.Data2.AccessHash)

	return x.buf
}

func (m *TLInputUser) Decode(dbuf *DecodeBuf) error {
	m.Data2.UserId = x.Int()
	m.Data2.AccessHash = x.Long()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Message <--
//  + TL_MessageEmpty
//  + TL_Message
//  + TL_MessageService
//

func (m *Message) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_messageEmpty:
		t := m.To_MessageEmpty()
		return t.Encode()
	case TLConstructor_CRC32_message:
		t := m.To_Message()
		return t.Encode()
	case TLConstructor_CRC32_messageService:
		t := m.To_MessageService()
		return t.Encode()

	default:
		return []byte{}
	}
}

// messageEmpty#83e5de54 id:int = Message;
func (m *Message) To_MessageEmpty() *TLMessageEmpty {
	return &TLMessageEmpty{
		Data2: m.Data2,
	}
}

// message#90dddc11 flags:# out:flags.1?true mentioned:flags.4?true media_unread:flags.5?true silent:flags.13?true post:flags.14?true id:int from_id:flags.8?int to_id:Peer fwd_from:flags.2?MessageFwdHeader via_bot_id:flags.11?int reply_to_msg_id:flags.3?int date:int message:string media:flags.9?MessageMedia reply_markup:flags.6?ReplyMarkup entities:flags.7?Vector<MessageEntity> views:flags.10?int edit_date:flags.15?int post_author:flags.16?string = Message;
func (m *Message) To_Message() *TLMessage {
	return &TLMessage{
		Data2: m.Data2,
	}
}

// messageService#9e19a1f6 flags:# out:flags.1?true mentioned:flags.4?true media_unread:flags.5?true silent:flags.13?true post:flags.14?true id:int from_id:flags.8?int to_id:Peer reply_to_msg_id:flags.3?int date:int action:MessageAction = Message;
func (m *Message) To_MessageService() *TLMessageService {
	return &TLMessageService{
		Data2: m.Data2,
	}
}

// messageEmpty#83e5de54 id:int = Message;
func (m *TLMessageEmpty) To_Message() *Message {
	return &Message{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessageEmpty) SetId(v int32) { m.Data2.Id = v }
func (m *TLMessageEmpty) GetId() int32  { return m.Data2.Id }

func (m *TLMessageEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageEmpty))

	x.Int(m.Data2.Id)

	return x.buf
}

func (m *TLMessageEmpty) Decode(dbuf *DecodeBuf) error {
	m.Data2.Id = x.Int()

	return dbuf.err
}

// message#90dddc11 flags:# out:flags.1?true mentioned:flags.4?true media_unread:flags.5?true silent:flags.13?true post:flags.14?true id:int from_id:flags.8?int to_id:Peer fwd_from:flags.2?MessageFwdHeader via_bot_id:flags.11?int reply_to_msg_id:flags.3?int date:int message:string media:flags.9?MessageMedia reply_markup:flags.6?ReplyMarkup entities:flags.7?Vector<MessageEntity> views:flags.10?int edit_date:flags.15?int post_author:flags.16?string = Message;
func (m *TLMessage) To_Message() *Message {
	return &Message{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessage) SetOut(v bool) { m.Data2.Out = v }
func (m *TLMessage) GetOut() bool  { return m.Data2.Out }

func (m *TLMessage) SetMentioned(v bool) { m.Data2.Mentioned = v }
func (m *TLMessage) GetMentioned() bool  { return m.Data2.Mentioned }

func (m *TLMessage) SetMediaUnread(v bool) { m.Data2.MediaUnread = v }
func (m *TLMessage) GetMediaUnread() bool  { return m.Data2.MediaUnread }

func (m *TLMessage) SetSilent(v bool) { m.Data2.Silent = v }
func (m *TLMessage) GetSilent() bool  { return m.Data2.Silent }

func (m *TLMessage) SetPost(v bool) { m.Data2.Post = v }
func (m *TLMessage) GetPost() bool  { return m.Data2.Post }

func (m *TLMessage) SetId(v int32) { m.Data2.Id = v }
func (m *TLMessage) GetId() int32  { return m.Data2.Id }

func (m *TLMessage) SetFromId(v int32) { m.Data2.FromId = v }
func (m *TLMessage) GetFromId() int32  { return m.Data2.FromId }

func (m *TLMessage) SetToId(v *Peer) { m.Data2.ToId = v }
func (m *TLMessage) GetToId() *Peer  { return m.Data2.ToId }

func (m *TLMessage) SetFwdFrom(v *MessageFwdHeader) { m.Data2.FwdFrom = v }
func (m *TLMessage) GetFwdFrom() *MessageFwdHeader  { return m.Data2.FwdFrom }

func (m *TLMessage) SetViaBotId(v int32) { m.Data2.ViaBotId = v }
func (m *TLMessage) GetViaBotId() int32  { return m.Data2.ViaBotId }

func (m *TLMessage) SetReplyToMsgId(v int32) { m.Data2.ReplyToMsgId = v }
func (m *TLMessage) GetReplyToMsgId() int32  { return m.Data2.ReplyToMsgId }

func (m *TLMessage) SetDate(v int32) { m.Data2.Date = v }
func (m *TLMessage) GetDate() int32  { return m.Data2.Date }

func (m *TLMessage) SetMessage(v string) { m.Data2.Message = v }
func (m *TLMessage) GetMessage() string  { return m.Data2.Message }

func (m *TLMessage) SetMedia(v *MessageMedia) { m.Data2.Media = v }
func (m *TLMessage) GetMedia() *MessageMedia  { return m.Data2.Media }

func (m *TLMessage) SetReplyMarkup(v *ReplyMarkup) { m.Data2.ReplyMarkup = v }
func (m *TLMessage) GetReplyMarkup() *ReplyMarkup  { return m.Data2.ReplyMarkup }

func (m *TLMessage) SetEntities(v []*MessageEntity) { m.Data2.Entities = v }
func (m *TLMessage) GetEntities() []*MessageEntity  { return m.Data2.Entities }

func (m *TLMessage) SetViews(v int32) { m.Data2.Views = v }
func (m *TLMessage) GetViews() int32  { return m.Data2.Views }

func (m *TLMessage) SetEditDate(v int32) { m.Data2.EditDate = v }
func (m *TLMessage) GetEditDate() int32  { return m.Data2.EditDate }

func (m *TLMessage) SetPostAuthor(v string) { m.Data2.PostAuthor = v }
func (m *TLMessage) GetPostAuthor() string  { return m.Data2.PostAuthor }

func (m *TLMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_message))

	// flags
	var flags uint32 = 0
	if m.GetOut() == true {
		flags |= 1 << 1
	}
	if m.GetMentioned() == true {
		flags |= 1 << 2
	}
	if m.GetMediaUnread() == true {
		flags |= 1 << 3
	}
	if m.GetSilent() == true {
		flags |= 1 << 4
	}
	if m.GetPost() == true {
		flags |= 1 << 5
	}
	if m.GetFromId() != 0 {
		flags |= 1 << 7
	}
	if m.GetFwdFrom() != nil {
		flags |= 1 << 9
	}
	if m.GetViaBotId() != 0 {
		flags |= 1 << 10
	}
	if m.GetReplyToMsgId() != 0 {
		flags |= 1 << 11
	}
	if m.GetMedia() != nil {
		flags |= 1 << 14
	}
	if m.GetReplyMarkup() != nil {
		flags |= 1 << 15
	}
	if m.GetEntities() != nil {
		flags |= 1 << 16
	}
	if m.GetViews() != 0 {
		flags |= 1 << 17
	}
	if m.GetEditDate() != 0 {
		flags |= 1 << 18
	}
	if m.GetPostAuthor() != "" {
		flags |= 1 << 19
	}
	x.UInt(flags)

	x.Int(m.Data2.Id)
	if m.GetFromId() != 0 {
		x.Int(m.Data2.FromId)
	}
	x.Bytes(m.Data2ToId.Encode())
	if m.GetFwdFrom() != 0 {
		x.Bytes(m.Data2FwdFrom.Encode())
	}
	if m.GetViaBotId() != 0 {
		x.Int(m.Data2.ViaBotId)
	}
	if m.GetReplyToMsgId() != 0 {
		x.Int(m.Data2.ReplyToMsgId)
	}
	x.Int(m.Data2.Date)
	x.StringBytes(m.Data2.Message)
	if m.GetMedia() != 0 {
		x.Bytes(m.Data2Media.Encode())
	}
	if m.GetReplyMarkup() != 0 {
		x.Bytes(m.Data2ReplyMarkup.Encode())
	}
	if m.Data2.Entities != 0 {

	}
	if m.GetViews() != 0 {
		x.Int(m.Data2.Views)
	}
	if m.GetEditDate() != 0 {
		x.Int(m.Data2.EditDate)
	}
	if m.GetPostAuthor() != 0 {
		x.StringBytes(m.Data2.PostAuthor)
	}

	return x.buf
}

func (m *TLMessage) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m.Data2.Out = true
	}
	if (flags & (1 << 3)) != 0 {
		m.Data2.Mentioned = true
	}
	if (flags & (1 << 4)) != 0 {
		m.Data2.MediaUnread = true
	}
	if (flags & (1 << 5)) != 0 {
		m.Data2.Silent = true
	}
	if (flags & (1 << 6)) != 0 {
		m.Data2.Post = true
	}
	m.Data2.Id = x.Int()
	if (flags & (1 << 8)) != 0 {
		m.Data2.FromId = x.Int()
	}
	m9 := &ToId{}
	m9.Decode(dbuf)
	m.SetToId(m9)
	if (flags & (1 << 10)) != 0 {
		m10 := &FwdFrom{}
		m10.Decode(dbuf)
		m.SetFwdFrom(m10)
	}
	if (flags & (1 << 11)) != 0 {
		m.Data2.ViaBotId = x.Int()
	}
	if (flags & (1 << 12)) != 0 {
		m.Data2.ReplyToMsgId = x.Int()
	}
	m.Data2.Date = x.Int()
	m.Data2.Message = x.StringBytes()
	if (flags & (1 << 15)) != 0 {
		m15 := &Media{}
		m15.Decode(dbuf)
		m.SetMedia(m15)
	}
	if (flags & (1 << 16)) != 0 {
		m16 := &ReplyMarkup{}
		m16.Decode(dbuf)
		m.SetReplyMarkup(m16)
	}
	if (flags & (1 << 17)) != 0 {

	}
	if (flags & (1 << 18)) != 0 {
		m.Data2.Views = x.Int()
	}
	if (flags & (1 << 19)) != 0 {
		m.Data2.EditDate = x.Int()
	}
	if (flags & (1 << 20)) != 0 {
		m.Data2.PostAuthor = x.StringBytes()
	}

	return dbuf.err
}

// messageService#9e19a1f6 flags:# out:flags.1?true mentioned:flags.4?true media_unread:flags.5?true silent:flags.13?true post:flags.14?true id:int from_id:flags.8?int to_id:Peer reply_to_msg_id:flags.3?int date:int action:MessageAction = Message;
func (m *TLMessageService) To_Message() *Message {
	return &Message{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessageService) SetOut(v bool) { m.Data2.Out = v }
func (m *TLMessageService) GetOut() bool  { return m.Data2.Out }

func (m *TLMessageService) SetMentioned(v bool) { m.Data2.Mentioned = v }
func (m *TLMessageService) GetMentioned() bool  { return m.Data2.Mentioned }

func (m *TLMessageService) SetMediaUnread(v bool) { m.Data2.MediaUnread = v }
func (m *TLMessageService) GetMediaUnread() bool  { return m.Data2.MediaUnread }

func (m *TLMessageService) SetSilent(v bool) { m.Data2.Silent = v }
func (m *TLMessageService) GetSilent() bool  { return m.Data2.Silent }

func (m *TLMessageService) SetPost(v bool) { m.Data2.Post = v }
func (m *TLMessageService) GetPost() bool  { return m.Data2.Post }

func (m *TLMessageService) SetId(v int32) { m.Data2.Id = v }
func (m *TLMessageService) GetId() int32  { return m.Data2.Id }

func (m *TLMessageService) SetFromId(v int32) { m.Data2.FromId = v }
func (m *TLMessageService) GetFromId() int32  { return m.Data2.FromId }

func (m *TLMessageService) SetToId(v *Peer) { m.Data2.ToId = v }
func (m *TLMessageService) GetToId() *Peer  { return m.Data2.ToId }

func (m *TLMessageService) SetReplyToMsgId(v int32) { m.Data2.ReplyToMsgId = v }
func (m *TLMessageService) GetReplyToMsgId() int32  { return m.Data2.ReplyToMsgId }

func (m *TLMessageService) SetDate(v int32) { m.Data2.Date = v }
func (m *TLMessageService) GetDate() int32  { return m.Data2.Date }

func (m *TLMessageService) SetAction(v *MessageAction) { m.Data2.Action = v }
func (m *TLMessageService) GetAction() *MessageAction  { return m.Data2.Action }

func (m *TLMessageService) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messageService))

	// flags
	var flags uint32 = 0
	if m.GetOut() == true {
		flags |= 1 << 1
	}
	if m.GetMentioned() == true {
		flags |= 1 << 2
	}
	if m.GetMediaUnread() == true {
		flags |= 1 << 3
	}
	if m.GetSilent() == true {
		flags |= 1 << 4
	}
	if m.GetPost() == true {
		flags |= 1 << 5
	}
	if m.GetFromId() != 0 {
		flags |= 1 << 7
	}
	if m.GetReplyToMsgId() != 0 {
		flags |= 1 << 9
	}
	x.UInt(flags)

	x.Int(m.Data2.Id)
	if m.GetFromId() != 0 {
		x.Int(m.Data2.FromId)
	}
	x.Bytes(m.Data2ToId.Encode())
	if m.GetReplyToMsgId() != 0 {
		x.Int(m.Data2.ReplyToMsgId)
	}
	x.Int(m.Data2.Date)
	x.Bytes(m.Data2Action.Encode())

	return x.buf
}

func (m *TLMessageService) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m.Data2.Out = true
	}
	if (flags & (1 << 3)) != 0 {
		m.Data2.Mentioned = true
	}
	if (flags & (1 << 4)) != 0 {
		m.Data2.MediaUnread = true
	}
	if (flags & (1 << 5)) != 0 {
		m.Data2.Silent = true
	}
	if (flags & (1 << 6)) != 0 {
		m.Data2.Post = true
	}
	m.Data2.Id = x.Int()
	if (flags & (1 << 8)) != 0 {
		m.Data2.FromId = x.Int()
	}
	m9 := &ToId{}
	m9.Decode(dbuf)
	m.SetToId(m9)
	if (flags & (1 << 10)) != 0 {
		m.Data2.ReplyToMsgId = x.Int()
	}
	m.Data2.Date = x.Int()
	m12 := &Action{}
	m12.Decode(dbuf)
	m.SetAction(m12)

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// GeoPoint <--
//  + TL_GeoPointEmpty
//  + TL_GeoPoint
//

func (m *GeoPoint) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_geoPointEmpty:
		t := m.To_GeoPointEmpty()
		return t.Encode()
	case TLConstructor_CRC32_geoPoint:
		t := m.To_GeoPoint()
		return t.Encode()

	default:
		return []byte{}
	}
}

// geoPointEmpty#1117dd5f = GeoPoint;
func (m *GeoPoint) To_GeoPointEmpty() *TLGeoPointEmpty {
	return &TLGeoPointEmpty{
		Data2: m.Data2,
	}
}

// geoPoint#2049d70c long:double lat:double = GeoPoint;
func (m *GeoPoint) To_GeoPoint() *TLGeoPoint {
	return &TLGeoPoint{
		Data2: m.Data2,
	}
}

// geoPointEmpty#1117dd5f = GeoPoint;
func (m *TLGeoPointEmpty) To_GeoPoint() *GeoPoint {
	return &GeoPoint{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLGeoPointEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_geoPointEmpty))

	return x.buf
}

func (m *TLGeoPointEmpty) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// geoPoint#2049d70c long:double lat:double = GeoPoint;
func (m *TLGeoPoint) To_GeoPoint() *GeoPoint {
	return &GeoPoint{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLGeoPoint) SetLong(v float64) { m.Data2.Long = v }
func (m *TLGeoPoint) GetLong() float64  { return m.Data2.Long }

func (m *TLGeoPoint) SetLat(v float64) { m.Data2.Lat = v }
func (m *TLGeoPoint) GetLat() float64  { return m.Data2.Lat }

func (m *TLGeoPoint) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_geoPoint))

	x.Double(m.Data2.Long)
	x.Double(m.Data2.Lat)

	return x.buf
}

func (m *TLGeoPoint) Decode(dbuf *DecodeBuf) error {
	m.Data2.Long = x.Double()
	m.Data2.Lat = x.Double()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// ContactBlocked <--
//  + TL_ContactBlocked
//

func (m *ContactBlocked) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_contactBlocked:
		t := m.To_ContactBlocked()
		return t.Encode()

	default:
		return []byte{}
	}
}

// contactBlocked#561bc879 user_id:int date:int = ContactBlocked;
func (m *ContactBlocked) To_ContactBlocked() *TLContactBlocked {
	return &TLContactBlocked{
		Data2: m.Data2,
	}
}

// contactBlocked#561bc879 user_id:int date:int = ContactBlocked;
func (m *TLContactBlocked) To_ContactBlocked() *ContactBlocked {
	return &ContactBlocked{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLContactBlocked) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLContactBlocked) GetUserId() int32  { return m.Data2.UserId }

func (m *TLContactBlocked) SetDate(v int32) { m.Data2.Date = v }
func (m *TLContactBlocked) GetDate() int32  { return m.Data2.Date }

func (m *TLContactBlocked) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_contactBlocked))

	x.Int(m.Data2.UserId)
	x.Int(m.Data2.Date)

	return x.buf
}

func (m *TLContactBlocked) Decode(dbuf *DecodeBuf) error {
	m.Data2.UserId = x.Int()
	m.Data2.Date = x.Int()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// CdnFileHash <--
//  + TL_CdnFileHash
//

func (m *CdnFileHash) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_cdnFileHash:
		t := m.To_CdnFileHash()
		return t.Encode()

	default:
		return []byte{}
	}
}

// cdnFileHash#77eec38f offset:int limit:int hash:bytes = CdnFileHash;
func (m *CdnFileHash) To_CdnFileHash() *TLCdnFileHash {
	return &TLCdnFileHash{
		Data2: m.Data2,
	}
}

// cdnFileHash#77eec38f offset:int limit:int hash:bytes = CdnFileHash;
func (m *TLCdnFileHash) To_CdnFileHash() *CdnFileHash {
	return &CdnFileHash{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLCdnFileHash) SetOffset(v int32) { m.Data2.Offset = v }
func (m *TLCdnFileHash) GetOffset() int32  { return m.Data2.Offset }

func (m *TLCdnFileHash) SetLimit(v int32) { m.Data2.Limit = v }
func (m *TLCdnFileHash) GetLimit() int32  { return m.Data2.Limit }

func (m *TLCdnFileHash) SetHash(v []byte) { m.Data2.Hash = v }
func (m *TLCdnFileHash) GetHash() []byte  { return m.Data2.Hash }

func (m *TLCdnFileHash) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_cdnFileHash))

	x.Int(m.Data2.Offset)
	x.Int(m.Data2.Limit)
	x.StringBytes(m.Data2.Hash)

	return x.buf
}

func (m *TLCdnFileHash) Decode(dbuf *DecodeBuf) error {
	m.Data2.Offset = x.Int()
	m.Data2.Limit = x.Int()
	m.Data2.Hash = x.StringBytes()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// InputPrivacyKey <--
//  + TL_InputPrivacyKeyStatusTimestamp
//  + TL_InputPrivacyKeyChatInvite
//  + TL_InputPrivacyKeyPhoneCall
//

func (m *InputPrivacyKey) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_inputPrivacyKeyStatusTimestamp:
		t := m.To_InputPrivacyKeyStatusTimestamp()
		return t.Encode()
	case TLConstructor_CRC32_inputPrivacyKeyChatInvite:
		t := m.To_InputPrivacyKeyChatInvite()
		return t.Encode()
	case TLConstructor_CRC32_inputPrivacyKeyPhoneCall:
		t := m.To_InputPrivacyKeyPhoneCall()
		return t.Encode()

	default:
		return []byte{}
	}
}

// inputPrivacyKeyStatusTimestamp#4f96cb18 = InputPrivacyKey;
func (m *InputPrivacyKey) To_InputPrivacyKeyStatusTimestamp() *TLInputPrivacyKeyStatusTimestamp {
	return &TLInputPrivacyKeyStatusTimestamp{
		Data2: m.Data2,
	}
}

// inputPrivacyKeyChatInvite#bdfb0426 = InputPrivacyKey;
func (m *InputPrivacyKey) To_InputPrivacyKeyChatInvite() *TLInputPrivacyKeyChatInvite {
	return &TLInputPrivacyKeyChatInvite{
		Data2: m.Data2,
	}
}

// inputPrivacyKeyPhoneCall#fabadc5f = InputPrivacyKey;
func (m *InputPrivacyKey) To_InputPrivacyKeyPhoneCall() *TLInputPrivacyKeyPhoneCall {
	return &TLInputPrivacyKeyPhoneCall{
		Data2: m.Data2,
	}
}

// inputPrivacyKeyStatusTimestamp#4f96cb18 = InputPrivacyKey;
func (m *TLInputPrivacyKeyStatusTimestamp) To_InputPrivacyKey() *InputPrivacyKey {
	return &InputPrivacyKey{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputPrivacyKeyStatusTimestamp) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputPrivacyKeyStatusTimestamp))

	return x.buf
}

func (m *TLInputPrivacyKeyStatusTimestamp) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// inputPrivacyKeyChatInvite#bdfb0426 = InputPrivacyKey;
func (m *TLInputPrivacyKeyChatInvite) To_InputPrivacyKey() *InputPrivacyKey {
	return &InputPrivacyKey{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputPrivacyKeyChatInvite) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputPrivacyKeyChatInvite))

	return x.buf
}

func (m *TLInputPrivacyKeyChatInvite) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// inputPrivacyKeyPhoneCall#fabadc5f = InputPrivacyKey;
func (m *TLInputPrivacyKeyPhoneCall) To_InputPrivacyKey() *InputPrivacyKey {
	return &InputPrivacyKey{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputPrivacyKeyPhoneCall) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputPrivacyKeyPhoneCall))

	return x.buf
}

func (m *TLInputPrivacyKeyPhoneCall) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// InputPrivacyRule <--
//  + TL_InputPrivacyValueAllowContacts
//  + TL_InputPrivacyValueAllowAll
//  + TL_InputPrivacyValueAllowUsers
//  + TL_InputPrivacyValueDisallowContacts
//  + TL_InputPrivacyValueDisallowAll
//  + TL_InputPrivacyValueDisallowUsers
//

func (m *InputPrivacyRule) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_inputPrivacyValueAllowContacts:
		t := m.To_InputPrivacyValueAllowContacts()
		return t.Encode()
	case TLConstructor_CRC32_inputPrivacyValueAllowAll:
		t := m.To_InputPrivacyValueAllowAll()
		return t.Encode()
	case TLConstructor_CRC32_inputPrivacyValueAllowUsers:
		t := m.To_InputPrivacyValueAllowUsers()
		return t.Encode()
	case TLConstructor_CRC32_inputPrivacyValueDisallowContacts:
		t := m.To_InputPrivacyValueDisallowContacts()
		return t.Encode()
	case TLConstructor_CRC32_inputPrivacyValueDisallowAll:
		t := m.To_InputPrivacyValueDisallowAll()
		return t.Encode()
	case TLConstructor_CRC32_inputPrivacyValueDisallowUsers:
		t := m.To_InputPrivacyValueDisallowUsers()
		return t.Encode()

	default:
		return []byte{}
	}
}

// inputPrivacyValueAllowContacts#d09e07b = InputPrivacyRule;
func (m *InputPrivacyRule) To_InputPrivacyValueAllowContacts() *TLInputPrivacyValueAllowContacts {
	return &TLInputPrivacyValueAllowContacts{
		Data2: m.Data2,
	}
}

// inputPrivacyValueAllowAll#184b35ce = InputPrivacyRule;
func (m *InputPrivacyRule) To_InputPrivacyValueAllowAll() *TLInputPrivacyValueAllowAll {
	return &TLInputPrivacyValueAllowAll{
		Data2: m.Data2,
	}
}

// inputPrivacyValueAllowUsers#131cc67f users:Vector<InputUser> = InputPrivacyRule;
func (m *InputPrivacyRule) To_InputPrivacyValueAllowUsers() *TLInputPrivacyValueAllowUsers {
	return &TLInputPrivacyValueAllowUsers{
		Data2: m.Data2,
	}
}

// inputPrivacyValueDisallowContacts#ba52007 = InputPrivacyRule;
func (m *InputPrivacyRule) To_InputPrivacyValueDisallowContacts() *TLInputPrivacyValueDisallowContacts {
	return &TLInputPrivacyValueDisallowContacts{
		Data2: m.Data2,
	}
}

// inputPrivacyValueDisallowAll#d66b66c9 = InputPrivacyRule;
func (m *InputPrivacyRule) To_InputPrivacyValueDisallowAll() *TLInputPrivacyValueDisallowAll {
	return &TLInputPrivacyValueDisallowAll{
		Data2: m.Data2,
	}
}

// inputPrivacyValueDisallowUsers#90110467 users:Vector<InputUser> = InputPrivacyRule;
func (m *InputPrivacyRule) To_InputPrivacyValueDisallowUsers() *TLInputPrivacyValueDisallowUsers {
	return &TLInputPrivacyValueDisallowUsers{
		Data2: m.Data2,
	}
}

// inputPrivacyValueAllowContacts#d09e07b = InputPrivacyRule;
func (m *TLInputPrivacyValueAllowContacts) To_InputPrivacyRule() *InputPrivacyRule {
	return &InputPrivacyRule{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputPrivacyValueAllowContacts) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputPrivacyValueAllowContacts))

	return x.buf
}

func (m *TLInputPrivacyValueAllowContacts) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// inputPrivacyValueAllowAll#184b35ce = InputPrivacyRule;
func (m *TLInputPrivacyValueAllowAll) To_InputPrivacyRule() *InputPrivacyRule {
	return &InputPrivacyRule{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputPrivacyValueAllowAll) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputPrivacyValueAllowAll))

	return x.buf
}

func (m *TLInputPrivacyValueAllowAll) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// inputPrivacyValueAllowUsers#131cc67f users:Vector<InputUser> = InputPrivacyRule;
func (m *TLInputPrivacyValueAllowUsers) To_InputPrivacyRule() *InputPrivacyRule {
	return &InputPrivacyRule{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputPrivacyValueAllowUsers) SetUsers(v []*InputUser) { m.Data2.Users = v }
func (m *TLInputPrivacyValueAllowUsers) GetUsers() []*InputUser  { return m.Data2.Users }

func (m *TLInputPrivacyValueAllowUsers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputPrivacyValueAllowUsers))

	return x.buf
}

func (m *TLInputPrivacyValueAllowUsers) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// inputPrivacyValueDisallowContacts#ba52007 = InputPrivacyRule;
func (m *TLInputPrivacyValueDisallowContacts) To_InputPrivacyRule() *InputPrivacyRule {
	return &InputPrivacyRule{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputPrivacyValueDisallowContacts) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputPrivacyValueDisallowContacts))

	return x.buf
}

func (m *TLInputPrivacyValueDisallowContacts) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// inputPrivacyValueDisallowAll#d66b66c9 = InputPrivacyRule;
func (m *TLInputPrivacyValueDisallowAll) To_InputPrivacyRule() *InputPrivacyRule {
	return &InputPrivacyRule{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputPrivacyValueDisallowAll) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputPrivacyValueDisallowAll))

	return x.buf
}

func (m *TLInputPrivacyValueDisallowAll) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// inputPrivacyValueDisallowUsers#90110467 users:Vector<InputUser> = InputPrivacyRule;
func (m *TLInputPrivacyValueDisallowUsers) To_InputPrivacyRule() *InputPrivacyRule {
	return &InputPrivacyRule{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputPrivacyValueDisallowUsers) SetUsers(v []*InputUser) { m.Data2.Users = v }
func (m *TLInputPrivacyValueDisallowUsers) GetUsers() []*InputUser  { return m.Data2.Users }

func (m *TLInputPrivacyValueDisallowUsers) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputPrivacyValueDisallowUsers))

	return x.buf
}

func (m *TLInputPrivacyValueDisallowUsers) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Authorization <--
//  + TL_Authorization
//

func (m *Authorization) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_authorization:
		t := m.To_Authorization()
		return t.Encode()

	default:
		return []byte{}
	}
}

// authorization#7bf2e6f6 hash:long flags:int device_model:string platform:string system_version:string api_id:int app_name:string app_version:string date_created:int date_active:int ip:string country:string region:string = Authorization;
func (m *Authorization) To_Authorization() *TLAuthorization {
	return &TLAuthorization{
		Data2: m.Data2,
	}
}

// authorization#7bf2e6f6 hash:long flags:int device_model:string platform:string system_version:string api_id:int app_name:string app_version:string date_created:int date_active:int ip:string country:string region:string = Authorization;
func (m *TLAuthorization) To_Authorization() *Authorization {
	return &Authorization{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLAuthorization) SetHash(v int64) { m.Data2.Hash = v }
func (m *TLAuthorization) GetHash() int64  { return m.Data2.Hash }

func (m *TLAuthorization) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLAuthorization) GetFlags() int32  { return m.Data2.Flags }

func (m *TLAuthorization) SetDeviceModel(v string) { m.Data2.DeviceModel = v }
func (m *TLAuthorization) GetDeviceModel() string  { return m.Data2.DeviceModel }

func (m *TLAuthorization) SetPlatform(v string) { m.Data2.Platform = v }
func (m *TLAuthorization) GetPlatform() string  { return m.Data2.Platform }

func (m *TLAuthorization) SetSystemVersion(v string) { m.Data2.SystemVersion = v }
func (m *TLAuthorization) GetSystemVersion() string  { return m.Data2.SystemVersion }

func (m *TLAuthorization) SetApiId(v int32) { m.Data2.ApiId = v }
func (m *TLAuthorization) GetApiId() int32  { return m.Data2.ApiId }

func (m *TLAuthorization) SetAppName(v string) { m.Data2.AppName = v }
func (m *TLAuthorization) GetAppName() string  { return m.Data2.AppName }

func (m *TLAuthorization) SetAppVersion(v string) { m.Data2.AppVersion = v }
func (m *TLAuthorization) GetAppVersion() string  { return m.Data2.AppVersion }

func (m *TLAuthorization) SetDateCreated(v int32) { m.Data2.DateCreated = v }
func (m *TLAuthorization) GetDateCreated() int32  { return m.Data2.DateCreated }

func (m *TLAuthorization) SetDateActive(v int32) { m.Data2.DateActive = v }
func (m *TLAuthorization) GetDateActive() int32  { return m.Data2.DateActive }

func (m *TLAuthorization) SetIp(v string) { m.Data2.Ip = v }
func (m *TLAuthorization) GetIp() string  { return m.Data2.Ip }

func (m *TLAuthorization) SetCountry(v string) { m.Data2.Country = v }
func (m *TLAuthorization) GetCountry() string  { return m.Data2.Country }

func (m *TLAuthorization) SetRegion(v string) { m.Data2.Region = v }
func (m *TLAuthorization) GetRegion() string  { return m.Data2.Region }

func (m *TLAuthorization) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_authorization))

	x.Long(m.Data2.Hash)
	x.Int(m.Data2.Flags)
	x.StringBytes(m.Data2.DeviceModel)
	x.StringBytes(m.Data2.Platform)
	x.StringBytes(m.Data2.SystemVersion)
	x.Int(m.Data2.ApiId)
	x.StringBytes(m.Data2.AppName)
	x.StringBytes(m.Data2.AppVersion)
	x.Int(m.Data2.DateCreated)
	x.Int(m.Data2.DateActive)
	x.StringBytes(m.Data2.Ip)
	x.StringBytes(m.Data2.Country)
	x.StringBytes(m.Data2.Region)

	return x.buf
}

func (m *TLAuthorization) Decode(dbuf *DecodeBuf) error {
	m.Data2.Hash = x.Long()
	m.Data2.Flags = x.Int()
	m.Data2.DeviceModel = x.StringBytes()
	m.Data2.Platform = x.StringBytes()
	m.Data2.SystemVersion = x.StringBytes()
	m.Data2.ApiId = x.Int()
	m.Data2.AppName = x.StringBytes()
	m.Data2.AppVersion = x.StringBytes()
	m.Data2.DateCreated = x.Int()
	m.Data2.DateActive = x.Int()
	m.Data2.Ip = x.StringBytes()
	m.Data2.Country = x.StringBytes()
	m.Data2.Region = x.StringBytes()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// TopPeerCategory <--
//  + TL_TopPeerCategoryBotsPM
//  + TL_TopPeerCategoryBotsInline
//  + TL_TopPeerCategoryCorrespondents
//  + TL_TopPeerCategoryGroups
//  + TL_TopPeerCategoryChannels
//  + TL_TopPeerCategoryPhoneCalls
//

func (m *TopPeerCategory) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_topPeerCategoryBotsPM:
		t := m.To_TopPeerCategoryBotsPM()
		return t.Encode()
	case TLConstructor_CRC32_topPeerCategoryBotsInline:
		t := m.To_TopPeerCategoryBotsInline()
		return t.Encode()
	case TLConstructor_CRC32_topPeerCategoryCorrespondents:
		t := m.To_TopPeerCategoryCorrespondents()
		return t.Encode()
	case TLConstructor_CRC32_topPeerCategoryGroups:
		t := m.To_TopPeerCategoryGroups()
		return t.Encode()
	case TLConstructor_CRC32_topPeerCategoryChannels:
		t := m.To_TopPeerCategoryChannels()
		return t.Encode()
	case TLConstructor_CRC32_topPeerCategoryPhoneCalls:
		t := m.To_TopPeerCategoryPhoneCalls()
		return t.Encode()

	default:
		return []byte{}
	}
}

// topPeerCategoryBotsPM#ab661b5b = TopPeerCategory;
func (m *TopPeerCategory) To_TopPeerCategoryBotsPM() *TLTopPeerCategoryBotsPM {
	return &TLTopPeerCategoryBotsPM{
		Data2: m.Data2,
	}
}

// topPeerCategoryBotsInline#148677e2 = TopPeerCategory;
func (m *TopPeerCategory) To_TopPeerCategoryBotsInline() *TLTopPeerCategoryBotsInline {
	return &TLTopPeerCategoryBotsInline{
		Data2: m.Data2,
	}
}

// topPeerCategoryCorrespondents#637b7ed = TopPeerCategory;
func (m *TopPeerCategory) To_TopPeerCategoryCorrespondents() *TLTopPeerCategoryCorrespondents {
	return &TLTopPeerCategoryCorrespondents{
		Data2: m.Data2,
	}
}

// topPeerCategoryGroups#bd17a14a = TopPeerCategory;
func (m *TopPeerCategory) To_TopPeerCategoryGroups() *TLTopPeerCategoryGroups {
	return &TLTopPeerCategoryGroups{
		Data2: m.Data2,
	}
}

// topPeerCategoryChannels#161d9628 = TopPeerCategory;
func (m *TopPeerCategory) To_TopPeerCategoryChannels() *TLTopPeerCategoryChannels {
	return &TLTopPeerCategoryChannels{
		Data2: m.Data2,
	}
}

// topPeerCategoryPhoneCalls#1e76a78c = TopPeerCategory;
func (m *TopPeerCategory) To_TopPeerCategoryPhoneCalls() *TLTopPeerCategoryPhoneCalls {
	return &TLTopPeerCategoryPhoneCalls{
		Data2: m.Data2,
	}
}

// topPeerCategoryBotsPM#ab661b5b = TopPeerCategory;
func (m *TLTopPeerCategoryBotsPM) To_TopPeerCategory() *TopPeerCategory {
	return &TopPeerCategory{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLTopPeerCategoryBotsPM) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_topPeerCategoryBotsPM))

	return x.buf
}

func (m *TLTopPeerCategoryBotsPM) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// topPeerCategoryBotsInline#148677e2 = TopPeerCategory;
func (m *TLTopPeerCategoryBotsInline) To_TopPeerCategory() *TopPeerCategory {
	return &TopPeerCategory{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLTopPeerCategoryBotsInline) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_topPeerCategoryBotsInline))

	return x.buf
}

func (m *TLTopPeerCategoryBotsInline) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// topPeerCategoryCorrespondents#637b7ed = TopPeerCategory;
func (m *TLTopPeerCategoryCorrespondents) To_TopPeerCategory() *TopPeerCategory {
	return &TopPeerCategory{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLTopPeerCategoryCorrespondents) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_topPeerCategoryCorrespondents))

	return x.buf
}

func (m *TLTopPeerCategoryCorrespondents) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// topPeerCategoryGroups#bd17a14a = TopPeerCategory;
func (m *TLTopPeerCategoryGroups) To_TopPeerCategory() *TopPeerCategory {
	return &TopPeerCategory{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLTopPeerCategoryGroups) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_topPeerCategoryGroups))

	return x.buf
}

func (m *TLTopPeerCategoryGroups) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// topPeerCategoryChannels#161d9628 = TopPeerCategory;
func (m *TLTopPeerCategoryChannels) To_TopPeerCategory() *TopPeerCategory {
	return &TopPeerCategory{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLTopPeerCategoryChannels) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_topPeerCategoryChannels))

	return x.buf
}

func (m *TLTopPeerCategoryChannels) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// topPeerCategoryPhoneCalls#1e76a78c = TopPeerCategory;
func (m *TLTopPeerCategoryPhoneCalls) To_TopPeerCategory() *TopPeerCategory {
	return &TopPeerCategory{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLTopPeerCategoryPhoneCalls) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_topPeerCategoryPhoneCalls))

	return x.buf
}

func (m *TLTopPeerCategoryPhoneCalls) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// ChannelAdminLogEvent <--
//  + TL_ChannelAdminLogEvent
//

func (m *ChannelAdminLogEvent) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_channelAdminLogEvent:
		t := m.To_ChannelAdminLogEvent()
		return t.Encode()

	default:
		return []byte{}
	}
}

// channelAdminLogEvent#3b5a3e40 id:long date:int user_id:int action:ChannelAdminLogEventAction = ChannelAdminLogEvent;
func (m *ChannelAdminLogEvent) To_ChannelAdminLogEvent() *TLChannelAdminLogEvent {
	return &TLChannelAdminLogEvent{
		Data2: m.Data2,
	}
}

// channelAdminLogEvent#3b5a3e40 id:long date:int user_id:int action:ChannelAdminLogEventAction = ChannelAdminLogEvent;
func (m *TLChannelAdminLogEvent) To_ChannelAdminLogEvent() *ChannelAdminLogEvent {
	return &ChannelAdminLogEvent{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLChannelAdminLogEvent) SetId(v int64) { m.Data2.Id = v }
func (m *TLChannelAdminLogEvent) GetId() int64  { return m.Data2.Id }

func (m *TLChannelAdminLogEvent) SetDate(v int32) { m.Data2.Date = v }
func (m *TLChannelAdminLogEvent) GetDate() int32  { return m.Data2.Date }

func (m *TLChannelAdminLogEvent) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLChannelAdminLogEvent) GetUserId() int32  { return m.Data2.UserId }

func (m *TLChannelAdminLogEvent) SetAction(v *ChannelAdminLogEventAction) { m.Data2.Action = v }
func (m *TLChannelAdminLogEvent) GetAction() *ChannelAdminLogEventAction  { return m.Data2.Action }

func (m *TLChannelAdminLogEvent) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelAdminLogEvent))

	x.Long(m.Data2.Id)
	x.Int(m.Data2.Date)
	x.Int(m.Data2.UserId)
	x.Bytes(m.Data2Action.Encode())

	return x.buf
}

func (m *TLChannelAdminLogEvent) Decode(dbuf *DecodeBuf) error {
	m.Data2.Id = x.Long()
	m.Data2.Date = x.Int()
	m.Data2.UserId = x.Int()
	m4 := &Action{}
	m4.Decode(dbuf)
	m.SetAction(m4)

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// InputFile <--
//  + TL_InputFile
//  + TL_InputFileBig
//

func (m *InputFile) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_inputFile:
		t := m.To_InputFile()
		return t.Encode()
	case TLConstructor_CRC32_inputFileBig:
		t := m.To_InputFileBig()
		return t.Encode()

	default:
		return []byte{}
	}
}

// inputFile#f52ff27f id:long parts:int name:string md5_checksum:string = InputFile;
func (m *InputFile) To_InputFile() *TLInputFile {
	return &TLInputFile{
		Data2: m.Data2,
	}
}

// inputFileBig#fa4f0bb5 id:long parts:int name:string = InputFile;
func (m *InputFile) To_InputFileBig() *TLInputFileBig {
	return &TLInputFileBig{
		Data2: m.Data2,
	}
}

// inputFile#f52ff27f id:long parts:int name:string md5_checksum:string = InputFile;
func (m *TLInputFile) To_InputFile() *InputFile {
	return &InputFile{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputFile) SetId(v int64) { m.Data2.Id = v }
func (m *TLInputFile) GetId() int64  { return m.Data2.Id }

func (m *TLInputFile) SetParts(v int32) { m.Data2.Parts = v }
func (m *TLInputFile) GetParts() int32  { return m.Data2.Parts }

func (m *TLInputFile) SetName(v string) { m.Data2.Name = v }
func (m *TLInputFile) GetName() string  { return m.Data2.Name }

func (m *TLInputFile) SetMd5Checksum(v string) { m.Data2.Md5Checksum = v }
func (m *TLInputFile) GetMd5Checksum() string  { return m.Data2.Md5Checksum }

func (m *TLInputFile) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputFile))

	x.Long(m.Data2.Id)
	x.Int(m.Data2.Parts)
	x.StringBytes(m.Data2.Name)
	x.StringBytes(m.Data2.Md5Checksum)

	return x.buf
}

func (m *TLInputFile) Decode(dbuf *DecodeBuf) error {
	m.Data2.Id = x.Long()
	m.Data2.Parts = x.Int()
	m.Data2.Name = x.StringBytes()
	m.Data2.Md5Checksum = x.StringBytes()

	return dbuf.err
}

// inputFileBig#fa4f0bb5 id:long parts:int name:string = InputFile;
func (m *TLInputFileBig) To_InputFile() *InputFile {
	return &InputFile{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputFileBig) SetId(v int64) { m.Data2.Id = v }
func (m *TLInputFileBig) GetId() int64  { return m.Data2.Id }

func (m *TLInputFileBig) SetParts(v int32) { m.Data2.Parts = v }
func (m *TLInputFileBig) GetParts() int32  { return m.Data2.Parts }

func (m *TLInputFileBig) SetName(v string) { m.Data2.Name = v }
func (m *TLInputFileBig) GetName() string  { return m.Data2.Name }

func (m *TLInputFileBig) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputFileBig))

	x.Long(m.Data2.Id)
	x.Int(m.Data2.Parts)
	x.StringBytes(m.Data2.Name)

	return x.buf
}

func (m *TLInputFileBig) Decode(dbuf *DecodeBuf) error {
	m.Data2.Id = x.Long()
	m.Data2.Parts = x.Int()
	m.Data2.Name = x.StringBytes()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Messages_ChatFull <--
//  + TL_MessagesChatFull
//

func (m *Messages_ChatFull) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_messages_chatFull:
		t := m.To_MessagesChatFull()
		return t.Encode()

	default:
		return []byte{}
	}
}

// messages.chatFull#e5d7d19c full_chat:ChatFull chats:Vector<Chat> users:Vector<User> = messages.ChatFull;
func (m *Messages_ChatFull) To_MessagesChatFull() *TLMessagesChatFull {
	return &TLMessagesChatFull{
		Data2: m.Data2,
	}
}

// messages.chatFull#e5d7d19c full_chat:ChatFull chats:Vector<Chat> users:Vector<User> = messages.ChatFull;
func (m *TLMessagesChatFull) To_Messages_ChatFull() *Messages_ChatFull {
	return &Messages_ChatFull{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMessagesChatFull) SetFullChat(v *ChatFull) { m.Data2.FullChat = v }
func (m *TLMessagesChatFull) GetFullChat() *ChatFull  { return m.Data2.FullChat }

func (m *TLMessagesChatFull) SetChats(v []*Chat) { m.Data2.Chats = v }
func (m *TLMessagesChatFull) GetChats() []*Chat  { return m.Data2.Chats }

func (m *TLMessagesChatFull) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLMessagesChatFull) GetUsers() []*User  { return m.Data2.Users }

func (m *TLMessagesChatFull) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_messages_chatFull))

	x.Bytes(m.Data2FullChat.Encode())

	return x.buf
}

func (m *TLMessagesChatFull) Decode(dbuf *DecodeBuf) error {
	m1 := &FullChat{}
	m1.Decode(dbuf)
	m.SetFullChat(m1)

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// InputDocument <--
//  + TL_InputDocumentEmpty
//  + TL_InputDocument
//

func (m *InputDocument) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_inputDocumentEmpty:
		t := m.To_InputDocumentEmpty()
		return t.Encode()
	case TLConstructor_CRC32_inputDocument:
		t := m.To_InputDocument()
		return t.Encode()

	default:
		return []byte{}
	}
}

// inputDocumentEmpty#72f0eaae = InputDocument;
func (m *InputDocument) To_InputDocumentEmpty() *TLInputDocumentEmpty {
	return &TLInputDocumentEmpty{
		Data2: m.Data2,
	}
}

// inputDocument#18798952 id:long access_hash:long = InputDocument;
func (m *InputDocument) To_InputDocument() *TLInputDocument {
	return &TLInputDocument{
		Data2: m.Data2,
	}
}

// inputDocumentEmpty#72f0eaae = InputDocument;
func (m *TLInputDocumentEmpty) To_InputDocument() *InputDocument {
	return &InputDocument{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputDocumentEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputDocumentEmpty))

	return x.buf
}

func (m *TLInputDocumentEmpty) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// inputDocument#18798952 id:long access_hash:long = InputDocument;
func (m *TLInputDocument) To_InputDocument() *InputDocument {
	return &InputDocument{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLInputDocument) SetId(v int64) { m.Data2.Id = v }
func (m *TLInputDocument) GetId() int64  { return m.Data2.Id }

func (m *TLInputDocument) SetAccessHash(v int64) { m.Data2.AccessHash = v }
func (m *TLInputDocument) GetAccessHash() int64  { return m.Data2.AccessHash }

func (m *TLInputDocument) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_inputDocument))

	x.Long(m.Data2.Id)
	x.Long(m.Data2.AccessHash)

	return x.buf
}

func (m *TLInputDocument) Decode(dbuf *DecodeBuf) error {
	m.Data2.Id = x.Long()
	m.Data2.AccessHash = x.Long()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// TopPeer <--
//  + TL_TopPeer
//

func (m *TopPeer) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_topPeer:
		t := m.To_TopPeer()
		return t.Encode()

	default:
		return []byte{}
	}
}

// topPeer#edcdc05b peer:Peer rating:double = TopPeer;
func (m *TopPeer) To_TopPeer() *TLTopPeer {
	return &TLTopPeer{
		Data2: m.Data2,
	}
}

// topPeer#edcdc05b peer:Peer rating:double = TopPeer;
func (m *TLTopPeer) To_TopPeer() *TopPeer {
	return &TopPeer{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLTopPeer) SetPeer(v *Peer) { m.Data2.Peer = v }
func (m *TLTopPeer) GetPeer() *Peer  { return m.Data2.Peer }

func (m *TLTopPeer) SetRating(v float64) { m.Data2.Rating = v }
func (m *TLTopPeer) GetRating() float64  { return m.Data2.Rating }

func (m *TLTopPeer) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_topPeer))

	x.Bytes(m.Data2Peer.Encode())
	x.Double(m.Data2.Rating)

	return x.buf
}

func (m *TLTopPeer) Decode(dbuf *DecodeBuf) error {
	m1 := &Peer{}
	m1.Decode(dbuf)
	m.SetPeer(m1)
	m.Data2.Rating = x.Double()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Config <--
//  + TL_Config
//

func (m *Config) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_config:
		t := m.To_Config()
		return t.Encode()

	default:
		return []byte{}
	}
}

// config#8df376a4 flags:# phonecalls_enabled:flags.1?true date:int expires:int test_mode:Bool this_dc:int dc_options:Vector<DcOption> chat_size_max:int megagroup_size_max:int forwarded_count_max:int online_update_period_ms:int offline_blur_timeout_ms:int offline_idle_timeout_ms:int online_cloud_timeout_ms:int notify_cloud_delay_ms:int notify_default_delay_ms:int chat_big_size:int push_chat_period_ms:int push_chat_limit:int saved_gifs_limit:int edit_time_limit:int rating_e_decay:int stickers_recent_limit:int stickers_faved_limit:int tmp_sessions:flags.0?int pinned_dialogs_count_max:int call_receive_timeout_ms:int call_ring_timeout_ms:int call_connect_timeout_ms:int call_packet_timeout_ms:int me_url_prefix:string suggested_lang_code:flags.2?string lang_pack_version:flags.2?int disabled_features:Vector<DisabledFeature> = Config;
func (m *Config) To_Config() *TLConfig {
	return &TLConfig{
		Data2: m.Data2,
	}
}

// config#8df376a4 flags:# phonecalls_enabled:flags.1?true date:int expires:int test_mode:Bool this_dc:int dc_options:Vector<DcOption> chat_size_max:int megagroup_size_max:int forwarded_count_max:int online_update_period_ms:int offline_blur_timeout_ms:int offline_idle_timeout_ms:int online_cloud_timeout_ms:int notify_cloud_delay_ms:int notify_default_delay_ms:int chat_big_size:int push_chat_period_ms:int push_chat_limit:int saved_gifs_limit:int edit_time_limit:int rating_e_decay:int stickers_recent_limit:int stickers_faved_limit:int tmp_sessions:flags.0?int pinned_dialogs_count_max:int call_receive_timeout_ms:int call_ring_timeout_ms:int call_connect_timeout_ms:int call_packet_timeout_ms:int me_url_prefix:string suggested_lang_code:flags.2?string lang_pack_version:flags.2?int disabled_features:Vector<DisabledFeature> = Config;
func (m *TLConfig) To_Config() *Config {
	return &Config{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLConfig) SetPhonecallsEnabled(v bool) { m.Data2.PhonecallsEnabled = v }
func (m *TLConfig) GetPhonecallsEnabled() bool  { return m.Data2.PhonecallsEnabled }

func (m *TLConfig) SetDate(v int32) { m.Data2.Date = v }
func (m *TLConfig) GetDate() int32  { return m.Data2.Date }

func (m *TLConfig) SetExpires(v int32) { m.Data2.Expires = v }
func (m *TLConfig) GetExpires() int32  { return m.Data2.Expires }

func (m *TLConfig) SetTestMode(v *Bool) { m.Data2.TestMode = v }
func (m *TLConfig) GetTestMode() *Bool  { return m.Data2.TestMode }

func (m *TLConfig) SetThisDc(v int32) { m.Data2.ThisDc = v }
func (m *TLConfig) GetThisDc() int32  { return m.Data2.ThisDc }

func (m *TLConfig) SetDcOptions(v []*DcOption) { m.Data2.DcOptions = v }
func (m *TLConfig) GetDcOptions() []*DcOption  { return m.Data2.DcOptions }

func (m *TLConfig) SetChatSizeMax(v int32) { m.Data2.ChatSizeMax = v }
func (m *TLConfig) GetChatSizeMax() int32  { return m.Data2.ChatSizeMax }

func (m *TLConfig) SetMegagroupSizeMax(v int32) { m.Data2.MegagroupSizeMax = v }
func (m *TLConfig) GetMegagroupSizeMax() int32  { return m.Data2.MegagroupSizeMax }

func (m *TLConfig) SetForwardedCountMax(v int32) { m.Data2.ForwardedCountMax = v }
func (m *TLConfig) GetForwardedCountMax() int32  { return m.Data2.ForwardedCountMax }

func (m *TLConfig) SetOnlineUpdatePeriodMs(v int32) { m.Data2.OnlineUpdatePeriodMs = v }
func (m *TLConfig) GetOnlineUpdatePeriodMs() int32  { return m.Data2.OnlineUpdatePeriodMs }

func (m *TLConfig) SetOfflineBlurTimeoutMs(v int32) { m.Data2.OfflineBlurTimeoutMs = v }
func (m *TLConfig) GetOfflineBlurTimeoutMs() int32  { return m.Data2.OfflineBlurTimeoutMs }

func (m *TLConfig) SetOfflineIdleTimeoutMs(v int32) { m.Data2.OfflineIdleTimeoutMs = v }
func (m *TLConfig) GetOfflineIdleTimeoutMs() int32  { return m.Data2.OfflineIdleTimeoutMs }

func (m *TLConfig) SetOnlineCloudTimeoutMs(v int32) { m.Data2.OnlineCloudTimeoutMs = v }
func (m *TLConfig) GetOnlineCloudTimeoutMs() int32  { return m.Data2.OnlineCloudTimeoutMs }

func (m *TLConfig) SetNotifyCloudDelayMs(v int32) { m.Data2.NotifyCloudDelayMs = v }
func (m *TLConfig) GetNotifyCloudDelayMs() int32  { return m.Data2.NotifyCloudDelayMs }

func (m *TLConfig) SetNotifyDefaultDelayMs(v int32) { m.Data2.NotifyDefaultDelayMs = v }
func (m *TLConfig) GetNotifyDefaultDelayMs() int32  { return m.Data2.NotifyDefaultDelayMs }

func (m *TLConfig) SetChatBigSize(v int32) { m.Data2.ChatBigSize = v }
func (m *TLConfig) GetChatBigSize() int32  { return m.Data2.ChatBigSize }

func (m *TLConfig) SetPushChatPeriodMs(v int32) { m.Data2.PushChatPeriodMs = v }
func (m *TLConfig) GetPushChatPeriodMs() int32  { return m.Data2.PushChatPeriodMs }

func (m *TLConfig) SetPushChatLimit(v int32) { m.Data2.PushChatLimit = v }
func (m *TLConfig) GetPushChatLimit() int32  { return m.Data2.PushChatLimit }

func (m *TLConfig) SetSavedGifsLimit(v int32) { m.Data2.SavedGifsLimit = v }
func (m *TLConfig) GetSavedGifsLimit() int32  { return m.Data2.SavedGifsLimit }

func (m *TLConfig) SetEditTimeLimit(v int32) { m.Data2.EditTimeLimit = v }
func (m *TLConfig) GetEditTimeLimit() int32  { return m.Data2.EditTimeLimit }

func (m *TLConfig) SetRatingEDecay(v int32) { m.Data2.RatingEDecay = v }
func (m *TLConfig) GetRatingEDecay() int32  { return m.Data2.RatingEDecay }

func (m *TLConfig) SetStickersRecentLimit(v int32) { m.Data2.StickersRecentLimit = v }
func (m *TLConfig) GetStickersRecentLimit() int32  { return m.Data2.StickersRecentLimit }

func (m *TLConfig) SetStickersFavedLimit(v int32) { m.Data2.StickersFavedLimit = v }
func (m *TLConfig) GetStickersFavedLimit() int32  { return m.Data2.StickersFavedLimit }

func (m *TLConfig) SetTmpSessions(v int32) { m.Data2.TmpSessions = v }
func (m *TLConfig) GetTmpSessions() int32  { return m.Data2.TmpSessions }

func (m *TLConfig) SetPinnedDialogsCountMax(v int32) { m.Data2.PinnedDialogsCountMax = v }
func (m *TLConfig) GetPinnedDialogsCountMax() int32  { return m.Data2.PinnedDialogsCountMax }

func (m *TLConfig) SetCallReceiveTimeoutMs(v int32) { m.Data2.CallReceiveTimeoutMs = v }
func (m *TLConfig) GetCallReceiveTimeoutMs() int32  { return m.Data2.CallReceiveTimeoutMs }

func (m *TLConfig) SetCallRingTimeoutMs(v int32) { m.Data2.CallRingTimeoutMs = v }
func (m *TLConfig) GetCallRingTimeoutMs() int32  { return m.Data2.CallRingTimeoutMs }

func (m *TLConfig) SetCallConnectTimeoutMs(v int32) { m.Data2.CallConnectTimeoutMs = v }
func (m *TLConfig) GetCallConnectTimeoutMs() int32  { return m.Data2.CallConnectTimeoutMs }

func (m *TLConfig) SetCallPacketTimeoutMs(v int32) { m.Data2.CallPacketTimeoutMs = v }
func (m *TLConfig) GetCallPacketTimeoutMs() int32  { return m.Data2.CallPacketTimeoutMs }

func (m *TLConfig) SetMeUrlPrefix(v string) { m.Data2.MeUrlPrefix = v }
func (m *TLConfig) GetMeUrlPrefix() string  { return m.Data2.MeUrlPrefix }

func (m *TLConfig) SetSuggestedLangCode(v string) { m.Data2.SuggestedLangCode = v }
func (m *TLConfig) GetSuggestedLangCode() string  { return m.Data2.SuggestedLangCode }

func (m *TLConfig) SetLangPackVersion(v int32) { m.Data2.LangPackVersion = v }
func (m *TLConfig) GetLangPackVersion() int32  { return m.Data2.LangPackVersion }

func (m *TLConfig) SetDisabledFeatures(v []*DisabledFeature) { m.Data2.DisabledFeatures = v }
func (m *TLConfig) GetDisabledFeatures() []*DisabledFeature  { return m.Data2.DisabledFeatures }

func (m *TLConfig) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_config))

	// flags
	var flags uint32 = 0
	if m.GetPhonecallsEnabled() == true {
		flags |= 1 << 1
	}
	if m.GetTmpSessions() != 0 {
		flags |= 1 << 24
	}
	if m.GetSuggestedLangCode() != "" {
		flags |= 1 << 31
	}
	if m.GetLangPackVersion() != 0 {
		flags |= 1 << 32
	}
	x.UInt(flags)

	x.Int(m.Data2.Date)
	x.Int(m.Data2.Expires)
	x.Bytes(m.Data2TestMode.Encode())
	x.Int(m.Data2.ThisDc)

	x.Int(m.Data2.ChatSizeMax)
	x.Int(m.Data2.MegagroupSizeMax)
	x.Int(m.Data2.ForwardedCountMax)
	x.Int(m.Data2.OnlineUpdatePeriodMs)
	x.Int(m.Data2.OfflineBlurTimeoutMs)
	x.Int(m.Data2.OfflineIdleTimeoutMs)
	x.Int(m.Data2.OnlineCloudTimeoutMs)
	x.Int(m.Data2.NotifyCloudDelayMs)
	x.Int(m.Data2.NotifyDefaultDelayMs)
	x.Int(m.Data2.ChatBigSize)
	x.Int(m.Data2.PushChatPeriodMs)
	x.Int(m.Data2.PushChatLimit)
	x.Int(m.Data2.SavedGifsLimit)
	x.Int(m.Data2.EditTimeLimit)
	x.Int(m.Data2.RatingEDecay)
	x.Int(m.Data2.StickersRecentLimit)
	x.Int(m.Data2.StickersFavedLimit)
	if m.GetTmpSessions() != 0 {
		x.Int(m.Data2.TmpSessions)
	}
	x.Int(m.Data2.PinnedDialogsCountMax)
	x.Int(m.Data2.CallReceiveTimeoutMs)
	x.Int(m.Data2.CallRingTimeoutMs)
	x.Int(m.Data2.CallConnectTimeoutMs)
	x.Int(m.Data2.CallPacketTimeoutMs)
	x.StringBytes(m.Data2.MeUrlPrefix)
	if m.GetSuggestedLangCode() != 0 {
		x.StringBytes(m.Data2.SuggestedLangCode)
	}
	if m.GetLangPackVersion() != 0 {
		x.Int(m.Data2.LangPackVersion)
	}

	return x.buf
}

func (m *TLConfig) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m.Data2.PhonecallsEnabled = true
	}
	m.Data2.Date = x.Int()
	m.Data2.Expires = x.Int()
	m5 := &TestMode{}
	m5.Decode(dbuf)
	m.SetTestMode(m5)
	m.Data2.ThisDc = x.Int()

	m.Data2.ChatSizeMax = x.Int()
	m.Data2.MegagroupSizeMax = x.Int()
	m.Data2.ForwardedCountMax = x.Int()
	m.Data2.OnlineUpdatePeriodMs = x.Int()
	m.Data2.OfflineBlurTimeoutMs = x.Int()
	m.Data2.OfflineIdleTimeoutMs = x.Int()
	m.Data2.OnlineCloudTimeoutMs = x.Int()
	m.Data2.NotifyCloudDelayMs = x.Int()
	m.Data2.NotifyDefaultDelayMs = x.Int()
	m.Data2.ChatBigSize = x.Int()
	m.Data2.PushChatPeriodMs = x.Int()
	m.Data2.PushChatLimit = x.Int()
	m.Data2.SavedGifsLimit = x.Int()
	m.Data2.EditTimeLimit = x.Int()
	m.Data2.RatingEDecay = x.Int()
	m.Data2.StickersRecentLimit = x.Int()
	m.Data2.StickersFavedLimit = x.Int()
	if (flags & (1 << 25)) != 0 {
		m.Data2.TmpSessions = x.Int()
	}
	m.Data2.PinnedDialogsCountMax = x.Int()
	m.Data2.CallReceiveTimeoutMs = x.Int()
	m.Data2.CallRingTimeoutMs = x.Int()
	m.Data2.CallConnectTimeoutMs = x.Int()
	m.Data2.CallPacketTimeoutMs = x.Int()
	m.Data2.MeUrlPrefix = x.StringBytes()
	if (flags & (1 << 32)) != 0 {
		m.Data2.SuggestedLangCode = x.StringBytes()
	}
	if (flags & (1 << 33)) != 0 {
		m.Data2.LangPackVersion = x.Int()
	}

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// ReceivedNotifyMessage <--
//  + TL_ReceivedNotifyMessage
//

func (m *ReceivedNotifyMessage) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_receivedNotifyMessage:
		t := m.To_ReceivedNotifyMessage()
		return t.Encode()

	default:
		return []byte{}
	}
}

// receivedNotifyMessage#a384b779 id:int flags:int = ReceivedNotifyMessage;
func (m *ReceivedNotifyMessage) To_ReceivedNotifyMessage() *TLReceivedNotifyMessage {
	return &TLReceivedNotifyMessage{
		Data2: m.Data2,
	}
}

// receivedNotifyMessage#a384b779 id:int flags:int = ReceivedNotifyMessage;
func (m *TLReceivedNotifyMessage) To_ReceivedNotifyMessage() *ReceivedNotifyMessage {
	return &ReceivedNotifyMessage{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLReceivedNotifyMessage) SetId(v int32) { m.Data2.Id = v }
func (m *TLReceivedNotifyMessage) GetId() int32  { return m.Data2.Id }

func (m *TLReceivedNotifyMessage) SetFlags(v int32) { m.Data2.Flags = v }
func (m *TLReceivedNotifyMessage) GetFlags() int32  { return m.Data2.Flags }

func (m *TLReceivedNotifyMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_receivedNotifyMessage))

	x.Int(m.Data2.Id)
	x.Int(m.Data2.Flags)

	return x.buf
}

func (m *TLReceivedNotifyMessage) Decode(dbuf *DecodeBuf) error {
	m.Data2.Id = x.Int()
	m.Data2.Flags = x.Int()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// BotCommand <--
//  + TL_BotCommand
//

func (m *BotCommand) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_botCommand:
		t := m.To_BotCommand()
		return t.Encode()

	default:
		return []byte{}
	}
}

// botCommand#c27ac8c7 command:string description:string = BotCommand;
func (m *BotCommand) To_BotCommand() *TLBotCommand {
	return &TLBotCommand{
		Data2: m.Data2,
	}
}

// botCommand#c27ac8c7 command:string description:string = BotCommand;
func (m *TLBotCommand) To_BotCommand() *BotCommand {
	return &BotCommand{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLBotCommand) SetCommand(v string) { m.Data2.Command = v }
func (m *TLBotCommand) GetCommand() string  { return m.Data2.Command }

func (m *TLBotCommand) SetDescription(v string) { m.Data2.Description = v }
func (m *TLBotCommand) GetDescription() string  { return m.Data2.Description }

func (m *TLBotCommand) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_botCommand))

	x.StringBytes(m.Data2.Command)
	x.StringBytes(m.Data2.Description)

	return x.buf
}

func (m *TLBotCommand) Decode(dbuf *DecodeBuf) error {
	m.Data2.Command = x.StringBytes()
	m.Data2.Description = x.StringBytes()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// ChannelParticipant <--
//  + TL_ChannelParticipant
//  + TL_ChannelParticipantSelf
//  + TL_ChannelParticipantCreator
//  + TL_ChannelParticipantAdmin
//  + TL_ChannelParticipantBanned
//

func (m *ChannelParticipant) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_channelParticipant:
		t := m.To_ChannelParticipant()
		return t.Encode()
	case TLConstructor_CRC32_channelParticipantSelf:
		t := m.To_ChannelParticipantSelf()
		return t.Encode()
	case TLConstructor_CRC32_channelParticipantCreator:
		t := m.To_ChannelParticipantCreator()
		return t.Encode()
	case TLConstructor_CRC32_channelParticipantAdmin:
		t := m.To_ChannelParticipantAdmin()
		return t.Encode()
	case TLConstructor_CRC32_channelParticipantBanned:
		t := m.To_ChannelParticipantBanned()
		return t.Encode()

	default:
		return []byte{}
	}
}

// channelParticipant#15ebac1d user_id:int date:int = ChannelParticipant;
func (m *ChannelParticipant) To_ChannelParticipant() *TLChannelParticipant {
	return &TLChannelParticipant{
		Data2: m.Data2,
	}
}

// channelParticipantSelf#a3289a6d user_id:int inviter_id:int date:int = ChannelParticipant;
func (m *ChannelParticipant) To_ChannelParticipantSelf() *TLChannelParticipantSelf {
	return &TLChannelParticipantSelf{
		Data2: m.Data2,
	}
}

// channelParticipantCreator#e3e2e1f9 user_id:int = ChannelParticipant;
func (m *ChannelParticipant) To_ChannelParticipantCreator() *TLChannelParticipantCreator {
	return &TLChannelParticipantCreator{
		Data2: m.Data2,
	}
}

// channelParticipantAdmin#a82fa898 flags:# can_edit:flags.0?true user_id:int inviter_id:int promoted_by:int date:int admin_rights:ChannelAdminRights = ChannelParticipant;
func (m *ChannelParticipant) To_ChannelParticipantAdmin() *TLChannelParticipantAdmin {
	return &TLChannelParticipantAdmin{
		Data2: m.Data2,
	}
}

// channelParticipantBanned#222c1886 flags:# left:flags.0?true user_id:int kicked_by:int date:int banned_rights:ChannelBannedRights = ChannelParticipant;
func (m *ChannelParticipant) To_ChannelParticipantBanned() *TLChannelParticipantBanned {
	return &TLChannelParticipantBanned{
		Data2: m.Data2,
	}
}

// channelParticipant#15ebac1d user_id:int date:int = ChannelParticipant;
func (m *TLChannelParticipant) To_ChannelParticipant() *ChannelParticipant {
	return &ChannelParticipant{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLChannelParticipant) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLChannelParticipant) GetUserId() int32  { return m.Data2.UserId }

func (m *TLChannelParticipant) SetDate(v int32) { m.Data2.Date = v }
func (m *TLChannelParticipant) GetDate() int32  { return m.Data2.Date }

func (m *TLChannelParticipant) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelParticipant))

	x.Int(m.Data2.UserId)
	x.Int(m.Data2.Date)

	return x.buf
}

func (m *TLChannelParticipant) Decode(dbuf *DecodeBuf) error {
	m.Data2.UserId = x.Int()
	m.Data2.Date = x.Int()

	return dbuf.err
}

// channelParticipantSelf#a3289a6d user_id:int inviter_id:int date:int = ChannelParticipant;
func (m *TLChannelParticipantSelf) To_ChannelParticipant() *ChannelParticipant {
	return &ChannelParticipant{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLChannelParticipantSelf) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLChannelParticipantSelf) GetUserId() int32  { return m.Data2.UserId }

func (m *TLChannelParticipantSelf) SetInviterId(v int32) { m.Data2.InviterId = v }
func (m *TLChannelParticipantSelf) GetInviterId() int32  { return m.Data2.InviterId }

func (m *TLChannelParticipantSelf) SetDate(v int32) { m.Data2.Date = v }
func (m *TLChannelParticipantSelf) GetDate() int32  { return m.Data2.Date }

func (m *TLChannelParticipantSelf) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelParticipantSelf))

	x.Int(m.Data2.UserId)
	x.Int(m.Data2.InviterId)
	x.Int(m.Data2.Date)

	return x.buf
}

func (m *TLChannelParticipantSelf) Decode(dbuf *DecodeBuf) error {
	m.Data2.UserId = x.Int()
	m.Data2.InviterId = x.Int()
	m.Data2.Date = x.Int()

	return dbuf.err
}

// channelParticipantCreator#e3e2e1f9 user_id:int = ChannelParticipant;
func (m *TLChannelParticipantCreator) To_ChannelParticipant() *ChannelParticipant {
	return &ChannelParticipant{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLChannelParticipantCreator) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLChannelParticipantCreator) GetUserId() int32  { return m.Data2.UserId }

func (m *TLChannelParticipantCreator) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelParticipantCreator))

	x.Int(m.Data2.UserId)

	return x.buf
}

func (m *TLChannelParticipantCreator) Decode(dbuf *DecodeBuf) error {
	m.Data2.UserId = x.Int()

	return dbuf.err
}

// channelParticipantAdmin#a82fa898 flags:# can_edit:flags.0?true user_id:int inviter_id:int promoted_by:int date:int admin_rights:ChannelAdminRights = ChannelParticipant;
func (m *TLChannelParticipantAdmin) To_ChannelParticipant() *ChannelParticipant {
	return &ChannelParticipant{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLChannelParticipantAdmin) SetCanEdit(v bool) { m.Data2.CanEdit = v }
func (m *TLChannelParticipantAdmin) GetCanEdit() bool  { return m.Data2.CanEdit }

func (m *TLChannelParticipantAdmin) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLChannelParticipantAdmin) GetUserId() int32  { return m.Data2.UserId }

func (m *TLChannelParticipantAdmin) SetInviterId(v int32) { m.Data2.InviterId = v }
func (m *TLChannelParticipantAdmin) GetInviterId() int32  { return m.Data2.InviterId }

func (m *TLChannelParticipantAdmin) SetPromotedBy(v int32) { m.Data2.PromotedBy = v }
func (m *TLChannelParticipantAdmin) GetPromotedBy() int32  { return m.Data2.PromotedBy }

func (m *TLChannelParticipantAdmin) SetDate(v int32) { m.Data2.Date = v }
func (m *TLChannelParticipantAdmin) GetDate() int32  { return m.Data2.Date }

func (m *TLChannelParticipantAdmin) SetAdminRights(v *ChannelAdminRights) { m.Data2.AdminRights = v }
func (m *TLChannelParticipantAdmin) GetAdminRights() *ChannelAdminRights  { return m.Data2.AdminRights }

func (m *TLChannelParticipantAdmin) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelParticipantAdmin))

	// flags
	var flags uint32 = 0
	if m.GetCanEdit() == true {
		flags |= 1 << 1
	}
	x.UInt(flags)

	x.Int(m.Data2.UserId)
	x.Int(m.Data2.InviterId)
	x.Int(m.Data2.PromotedBy)
	x.Int(m.Data2.Date)
	x.Bytes(m.Data2AdminRights.Encode())

	return x.buf
}

func (m *TLChannelParticipantAdmin) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m.Data2.CanEdit = true
	}
	m.Data2.UserId = x.Int()
	m.Data2.InviterId = x.Int()
	m.Data2.PromotedBy = x.Int()
	m.Data2.Date = x.Int()
	m7 := &AdminRights{}
	m7.Decode(dbuf)
	m.SetAdminRights(m7)

	return dbuf.err
}

// channelParticipantBanned#222c1886 flags:# left:flags.0?true user_id:int kicked_by:int date:int banned_rights:ChannelBannedRights = ChannelParticipant;
func (m *TLChannelParticipantBanned) To_ChannelParticipant() *ChannelParticipant {
	return &ChannelParticipant{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLChannelParticipantBanned) SetLeft(v bool) { m.Data2.Left = v }
func (m *TLChannelParticipantBanned) GetLeft() bool  { return m.Data2.Left }

func (m *TLChannelParticipantBanned) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLChannelParticipantBanned) GetUserId() int32  { return m.Data2.UserId }

func (m *TLChannelParticipantBanned) SetKickedBy(v int32) { m.Data2.KickedBy = v }
func (m *TLChannelParticipantBanned) GetKickedBy() int32  { return m.Data2.KickedBy }

func (m *TLChannelParticipantBanned) SetDate(v int32) { m.Data2.Date = v }
func (m *TLChannelParticipantBanned) GetDate() int32  { return m.Data2.Date }

func (m *TLChannelParticipantBanned) SetBannedRights(v *ChannelBannedRights) { m.Data2.BannedRights = v }
func (m *TLChannelParticipantBanned) GetBannedRights() *ChannelBannedRights {
	return m.Data2.BannedRights
}

func (m *TLChannelParticipantBanned) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_channelParticipantBanned))

	// flags
	var flags uint32 = 0
	if m.GetLeft() == true {
		flags |= 1 << 1
	}
	x.UInt(flags)

	x.Int(m.Data2.UserId)
	x.Int(m.Data2.KickedBy)
	x.Int(m.Data2.Date)
	x.Bytes(m.Data2BannedRights.Encode())

	return x.buf
}

func (m *TLChannelParticipantBanned) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m.Data2.Left = true
	}
	m.Data2.UserId = x.Int()
	m.Data2.KickedBy = x.Int()
	m.Data2.Date = x.Int()
	m6 := &BannedRights{}
	m6.Decode(dbuf)
	m.SetBannedRights(m6)

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// ChatParticipant <--
//  + TL_ChatParticipant
//  + TL_ChatParticipantCreator
//  + TL_ChatParticipantAdmin
//

func (m *ChatParticipant) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_chatParticipant:
		t := m.To_ChatParticipant()
		return t.Encode()
	case TLConstructor_CRC32_chatParticipantCreator:
		t := m.To_ChatParticipantCreator()
		return t.Encode()
	case TLConstructor_CRC32_chatParticipantAdmin:
		t := m.To_ChatParticipantAdmin()
		return t.Encode()

	default:
		return []byte{}
	}
}

// chatParticipant#c8d7493e user_id:int inviter_id:int date:int = ChatParticipant;
func (m *ChatParticipant) To_ChatParticipant() *TLChatParticipant {
	return &TLChatParticipant{
		Data2: m.Data2,
	}
}

// chatParticipantCreator#da13538a user_id:int = ChatParticipant;
func (m *ChatParticipant) To_ChatParticipantCreator() *TLChatParticipantCreator {
	return &TLChatParticipantCreator{
		Data2: m.Data2,
	}
}

// chatParticipantAdmin#e2d6e436 user_id:int inviter_id:int date:int = ChatParticipant;
func (m *ChatParticipant) To_ChatParticipantAdmin() *TLChatParticipantAdmin {
	return &TLChatParticipantAdmin{
		Data2: m.Data2,
	}
}

// chatParticipant#c8d7493e user_id:int inviter_id:int date:int = ChatParticipant;
func (m *TLChatParticipant) To_ChatParticipant() *ChatParticipant {
	return &ChatParticipant{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLChatParticipant) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLChatParticipant) GetUserId() int32  { return m.Data2.UserId }

func (m *TLChatParticipant) SetInviterId(v int32) { m.Data2.InviterId = v }
func (m *TLChatParticipant) GetInviterId() int32  { return m.Data2.InviterId }

func (m *TLChatParticipant) SetDate(v int32) { m.Data2.Date = v }
func (m *TLChatParticipant) GetDate() int32  { return m.Data2.Date }

func (m *TLChatParticipant) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_chatParticipant))

	x.Int(m.Data2.UserId)
	x.Int(m.Data2.InviterId)
	x.Int(m.Data2.Date)

	return x.buf
}

func (m *TLChatParticipant) Decode(dbuf *DecodeBuf) error {
	m.Data2.UserId = x.Int()
	m.Data2.InviterId = x.Int()
	m.Data2.Date = x.Int()

	return dbuf.err
}

// chatParticipantCreator#da13538a user_id:int = ChatParticipant;
func (m *TLChatParticipantCreator) To_ChatParticipant() *ChatParticipant {
	return &ChatParticipant{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLChatParticipantCreator) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLChatParticipantCreator) GetUserId() int32  { return m.Data2.UserId }

func (m *TLChatParticipantCreator) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_chatParticipantCreator))

	x.Int(m.Data2.UserId)

	return x.buf
}

func (m *TLChatParticipantCreator) Decode(dbuf *DecodeBuf) error {
	m.Data2.UserId = x.Int()

	return dbuf.err
}

// chatParticipantAdmin#e2d6e436 user_id:int inviter_id:int date:int = ChatParticipant;
func (m *TLChatParticipantAdmin) To_ChatParticipant() *ChatParticipant {
	return &ChatParticipant{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLChatParticipantAdmin) SetUserId(v int32) { m.Data2.UserId = v }
func (m *TLChatParticipantAdmin) GetUserId() int32  { return m.Data2.UserId }

func (m *TLChatParticipantAdmin) SetInviterId(v int32) { m.Data2.InviterId = v }
func (m *TLChatParticipantAdmin) GetInviterId() int32  { return m.Data2.InviterId }

func (m *TLChatParticipantAdmin) SetDate(v int32) { m.Data2.Date = v }
func (m *TLChatParticipantAdmin) GetDate() int32  { return m.Data2.Date }

func (m *TLChatParticipantAdmin) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_chatParticipantAdmin))

	x.Int(m.Data2.UserId)
	x.Int(m.Data2.InviterId)
	x.Int(m.Data2.Date)

	return x.buf
}

func (m *TLChatParticipantAdmin) Decode(dbuf *DecodeBuf) error {
	m.Data2.UserId = x.Int()
	m.Data2.InviterId = x.Int()
	m.Data2.Date = x.Int()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// PhotoSize <--
//  + TL_PhotoSizeEmpty
//  + TL_PhotoSize
//  + TL_PhotoCachedSize
//

func (m *PhotoSize) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_photoSizeEmpty:
		t := m.To_PhotoSizeEmpty()
		return t.Encode()
	case TLConstructor_CRC32_photoSize:
		t := m.To_PhotoSize()
		return t.Encode()
	case TLConstructor_CRC32_photoCachedSize:
		t := m.To_PhotoCachedSize()
		return t.Encode()

	default:
		return []byte{}
	}
}

// photoSizeEmpty#e17e23c type:string = PhotoSize;
func (m *PhotoSize) To_PhotoSizeEmpty() *TLPhotoSizeEmpty {
	return &TLPhotoSizeEmpty{
		Data2: m.Data2,
	}
}

// photoSize#77bfb61b type:string location:FileLocation w:int h:int size:int = PhotoSize;
func (m *PhotoSize) To_PhotoSize() *TLPhotoSize {
	return &TLPhotoSize{
		Data2: m.Data2,
	}
}

// photoCachedSize#e9a734fa type:string location:FileLocation w:int h:int bytes:bytes = PhotoSize;
func (m *PhotoSize) To_PhotoCachedSize() *TLPhotoCachedSize {
	return &TLPhotoCachedSize{
		Data2: m.Data2,
	}
}

// photoSizeEmpty#e17e23c type:string = PhotoSize;
func (m *TLPhotoSizeEmpty) To_PhotoSize() *PhotoSize {
	return &PhotoSize{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPhotoSizeEmpty) SetType(v string) { m.Data2.Type = v }
func (m *TLPhotoSizeEmpty) GetType() string  { return m.Data2.Type }

func (m *TLPhotoSizeEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_photoSizeEmpty))

	x.StringBytes(m.Data2.Type)

	return x.buf
}

func (m *TLPhotoSizeEmpty) Decode(dbuf *DecodeBuf) error {
	m.Data2.Type = x.StringBytes()

	return dbuf.err
}

// photoSize#77bfb61b type:string location:FileLocation w:int h:int size:int = PhotoSize;
func (m *TLPhotoSize) To_PhotoSize() *PhotoSize {
	return &PhotoSize{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPhotoSize) SetType(v string) { m.Data2.Type = v }
func (m *TLPhotoSize) GetType() string  { return m.Data2.Type }

func (m *TLPhotoSize) SetLocation(v *FileLocation) { m.Data2.Location = v }
func (m *TLPhotoSize) GetLocation() *FileLocation  { return m.Data2.Location }

func (m *TLPhotoSize) SetW(v int32) { m.Data2.W = v }
func (m *TLPhotoSize) GetW() int32  { return m.Data2.W }

func (m *TLPhotoSize) SetH(v int32) { m.Data2.H = v }
func (m *TLPhotoSize) GetH() int32  { return m.Data2.H }

func (m *TLPhotoSize) SetSize(v int32) { m.Data2.Size = v }
func (m *TLPhotoSize) GetSize() int32  { return m.Data2.Size }

func (m *TLPhotoSize) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_photoSize))

	x.StringBytes(m.Data2.Type)
	x.Bytes(m.Data2Location.Encode())
	x.Int(m.Data2.W)
	x.Int(m.Data2.H)
	x.Int(m.Data2.Size)

	return x.buf
}

func (m *TLPhotoSize) Decode(dbuf *DecodeBuf) error {
	m.Data2.Type = x.StringBytes()
	m2 := &Location{}
	m2.Decode(dbuf)
	m.SetLocation(m2)
	m.Data2.W = x.Int()
	m.Data2.H = x.Int()
	m.Data2.Size = x.Int()

	return dbuf.err
}

// photoCachedSize#e9a734fa type:string location:FileLocation w:int h:int bytes:bytes = PhotoSize;
func (m *TLPhotoCachedSize) To_PhotoSize() *PhotoSize {
	return &PhotoSize{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLPhotoCachedSize) SetType(v string) { m.Data2.Type = v }
func (m *TLPhotoCachedSize) GetType() string  { return m.Data2.Type }

func (m *TLPhotoCachedSize) SetLocation(v *FileLocation) { m.Data2.Location = v }
func (m *TLPhotoCachedSize) GetLocation() *FileLocation  { return m.Data2.Location }

func (m *TLPhotoCachedSize) SetW(v int32) { m.Data2.W = v }
func (m *TLPhotoCachedSize) GetW() int32  { return m.Data2.W }

func (m *TLPhotoCachedSize) SetH(v int32) { m.Data2.H = v }
func (m *TLPhotoCachedSize) GetH() int32  { return m.Data2.H }

func (m *TLPhotoCachedSize) SetBytes(v []byte) { m.Data2.Bytes = v }
func (m *TLPhotoCachedSize) GetBytes() []byte  { return m.Data2.Bytes }

func (m *TLPhotoCachedSize) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_photoCachedSize))

	x.StringBytes(m.Data2.Type)
	x.Bytes(m.Data2Location.Encode())
	x.Int(m.Data2.W)
	x.Int(m.Data2.H)
	x.StringBytes(m.Data2.Bytes)

	return x.buf
}

func (m *TLPhotoCachedSize) Decode(dbuf *DecodeBuf) error {
	m.Data2.Type = x.StringBytes()
	m2 := &Location{}
	m2.Decode(dbuf)
	m.SetLocation(m2)
	m.Data2.W = x.Int()
	m.Data2.H = x.Int()
	m.Data2.Bytes = x.StringBytes()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Contacts_Blocked <--
//  + TL_ContactsBlocked
//  + TL_ContactsBlockedSlice
//

func (m *Contacts_Blocked) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_contacts_blocked:
		t := m.To_ContactsBlocked()
		return t.Encode()
	case TLConstructor_CRC32_contacts_blockedSlice:
		t := m.To_ContactsBlockedSlice()
		return t.Encode()

	default:
		return []byte{}
	}
}

// contacts.blocked#1c138d15 blocked:Vector<ContactBlocked> users:Vector<User> = contacts.Blocked;
func (m *Contacts_Blocked) To_ContactsBlocked() *TLContactsBlocked {
	return &TLContactsBlocked{
		Data2: m.Data2,
	}
}

// contacts.blockedSlice#900802a1 count:int blocked:Vector<ContactBlocked> users:Vector<User> = contacts.Blocked;
func (m *Contacts_Blocked) To_ContactsBlockedSlice() *TLContactsBlockedSlice {
	return &TLContactsBlockedSlice{
		Data2: m.Data2,
	}
}

// contacts.blocked#1c138d15 blocked:Vector<ContactBlocked> users:Vector<User> = contacts.Blocked;
func (m *TLContactsBlocked) To_Contacts_Blocked() *Contacts_Blocked {
	return &Contacts_Blocked{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLContactsBlocked) SetBlocked(v []*ContactBlocked) { m.Data2.Blocked = v }
func (m *TLContactsBlocked) GetBlocked() []*ContactBlocked  { return m.Data2.Blocked }

func (m *TLContactsBlocked) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLContactsBlocked) GetUsers() []*User  { return m.Data2.Users }

func (m *TLContactsBlocked) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_contacts_blocked))

	return x.buf
}

func (m *TLContactsBlocked) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// contacts.blockedSlice#900802a1 count:int blocked:Vector<ContactBlocked> users:Vector<User> = contacts.Blocked;
func (m *TLContactsBlockedSlice) To_Contacts_Blocked() *Contacts_Blocked {
	return &Contacts_Blocked{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLContactsBlockedSlice) SetCount(v int32) { m.Data2.Count = v }
func (m *TLContactsBlockedSlice) GetCount() int32  { return m.Data2.Count }

func (m *TLContactsBlockedSlice) SetBlocked(v []*ContactBlocked) { m.Data2.Blocked = v }
func (m *TLContactsBlockedSlice) GetBlocked() []*ContactBlocked  { return m.Data2.Blocked }

func (m *TLContactsBlockedSlice) SetUsers(v []*User) { m.Data2.Users = v }
func (m *TLContactsBlockedSlice) GetUsers() []*User  { return m.Data2.Users }

func (m *TLContactsBlockedSlice) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_contacts_blockedSlice))

	x.Int(m.Data2.Count)

	return x.buf
}

func (m *TLContactsBlockedSlice) Decode(dbuf *DecodeBuf) error {
	m.Data2.Count = x.Int()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// Upload_File <--
//  + TL_UploadFile
//  + TL_UploadFileCdnRedirect
//

func (m *Upload_File) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_upload_file:
		t := m.To_UploadFile()
		return t.Encode()
	case TLConstructor_CRC32_upload_fileCdnRedirect:
		t := m.To_UploadFileCdnRedirect()
		return t.Encode()

	default:
		return []byte{}
	}
}

// upload.file#96a18d5 type:storage.FileType mtime:int bytes:bytes = upload.File;
func (m *Upload_File) To_UploadFile() *TLUploadFile {
	return &TLUploadFile{
		Data2: m.Data2,
	}
}

// upload.fileCdnRedirect#ea52fe5a dc_id:int file_token:bytes encryption_key:bytes encryption_iv:bytes cdn_file_hashes:Vector<CdnFileHash> = upload.File;
func (m *Upload_File) To_UploadFileCdnRedirect() *TLUploadFileCdnRedirect {
	return &TLUploadFileCdnRedirect{
		Data2: m.Data2,
	}
}

// upload.file#96a18d5 type:storage.FileType mtime:int bytes:bytes = upload.File;
func (m *TLUploadFile) To_Upload_File() *Upload_File {
	return &Upload_File{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUploadFile) SetType(v *Storage_FileType) { m.Data2.Type = v }
func (m *TLUploadFile) GetType() *Storage_FileType  { return m.Data2.Type }

func (m *TLUploadFile) SetMtime(v int32) { m.Data2.Mtime = v }
func (m *TLUploadFile) GetMtime() int32  { return m.Data2.Mtime }

func (m *TLUploadFile) SetBytes(v []byte) { m.Data2.Bytes = v }
func (m *TLUploadFile) GetBytes() []byte  { return m.Data2.Bytes }

func (m *TLUploadFile) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_upload_file))

	x.Bytes(m.Data2Type.Encode())
	x.Int(m.Data2.Mtime)
	x.StringBytes(m.Data2.Bytes)

	return x.buf
}

func (m *TLUploadFile) Decode(dbuf *DecodeBuf) error {
	m1 := &Type{}
	m1.Decode(dbuf)
	m.SetType(m1)
	m.Data2.Mtime = x.Int()
	m.Data2.Bytes = x.StringBytes()

	return dbuf.err
}

// upload.fileCdnRedirect#ea52fe5a dc_id:int file_token:bytes encryption_key:bytes encryption_iv:bytes cdn_file_hashes:Vector<CdnFileHash> = upload.File;
func (m *TLUploadFileCdnRedirect) To_Upload_File() *Upload_File {
	return &Upload_File{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLUploadFileCdnRedirect) SetDcId(v int32) { m.Data2.DcId = v }
func (m *TLUploadFileCdnRedirect) GetDcId() int32  { return m.Data2.DcId }

func (m *TLUploadFileCdnRedirect) SetFileToken(v []byte) { m.Data2.FileToken = v }
func (m *TLUploadFileCdnRedirect) GetFileToken() []byte  { return m.Data2.FileToken }

func (m *TLUploadFileCdnRedirect) SetEncryptionKey(v []byte) { m.Data2.EncryptionKey = v }
func (m *TLUploadFileCdnRedirect) GetEncryptionKey() []byte  { return m.Data2.EncryptionKey }

func (m *TLUploadFileCdnRedirect) SetEncryptionIv(v []byte) { m.Data2.EncryptionIv = v }
func (m *TLUploadFileCdnRedirect) GetEncryptionIv() []byte  { return m.Data2.EncryptionIv }

func (m *TLUploadFileCdnRedirect) SetCdnFileHashes(v []*CdnFileHash) { m.Data2.CdnFileHashes = v }
func (m *TLUploadFileCdnRedirect) GetCdnFileHashes() []*CdnFileHash  { return m.Data2.CdnFileHashes }

func (m *TLUploadFileCdnRedirect) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_upload_fileCdnRedirect))

	x.Int(m.Data2.DcId)
	x.StringBytes(m.Data2.FileToken)
	x.StringBytes(m.Data2.EncryptionKey)
	x.StringBytes(m.Data2.EncryptionIv)

	return x.buf
}

func (m *TLUploadFileCdnRedirect) Decode(dbuf *DecodeBuf) error {
	m.Data2.DcId = x.Int()
	m.Data2.FileToken = x.StringBytes()
	m.Data2.EncryptionKey = x.StringBytes()
	m.Data2.EncryptionIv = x.StringBytes()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// DraftMessage <--
//  + TL_DraftMessageEmpty
//  + TL_DraftMessage
//

func (m *DraftMessage) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_draftMessageEmpty:
		t := m.To_DraftMessageEmpty()
		return t.Encode()
	case TLConstructor_CRC32_draftMessage:
		t := m.To_DraftMessage()
		return t.Encode()

	default:
		return []byte{}
	}
}

// draftMessageEmpty#ba4baec5 = DraftMessage;
func (m *DraftMessage) To_DraftMessageEmpty() *TLDraftMessageEmpty {
	return &TLDraftMessageEmpty{
		Data2: m.Data2,
	}
}

// draftMessage#fd8e711f flags:# no_webpage:flags.1?true reply_to_msg_id:flags.0?int message:string entities:flags.3?Vector<MessageEntity> date:int = DraftMessage;
func (m *DraftMessage) To_DraftMessage() *TLDraftMessage {
	return &TLDraftMessage{
		Data2: m.Data2,
	}
}

// draftMessageEmpty#ba4baec5 = DraftMessage;
func (m *TLDraftMessageEmpty) To_DraftMessage() *DraftMessage {
	return &DraftMessage{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLDraftMessageEmpty) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_draftMessageEmpty))

	return x.buf
}

func (m *TLDraftMessageEmpty) Decode(dbuf *DecodeBuf) error {

	return dbuf.err
}

// draftMessage#fd8e711f flags:# no_webpage:flags.1?true reply_to_msg_id:flags.0?int message:string entities:flags.3?Vector<MessageEntity> date:int = DraftMessage;
func (m *TLDraftMessage) To_DraftMessage() *DraftMessage {
	return &DraftMessage{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLDraftMessage) SetNoWebpage(v bool) { m.Data2.NoWebpage = v }
func (m *TLDraftMessage) GetNoWebpage() bool  { return m.Data2.NoWebpage }

func (m *TLDraftMessage) SetReplyToMsgId(v int32) { m.Data2.ReplyToMsgId = v }
func (m *TLDraftMessage) GetReplyToMsgId() int32  { return m.Data2.ReplyToMsgId }

func (m *TLDraftMessage) SetMessage(v string) { m.Data2.Message = v }
func (m *TLDraftMessage) GetMessage() string  { return m.Data2.Message }

func (m *TLDraftMessage) SetEntities(v []*MessageEntity) { m.Data2.Entities = v }
func (m *TLDraftMessage) GetEntities() []*MessageEntity  { return m.Data2.Entities }

func (m *TLDraftMessage) SetDate(v int32) { m.Data2.Date = v }
func (m *TLDraftMessage) GetDate() int32  { return m.Data2.Date }

func (m *TLDraftMessage) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_draftMessage))

	// flags
	var flags uint32 = 0
	if m.GetNoWebpage() == true {
		flags |= 1 << 1
	}
	if m.GetReplyToMsgId() != 0 {
		flags |= 1 << 2
	}
	if m.GetEntities() != nil {
		flags |= 1 << 4
	}
	x.UInt(flags)

	if m.GetReplyToMsgId() != 0 {
		x.Int(m.Data2.ReplyToMsgId)
	}
	x.StringBytes(m.Data2.Message)
	if m.Data2.Entities != 0 {

	}
	x.Int(m.Data2.Date)

	return x.buf
}

func (m *TLDraftMessage) Decode(dbuf *DecodeBuf) error {
	flags := dbuf.UInt()
	if (flags & (1 << 2)) != 0 {
		m.Data2.NoWebpage = true
	}
	if (flags & (1 << 3)) != 0 {
		m.Data2.ReplyToMsgId = x.Int()
	}
	m.Data2.Message = x.StringBytes()
	if (flags & (1 << 5)) != 0 {

	}
	m.Data2.Date = x.Int()

	return dbuf.err
}

///////////////////////////////////////////////////////////////////////////////
// MaskCoords <--
//  + TL_MaskCoords
//

func (m *MaskCoords) Encode() []byte {
	switch m.GetConstructor() {
	case TLConstructor_CRC32_maskCoords:
		t := m.To_MaskCoords()
		return t.Encode()

	default:
		return []byte{}
	}
}

// maskCoords#aed6dbb2 n:int x:double y:double zoom:double = MaskCoords;
func (m *MaskCoords) To_MaskCoords() *TLMaskCoords {
	return &TLMaskCoords{
		Data2: m.Data2,
	}
}

// maskCoords#aed6dbb2 n:int x:double y:double zoom:double = MaskCoords;
func (m *TLMaskCoords) To_MaskCoords() *MaskCoords {
	return &MaskCoords{
		Constructor: TLConstructor_CRC32_error,
		Data2:       m.Data2,
	}
}

func (m *TLMaskCoords) SetN(v int32) { m.Data2.N = v }
func (m *TLMaskCoords) GetN() int32  { return m.Data2.N }

func (m *TLMaskCoords) SetX(v float64) { m.Data2.X = v }
func (m *TLMaskCoords) GetX() float64  { return m.Data2.X }

func (m *TLMaskCoords) SetY(v float64) { m.Data2.Y = v }
func (m *TLMaskCoords) GetY() float64  { return m.Data2.Y }

func (m *TLMaskCoords) SetZoom(v float64) { m.Data2.Zoom = v }
func (m *TLMaskCoords) GetZoom() float64  { return m.Data2.Zoom }

func (m *TLMaskCoords) Encode() []byte {
	x := NewEncodeBuf(512)
	x.Int(int32(TLConstructor_CRC32_maskCoords))

	x.Int(m.Data2.N)
	x.Double(m.Data2.X)
	x.Double(m.Data2.Y)
	x.Double(m.Data2.Zoom)

	return x.buf
}

func (m *TLMaskCoords) Decode(dbuf *DecodeBuf) error {
	m.Data2.N = x.Int()
	m.Data2.X = x.Double()
	m.Data2.Y = x.Double()
	m.Data2.Zoom = x.Double()

	return dbuf.err
}
