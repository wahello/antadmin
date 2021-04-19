import React from 'react';
import { createForm } from '@formily/core';
import { Field } from '@formily/react';
import {
  Form,
  FormItem,
  Input,
  Select,
  Submit,
  FormButtonGroup,
} from '@formily/antd';

export default (): React.ReactNode => {
  const form = createForm({
    validateFirst: true,
  });

  return (
    <div
      style={{
        paddingTop: '12px',
      }}
    >
      <Form
        form={form}
        labelAlign="left"
        labelCol={2}
        wrapperCol={12}
        onAutoSubmit={console.log}
      >
        <Field
          name="username"
          title="用户名"
          required
          decorator={[FormItem]}
          component={[Input]}
          required
        />
        <Field
          name="email"
          title="邮箱"
          required
          validator="email"
          decorator={[FormItem]}
          component={[Input]}
        />
        <Field
          name="phone"
          title="电话"
          required
          validator="phone"
          decorator={[FormItem]}
          component={[Input]}
        />
        <Field
          name="gender"
          title="性别"
          decorator={[FormItem]}
          component={[Select]}
          dataSource={[
            {
              label: '保密',
              value: 0,
            },
            {
              label: '男',
              value: 1,
            },
            {
              label: '女',
              value: 2,
            },
          ]}
          required
        />
        <FormButtonGroup.FormItem>
          <Submit block size="large">
            更新信息
          </Submit>
        </FormButtonGroup.FormItem>
      </Form>
    </div>
  )
}
