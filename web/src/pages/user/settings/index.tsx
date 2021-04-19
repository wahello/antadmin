import React, { useState } from 'react';
import { GridContent } from "@ant-design/pro-layout";
import { Menu } from "antd";
import ProfileView from './profile';
import styles from './index.less';

export default (): React.ReactNode => {
  const [selectKey, setSelectKey] = useState('profile');
  const menuMap = {
    profile: '个人信息',
    password: '修改密码',
  }

  const renderChildren = () => {
    switch (selectKey) {
      case 'profile':
        return <ProfileView />;
      default:
        break;
    }
    return null;
  };

  const getMenu = () =>
    Object.keys(menuMap).map((item) => <Menu.Item key={item}>{menuMap[item]}</Menu.Item>);

  return (
    <GridContent>
      <div className={styles.main}>
        <div className={styles.leftMenu}>
          <Menu
            mode={'inline'}
            selectedKeys={[selectKey]}
            onClick={({ key }) => {
              setSelectKey(key);
            }}
          >
            {getMenu()}
          </Menu>
        </div>
        <div className={styles.right}>
          <div className={styles.title}>{menuMap[selectKey]}</div>
          {renderChildren()}
        </div>
      </div>
    </GridContent>
  );
};
