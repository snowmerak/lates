// Code generated by templ@v0.2.186 DO NOT EDIT.

package made

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import "context"
import "io"
import "bytes"

func Join(comp ...templ.Component) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = new(bytes.Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_1 := templ.GetChildren(ctx)
		if var_1 == nil {
			var_1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		// For
		for _, c := range comp {
			// CallTemplate
			err = c.Render(ctx, templBuffer)
			if err != nil {
				return err
			}
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}

func OrderedList(comp ...templ.Component) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = new(bytes.Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_2 := templ.GetChildren(ctx)
		if var_2 == nil {
			var_2 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		// Element (standard)
		_, err = templBuffer.WriteString("<ol>")
		if err != nil {
			return err
		}
		// For
		for _, c := range comp {
			// Element (standard)
			_, err = templBuffer.WriteString("<li>")
			if err != nil {
				return err
			}
			// CallTemplate
			err = c.Render(ctx, templBuffer)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</li>")
			if err != nil {
				return err
			}
		}
		_, err = templBuffer.WriteString("</ol>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}

func UnorderedList(comp ...templ.Component) templ.Component {
	return templ.ComponentFunc(func(ctx context.Context, w io.Writer) (err error) {
		templBuffer, templIsBuffer := w.(*bytes.Buffer)
		if !templIsBuffer {
			templBuffer = new(bytes.Buffer)
		}
		ctx = templ.InitializeContext(ctx)
		var_3 := templ.GetChildren(ctx)
		if var_3 == nil {
			var_3 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		// Element (standard)
		_, err = templBuffer.WriteString("<ul>")
		if err != nil {
			return err
		}
		// For
		for _, c := range comp {
			// Element (standard)
			_, err = templBuffer.WriteString("<li>")
			if err != nil {
				return err
			}
			// CallTemplate
			err = c.Render(ctx, templBuffer)
			if err != nil {
				return err
			}
			_, err = templBuffer.WriteString("</li>")
			if err != nil {
				return err
			}
		}
		_, err = templBuffer.WriteString("</ul>")
		if err != nil {
			return err
		}
		if !templIsBuffer {
			_, err = io.Copy(w, templBuffer)
		}
		return err
	})
}

