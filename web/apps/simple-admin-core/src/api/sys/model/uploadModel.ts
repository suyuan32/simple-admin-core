export interface UploadInfo {
  name: string;
  url: string;
}

export interface UploadApiResp {
  msg: string;
  code: number;
  data: UploadInfo;
}
