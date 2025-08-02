<template>
  <div
    class="min-h-screen flex flex-col items-center justify-between"
    :style="{
      background: tgTheme.bg_color,
      color: tgTheme.text_color
    }"
  >
    <!-- Top Panel -->
    <div 
      class="w-full max-w-md p-4 mb-4 rounded-lg shadow-sm"
      :style="{
        background: tgTheme.secondary_bg_color,
        color: tgTheme.text_color,
        borderColor: tgTheme.hint_color
      }"
    >
      <div class="text-lg font-bold text-center mb-2">
        Пикник планер
      </div>
      <div class="text-sm font-medium text-center">
        {{ userName }}
        <span v-if="userId" class="text-xs opacity-60">({{ userId }})</span>
      </div>
    </div>

    <div v-if="loading" class="w-full max-w-md flex-1 flex flex-col items-center justify-center">
      <div class="animate-spin rounded-full h-12 w-12 border-t-2 border-b-2" :style="{ borderColor: tgTheme.button_color }"></div>
      <div class="mt-4">Загрузка...</div>
    </div>

    <div v-else class="w-full max-w-md flex-1 flex flex-col justify-center">
      <Calendar
        v-if="authorized"
        v-model="selectedDates"
        :min-date="'2023-01-01T00:00:00Z'"
        :max-date="'2025-12-31T23:59:59Z'"
      />
      <Error
        v-else
        :error="'Для использования этого приложения, пожалуйста, откройте его через Telegram'"
      />
    </div>

    <div v-if="authorized" class="w-full max-w-md px-4 py-2 mb-2">
      <div class="text-center" :style="{ color: tgTheme.hint_color }">
        Выбрано дат {{ selectedDates.length}}
      </div>
      <button
        class="mt-4 w-full py-2 rounded-lg font-semibold shadow transition-colors"
        :style="{
          background: tgTheme.button_color,
          color: tgTheme.button_text_color,
          opacity: selectedDates.length > 0 ? 1 : 0.5,
          cursor: selectedDates.length > 0 ? 'pointer' : 'not-allowed'
        }"
        :disabled="selectedDates.length == 0 || loading"
        @click="confirmSelection"
      >
        Подтвердить
      </button>
    </div>

    <!-- Feedback message -->
    <div v-if="feedbackMessage" class="w-full max-w-md px-4 py-2 mb-2 text-center"
      :style="{ color: isError ? '#FF3B30' : '#34C759' }">
      {{ feedbackMessage }}
    </div>
  </div>
</template>

<script setup>
import { ref, computed, onMounted } from 'vue';
import Calendar from '@/components/Calendar.vue';
import Error from '@/components/Error.vue';

const initdata = ref(null);
const authorized = ref(false);
const loading = ref(true);
const selectedDate = ref(null);
const selectedDates = ref([]);
const feedbackMessage = ref('');
const isError = ref(false);

// Telegram WebApp instance
const tg = window.Telegram?.WebApp;

onMounted(() => {
  // Initialize Telegram WebApp
  if (tg) {
    tg.ready();
    tg.expand();
    initdata.value = tg.initDataUnsafe;
    if (initdata.value?.user?.id) {
      authorized.value = true;
    }
  }
  loading.value = false;
});

// Computed properties for user info
const userName = computed(() => {
  if (!initdata.value?.user) return 'Гость';
  return `${initdata.value.user.first_name} ${initdata.value.user.last_name || ''}`.trim();
});

const userId = computed(() => initdata.value?.user?.id || null);

// Extract Telegram theme colors with fallback
const tgTheme = computed(() => ({
  bg_color: tg?.themeParams?.bg_color || '#ffffff',
  text_color: tg?.themeParams?.text_color || '#222222',
  secondary_bg_color: tg?.themeParams?.secondary_bg_color || '#f2f2f2',
  button_color: tg?.themeParams?.button_color || '#2aabee',
  button_text_color: tg?.themeParams?.button_text_color || '#ffffff',
  hint_color: tg?.themeParams?.hint_color || '#999999',
}));

// Handle date selection
function handleDateSelection(date) {
  selectedDate.value = date;
}

// Format date for display
function formatDate(date) {
  if (!date) return '';
  return new Date(date).toLocaleDateString('ru-RU', {
    day: 'numeric',
    month: 'long',
    year: 'numeric'
  });
}

// Handle confirmation (send to backend, then close WebApp)
async function confirmSelection() {
  if (!selectedDate.value) {
    showFeedback('Пожалуйста, выберите дату', true);
    return;
  }
  loading.value = true;
  try {
    // Отправка данных на бэкенд
    const response = await fetch('/api/choose-date', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify({
        date: selectedDate.value,
        user_id: userId.value
      })
    });
    if (!response.ok) throw new Error('Ошибка ответа сервера');
    showFeedback('Дата подтверждена!');
    setTimeout(() => {
      tg?.close();
    }, 1000);
  } catch (error) {
    showFeedback('Ошибка при отправке данных', true);
    console.error(error);
  } finally {
    loading.value = false;
  }
}

// Show feedback message
function showFeedback(message, error = false) {
  feedbackMessage.value = message;
  isError.value = error;
  
  setTimeout(() => {
    feedbackMessage.value = '';
  }, 3000);
}
</script>