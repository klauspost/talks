    // Create a decoder
    d, err := zstd.NewReader(nil)
    if err != nil {
      return err
    }
    defer d.Close()

    // Decode block:
    dst, err := decoder.DecodeAll(src, nil)

    // Or decompress a stream.
    d.Reset(reader)
    _, err = io.Copy(ioutil.Discard, d)
