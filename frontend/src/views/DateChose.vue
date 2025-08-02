<template>
  <div>{{ initdata }}</div>
  <div> {{ initdata?.user?.username }}</div>
  <div v-if="authorized">
    <Calendar />
  </div>
  <div v-else>
    <Error error="Кажется вы не заходите не через телеграм, обновите страничку или откройте приложение через телеграм" />
  </div>
</template>

<script setup>
import { ref } from 'vue';
import Calendar from '@/components/Calendar.vue';
import Error from '@/components/Error.vue';

const initdata = ref();
const authorized = ref(false);

let tg = window.Telegram.WebApp;
initdata.value = tg.initDataUnsafe;
if (tg.initDataUnsafe && tg.initDataUnsafe.user && tg.initDataUnsafe.user.id) {
  authorized.value = true;
}
</script>