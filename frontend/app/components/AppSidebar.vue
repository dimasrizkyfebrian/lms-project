<template>
  <aside class="w-64 flex-shrink-0 bg-gray-800 text-gray-200 flex flex-col">
    <div
      class="h-16 flex items-center justify-center text-xl font-bold border-b border-gray-700"
    >
      LMS Platform
    </div>

    <nav class="flex-1 px-4 py-6 space-y-2">
      <NuxtLink
        v-for="link in navigation"
        :key="link.to"
        :to="link.to"
        class="flex items-center px-4 py-2 rounded-md text-sm font-medium hover:bg-gray-700"
        active-class="bg-primary-500 text-white"
      >
        <Icon :name="link.icon" class="mr-3 h-5 w-5" />
        <span>{{ link.label }}</span>
      </NuxtLink>
    </nav>

    <div v-if="user" class="p-4 border-t border-gray-700">
      <div class="flex items-center">
        <UAvatar size="sm" icon="i-heroicons-user-circle" />
        <span class="ml-2 font-medium">{{ user.username }}</span>
      </div>
    </div>
  </aside>
</template>

<script setup>
const user = useUser();

// Navigation links with Lucide icon names
const navigation = [
  { label: "Dashboard", icon: "i-lucide-layout-dashboard", to: "/dashboard" },
  { label: "Courses", icon: "i-lucide-book-open", to: "/courses" },
  { label: "Profile", icon: "i-lucide-user", to: "/profile" },
];
</script>
