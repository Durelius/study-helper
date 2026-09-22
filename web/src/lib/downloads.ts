// Which episodes this browser has saved to disk.
//
// This is per-device on purpose: downloading on a laptop tells you nothing about what
// is on the phone you will actually listen on, so unlike listening progress it does
// not belong on the server.

const KEY = 'vc_downloads'

export type DownloadRecord = { topic: string; at: number; bytes: number }

export function loadDownloads(): Record<string, DownloadRecord> {
  try {
    return JSON.parse(localStorage.getItem(KEY) ?? '{}') as Record<string, DownloadRecord>
  } catch {
    return {}
  }
}

export function rememberDownload(topic: string, bytes: number): Record<string, DownloadRecord> {
  const all = loadDownloads()
  all[topic] = { topic, at: Date.now(), bytes }
  try {
    localStorage.setItem(KEY, JSON.stringify(all))
  } catch {
    /* a private window just forgets; the file is still saved */
  }
  return all
}

export function forgetDownload(topic: string): Record<string, DownloadRecord> {
  const all = loadDownloads()
  delete all[topic]
  try {
    localStorage.setItem(KEY, JSON.stringify(all))
  } catch {
    /* nothing to do */
  }
  return all
}
