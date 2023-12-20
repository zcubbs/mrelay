import type { Ref } from 'vue'

export type Store<T> = {
  store: Ref<T>
  $setStore: (storeData: any) => void
  $resetStore: () => void
  $cookieExists: () => boolean
}
