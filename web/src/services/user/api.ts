// @ts-ignore
/* eslint-disable */
import { request } from 'umi';

/** 获取当前的用户 GET /api/user/profile */
export async function getCurrentUser(options?: { [key: string]: any }) {
  return request<API.CurrentUser>('/api/user/profile', {
    method: 'GET',
    ...(options || {}),
  });
}

/** 登录接口 POST /api/user/signin */
export async function signin(options?: { [key: string]: any }) {
  return request<Record<string, any>>('/api/user/signin', {
    method: 'POST',
    ...(options || {}),
  });
}

/** 登出接口 POST /api/user/signout */
export async function signout(options?: { [key: string]: any }) {
  return request<Record<string, any>>('/api/user/signout', {
    method: 'POST',
    ...(options || {}),
  });
}

/** 发送验证码 POST /api/login/captcha */
export async function getFakeCaptcha(
  params: {
    // query
    /** 手机号 */
    phone?: string;
  },
  options?: { [key: string]: any },
) {
  return request<API.FakeCaptcha>('/api/login/captcha', {
    method: 'POST',
    params: {
      ...params,
    },
    ...(options || {}),
  });
}

/** 此处后端没有提供注释 GET /api/notices */
export async function getNotices(options?: { [key: string]: any }) {
  return request<API.NoticeIconList>('/api/notices', {
    method: 'GET',
    ...(options || {}),
  });
}
