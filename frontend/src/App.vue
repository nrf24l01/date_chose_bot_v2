<template>
  <router-view />
</template>

<script setup>
import { onMounted } from 'vue'

onMounted(() => {
  const theme = window?.Telegram?.WebApp?.themeParams || {}
  const root = document.body.style

  // Применяем основные цвета, если они переданы
  if (theme.bg_color) root.setProperty('--tg-bg-color', theme.bg_color)
  if (theme.text_color) root.setProperty('--tg-text-color', theme.text_color)
  if (theme.hint_color) root.setProperty('--tg-hint-color', theme.hint_color)
  if (theme.link_color) root.setProperty('--tg-link-color', theme.link_color)
  if (theme.button_color) root.setProperty('--tg-button-color', theme.button_color)
  if (theme.button_text_color) root.setProperty('--tg-button-text-color', theme.button_text_color)
  // Добавьте другие цвета по необходимости
})
</script>

<style>
body {
  background: var(--tg-bg-color, #fff);
  color: var(--tg-text-color, #222);
}
a {
  color: var(--tg-link-color, #2aabee);
}
button {
  background: var(--tg-button-color, #2aabee);
  color: var(--tg-button-text-color, #fff);
}
::placeholder {
  color: var(--tg-hint-color, #999);
}
</style>