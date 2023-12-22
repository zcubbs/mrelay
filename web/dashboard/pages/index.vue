<template>
  <div class="flex flex-col gap-y-5">
    <div class="grid grid-cols-1 gap-4 md:grid-cols-2 lg:grid-cols-6">
      <shad-card class="border-primary bg-muted shadow dark:border-border dark:bg-background">
        <shad-card-header>
          <div class="flex justify-between">
            <h4 class="text-lg">{{ $t('stats.card.totalEmailSend.title') }}</h4>
            <PhosphorIconPulse size="20" />
          </div>
        </shad-card-header>
        <shad-card-content>
          <p class="text-lg font-bold">827</p>
        </shad-card-content>
      </shad-card>
      <shad-card class="shadow">
        <shad-card-header>
          <div class="flex justify-between">
            <div>
              <h4 class="text-lg">{{ $t('stats.card.topService.title') }}</h4>
            </div>
            <PhosphorIconShootingStar size="20" />
          </div>
        </shad-card-header>
        <shad-card-content>
          <div class="flex justify-between">
            <p>Alerts BL.Predict</p>
            <p class="text-lg font-bold">418</p>
          </div>
        </shad-card-content>
      </shad-card>
    </div>
    <div class="grid grid-cols-1 gap-4 lg:grid-cols-3">
      <ui-stats-card-emails-overview />
      <ui-stats-card-emails-count />
    </div>
  </div>
</template>
