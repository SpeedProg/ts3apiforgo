// ClientJoinEvent
package ts3api

import ()

type ClientJoinEvent struct {
	// TODO: notifycliententerview cfid=0 ctid=1 reasonid=0 clid=18 client_unique_identifier=ServerQuery client_nickname=Unknown\sfrom\s192.168.0.1:57329 client_input_muted=0 client_output_muted=0 client_outputonly_muted=0 client_input_hardware=0 client_output_hardware=0 client_meta_data client_is_recording=0 client_database_id=31 client_channel_group_id=8 client_servergroups=8 client_away=0 client_away_message client_type=1 client_flag_avatar client_talk_power=0 client_talk_request=0 client_talk_request_msg client_description client_is_talker=0 client_is_priority_speaker=0 client_unread_messages=0 client_nickname_phonetic client_needed_serverquery_view_power=75 client_icon_id=0 client_is_channel_commander=0 client_country client_channel_group_inherited_channel_id=1 client_badges
	cId int
	api *TS3Api
}

func (event ClientJoinEvent) ClientId() int {
	return event.cId
}

func (event ClientJoinEvent) Api() *TS3Api {
	return event.api
}
