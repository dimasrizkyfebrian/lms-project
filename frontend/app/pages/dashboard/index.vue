<template>
  <div>
    <h1 class="text-3xl font-bold mb-4">Dashboard</h1>
    <div v-if="user" class="bg-white dark:bg-gray-800 p-6 rounded-lg shadow">
      <h2 class="text-2xl">
        Welcome, <span class="font-bold text-primary">{{ user.username }}</span
        >!
      </h2>
      <p class="text-gray-500 mt-2">Your User ID is: {{ user.id }}</p>
    </div>
  </div>
</template>

<script setup>
// Script ini tidak perlu diubah, karena sudah mengambil data user
const user = useUser();
const token = useCookie("token");

// Fetch the user data if it's not already present
if (!user.value) {
  await useFetch("/api/v1/me", {
    baseURL: "http://localhost:8080",
    headers: { Authorization: `Bearer ${token.value}` },
    onResponse({ response }) {
      if (response.ok) {
        user.value = response._data;
      }
    },
  });
}
</script>
