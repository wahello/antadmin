import React from 'react';
import { PageContainer } from '@ant-design/pro-layout';
import { Card, Alert } from 'antd';

export default (): React.ReactNode => {
  return (
    <PageContainer>
      <Card>
        <Alert
          message="AntAdmin"
          type="success"
          showIcon
          banner
          style={{
            margin: -12,
            marginBottom: 24,
          }}
        />
      </Card>
    </PageContainer>
  );
};
