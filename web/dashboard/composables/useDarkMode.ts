export const useDarkMode = () => {
  const { store, $setStore } = defineStore('dark_mode', false, true)

  // Mutations
  const setDarkMode = (darkMode: boolean) => $setStore(darkMode)

  return { isDarkMode: store, setDarkMode }
}
