<template>
  <div class="flex items-center justify-center min-h-screen">
    <UCard class="w-full max-w-sm">
      <template #header>
        <h1 class="text-2xl font-bold text-center">Log In to Your Account</h1>
      </template>

      <UForm :state="state" @submit="submit">
        <UFormField label="Username" name="username" class="mb-4">
          <UInput
            v-model="state.username"
            placeholder="Enter your username"
            class="w-full"
          />
        </UFormField>

        <UFormField label="Password" name="password">
          <UInput
            v-model="state.password"
            type="password"
            placeholder="Enter your password"
            class="w-full"
          />
        </UFormField>

        <UButton type="submit" class="mt-6 w-full" block :loading="isLoading">
          Log In
        </UButton>
      </UForm>

      <template #footer>
        <p class="text-sm text-center text-gray-500">
          Don't have an account?
          <NuxtLink to="/register" class="text-primary font-medium"
            >Register here</NuxtLink
          >
        </p>
      </template>
    </UCard>
  </div>
</template>

<script setup>
import { ref } from "vue";

// `useToast` for showing notifications
const toast = useToast();
// `useCookie` to store our auth token
const token = useCookie("token");
// Loading state for the button
const isLoading = ref(false);

// Reactive state object for our form inputs
const state = ref({
  username: "",
  password: "",
});

// Function that runs when the form is submitted
async function submit() {
  isLoading.value = true;
  try {
    // Use $fetch to send a POST request to our Go backend
    const response = await $fetch("/api/v1/login", {
      method: "POST",
      baseURL: "http://localhost:8080", // Backend URL
      body: {
        username: state.value.username,
        password: state.value.password,
      },
    });

    // Store the token in a cookie.
    // This makes it available across the entire app and for future requests.
    token.value = response.token;

    toast.add({ title: "Login successful!" });

    // Redirect to the dashboard page
    await navigateTo("/dashboard");
  } catch (error) {
    // Show an error notification if the request fails
    toast.add({
      title: "Login Failed",
      description: error.data?.error || "Invalid username or password.",
      color: "error",
    });
  } finally {
    isLoading.value = false;
  }
}
</script>
