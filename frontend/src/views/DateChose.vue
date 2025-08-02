<template>
  <div
    class="min-h-screen flex flex-col items-center justify-between"
    :style="{
      background: tgTheme.bg_color,
      color: tgTheme.text_color
    }"
  >
    <!-- Top Panel -->
    <div class="w-full max-w-md border-b p-4 mb-4"
      :style="{
      background: tgTheme.secondary_bg_color,
      color: tgTheme.text_color,
      borderColor: '#ccc'
      }"
    >
      <div class="text-lg font-bold text-center mb-2">
      Пикник планер
      </div>
      <div class="text-sm font-medium text-center">
      {{ initdata?.user?.first_name }} {{ initdata?.user?.last_name || 'Guest' }}
      <span v-if="initdata?.user?.id" class="text-xs text-gray-400">({{ initdata?.user?.id || '-52' }})</span>
      </div>
    </div>

    <!-- Calendar or Error -->
    <div class="w-full max-w-md flex-1 flex flex-col justify-center">
      <Calendar v-if="authorized" />
      <Error
        v-else
        :error="'Кажется вы не заходите не через телеграм, обновите страничку или откройте приложение через телеграм. Бюджет на эту штуку закончился(бутылка псыжа закончилась), так что нормальной штуки тут не будет'"
      />
    </div>

    <!-- Confirm Button -->
    <div class="w-full max-w-md mt-6">
      <button
        class="w-full py-3 rounded-lg font-semibold"
        :style="{
          background: tgTheme.button_color,
          color: tgTheme.text_color,
          border: '1px solid #ccc'
        }"
      >
        Подтвердить
      </button>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue';
import Calendar from '@/components/Calendar.vue';
import Error from '@/components/Error.vue';

const initdata = ref();
const authorized = ref(false);

let tg = window.Telegram.WebApp;
initdata.value = tg.initDataUnsafe;
if (tg.initDataUnsafe && tg.initDataUnsafe.user && tg.initDataUnsafe.user.id) {
  authorized.value = true;
}

// Extract Telegram theme colors with fallback
const tgTheme = computed(() => ({
  bg_color: tg.themeParams?.bg_color || '#ffffff',
  text_color: tg.themeParams?.text_color || '#222222',
  secondary_bg_color: tg.themeParams?.secondary_bg_color || '#a1a1a1',
  button_color: tg.themeParams?.button_color || '#2aabee',
}));
</script>