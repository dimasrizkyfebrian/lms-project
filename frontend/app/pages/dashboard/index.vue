<template>
  <div class="dashboard-container">
    <UCard class="w-full max-w-lg">
      <template #header>
        <div class="flex justify-between items-center">
          <h1 class="text-xl font-bold">Dashboard</h1>
          <UButton color="red" variant="soft" @click="handleLogout">
            Log Out
          </UButton>
        </div>
      </template>

      <div v-if="pending" class="space-y-4">
        <USkeleton class="h-8 w-1/2" />
        <USkeleton class="h-4 w-full" />
        <USkeleton class="h-4 w-3/4" />
      </div>

      <div v-else-if="error">
        <p class="text-red-500">
          Could not load user data. Please try logging in again.
        </p>
      </div>

      <div v-else-if="user" class="space-y-2">
        <h2 class="text-2xl">
          Welcome,
          <span class="font-bold text-primary">{{ user.username }}</span
          >!
        </h2>
        <p class="text-gray-500">
          You have successfully logged in and accessed a protected route.
        </p>
        <p>Your User ID is: {{ user.id }}</p>
      </div>
    </UCard>
  </div>
</template>

<script setup>
// Get a reference to the token cookie
const token = useCookie("token");

// useFetch will automatically run on the client side when the page loads
const {
  data: user,
  pending,
  error,
} = await useFetch("/api/v1/me", {
  baseURL: "http://localhost:8080",
  // Add the Authorization header to the request.
  // This is how the backend knows who we are.
  headers: {
    Authorization: `Bearer ${token.value}`,
  },
});

// Function to handle logging out
async function handleLogout() {
  // Clear the token cookie
  token.value = null;
  // Redirect to the login page
  await navigateTo("/login");
}
</script>

<style scoped>
.dashboard-container {
  display: flex;
  justify-content: center;
  align-items: center;
  min-height: 100vh;
  background-color: #f0f2f5;
}
</style>
