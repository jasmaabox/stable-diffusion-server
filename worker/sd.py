import torch
from diffusers import StableDiffusionPipeline, LMSDiscreteScheduler

MODEL_NAME = "hakurei/waifu-diffusion"
DISABLE_SAFTEY_CHECKER = True  # enabling safety filter will cause memory issues
DEVICE = "cuda" if torch.cuda.is_available() else "cpu"

lms = LMSDiscreteScheduler(
    beta_start=0.00085,
    beta_end=0.012,
    beta_schedule="scaled_linear"
)

txt2img_pipe = StableDiffusionPipeline.from_pretrained(
    MODEL_NAME,
    scheduler=lms,
    use_auth_token=True
).to(DEVICE)

if DISABLE_SAFTEY_CHECKER:
    txt2img_pipe.safety_checker = lambda images, **kwargs: (images, False)
