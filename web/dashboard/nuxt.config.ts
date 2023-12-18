export default defineNuxtConfig({
  devtools: { enabled: true },
  modules: ['@nuxtjs/tailwindcss', 'shadcn-nuxt'],
  shadcn: {
    /**
     * Prefix for all the imported component, eg: shad-button
     */

    prefix: 'Shad',
    /**
     * Directory that the component lives in.
     * @default "./components/shad"
     */
    componentDir: './components/shad',
  },
})
