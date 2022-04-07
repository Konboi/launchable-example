import os
import hashlib
from pathlib import Path

DEFAULT_SESSION_DIR = '~/.config/launchable/sessions/'
SESSION_DIR_KEY = 'LAUNCHABLE_SESSION_DIR'


def _session_file_dir() -> Path:
    Path(os.environ.get(SESSION_DIR_KEY) or DEFAULT_SESSION_DIR).expanduser()


def _session_file_path(build_name: str) -> Path:
    print("{}:{}".format(build_name, os.getsid(os.getpid())))
    return str(_session_file_dir()) + "/" + (hashlib.sha1("{}:{}".format(build_name, os.getsid(os.getpid())).encode()).hexdigest() + ".txt")


print(_session_file_path(os.environ.get('CIRCLE_WORKFLOW_ID')))
