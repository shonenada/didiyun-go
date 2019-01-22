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
