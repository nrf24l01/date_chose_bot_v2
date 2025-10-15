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
        Даты выбиралка
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
        :min-date="minDate"
        :max-date="maxDate"
        :disabled="selectedDates[0] === unavailableValue"
      />
      <Error
        v-else
        :error="'Для использования этого приложения, пожалуйста, откройте его через Telegram. ' + minDate + ' - ' + maxDate"
      />
    </div>

    <div v-if="authorized" class="w-full max-w-md px-4 py-2 mb-2">
      <div class="text-center" :style="{ color: tgTheme.hint_color }">
        <template v-if="selectedDates[0] !== unavailableValue">
          Выбрано дат {{ selectedDates.length }}
        </template>
        <template v-else>
          Вы не можете пойти :(
        </template>
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
      <!-- Switcher button for unavailable -->
      <button
        class="mt-2 w-full py-2 rounded-lg font-semibold shadow transition-colors border flex items-center justify-center"
        :style="{
          background: selectedDates[0] === unavailableValue ? tgTheme.button_color : tgTheme.secondary_bg_color,
          color: selectedDates[0] === unavailableValue ? tgTheme.button_text_color : tgTheme.button_color,
          borderColor: tgTheme.button_color,
          opacity: 1,
        }"
        type="button"
        @click="toggleUnavailable"
      >
        <span>
          {{ selectedDates[0] === unavailableValue ? 'Я не смогу попасть' : 'Я смогу попасть' }}
        </span>
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
const unavailableValue = '0001-01-01';

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
      getPreviousDates();
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

function safe(themeValue, fallback) {
  return themeValue && typeof themeValue === 'string' && themeValue.trim() !== '' ? themeValue : fallback;
}

// Fetch interval from backend (no auth required)
const minDate = ref(null);
const maxDate = ref(null);

async function fetchInterval() {
  try {
    const res = await fetch(import.meta.env.VITE_BACKEND_URL + '/date/interval', {
      method: 'GET',
      headers: {
        'Content-Type': 'application/json'
      }
    });

    if (!res.ok) {
      const text = await res.text();
      showFeedback(`Ошибка ответа сервера (interval): ${res.status} ${text}`, true);
      return;
    }

    const data = await res.json();
    if (data.start_date && data.end_date) {
      // keep same format as previous env-based values (ISO datetimes)
      minDate.value = data.start_date + 'T00:00:00Z';
      maxDate.value = data.end_date + 'T23:59:59Z';
    } else {
      showFeedback('Неверный формат данных интервала от сервера', true);
    }
  } catch (err) {
    showFeedback(`Ошибка при получении интервала: ${err?.message || err}`, true);
    console.error(err);
  }
}

// Fetch immediately (no auth gating)
fetchInterval();

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
  if (selectedDates.length === 0) {
    showFeedback('Пожалуйста, выберите дату', true);
    return;
  }
  loading.value = true;
  try {
    const response = await fetch(import.meta.env.VITE_BACKEND_URL + '/date/choice', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
        "Authorization": `tma ${tg.initData}`
      },
      body: JSON.stringify({
        dates: selectedDates.value.map(date => {
          if (date === unavailableValue) return unavailableValue;
          const d = typeof date === 'string' ? new Date(date) : date;
          const gmt3 = new Date(d.getTime() + 3 * 60 * 60 * 1000);
          return gmt3.toISOString().slice(0, 10);
        })
      })
    });
    if (!response.ok) {
      const errorText = await response.text();
      showFeedback(`Ошибка ответа сервера: ${response.status} ${errorText}`, true);
      throw `Ошибка ответа сервера: ${response.status} ${errorText}`;
    }
    showFeedback('Дата подтверждена!');
    } catch (error) {
    showFeedback(`Ошибка при отправке данных: ${error?.message || error}`, true);
    console.error(error);
    } finally {
    loading.value = false;
    }
}

// Mark/unmark as unavailable (switcher)
function toggleUnavailable() {
  if (selectedDates.value[0] === unavailableValue) {
    selectedDates.value = [];
  } else {
    selectedDates.value = [unavailableValue];
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

async function getPreviousDates() {
  const response = await fetch(import.meta.env.VITE_BACKEND_URL + '/date/choiced', {
    method: 'GET',
    headers: {
      'Content-Type': 'application/json',
      "Authorization": `tma ${tg.initData}`
    }
  });
  if (!response.ok) {
    const errorText = await response.text();
    showFeedback(`Ошибка ответа сервера: ${response.status} ${errorText}`, true);
    throw `Ошибка ответа сервера: ${response.status} ${errorText}`;
  }
  const data = await response.json();
  selectedDates.value = Array.isArray(data.dates)
    ? data.dates.map(date =>
        date === unavailableValue
          ? unavailableValue
          : new Date(date + 'T00:00:00+03:00')
      )
    : [];
}
</script>

<style>
.dp__theme_light {
  --dp-background-color: var(--tg-theme-bg-color, #fff);
  --dp-text-color: var(--tg-theme-text-color, #212121);
  --dp-hover-color: var(--tg-theme-secondary-bg-color, #f3f3f3);
  --dp-hover-text-color: var(--tg-theme-text-color, #212121);
  --dp-hover-icon-color: var(--tg-theme-hint-color, #959595);
  --dp-primary-color: var(--tg-theme-button-color, #1976d2);
  --dp-primary-disabled-color: var(--tg-theme-hint-color, #6bacea);
  --dp-primary-text-color: var(--tg-theme-button-text-color, #f8f5f5);
  --dp-secondary-color: var(--tg-theme-secondary-bg-color, #c0c4cc);
  --dp-border-color: var(--tg-theme-bg-color, #fff);
  --dp-menu-border-color: var(--tg-theme-bg-color, #fff);
  --dp-border-color-hover: var(--tg-theme-bg-color, #fff);
  --dp-border-color-focus: var(--tg-theme-bg-color, #fff);
  --dp-disabled-color: var(--tg-theme-secondary-bg-color, #f6f6f6);
  --dp-scroll-bar-background: var(--tg-theme-secondary-bg-color, #f3f3f3);
  --dp-scroll-bar-color: var(--tg-theme-hint-color, #959595);
  --dp-success-color: #76d275;
  --dp-success-color-disabled: #a3d9b1;
  --dp-icon-color: var(--tg-theme-hint-color, #959595);
  --dp-danger-color: var(--tg-theme-hint-color, #ff6f60);
  --dp-marker-color: var(--tg-theme-button-color, #ff6f60);
  --dp-tooltip-color: var(--tg-theme-secondary-bg-color, #fafafa);
  --dp-disabled-color-text: var(--tg-theme-hint-bg-color, #8e8e8e);
  --dp-highlight-color: color-mix(in srgb, var(--tg-theme-button-color, #1976d2) 10%, transparent);
  --dp-range-between-dates-background-color: var(--tg-theme-secondary-bg-color, #f3f3f3);
  --dp-range-between-dates-text-color: var(--tg-theme-text-color, #212121);
  --dp-range-between-border-color: var(--tg-theme-secondary-bg-color, #f3f3f3);
}
</style>