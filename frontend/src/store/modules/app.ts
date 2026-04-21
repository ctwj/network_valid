import {defineStore} from 'pinia';
import {store} from '@/store';
import {cardList, getAgent, projectList} from "@/api";
import {isArray} from "@/utils/is";


export interface IAppState {
  projectList: Array<any>,
  cardList: Array<any>,
  agentList: Array<any>
}

export const useAppStore = defineStore({
  id: 'app-app',
  state: (): IAppState => ({
    projectList: [],
    cardList: [],
    agentList: []
  }),
  getters: {
    getCardList(): Array<any> {
      return this.cardList;
    },
    getProjectList(): Array<any> {
      return this.projectList;
    },
    getAgentList(): Array<any>{
      return this.agentList
    }
  },
  actions: {
    setProjectList(data: Array<any>) {
      this.projectList = data
    },
    async fetchProjectList(update: Boolean) {
      if (update) {
        let {data} = await projectList()
        if (data !== null) {
          this.setProjectList(data)
        }
        return Promise.resolve(null);
      }
      if (isArray(this.projectList)) {
        if (this.projectList.length === 0) {
          let {data} = await projectList()
          if (data !== null) {
            this.setProjectList(data)
          }
        }
      }
      return Promise.resolve(null);
    },
    setCardList(data: Array<any>) {
      this.cardList = data
    },
    async fetchCardList(update: Boolean) {
      if (update) {

        let {data} = await cardList()
        if (data !== null) {
          this.setCardList(data)
        }

        return Promise.resolve(null);
      }
      if (isArray(this.cardList)) {
        if (this.cardList.length === 0) {
          let {data} = await cardList()
          if (data !== null) {
            this.setCardList(data)
          }
        }
      }
      return Promise.resolve(null);
    },
    setAgentList(data: Array<any>) {
      this.agentList = data
    },
    async fetchAgentList(update: Boolean) {
      if (update) {

        let data = await getAgent()
        if (data !== null) {
          this.setAgentList(data)
        }

        return Promise.resolve(null);
      }
      if (isArray(this.cardList)) {
        if (this.cardList.length === 0) {
          let data = await getAgent()
          if (data !== null) {
            this.setAgentList(data)
          }
        }
      }
      return Promise.resolve(null);
    },
  },
});

// Need to be used outside the setup
export function useUserStoreWidthOut() {
  return useAppStore(store);
}
