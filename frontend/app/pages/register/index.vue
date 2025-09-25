<template>
  <div class="flex items-center justify-center min-h-screen">
    <UCard class="w-full max-w-sm">
      <template #header>
        <h1 class="text-2xl font-bold text-center">Create an Account</h1>
      </template>

      <UForm :state="state" @submit="submit">
        <UFormField label="Username" name="username" class="mb-4">
          <UInput
            v-model="state.username"
            placeholder="Choose a username"
            class="w-full"
          />
        </UFormField>

        <UFormField label="Password" name="password">
          <UInput
            v-model="state.password"
            type="password"
            placeholder="Create a password"
            class="w-full"
          />
        </UFormField>

        <UButton type="submit" class="mt-6 w-full" block> Register </UButton>
      </UForm>

      <template #footer>
        <p class="text-sm text-center text-gray-500">
          Already have an account?
          <NuxtLink to="/login" class="text-primary font-medium"
            >Log in</NuxtLink
          >
        </p>
      </template>
    </UCard>
  </div>
</template>

<script setup>
// Import ref for reactive state and navigateTo for redirection
import { ref } from "vue";

// `useToast` is a composable from Nuxt UI for showing notifications
const toast = useToast();

definePageMeta({
  layout: "auth",
});

// Reactive state object for our form inputs
const state = ref({
  username: "",
  password: "",
});

// Function that runs when the form is submitted
async function submit() {
  try {
    // Use $fetch to send a POST request to our Go backend
    const response = await $fetch("/api/v1/register", {
      method: "POST",
      baseURL: "http://localhost:8080", // Backend URL
      body: {
        username: state.value.username,
        password: state.value.password,
      },
    });

    // Show a success notification
    toast.add({
      title: "Success!",
      description: "Your account has been created. Please log in.",
      color: "primary",
    });

    // Redirect to the login page after a short delay
    setTimeout(() => {
      navigateTo("/login");
    }, 1500);
  } catch (error) {
    // Show an error notification if the request fails
    toast.add({
      title: "Registration Failed!",
      description: error.data?.error || "An unknown error occurred.",
      color: "error",
    });
  }
}
</script>
