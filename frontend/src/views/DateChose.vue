<template>
  <div>{{ initdata }}</div>
  <div> {{ initdata?.user?.username }}</div>
  <div>
    <Calendar />
  </div>
</template>

<script setup>
import { ref } from 'vue';
import Calendar from '@/components/Calendar.vue';

const initdata = ref();

let tg = window.Telegram.WebApp;
initdata.value = tg.initDataUnsafe;
</script>