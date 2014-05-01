package ts3api

import (
	"errors"
	"strconv"
	"strings"
)

type Serverinfo struct {
	VirtualserverUniqueIdentifier                       string
	VirtualserverName                                   string
	VirtualserverWelcomemessage                         string
	VirtualserverPlatform                               string
	VirtualserverVersion                                string
	VirtualserverMaxclients                             int
	VirtualserverPassword                               string
	VirtualserverClientsonline                          int
	VirtualserverChannelsonline                         int
	VirtualserverCreated                                bool
	VirtualserverUptime                                 int
	VirtualserverCodecEncryptionMode                    int
	VirtualserverHostmessage                            string
	VirtualserverHostmessageMode                        int
	VirtualserverFilebase                               string
	VirtualserverDefaultServerGroup                     int
	VirtualserverDefaultChannelGroup                    int
	VirtualserverFlagPassword                           bool
	VirtualserverDefaultChannelAdminGroup               int
	VirtualserverMaxDownloadTotalBandwidth              uint64
	VirtualserverMaxUploadTotalBandwidth                uint64
	VirtualserverHostbannerUrl                          string
	VirtualserverHostbannerGfxUrl                       string
	VirtualserverHostbannerGfxInterval                  int
	VirtualserverComplainAutobanCount                   int
	VirtualserverComplainAutobanTime                    int
	VirtualserverComplainRemoveTime                     int
	VirtualserverMinClientsInChannelBeforeForcedSilence int
	VirtualserverPrioritySpeakerDimmModificator         float32
	VirtualserverId                                     int
	VirtualserverAntifloodPointsTickReduce              int
	VirtualserverAntifloodPointsNeededCommandBlock      int
	VirtualserverAntifloodPointsNeededIpBlock           int
	VirtualserverClientConnections                      int
	VirtualserverQueryClientConnections                 int
	VirtualserverHostbuttonTooltip                      string
	VirtualserverHostbuttonUrl                          string
	VirtualserverHostbuttonGfxUrl                       string
	VirtualserverQueryclientsonline                     uint64
	VirtualserverDownloadQuota                          uint64
	VirtualserverUploadQuota                            uint64
	VirtualserverMonthBytesDownloaded                   int
	VirtualserverMonthBytesUploaded                     int
	VirtualserverTotalBytesDownloaded                   uint64
	VirtualserverTotalBytesUploaded                     uint64
	VirtualserverPort                                   int
	VirtualserverAutostart                              bool
	VirtualserverMachineId                              int
	VirtualserverNeededIdentitySecurityLevel            int
	VirtualserverLogClient                              bool
	VirtualserverLogQuery                               bool
	VirtualserverLogChannel                             bool
	VirtualserverLogPermissions                         bool
	VirtualserverLogServer                              bool
	VirtualserverLogFiletransfer                        bool
	VirtualserverMinClientVersion                       uint64
	VirtualserverNamePhonetic                           string
	VirtualserverIconId                                 int
	VirtualserverReservedSlots                          uint64
	VirtualserverTotalPacketlossSpeech                  float32
	VirtualserverTotalPacketlossKeepalive               float32
	VirtualserverTotalPacketlossControl                 float32
	VirtualserverTotalPacketlossTotal                   float32
	VirtualserverTotalPing                              float32
	VirtualserverIp                                     string
	VirtualserverWeblistEnabled                         bool
	VirtualserverAskForPrivilegekey                     bool
	VirtualserverHostbannerMode                         int
	VirtualserverChannelTempDeleteDelayDefault          int
	VirtualserverStatus                                 string
	ConnectionFiletransferBandwidthSent                 uint64
	ConnectionFiletransferBandwidthReceived             uint64
	ConnectionFiletransferBytesSentTotal                int
	ConnectionFiletransferBytesReceivedTotal            int
	ConnectionPacketsSentSpeech                         int
	ConnectionBytesSentSpeech                           uint64
	ConnectionPacketsReceivedSpeech                     int
	ConnectionBytesReceivedSpeech                       int
	ConnectionPacketsSentKeepalive                      int
	ConnectionBytesSentKeepalive                        int
	ConnectionPacketsReceivedKeepalive                  int
	ConnectionBytesReceivedKeepalive                    int
	ConnectionPacketsSentControl                        int
	ConnectionBytesSentControl                          int
	ConnectionPacketsReceivedControl                    int
	ConnectionBytesReceivedControl                      int
	ConnectionPacketsSentTotal                          int
	ConnectionBytesSentTotal                            uint64
	ConnectionPacketsReceivedTotal                      int
	ConnectionBytesReceivedTotal                        uint64
	ConnectionBandwidthSentLastSecondTotal              int
	ConnectionBandwidthSentLastMinuteTotal              int
	ConnectionBandwidthReceivedLastSecondTotal          int
	ConnectionBandwidthReceivedLastMinuteTotal          int
}

func (this *Serverinfo) parseMsg(s string) {
	params := strings.Split(s, " ")
	for _, param := range params {
		if strings.Contains(param, "=") {
			kvpair := strings.SplitN(param, "=", 2)
			this.parseParam(kvpair[0], kvpair[1])
		} else {
			this.parseParam(param, "")
		}
	}
}
func (this *Serverinfo) parseParam(k, v string) (err error) {
	switch k {
	case "virtualserver_unique_identifier":
		this.VirtualserverUniqueIdentifier = v
	case "virtualserver_name":
		this.VirtualserverName = v
	case "virtualserver_welcomemessage":
		this.VirtualserverWelcomemessage = v
	case "virtualserver_platform":
		this.VirtualserverPlatform = v
	case "virtualserver_version":
		this.VirtualserverVersion = v
	case "virtualserver_maxclients":
		this.VirtualserverMaxclients, err = strconv.Atoi(v)
		if err != nil {
			this.VirtualserverMaxclients = -1
			return
		}

	case "virtualserver_password":
		this.VirtualserverPassword = v
	case "virtualserver_clientsonline":
		this.VirtualserverClientsonline, err = strconv.Atoi(v)
		if err != nil {
			this.VirtualserverClientsonline = -1
			return
		}

	case "virtualserver_channelsonline":
		this.VirtualserverChannelsonline, err = strconv.Atoi(v)
		if err != nil {
			this.VirtualserverChannelsonline = -1
			return
		}

	case "virtualserver_created":
		this.VirtualserverCreated, err = getBoolFromString(v)
	case "virtualserver_uptime":
		this.VirtualserverUptime, err = strconv.Atoi(v)
		if err != nil {
			this.VirtualserverUptime = -1
			return
		}

	case "virtualserver_codec_encryption_mode":
		this.VirtualserverCodecEncryptionMode, err = strconv.Atoi(v)
		if err != nil {
			this.VirtualserverCodecEncryptionMode = -1
			return
		}

	case "virtualserver_hostmessage":
		this.VirtualserverHostmessage = v
	case "virtualserver_hostmessage_mode":
		this.VirtualserverHostmessageMode, err = strconv.Atoi(v)
		if err != nil {
			this.VirtualserverHostmessageMode = -1
			return
		}

	case "virtualserver_filebase":
		this.VirtualserverFilebase = v
	case "virtualserver_default_server_group":
		this.VirtualserverDefaultServerGroup, err = strconv.Atoi(v)
		if err != nil {
			this.VirtualserverDefaultServerGroup = -1
			return
		}

	case "virtualserver_default_channel_group":
		this.VirtualserverDefaultChannelGroup, err = strconv.Atoi(v)
		if err != nil {
			this.VirtualserverDefaultChannelGroup = -1
			return
		}

	case "virtualserver_flag_password":
		this.VirtualserverFlagPassword, err = getBoolFromString(v)
	case "virtualserver_default_channel_admin_group":
		this.VirtualserverDefaultChannelAdminGroup, err = strconv.Atoi(v)
		if err != nil {
			this.VirtualserverDefaultChannelAdminGroup = -1
			return
		}

	case "virtualserver_max_download_total_bandwidth":
		this.VirtualserverMaxDownloadTotalBandwidth, err = strconv.ParseUint(v, 10, 64)
		if err != nil {
			this.VirtualserverMaxDownloadTotalBandwidth = 0
			return
		}

	case "virtualserver_max_upload_total_bandwidth":
		this.VirtualserverMaxUploadTotalBandwidth, err = strconv.ParseUint(v, 10, 64)
		if err != nil {
			this.VirtualserverMaxUploadTotalBandwidth = 0
			return
		}

	case "virtualserver_hostbanner_url":
		this.VirtualserverHostbannerUrl = v
	case "virtualserver_hostbanner_gfx_url":
		this.VirtualserverHostbannerGfxUrl = v
	case "virtualserver_hostbanner_gfx_interval":
		this.VirtualserverHostbannerGfxInterval, err = strconv.Atoi(v)
		if err != nil {
			this.VirtualserverHostbannerGfxInterval = -1
			return
		}

	case "virtualserver_complain_autoban_count":
		this.VirtualserverComplainAutobanCount, err = strconv.Atoi(v)
		if err != nil {
			this.VirtualserverComplainAutobanCount = -1
			return
		}

	case "virtualserver_complain_autoban_time":
		this.VirtualserverComplainAutobanTime, err = strconv.Atoi(v)
		if err != nil {
			this.VirtualserverComplainAutobanTime = -1
			return
		}

	case "virtualserver_complain_remove_time":
		this.VirtualserverComplainRemoveTime, err = strconv.Atoi(v)
		if err != nil {
			this.VirtualserverComplainRemoveTime = -1
			return
		}

	case "virtualserver_min_clients_in_channel_before_forced_silence":
		this.VirtualserverMinClientsInChannelBeforeForcedSilence, err = strconv.Atoi(v)
		if err != nil {
			this.VirtualserverMinClientsInChannelBeforeForcedSilence = -1
			return
		}

	case "virtualserver_priority_speaker_dimm_modificator":
		var g float64
		g, err = strconv.ParseFloat(v, 32)
		if err != nil {
			this.VirtualserverPrioritySpeakerDimmModificator = 0.0
			return
		}
		this.VirtualserverPrioritySpeakerDimmModificator = float32(g)

	case "virtualserver_id":
		this.VirtualserverId, err = strconv.Atoi(v)
		if err != nil {
			this.VirtualserverId = -1
			return
		}

	case "virtualserver_antiflood_points_tick_reduce":
		this.VirtualserverAntifloodPointsTickReduce, err = strconv.Atoi(v)
		if err != nil {
			this.VirtualserverAntifloodPointsTickReduce = -1
			return
		}

	case "virtualserver_antiflood_points_needed_command_block":
		this.VirtualserverAntifloodPointsNeededCommandBlock, err = strconv.Atoi(v)
		if err != nil {
			this.VirtualserverAntifloodPointsNeededCommandBlock = -1
			return
		}

	case "virtualserver_antiflood_points_needed_ip_block":
		this.VirtualserverAntifloodPointsNeededIpBlock, err = strconv.Atoi(v)
		if err != nil {
			this.VirtualserverAntifloodPointsNeededIpBlock = -1
			return
		}

	case "virtualserver_client_connections":
		this.VirtualserverClientConnections, err = strconv.Atoi(v)
		if err != nil {
			this.VirtualserverClientConnections = -1
			return
		}

	case "virtualserver_query_client_connections":
		this.VirtualserverQueryClientConnections, err = strconv.Atoi(v)
		if err != nil {
			this.VirtualserverQueryClientConnections = -1
			return
		}

	case "virtualserver_hostbutton_tooltip":
		this.VirtualserverHostbuttonTooltip = v
	case "virtualserver_hostbutton_url":
		this.VirtualserverHostbuttonUrl = v
	case "virtualserver_hostbutton_gfx_url":
		this.VirtualserverHostbuttonGfxUrl = v
	case "virtualserver_queryclientsonline":
		this.VirtualserverQueryclientsonline, err = strconv.ParseUint(v, 10, 64)
		if err != nil {
			this.VirtualserverQueryclientsonline = 0
			return
		}

	case "virtualserver_download_quota":
		this.VirtualserverDownloadQuota, err = strconv.ParseUint(v, 10, 64)
		if err != nil {
			this.VirtualserverDownloadQuota = 0
			return
		}

	case "virtualserver_upload_quota":
		this.VirtualserverUploadQuota, err = strconv.ParseUint(v, 10, 64)
		if err != nil {
			this.VirtualserverUploadQuota = 0
			return
		}

	case "virtualserver_month_bytes_downloaded":
		this.VirtualserverMonthBytesDownloaded, err = strconv.Atoi(v)
		if err != nil {
			this.VirtualserverMonthBytesDownloaded = -1
			return
		}

	case "virtualserver_month_bytes_uploaded":
		this.VirtualserverMonthBytesUploaded, err = strconv.Atoi(v)
		if err != nil {
			this.VirtualserverMonthBytesUploaded = -1
			return
		}

	case "virtualserver_total_bytes_downloaded":
		this.VirtualserverTotalBytesDownloaded, err = strconv.ParseUint(v, 10, 64)
		if err != nil {
			this.VirtualserverTotalBytesDownloaded = 0
			return
		}

	case "virtualserver_total_bytes_uploaded":
		this.VirtualserverTotalBytesUploaded, err = strconv.ParseUint(v, 10, 64)
		if err != nil {
			this.VirtualserverTotalBytesUploaded = 0
			return
		}

	case "virtualserver_port":
		this.VirtualserverPort, err = strconv.Atoi(v)
		if err != nil {
			this.VirtualserverPort = -1
			return
		}

	case "virtualserver_autostart":
		this.VirtualserverAutostart, err = getBoolFromString(v)
	case "virtualserver_machine_id":
		this.VirtualserverMachineId, err = strconv.Atoi(v)
		if err != nil {
			this.VirtualserverMachineId = -1
			return
		}

	case "virtualserver_needed_identity_security_level":
		this.VirtualserverNeededIdentitySecurityLevel, err = strconv.Atoi(v)
		if err != nil {
			this.VirtualserverNeededIdentitySecurityLevel = -1
			return
		}

	case "virtualserver_log_client":
		this.VirtualserverLogClient, err = getBoolFromString(v)
	case "virtualserver_log_query":
		this.VirtualserverLogQuery, err = getBoolFromString(v)
	case "virtualserver_log_channel":
		this.VirtualserverLogChannel, err = getBoolFromString(v)
	case "virtualserver_log_permissions":
		this.VirtualserverLogPermissions, err = getBoolFromString(v)
	case "virtualserver_log_server":
		this.VirtualserverLogServer, err = getBoolFromString(v)
	case "virtualserver_log_filetransfer":
		this.VirtualserverLogFiletransfer, err = getBoolFromString(v)
	case "virtualserver_min_client_version":
		this.VirtualserverMinClientVersion, err = strconv.ParseUint(v, 10, 64)
		if err != nil {
			this.VirtualserverMinClientVersion = 0
			return
		}

	case "virtualserver_name_phonetic":
		this.VirtualserverNamePhonetic = v
	case "virtualserver_icon_id":
		this.VirtualserverIconId, err = strconv.Atoi(v)
		if err != nil {
			this.VirtualserverIconId = -1
			return
		}

	case "virtualserver_reserved_slots":
		this.VirtualserverReservedSlots, err = strconv.ParseUint(v, 10, 64)
		if err != nil {
			this.VirtualserverReservedSlots = 0
			return
		}

	case "virtualserver_total_packetloss_speech":
		var g float64
		g, err = strconv.ParseFloat(v, 32)
		if err != nil {
			this.VirtualserverTotalPacketlossSpeech = 0.0
			return
		}
		this.VirtualserverTotalPacketlossSpeech = float32(g)

	case "virtualserver_total_packetloss_keepalive":
		var g float64
		g, err = strconv.ParseFloat(v, 32)
		if err != nil {
			this.VirtualserverTotalPacketlossKeepalive = 0.0
			return
		}
		this.VirtualserverTotalPacketlossKeepalive = float32(g)

	case "virtualserver_total_packetloss_control":
		var g float64
		g, err = strconv.ParseFloat(v, 32)
		if err != nil {
			this.VirtualserverTotalPacketlossControl = 0.0
			return
		}
		this.VirtualserverTotalPacketlossControl = float32(g)

	case "virtualserver_total_packetloss_total":
		var g float64
		g, err = strconv.ParseFloat(v, 32)
		if err != nil {
			this.VirtualserverTotalPacketlossTotal = 0.0
			return
		}
		this.VirtualserverTotalPacketlossTotal = float32(g)

	case "virtualserver_total_ping":
		var g float64
		g, err = strconv.ParseFloat(v, 32)
		if err != nil {
			this.VirtualserverTotalPing = 0.0
			return
		}
		this.VirtualserverTotalPing = float32(g)

	case "virtualserver_ip":
		this.VirtualserverIp = v
	case "virtualserver_weblist_enabled":
		this.VirtualserverWeblistEnabled, err = getBoolFromString(v)
	case "virtualserver_ask_for_privilegekey":
		this.VirtualserverAskForPrivilegekey, err = getBoolFromString(v)
	case "virtualserver_hostbanner_mode":
		this.VirtualserverHostbannerMode, err = strconv.Atoi(v)
		if err != nil {
			this.VirtualserverHostbannerMode = -1
			return
		}

	case "virtualserver_channel_temp_delete_delay_default":
		this.VirtualserverChannelTempDeleteDelayDefault, err = strconv.Atoi(v)
		if err != nil {
			this.VirtualserverChannelTempDeleteDelayDefault = -1
			return
		}

	case "virtualserver_status":
		this.VirtualserverStatus = v
	case "connection_filetransfer_bandwidth_sent":
		this.ConnectionFiletransferBandwidthSent, err = strconv.ParseUint(v, 10, 64)
		if err != nil {
			this.ConnectionFiletransferBandwidthSent = 0
			return
		}

	case "connection_filetransfer_bandwidth_received":
		this.ConnectionFiletransferBandwidthReceived, err = strconv.ParseUint(v, 10, 64)
		if err != nil {
			this.ConnectionFiletransferBandwidthReceived = 0
			return
		}

	case "connection_filetransfer_bytes_sent_total":
		this.ConnectionFiletransferBytesSentTotal, err = strconv.Atoi(v)
		if err != nil {
			this.ConnectionFiletransferBytesSentTotal = -1
			return
		}

	case "connection_filetransfer_bytes_received_total":
		this.ConnectionFiletransferBytesReceivedTotal, err = strconv.Atoi(v)
		if err != nil {
			this.ConnectionFiletransferBytesReceivedTotal = -1
			return
		}

	case "connection_packets_sent_speech":
		this.ConnectionPacketsSentSpeech, err = strconv.Atoi(v)
		if err != nil {
			this.ConnectionPacketsSentSpeech = -1
			return
		}

	case "connection_bytes_sent_speech":
		this.ConnectionBytesSentSpeech, err = strconv.ParseUint(v, 10, 64)
		if err != nil {
			this.ConnectionBytesSentSpeech = 0
			return
		}

	case "connection_packets_received_speech":
		this.ConnectionPacketsReceivedSpeech, err = strconv.Atoi(v)
		if err != nil {
			this.ConnectionPacketsReceivedSpeech = -1
			return
		}

	case "connection_bytes_received_speech":
		this.ConnectionBytesReceivedSpeech, err = strconv.Atoi(v)
		if err != nil {
			this.ConnectionBytesReceivedSpeech = -1
			return
		}

	case "connection_packets_sent_keepalive":
		this.ConnectionPacketsSentKeepalive, err = strconv.Atoi(v)
		if err != nil {
			this.ConnectionPacketsSentKeepalive = -1
			return
		}

	case "connection_bytes_sent_keepalive":
		this.ConnectionBytesSentKeepalive, err = strconv.Atoi(v)
		if err != nil {
			this.ConnectionBytesSentKeepalive = -1
			return
		}

	case "connection_packets_received_keepalive":
		this.ConnectionPacketsReceivedKeepalive, err = strconv.Atoi(v)
		if err != nil {
			this.ConnectionPacketsReceivedKeepalive = -1
			return
		}

	case "connection_bytes_received_keepalive":
		this.ConnectionBytesReceivedKeepalive, err = strconv.Atoi(v)
		if err != nil {
			this.ConnectionBytesReceivedKeepalive = -1
			return
		}

	case "connection_packets_sent_control":
		this.ConnectionPacketsSentControl, err = strconv.Atoi(v)
		if err != nil {
			this.ConnectionPacketsSentControl = -1
			return
		}

	case "connection_bytes_sent_control":
		this.ConnectionBytesSentControl, err = strconv.Atoi(v)
		if err != nil {
			this.ConnectionBytesSentControl = -1
			return
		}

	case "connection_packets_received_control":
		this.ConnectionPacketsReceivedControl, err = strconv.Atoi(v)
		if err != nil {
			this.ConnectionPacketsReceivedControl = -1
			return
		}

	case "connection_bytes_received_control":
		this.ConnectionBytesReceivedControl, err = strconv.Atoi(v)
		if err != nil {
			this.ConnectionBytesReceivedControl = -1
			return
		}

	case "connection_packets_sent_total":
		this.ConnectionPacketsSentTotal, err = strconv.Atoi(v)
		if err != nil {
			this.ConnectionPacketsSentTotal = -1
			return
		}

	case "connection_bytes_sent_total":
		this.ConnectionBytesSentTotal, err = strconv.ParseUint(v, 10, 64)
		if err != nil {
			this.ConnectionBytesSentTotal = 0
			return
		}

	case "connection_packets_received_total":
		this.ConnectionPacketsReceivedTotal, err = strconv.Atoi(v)
		if err != nil {
			this.ConnectionPacketsReceivedTotal = -1
			return
		}

	case "connection_bytes_received_total":
		this.ConnectionBytesReceivedTotal, err = strconv.ParseUint(v, 10, 64)
		if err != nil {
			this.ConnectionBytesReceivedTotal = 0
			return
		}

	case "connection_bandwidth_sent_last_second_total":
		this.ConnectionBandwidthSentLastSecondTotal, err = strconv.Atoi(v)
		if err != nil {
			this.ConnectionBandwidthSentLastSecondTotal = -1
			return
		}

	case "connection_bandwidth_sent_last_minute_total":
		this.ConnectionBandwidthSentLastMinuteTotal, err = strconv.Atoi(v)
		if err != nil {
			this.ConnectionBandwidthSentLastMinuteTotal = -1
			return
		}

	case "connection_bandwidth_received_last_second_total":
		this.ConnectionBandwidthReceivedLastSecondTotal, err = strconv.Atoi(v)
		if err != nil {
			this.ConnectionBandwidthReceivedLastSecondTotal = -1
			return
		}

	case "connection_bandwidth_received_last_minute_total":
		this.ConnectionBandwidthReceivedLastMinuteTotal, err = strconv.Atoi(v)
		if err != nil {
			this.ConnectionBandwidthReceivedLastMinuteTotal = -1
			return
		}

	default:
		logger.Error("%s=%s is not a valid param of Serverinfo.", k, v)
		err = errors.New(k + "=" + v + " is not a valid param of Serverinfo.")
	}
	return
}
