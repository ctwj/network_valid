import {h} from 'vue';
import {NTime} from 'naive-ui';



export const roleColumns = [
  {
    title: '#',
    key: 'id',
    width: 100,
  },
  {
    title: '角色名称',
    key: 'title',
    width: 100,
  },
  {
    title: '描述',
    key: 'description',
    width: 150,
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

