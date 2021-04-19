export default [
  {
    path: '/user',
    layout: false,
    routes: [
      {
        path: '/user',
        routes: [
          {
            name: '登录',
            path: '/user/signin',
            component: './user/signin',
          },
        ],
      },
    ],
  },
  {
    path: '/account',
    hideInMenu: true,
    routes: [
      {
        name: '个人设置',
        path: '/account/settings',
        component: './user/settings',
      },
    ]
  },
  {
    name: '欢迎',
    path: '/welcome',
    icon: 'smile',
    component: './Welcome',
  },
  {
    path: '/',
    redirect: '/welcome',
  },
  {
    component: './404',
  },
];
