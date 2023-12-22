import type { Ref } from 'vue'
import type { Store } from '~/types/store'

/**
 * Creates a reactive store in Nuxt 3 with optional data persistence using cookies.
 *
 * @param {string} storeName - The name of the store.
 * @param {T} initStoreData - The initial data for the store.
 * @param {boolean} storePersist - Indicates whether the store data should be persisted in cookies.
 * @returns {object} An object containing the reactive store, the $setStore function to set store data, and the $resetStore function to reset the store.
 * @template T
 */
export default <T>(storeName: string, initStoreData: T, storePersist: boolean = false): Store<T> => {
  let storeCookie: Ref
  const initStoreDataRaw = toRaw(initStoreData)
  const EXPIRES_IN_7_DAYS = new Date(Date.now() + 1000 * 60 * 60 * 24 * 7)

  // Check if data should be persisted using cookies
  if (storePersist) {
    // Initialize a reactive cookie reference with an expiration time
    storeCookie = useCookie<T>(storeName, { expires: EXPIRES_IN_7_DAYS })
  }

  // Initialize a reactive store with a getter function
  const store = useState(storeName, (): T => {
    // Check if data should be loaded from the cookie
    if (storePersist && isDefined(storeCookie)) return storeCookie.value

    // Otherwise, use the initial data
    return initStoreDataRaw
  })

  // Watch for changes in the store and update the cookie if necessary
  if (storePersist) {
    watchDeep(store, () => {
      storeCookie.value = store.value
    })
  }

  const $setStore = (storeData: T) => {
    if (storePersist) store.value = initStoreDataRaw
    store.value = storeData
  }

  const $resetStore = () => {
    store.value = initStoreDataRaw
  }

  const $cookieExists = () => isDefined(storeCookie)

  return { store, $setStore, $resetStore, $cookieExists }
}
