<template>
  <div class="tz-demo-wrap">
    <div class="dp-container-wrap">
      <VueDatePicker 
        v-model="dates" 
        :dark="isDark"
        :timezone="tz" 
        inline 
        auto-apply
        :max-date="maxDate"
        :enable-time-picker="false"
        multi-dates
      />
    </div>
    <div class="tz-range-slider-wrap">
      <div>
        <span>Timezone: {{ activeTz.tz }}</span>
        <br />
        <span>Offset: {{ activeTz.offset > 0 ? `+${activeTz.offset}` : activeTz.offset }}</span>
      </div>
      <div>
        <input class="tz-range-slider" type="range" v-model="selectedTz" min="0" max="22" />
      </div>
    </div>
    <div>
      <label>
        <input type="checkbox" v-model="isDark" />
        Dark Mode
      </label>
    </div>
    <div class="selected-dates-list mt-4">
      <h3>Selected Dates:</h3>
      <ul>
        <li v-for="(date, idx) in dates" :key="idx">
          {{ date }}
        </li>
      </ul>
    </div>
  </div>
</template>

<script setup>
import { ref, computed } from 'vue';
import VueDatePicker from '@vuepic/vue-datepicker';
import '@vuepic/vue-datepicker/dist/main.css'
import { getMonth, getYear } from "date-fns";

const dates = ref([]);
const selectedTz = ref(11);

const timezone = ref({ timezone: undefined })
const isDark = ref(false);

const maxDate = computed(() => {
  const month = getMonth(new Date()) + 1 > 9 ? getMonth(new Date()) + 1 : `0${getMonth(new Date()) + 1}`;
  return `${getYear(new Date())}-${month}-15T01:00:00Z`;
});

const timezones = [
  { tz: 'Pacific/Midway', offset: -11 },
  { tz: 'America/Adak', offset: -10 }, 
  { tz: 'Pacific/Gambier', offset: -9 }, 
  { tz: 'America/Los_Angeles', offset: -8 }, 
  { tz: 'America/Denver', offset: -7 }, 
  { tz: 'America/Chicago', offset: -6 }, 
  { tz: 'America/New_York', offset: -5 }, 
  { tz: 'America/Santiago', offset: -4 }, 
  { tz: 'America/Sao_Paulo', offset: -3 }, 
  { tz: 'America/Noronha', offset: -2 }, 
  { tz: 'Atlantic/Cape_Verde', offset: -1 }, 
  { tz: 'UTC', offset: 0 },
  { tz: 'Europe/Brussels', offset: 1 }, 
  { tz: 'Africa/Cairo', offset: 2 }, 
  { tz: 'Europe/Minsk', offset: 3 }, 
  { tz: 'Europe/Moscow', offset: 4 },
  { tz: 'Asia/Tashkent', offset: 5 },
  { tz: 'Asia/Dhaka', offset: 6 },
  { tz: 'Asia/Novosibirsk', offset: 7 },
  { tz: 'Australia/Perth', offset: 8 }, 
  { tz: 'Asia/Tokyo', offset: 9 },
  { tz: 'Australia/Hobart', offset: 10 },
  { tz: 'Asia/Vladivostok', offset: 11 },
];

const activeTz = computed(() => timezones[selectedTz.value]);

const tz = computed(() => {
  return { ...timezone.value, timezone: activeTz.value.tz };
});
</script>