import { GithubOutlined } from '@ant-design/icons';
import { DefaultFooter } from '@ant-design/pro-layout';

export default () => (
  <DefaultFooter
    copyright="2021 AntBiz"
    links={[
      {
        key: 'antadmin',
        title: 'AntAdmin',
        href: 'https://github.com/antbiz/antadmin',
        blankTarget: true,
      },
      {
        key: 'github',
        title: <GithubOutlined />,
        href: 'https://github.com/BeanWei',
        blankTarget: true,
      },
      {
        key: 'antbiz',
        title: 'AntBiz',
        href: 'https://github.com/antbiz',
        blankTarget: true,
      },
    ]}
  />
);
