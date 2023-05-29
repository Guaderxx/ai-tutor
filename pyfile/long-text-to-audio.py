import os
import sys
from bark.generation import (
    generate_text_semantic,
    preload_models,
)
from bark.api import semantic_to_waveform
from bark import generate_audio, SAMPLE_RATE
from scipy.io.wavfile import write as write_wav
import nltk
import numpy as np

# use minimum model
os.environ["SUNO_OFFLOAD_CPU"] = "True"
os.environ["SUNO_USE_SMALL_MODELS"] = "True"
os.environ["CUDA_VISIBLE_DEVICES"] = "0"

# download and load all models
preload_models()

text = sys.argv[1]
output_path = sys.argv[2]
text_prompt = text.replace("\n", " ").strip()

sentences = nltk.sent_tokenize(text_prompt)

SPEAKER = "v2/en_speaker_9"
silence = np.zeros(int(0.25 * SAMPLE_RATE))  # quarter second of silence

pieces = []
for sen in sentences:
    audio_array = generate_audio(sen, history_prompt=SPEAKER)
    pieces += [audio_array, silence.copy()]

audio = np.concatenate(pieces)
# save audio to disk
write_wav(output_path, SAMPLE_RATE, audio)
