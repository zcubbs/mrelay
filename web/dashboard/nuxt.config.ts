export default defineNuxtConfig({
  devtools: { enabled: true },
  modules: ['@nuxtjs/tailwindcss'],
  components: [
    {
      path: '~/components/shad',
      // this is required else Nuxt will autoImport `.ts` file
      extensions: ['.vue'],
      // prefix for your components, eg: shad-button
      prefix: 'Shad',
    },
  ],
})
