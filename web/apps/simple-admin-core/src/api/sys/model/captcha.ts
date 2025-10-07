export interface CaptchaResp {
  captchaId: string;
  imgPath: string;
}

export interface GetEmailCaptchaReq {
  email: string;
}

export interface GetSmsCaptchaReq {
  phoneNumber: string;
}
