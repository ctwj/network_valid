import {h} from 'vue';
import {NTime} from 'naive-ui';



export const agentColumns = [
  {
    title: '#',
    key: 'id',
    width: 100,
  },
  {
    title: '用户名',
    key: 'user',
    width: 100,
  },
   {
    title: '邮箱',
    key: 'email',
    width: 100,
  },
  {
    title: '余额',
    key: 'money',
    width: 100,
  },
  {
    title: '创建时间',
    key: 'create_time',
    width: 180,
    render(row) {
      return h(
        NTime, {
          time: new Date(row.create_time),
          type: 'datetime'
        })
    }
  },
];

