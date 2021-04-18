import {
  AlipayCircleOutlined,
  LockOutlined,
  TaobaoCircleOutlined,
  UserOutlined,
  WeiboCircleOutlined,
} from '@ant-design/icons';
import { Alert, Space, message } from 'antd';
import React, { useState } from 'react';
import ProForm, { ProFormCheckbox, ProFormText } from '@ant-design/pro-form';
import { Link, history, useModel } from 'umi';
import Footer from '@/components/Footer';
import { signin } from '@/services/user/api';

import styles from './index.less';

const SigninMessage: React.FC<{
  content: string;
}> = ({ content }) => (
  <Alert
    style={{
      marginBottom: 24,
    }}
    message={content}
    type="error"
    showIcon
  />
);

/** 此方法会跳转到 redirect 参数所在的位置 */
const goto = () => {
  if (!history) return;
  setTimeout(() => {
    const { query } = history.location;
    const { redirect } = query as { redirect: string };
    history.push(redirect || '/');
  }, 10);
};

const Signin: React.FC = () => {
  const [submitting, setSubmitting] = useState(false);
  const [userSigninState, setUserSigninState] = useState<API.SigninResult>({});
  const { initialState, setInitialState } = useModel('@@initialState');

  const fetchUserInfo = async () => {
    const userInfo = await initialState?.fetchUserInfo?.();
    if (userInfo) {
      setInitialState({
        ...initialState,
        currentUser: userInfo,
      });
    }
  };

  const handleSubmit = async (values: API.SigninParams) => {
    setSubmitting(true);
    try {
      // 登录
      const msg = await signin({ ...values });
      if (msg.status === 'ok') {
        message.success('登录成功！');
        await fetchUserInfo();
        goto();
        return;
      }
      // 如果失败去设置用户错误信息
      setUserSigninState(msg);
    } catch (error) {
      message.error('登录失败，请重试！');
    }
    setSubmitting(false);
  };
  const { status } = userSigninState;

  return (
    <div className={styles.container}>
      <div className={styles.content}>
        <div className={styles.top}>
          <div className={styles.header}>
            <Link to="/">
              <img alt="logo" className={styles.logo} src="/logo.svg" />
              <span className={styles.title}>AntAdmin</span>
            </Link>
          </div>
          <div className={styles.desc}>🚀 通用的后台管理系统</div>
        </div>

        <div className={styles.main}>
          <ProForm
            initialValues={{
              autoSignin: true,
            }}
            submitter={{
              searchConfig: {
                submitText: '登录',
              },
              render: (_, dom) => dom.pop(),
              submitButtonProps: {
                loading: submitting,
                size: 'large',
                style: {
                  width: '100%',
                },
              },
            }}
            onFinish={async (values) => {
              handleSubmit(values as API.SigninParams);
            }}
          >
            {status === 'error' && (
              <SigninMessage
                content="账户或密码错误"
              />
            )}
            <ProFormText
              name="username"
              fieldProps={{
                size: 'large',
                prefix: <UserOutlined className={styles.prefixIcon} />,
              }}
              placeholder="用户名: "
              rules={[
                {
                  required: true,
                  message: "请输入用户名!",
                },
              ]}
            />
            <ProFormText.Password
              name="password"
              fieldProps={{
                size: 'large',
                prefix: <LockOutlined className={styles.prefixIcon} />,
              }}
              placeholder="密码: "
              rules={[
                {
                  required: true,
                  message: "请输入密码！",
                },
              ]}
            />
            <div
              style={{
                marginBottom: 24,
              }}
            >
              <ProFormCheckbox noStyle name="autoSignin">自动登录</ProFormCheckbox>
              <a
                style={{
                  float: 'right',
                }}
              >
                忘记密码
              </a>
            </div>
          </ProForm>
          <Space className={styles.other}>
            其他登录方式
            <AlipayCircleOutlined className={styles.icon} />
            <TaobaoCircleOutlined className={styles.icon} />
            <WeiboCircleOutlined className={styles.icon} />
          </Space>
        </div>
      </div>
      <Footer />
    </div>
  );
};

export default Signin;
