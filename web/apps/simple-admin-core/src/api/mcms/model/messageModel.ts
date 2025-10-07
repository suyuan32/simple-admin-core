/**
 *  @description: Send email message request
 */
export interface SendEmailReq {
  target: string;
  subject: string;
  content: string;
  provider?: string;
}

/**
 *  @description: Send sms message request
 */
export interface SendSmsReq {
  phoneNumber: string;
  params: string;
  templateId?: string;
  appId?: string;
  signName?: string;
  provider?: string;
}
