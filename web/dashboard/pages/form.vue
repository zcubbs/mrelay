<script setup lang="ts">
import { FormField } from '~/components/shad/form'
import { useToast } from '~/components/shad/toast/use-toast'

// Utils
const { z, createForm } = defineForm()

const { handleSubmit } = createForm({
  server: z.string(),
  email: z.string().email(),
  subject: z.string().min(1),
  body: z.string().min(1),
  file: z.any(),
})

// Composable
const { t } = useI18n()
const { toast } = useToast()

// Methods
const onSubmit = handleSubmit(async (value) => {
  console.log(value)
  toast({
    description: t('form.toast.success'),
  })
})

const onSubmitThrottle = useThrottleFn(() => onSubmit(), 1000)
</script>

<template>
  <div class="flex h-full items-center justify-center">
    <shad-card class="w-96">
      <shad-card-header>
        <div class="flex items-center justify-between">
          <h3>{{ $t('navbar.title.form') }}</h3>
          <shad-tooltip-provider>
            <shad-tooltip>
              <shad-tooltip-trigger><PhosphorIconQuestion size="20" /></shad-tooltip-trigger>
              <shad-tooltip-content>
                <p>{{ $t('form.tooltip') }}</p>
              </shad-tooltip-content>
            </shad-tooltip>
          </shad-tooltip-provider>
        </div>
      </shad-card-header>
      <shad-card-content>
        <form class="flex flex-col gap-y-5" @submit.prevent="onSubmitThrottle">
          <form-field v-slot="{ componentField }" name="server">
            <shad-form-item>
              <shad-form-label>{{ $t('form.input.serverLabel') }}</shad-form-label>
              <shad-form-control>
                <shad-select v-bind="componentField">
                  <shad-select-trigger>
                    <shad-select-value :placeholder="$t('form.input.serverPlaceholder')" />
                  </shad-select-trigger>
                  <shad-select-content>
                    <shad-select-group>
                      <shad-select-item value="predict-smtp"> predict-smtp </shad-select-item>
                    </shad-select-group>
                  </shad-select-content>
                </shad-select>
              </shad-form-control>
              <shad-form-message />
            </shad-form-item>
          </form-field>
          <form-field v-slot="{ componentField }" name="email">
            <shad-form-item>
              <shad-form-label>{{ $t('form.input.emailLabel') }}</shad-form-label>
              <shad-form-control>
                <shad-input v-bind="componentField" id="email" type="email" placeholder="email@test.com" />
              </shad-form-control>
              <shad-form-message />
            </shad-form-item>
          </form-field>
          <form-field v-slot="{ componentField }" name="subject">
            <shad-form-item>
              <shad-form-label>{{ $t('form.input.subjectLabel') }}</shad-form-label>
              <shad-form-control>
                <shad-input v-bind="componentField" id="subject" :placeholder="$t('form.input.subjectPlaceholder')" />
              </shad-form-control>
              <shad-form-message />
            </shad-form-item>
          </form-field>
          <form-field v-slot="{ componentField }" name="body">
            <shad-form-item>
              <shad-form-label>{{ $t('form.input.bodyLabel') }}</shad-form-label>
              <shad-form-control>
                <shad-textarea v-bind="componentField" id="body" class="max-h-60" :placeholder="$t('form.input.bodyPlaceholder')" />
              </shad-form-control>
              <shad-form-message />
            </shad-form-item>
          </form-field>
          <form-field v-slot="{ componentField }" name="file">
            <shad-form-item>
              <shad-form-label>{{ $t('form.input.fileLabel') }}</shad-form-label>
              <shad-form-control>
                <shad-input v-bind="componentField" id="file" type="file" :placeholder="$t('form.input.subjectPlaceholder')" />
              </shad-form-control>
              <shad-form-message />
            </shad-form-item>
          </form-field>
        </form>
      </shad-card-content>
      <shad-card-footer>
        <div class="flex justify-center">
          <shad-button class="mx-auto" @click.prevent="onSubmitThrottle">
            <PhosphorIconPaperPlaneTilt size="16" />
            <span class="ml-3">{{ $t('form.btnLabel') }}</span>
          </shad-button>
        </div>
      </shad-card-footer>
    </shad-card>
  </div>
</template>
