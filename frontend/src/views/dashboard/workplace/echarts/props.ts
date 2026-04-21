import { PropType } from 'vue';

export interface BasicProps {
  width: string;
  height: string;
}

export const basicProps = {
  width: {
    type: String as PropType<string>,
    default: '100%',
  },
  height: {
    type: String as PropType<string>,
    default: '340px',
  },
  login: {
    type: Array,
    default(){
      return []
    }
  },
  register: {
    type: Array,
    default(){
      return []
    }
  },
  range: {
    type: Array,
    default(){
      return []
    }
  },
  add: {
    type: Array,
    default(){
      return []
    }
  },
  active: {
    type: Array,
    default(){
      return []
    }
  }
};
