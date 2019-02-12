# Machine Learning Frameworks


<!-- @import "[TOC]" {cmd="toc" depthFrom=2 depthTo=6 orderedList=false} -->

<!-- code_chunk_output -->

* [Tensorflow](#tensorflow)
	* [MacOS](#macos)
	* [Test and Hello World](#test-and-hello-world)
* [Cuda (related issue)](#cuda-related-issue)
	* [Show Memory](#show-memory)
	* [Library not found.](#library-not-found)

<!-- /code_chunk_output -->

## Tensorflow

### MacOS

* (mac pro) GPU: You can try to use one of the [precompiled](https://github.com/lakshayg/tensorflow-build) [tensorflow](https://storage.googleapis.com/74thopen/tensorflow_osx/index.html) binaries or [compile it yourself](https://github.com/butuzov/tensorflow-gpu-macosx).

* (mac pro) CPU: **not supported** (atm).

  Build it yourself next time you will have time.


* (macbook 13") CPU: Multiple cpu issues

  * https://github.com/lakshayg/tensorflow-build
  * https://stackoverflow.com/questions/47068709

  ```bash
  python3.6 -m venv .venv
  . .venv/bin/activate
  pip install --ignore-installed --upgrade "https://github.com/lakshayg/tensorflow-build/releases/download/tf1.9.0-macos-py27-py36/tensorflow-1.9.0-cp36-cp36m-macosx_10_13_x86_64.whl"
  ```

### Test and Hello World

  ```python
  import tensorflow as tf
  print("GPU SUPPORTED?", tf.test.is_gpu_available())

  # list gpu devices
  from tensorflow.python.client import dl
  print("Supported Devices\n",
      [(x.device_type, x.name) for x in dl.list_local_devices()]
  )

  # Hello, TensorFlow!
  hello = tf.constant('Hello, TensorFlow!')
  sess = tf.Session()
  print(sess.run(hello))
  ```


## Cuda (related issue)

### Show Memory

  * https://github.com/micahstubbs/cuda-smi/releases/tag/v1.0
  * https://github.com/phvu/cuda-smi

### Library not found.
  Happend while practicing with [Gorgonia](https://github.com/gorgonia)

  ```bash
  > dyld: Library not loaded: @rpath/libcudart.10.0.dylib
  ```

  ```bash
  # install
  for i in /Developer/NVIDIA/CUDA-10.0/lib/*.a /Developer/NVIDIA/CUDA-10.0/lib/*.dylib; do
    ln -sv "$i" "/usr/local/lib/$(basename "$i")";
  done

  # unintstall
  for i in /Developer/NVIDIA/CUDA-10.0/lib/*.a /Developer/NVIDIA/CUDA-10.0/lib/*.dylib; do
    rm -v "/usr/local/lib/$(basename "$i")";
  done
  ```
