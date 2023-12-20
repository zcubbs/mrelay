export default () => {
  // Composable
  const { locale } = useI18n()

  const getCurrentTime = (format: string) => useDateFormat(useNow(), format, { locales: locale.value }).value

  const getTimestamp = () => useTimestamp().value

  const formatTime = (time: string, format: string) => useDateFormat(time, format, { locales: locale.value }).value

  return { getCurrentTime, getTimestamp, formatTime }
}
