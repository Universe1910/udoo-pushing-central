import axios from 'axios'

const service = axios.create()

export function Commits(page) {
  return service({
    url: 'https://api.github.com/repos/Universe1910/udoo-pushing-central/commits?page=' + page,
    method: 'get'
  })
}

// https://github.com/Universe1910/udoo-pushing-central.git

export function Members() {
  return service({
    url: 'https://api.github.com/orgs/FLIPPED-AURORA/members',
    method: 'get'
  })
}
