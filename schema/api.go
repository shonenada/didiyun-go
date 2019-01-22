package schema

const API_BASE = "https://open.didiyunapi.com/dicloud/i"
const LIST_REGION_URL = API_BASE + "/region/list"

const GET_DC2_URL = API_BASE + "/compute/dc2"
const LIST_DC2_URL = API_BASE + "/compute/dc2/list"
const COUNT_DC2_URL = API_BASE + "/compute/dc2/count"
const CREATE_DC2_URL = API_BASE + "/compute/dc2/assign"
const DELETE_DC2_URL = API_BASE + "/compute/dc2/delete"
const START_DC2_URL = API_BASE + "/compute/dc2/start"
const STOP_DC2_URL = API_BASE + "/compute/dc2/stop"
const REBOOT_DC2_URL = API_BASE + "/compute/dc2/reboot"
const REINSTALL_DC2_URL = API_BASE + "/compute/dc2/reinstall"
const CHANGE_PASSWORD_DC2_URL = API_BASE + "/compute/dc2/changePassword"
const CHANGE_NAME_DC2_URL = API_BASE + "/compute/dc2/changeName"
const CHANGE_SPEC_DC2_URL = API_BASE + "/compute/dc2/changeSpec"

const LIST_IMAGE_URL = API_BASE + "/image/list"

const LIST_SSHKEY_URL = API_BASE + "/sshkeys/list"
const CREATE_SSHKEY_URL = API_BASE + "/sshkeys/create"
const DELETE_SSHKEY_URL = API_BASE + "/sshkeys/delete"

const GET_EIP_URL = API_BASE + "/network/eip"
const LIST_EIP_URL = API_BASE + "/network/eip/list"
const COUNT_EIP_URL = API_BASE + "/network/eip/count"
const CREATE_EIP_URL = API_BASE + "/network/eip/assign"
const ATTACH_EIP_URL = API_BASE + "/network/eip/attach"
const DEATCH_EIP_URL = API_BASE + "/network/eip/deatch"
const DELETE_EIP_URL = API_BASE + "/network/eip/delete"
const CHANGE_BANDWIDTH_EIP_URL = API_BASE + "/network/eip/changeBandwidth"

const GET_EBS_URL = API_BASE + "/storage/ebs"
const LIST_EBS_URL = API_BASE + "/storage/ebs/list"
const COUNT_EBS_URL = API_BASE + "/storage/ebs/count"
const CREATE_EBS_URL = API_BASE + "/storage/ebs/assign"
const DELETE_EBS_URL = API_BASE + "/storage/ebs/delete"
const ATTACH_EBS_URL = API_BASE + "/storage/ebs/attach"
const DEATCH_EBS_URL = API_BASE + "/storage/ebs/deatch"
const CHANGE_NAME_EBS_URL = API_BASE + "/storage/ebs/changeName"
const CHANGE_SIZE_EBS_URL = API_BASE + "/storage/ebs/changeSize"

const LIST_SNAP_URL = API_BASE + "/storage/snapshot/list"
